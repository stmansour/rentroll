package main

import (
	"regexp"
	"strconv"
	"strings"
	"time"
)

type rpnCtx struct {
	xbiz  *XBusiness  // the biz associated with this assessment/payment
	m     *[]acctRule // READ-ONLY access to the rule array being created
	xu    XUnit       // the rentable associated with this rule, loaded only if needed
	rid   int64       // rentable id
	d1    *time.Time  // start of time range
	d2    *time.Time  // end of time range
	pf    float64     // proration factor
	stack []float64   // the stack used by the rpn calculator
}

var rpnVariable *regexp.Regexp
var rpnOperator *regexp.Regexp
var rpnNumber *regexp.Regexp
var rpnFunction *regexp.Regexp

func rpnInit() {
	rpnVariable = regexp.MustCompile("{([^}]+)}")
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

func rpnLoadRentable(ctx *rpnCtx) {
	// only load it if necessary
	if 0 == ctx.xu.U.UNITID {
		GetXUnit(ctx.rid, &ctx.xu)
	}
}

func rpnCreateCtx(xbiz *XBusiness, rid int64, d1, d2 *time.Time, m *[]acctRule, pf float64) rpnCtx {
	var ctx rpnCtx
	ctx.xbiz = xbiz
	ctx.m = m
	ctx.d1 = d1
	ctx.d2 = d2
	ctx.rid = rid
	ctx.stack = make([]float64, 0)
	ctx.pf = pf
	return ctx
}

func rpnFunctionResolve(ctx *rpnCtx, cmd, val string) float64 {
	switch {
	case cmd == "aval":
		for i := 0; i < len(*ctx.m); i++ {
			if (*ctx.m)[i].Account == val {
				// fmt.Printf("rpnFunctionResolve: returnning %f\n", (*ctx.m)[i].Amount)
				return (*ctx.m)[i].Amount
			}
		}
	default:
		ulog("rpnFunctionResolve: unrecognized function: %s\n", cmd)
	}
	return float64(0)
}

func varResolve(ctx *rpnCtx, s string) float64 {
	if s == "UMR" {
		rpnLoadRentable(ctx) // make sure it's loaded
		if ctx.xu.U.UNITID > 0 {
			return ctx.pf * GetUnitMarketRate(ctx.xbiz, &ctx.xu.U, ctx.d1, ctx.d2)
		}
		return ctx.pf * GetRentableMarketRate(ctx.xbiz, &ctx.xu.R, ctx.d1, ctx.d2)
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
	for len(s) > 0 {
		s = strings.TrimSpace(s)
		if s[0] == '$' { // is it a variable?
			m := rpnVariable.FindStringSubmatchIndex(s)
			if m != nil {
				match := s[m[2]:m[3]]
				n := varResolve(ctx, match)
				ctx.stack = append(ctx.stack, n)
				s = s[m[1]:]
			}
		} else if ('0' <= s[0] && s[0] <= '9') || '.' == s[0] { // is it a number?
			m := rpnNumber.FindStringSubmatchIndex(s)
			match := s[m[0]:m[1]]
			n, _ := strconv.ParseFloat(match, 64)
			ctx.stack = append(ctx.stack, n*ctx.pf)
			s = s[m[1]:]
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
			s = s[1:]
		} else {
			s = s[1:]
		}
	}
	z := rpnPop(ctx)
	return z
}
