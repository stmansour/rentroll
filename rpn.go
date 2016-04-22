package main

import (
	"fmt"
	"regexp"
	"rentroll/rlib"
	"strconv"
	"strings"
	"time"
)

type rpnCtx struct {
	xbiz   *rlib.XBusiness // the biz associated with this assessment/payment
	m      *[]acctRule     // READ-ONLY access to the rule array being created
	xu     rlib.XUnit      // the rentable associated with this rule, loaded only if needed
	rid    int64           // rentable id
	d1     *time.Time      // start of time range
	d2     *time.Time      // end of time range
	pf     float64         // proration factor
	amount float64         // the full amount of the assessment or payment
	stack  []float64       // the stack used by the rpn calculator
}

var rpnVariable *regexp.Regexp
var rpnOperator *regexp.Regexp
var rpnNumber *regexp.Regexp
var rpnFunction *regexp.Regexp

func rpnPrintStack(ctx *rpnCtx) {
	fmt.Printf("Stack --- size: %d\n", len(ctx.stack))
	for i := 0; i < len(ctx.stack); i++ {
		fmt.Printf("%2d: %f\n", i, ctx.stack[i])
	}
}

func rpnInit() {
	rpnVariable = regexp.MustCompile("{(.*)}")
	rpnOperator = regexp.MustCompile("[\\-+*/%]")
	rpnNumber = regexp.MustCompile("^\\d+\\.?[0-9]+")
	rpnFunction = regexp.MustCompile("([a-zA-Z]+)\\(([^\\)]+)\\)")
}

func rpnPop(ctx *rpnCtx) float64 {
	l := len(ctx.stack)
	if l > 0 {
		x := ctx.stack[l-1]
		ctx.stack = ctx.stack[0 : l-1]
		return x
	}
	return 0
}

func rpnPush(ctx *rpnCtx, x float64) {
	ctx.stack = append(ctx.stack, x*ctx.pf)
}

func rpnLoadRentable(ctx *rpnCtx) {
	// only load it if necessary
	if 0 == ctx.xu.R.RID {
		rlib.GetXRentable(ctx.rid, &ctx.xu)
	}
}

func rpnCreateCtx(xbiz *rlib.XBusiness, rid int64, d1, d2 *time.Time, m *[]acctRule, amount, pf float64) rpnCtx {
	var ctx rpnCtx
	ctx.xbiz = xbiz
	ctx.m = m
	ctx.d1 = d1
	ctx.d2 = d2
	ctx.rid = rid
	ctx.stack = make([]float64, 0)
	ctx.pf = pf
	ctx.amount = amount
	return ctx
}

func rpnFunctionResolve(ctx *rpnCtx, cmd, val string) float64 {
	switch {
	case cmd == "aval":
		if val[0] == '$' {
			val = doAcctSubstitution(ctx.xbiz.P.BID, val) // could be a substitution
		}
		for i := 0; i < len(*ctx.m); i++ {
			if (*ctx.m)[i].Account == val {
				// fmt.Printf("rpnFunctionResolve: returning %f\n", (*ctx.m)[i].Amount)
				return (*ctx.m)[i].Amount
			}
		}
	default:
		rlib.Ulog("rpnFunctionResolve: unrecognized function: %s\n", cmd)
	}
	return float64(0)
}

func varResolve(ctx *rpnCtx, s string) float64 {
	if s == "UMR" {
		rpnLoadRentable(ctx) // make sure it's loaded
		return ctx.pf * rlib.GetRentableMarketRate(ctx.xbiz, &ctx.xu.R, ctx.d1, ctx.d2)
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

func varParseAmount(ctx *rpnCtx, s string) float64 {
	t := strings.Split(s, " ")
	for i := 0; i < len(t); i++ {
		s = t[i]
		// fmt.Printf("\nfor loop parsing: %s\n", s)
		if s[0] == '$' { // is it a variable?
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
		// rpnPrintStack(ctx)
	}
	return rpnPop(ctx)
}
