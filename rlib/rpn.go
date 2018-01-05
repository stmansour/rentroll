package rlib

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// https://play.golang.org/p/p842UZpQaK

// RpnCtx defines the context structure needed by all the Rpn routines
type RpnCtx struct {
	xbiz   *XBusiness  // the biz associated with this assessment/payment
	m      *[]AcctRule // READ-ONLY access to the rule array being created
	xu     XRentable   // the Rentable associated with this rule, loaded only if needed
	rid    int64       // Rentable id
	d1     *time.Time  // start of time range
	d2     *time.Time  // end of time range
	pf     float64     // proration factor
	amount float64     // the full amount of the assessment or payment
	stack  []float64   // the stack used by the rpn calculator
	GSRset bool        // initially false, set to true after GSR is calculated
	GSR    float64     // this is a heavyweight calculation. If GSRset is true, then don't recalculate, just use current value
	r      *AcctRule   // the account rule in the process of being constructed
}

var rpnVariable *regexp.Regexp
var rpnOperator *regexp.Regexp
var rpnNumber *regexp.Regexp
var rpnFunction *regexp.Regexp
var rpnASM *regexp.Regexp
var rpnSUM *regexp.Regexp

func rpnPrintStack(rpnCtx *RpnCtx) {
	fmt.Printf("Stack --- size: %d\n", len(rpnCtx.stack))
	for i := 0; i < len(rpnCtx.stack); i++ {
		fmt.Printf("%2d: %f\n", i, rpnCtx.stack[i])
	}
}

// RpnInit initializes the Rpn calculator routines
func RpnInit() {
	rpnVariable = regexp.MustCompile("{(.*)}")
	rpnOperator = regexp.MustCompile(`[\-+*/%]`)
	rpnNumber = regexp.MustCompile(`^-?\d+\.?[0-9]+`)
	rpnFunction = regexp.MustCompile(`([a-zA-Z]+)\(([^\)]+)\)`)
	rpnASM = regexp.MustCompile(`^ASM\(([^)]+)\)`)
}

func rpnPop(rpnCtx *RpnCtx) float64 {
	l := len(rpnCtx.stack)
	if l > 0 {
		x := rpnCtx.stack[l-1]
		rpnCtx.stack = rpnCtx.stack[0 : l-1]
		return x
	}
	return 0
}

func rpnPush(rpnCtx *RpnCtx, x float64) {
	rpnCtx.stack = append(rpnCtx.stack, x*rpnCtx.pf)
}

func rpnLoadRentable(ctx context.Context, rpnCtx *RpnCtx) error {
	var (
		err error
	)

	// only load it if necessary
	if 0 == rpnCtx.xu.R.RID {
		err = GetXRentable(ctx, rpnCtx.rid, &rpnCtx.xu)
		if err != nil {
			return err
		}
	}

	return err
}

// RpnCreateCtx creates the context structure needed for use with all the Rpn functions
func RpnCreateCtx(xbiz *XBusiness, rid int64, d1, d2 *time.Time, m *[]AcctRule, amount, pf float64) RpnCtx {
	var rpnCtx RpnCtx
	rpnCtx.xbiz = xbiz
	rpnCtx.m = m
	rpnCtx.d1 = d1
	rpnCtx.d2 = d2
	rpnCtx.rid = rid
	rpnCtx.stack = make([]float64, 0)
	rpnCtx.pf = pf
	rpnCtx.amount = amount
	rpnCtx.GSRset = false
	return rpnCtx
}

func rpnFunctionResolve(rpnCtx *RpnCtx, cmd, val string) float64 {
	switch {
	case cmd == "aval":
		if val[0] == '$' {
			val = DoAcctSubstitution(rpnCtx.xbiz.P.BID, val) // could be a substitution
		}
		for i := 0; i < len(*rpnCtx.m); i++ {
			if (*rpnCtx.m)[i].Account == val {
				// fmt.Printf("rpnFunctionResolve: returning %f\n", (*rpnCtx.m)[i].Amount)
				return (*rpnCtx.m)[i].Amount
			}
		}
	default:
		Ulog("rpnFunctionResolve: unrecognized function: %s\n", cmd)
	}
	return float64(0)
}

