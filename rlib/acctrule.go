package rlib

import (
	"fmt"
	"strings"
	"time"
)

// AcctRule is a structure of the 3-tuple that makes up a whole part of an AcctRule
type AcctRule struct {
	Action      string  // "d" = debit, "c" = credit
	Account     string  // GL No for the account
	AccountOrig string  // account before substitution
	Amount      float64 // use the entire amount of the assessment or deposit, otherwise the amount to use
	ASMID       int64   // Used only for ReceiptAllocation; the assessment that caused this payment
	Expr        string  // the formula of the Amount
	AcctExpr    string  // the input Acct Expression -- may be the same as the GLNo or may be a ${ref}
}

// VarAcctResolve replaces string references with the appropriate values for variable account names
func VarAcctResolve(bid int64, s string) string {

	// fmt.Printf("VarAcctResolve( %d, %q )\n", bid, s)

	i := int64(0)
	switch {
	case s == "GLCASH":
		i = GLCASH
	case s == "GLGENRCV":
		i = GLGENRCV
	case s == "GLGSRENT":
		i = GLGSRENT
	case s == "GLLTL":
		i = GLLTL
	case s == "GLVAC":
		i = GLVAC
	// case s == "GLSECDEPRCV":
	// 	i = GLSECDEPRCV
	case s == "GLSECDEP":
		i = GLSECDEP
	case s == "GLOWNREQUITY":
		i = GLOWNREQUITY
	}
	if i > 0 {
		// fmt.Printf("VarAcctResolve: i = %d, returning %s\n", i, RRdb.BizTypes[bid].DefaultAccts[i].GLNumber)
		return RRdb.BizTypes[bid].DefaultAccts[i].GLNumber
	}
	// fmt.Printf("VarAcctResolve:  returning %s\n", s)
	return s
}

// DoAcctSubstitution replaces variables with their values
func DoAcctSubstitution(bid int64, s string) string {
	// fmt.Printf("Entering DoAcctSubstitution. bid=%d, s=%s\n", bid, s)
	if s[0] == '$' {
		m := rpnVariable.FindStringSubmatchIndex(s)
		if m != nil {
			match := s[m[2]:m[3]]
			return VarAcctResolve(bid, match)
		}
	}
	return s
}

// ParseAcctRule expands the supplied rule string into an array of AcctRule structs and replaces any variables/formulas
// with the final amounts.
// INPUTS:
//     xbiz - XBusiness struct for this business
//      rid - the associated Rentable ID (if needed)
//    d1,d2 - time period being examined
//     rule - the actual account rule to parse
//   amount - total amount of this transaction
//       pf - the proration factor.  1.0 if we're applying the amount for the entire period, otherwise (days applicable)/(days in period)
//
// RETURNS:
//     a slice of AcctRule structs that make up the account rule
func ParseAcctRule(xbiz *XBusiness, rid int64, d1, d2 *time.Time, rule string, amount, pf float64) []AcctRule {
	funcname := "ParseAcctRule"
	var m []AcctRule
	// fmt.Printf("%s:  rid = %d, d1 = %s, d2 = %s, rule = %s, amount = %f, pf = %f, xbiz.P.BID = %d\n", funcname, rid, d1.Format(RRDATEFMT4), d2.Format(RRDATEFMT4), rule, amount, pf, xbiz.P.BID)
	ctx := RpnCreateCtx(xbiz, rid, d1, d2, &m, amount, pf)
	// fmt.Printf("ctx.Amount = %f\n", ctx.amount)
	if len(rule) > 0 {
		sa := strings.Split(rule, ",")
		for k := 0; k < len(sa); k++ {
			// fmt.Printf("%s:  k = %d\n", funcname, k)
			// fmt.Printf("\tsa[k] = %s\n", sa[k])
			var r AcctRule
			t := strings.Join(strings.Fields(sa[k]), " ") // this puts 1 space between every field in sa[k]
			ta := strings.Split(t, " ")                   // an array of fields
			base := 0                                     // assume the main 3 fields start at index 0
			if strings.HasPrefix(ta[0], "ASM") {          // if the first string is of the form "ASM(x)" then we have 4 fields, otherwise we'll have 3
				base = 1                              // base moves by one
				a := rpnASM.FindStringSubmatch(ta[0]) // need to find the assessment id
				if len(a) != 2 {
					LogAndPrintError(funcname, fmt.Errorf("%s: invalid assessment identifier: %s", funcname, ta[0]))
				} else {
					var err error
					r.ASMID, err = IntFromString(a[1], "Invalid Assessment ID")
					CheckLogAndPrintError(funcname, err)
				}
			}
			r.Action = strings.ToLower(strings.TrimSpace(ta[base])) // action is at index base
			r.AcctExpr = strings.TrimSpace(ta[base+1])              // account is at base+1, this is the source
			r.Account = DoAcctSubstitution(xbiz.P.BID, r.AcctExpr)  // the is the substituted acct name
			ar := strings.Join(ta[base+2:], " ")                    // remaining fields make up the amount formula
			r.Expr = strings.TrimSpace(ar)                          // prepare the formula for the calculator
			ctx.r = &r                                              // the AcctRule in the process of being constructed.  Has the Assessment ID which may be needed.
			// fmt.Printf("ctx = %#v\n", ctx)
			// fmt.Printf("r.Expr = %s\n", r.Expr)
			x := RpnCalculateEquation(&ctx, r.Expr) // let the calculator compute the amount
			// fmt.Printf("\ncalc returned x = %8.2f\n\n", x)
			r.Amount = x     // set the Amount field
			m = append(m, r) // and we're done
		}
	}
	return m
}
