package rlib

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

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

func rpnPrintStack(ctx *RpnCtx) {
	fmt.Printf("Stack --- size: %d\n", len(ctx.stack))
	for i := 0; i < len(ctx.stack); i++ {
		fmt.Printf("%2d: %f\n", i, ctx.stack[i])
	}
}

// RpnInit initializes the Rpn calculator routines
func RpnInit() {
	rpnVariable = regexp.MustCompile("{(.*)}")
	rpnOperator = regexp.MustCompile("[\\-+*/%]")
	rpnNumber = regexp.MustCompile("^\\d+\\.?[0-9]+")
	rpnFunction = regexp.MustCompile("([a-zA-Z]+)\\(([^\\)]+)\\)")
	rpnASM = regexp.MustCompile("^ASM\\(([^)]+)\\)")
}

func rpnPop(ctx *RpnCtx) float64 {
	l := len(ctx.stack)
	if l > 0 {
		x := ctx.stack[l-1]
		ctx.stack = ctx.stack[0 : l-1]
		return x
	}
	return 0
}

func rpnPush(ctx *RpnCtx, x float64) {
	ctx.stack = append(ctx.stack, x*ctx.pf)
}

func rpnLoadRentable(ctx *RpnCtx) {
	// only load it if necessary
	if 0 == ctx.xu.R.RID {
		GetXRentable(ctx.rid, &ctx.xu)
	}
}

// RpnCreateCtx creates the context structure needed for use with all the Rpn functions
func RpnCreateCtx(xbiz *XBusiness, rid int64, d1, d2 *time.Time, m *[]AcctRule, amount, pf float64) RpnCtx {
	var ctx RpnCtx
	ctx.xbiz = xbiz
	ctx.m = m
	ctx.d1 = d1
	ctx.d2 = d2
	ctx.rid = rid
	ctx.stack = make([]float64, 0)
	ctx.pf = pf
	ctx.amount = amount
	ctx.GSRset = false
	return ctx
}

func rpnFunctionResolve(ctx *RpnCtx, cmd, val string) float64 {
	switch {
	case cmd == "aval":
		if val[0] == '$' {
			val = DoAcctSubstitution(ctx.xbiz.P.BID, val) // could be a substitution
		}
		for i := 0; i < len(*ctx.m); i++ {
			if (*ctx.m)[i].Account == val {
				// fmt.Printf("rpnFunctionResolve: returning %f\n", (*ctx.m)[i].Amount)
				return (*ctx.m)[i].Amount
			}
		}
	default:
		Ulog("rpnFunctionResolve: unrecognized function: %s\n", cmd)
	}
	return float64(0)
}

func varResolve(ctx *RpnCtx, s string) float64 {

	if s == "UMR" { // Unit MARKET RATE
		rpnLoadRentable(ctx) // make sure it's loaded
		return ctx.pf * GetRentableMarketRate(ctx.xbiz, &ctx.xu.R, ctx.d1, ctx.d2)
	}
	if s == "GSR" { // Gross Schedule Rent = Market Rate + Specialties
		if ctx.GSRset { // don't recalculate if already set
			return ctx.pf * ctx.GSR
		}
		rpnLoadRentable(ctx) // make sure it's loaded
		amt, _, _, err := CalculateLoadedGSR(&ctx.xu.R, ctx.d1, ctx.d2, ctx.xbiz)
		if err == nil {
			// fmt.Printf("varResolve: amt = %f, d1 = %s, d2 = %s\n", amt, ctx.d1.Format(RRDATEFMT4), ctx.d2.Format(RRDATEFMT4))
			ctx.GSR = amt
			ctx.GSRset = true
			return ctx.pf * ctx.GSR
		}
	}
	if s == "ASM.Amount" { // the amount of the associated assessment
		a, err := GetAssessment(ctx.r.ASMID)
		if nil != err {
			Ulog("varResolve: could not load Assessment %d. err = %s\n", ctx.r.ASMID, err.Error())
			return float64(0)
		}
		return ctx.pf * a.Amount
	}
	m1 := rpnFunction.FindAllStringSubmatchIndex(s, -1)
	if m1 != nil {
		m := m1[0]
		cmd := s[m[2]:m[3]]
		val := s[m[4]:m[5]]
		return rpnFunctionResolve(ctx, cmd, val)
	}

	return float64(0)
}

// RpnCalculateEquation takes a formula, parses and executes the formula and returns the number it calculates
func RpnCalculateEquation(ctx *RpnCtx, s string) float64 {
	// funcname := "RpnCalculateEquation"
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
					n := varResolve(ctx, match)
					ctx.stack = append(ctx.stack, n)
				}
			} else if s[0] == '_' {
				rpnPush(ctx, ctx.amount)
			} else if ('0' <= s[0] && s[0] <= '9') || '.' == s[0] { // is it a number?
				m := rpnNumber.FindStringSubmatchIndex(s)
				match := s[m[0]:m[1]]
				n, _ := strconv.ParseFloat(match, 64)
				ctx.stack = append(ctx.stack, n*ctx.pf)
			} else if s[0] == '-' || s[0] == '+' || s[0] == '*' || s[0] == '/' { // is it an operator?
				op := s[0:1]
				var x, y float64
				y = rpnPop(ctx)
				x = rpnPop(ctx)
				switch op {
				case "+":
					ctx.stack = append(ctx.stack, x+y)
				case "-":
					ctx.stack = append(ctx.stack, x-y)
				case "*":
					ctx.stack = append(ctx.stack, x*y)
				case "/":
					ctx.stack = append(ctx.stack, x/y)
				}
			}
		}
		// rpnPrintStack(ctx)
	}
	return rpnPop(ctx)
}