func varResolve(ctx context.Context, rpnCtx *RpnCtx, s string) (float64, error) {
	var (
		err error
		val float64
	)

	if s == "UMR" { // Unit MARKET RATE
		err = rpnLoadRentable(ctx, rpnCtx) // make sure it's loaded
		if err != nil {
			return val, err
		}

		mr, err := GetRentableMarketRate(ctx, rpnCtx.xbiz, rpnCtx.xu.R.RID, rpnCtx.d1, rpnCtx.d2)
		if err != nil {
			return val, err
		}
		return rpnCtx.pf * mr, err
	}

	if s == "GSR" { // Gross Schedule Rent = Market Rate + Specialties
		if rpnCtx.GSRset { // don't recalculate if already set
			return rpnCtx.pf * rpnCtx.GSR, err
		}
		err = rpnLoadRentable(ctx, rpnCtx) // make sure it's loaded
		if err != nil {
			return val, err
		}

		amt, _, _, err := CalculateLoadedGSR(ctx, rpnCtx.xu.R.BID, rpnCtx.xu.R.RID, rpnCtx.d1, rpnCtx.d2, rpnCtx.xbiz)
		if err != nil {
			return val, err
		}
		// fmt.Printf("varResolve: amt = %f, d1 = %s, d2 = %s\n", amt, rpnCtx.d1.Format(RRDATEFMT4), rpnCtx.d2.Format(RRDATEFMT4))
		rpnCtx.GSR = amt
		rpnCtx.GSRset = true
		return rpnCtx.pf * rpnCtx.GSR, err
	}

	if s == "ASM.Amount" { // the amount of the associated assessment
		a, err := GetAssessment(ctx, rpnCtx.r.ASMID)
		if nil != err {
			Ulog("varResolve: could not load Assessment %d. err = %s\n", rpnCtx.r.ASMID, err.Error())
			return val, err
		}

		return rpnCtx.pf * a.Amount, err
	}

	m1 := rpnFunction.FindAllStringSubmatchIndex(s, -1)
	if m1 != nil {
		m := m1[0]
		cmd := s[m[2]:m[3]]
		val := s[m[4]:m[5]]
		return rpnFunctionResolve(rpnCtx, cmd, val), err
	}

	return val, err
}

// RpnCalculateEquation takes a formula, parses and executes the formula and returns the number it calculates.
// This may be helpful: https://play.golang.org/p/p842UZpQaK
func RpnCalculateEquation(ctx context.Context, rpnCtx *RpnCtx, s string) (float64, error) {
	// funcname := "RpnCalculateEquation"

	var (
		err error
	)

	// fmt.Printf("%s: entered\n", funcname)
	t := strings.Split(s, " ")
	// fmt.Printf("%s: t = %#v\n", funcname, t)

	for i := 0; i < len(t); i++ {
		s = t[i]
		// fmt.Printf("\n%s: for loop parsing: %s\n", funcname, s)
		if len(s) > 0 {
			if s[0] == '$' { // is it a special notation?
				m := rpnVariable.FindStringSubmatchIndex(s)
				if m != nil {
					match := s[m[2]:m[3]]
					n, err := varResolve(ctx, rpnCtx, match)
					if err != nil {
						return float64(0), err
					}

					rpnCtx.stack = append(rpnCtx.stack, n)
				}
			} else if s[0] == '_' {
				// fmt.Printf("%s: found '_', pushing rpnCtx.amount = %8.2f\n", funcname, rpnCtx.amount)
				rpnPush(rpnCtx, rpnCtx.amount)
			} else if ('0' <= s[0] && s[0] <= '9') || '.' == s[0] { // is it a number?
				m := rpnNumber.FindStringSubmatchIndex(s)
				match := s[m[0]:m[1]]
				n, _ := strconv.ParseFloat(match, 64)
				rpnCtx.stack = append(rpnCtx.stack, n*rpnCtx.pf)
			} else if len(s) > 1 && s[0] == '-' && (('0' <= s[1] && s[1] <= '9') || '.' == s[1]) {
				m := rpnNumber.FindStringSubmatchIndex(s)
				match := s[m[0]:m[1]]
				n, _ := strconv.ParseFloat(match, 64)
				rpnCtx.stack = append(rpnCtx.stack, n*rpnCtx.pf)
			} else if s[0] == '-' || s[0] == '+' || s[0] == '*' || s[0] == '/' { // is it an operator?
				op := s[0:1]
				var x, y float64
				y = rpnPop(rpnCtx)
				x = rpnPop(rpnCtx)
				switch op {
				case "+":
					rpnCtx.stack = append(rpnCtx.stack, x+y)
				case "-":
					rpnCtx.stack = append(rpnCtx.stack, x-y)
				case "*":
					rpnCtx.stack = append(rpnCtx.stack, x*y)
				case "/":
					rpnCtx.stack = append(rpnCtx.stack, x/y)
				}
			}
		}
		// rpnPrintStack(rpnCtx)
	}
	return rpnPop(rpnCtx), err
}
