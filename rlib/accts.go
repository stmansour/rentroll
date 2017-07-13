package rlib

import "strings"

const (
	acctsRcv = string("accounts receivable")
	secDep   = string("security deposit")
)

// getAccounts returns a list of the specified account LIDs
func getAccounts(bid int64, s string) []int64 {
	var m []int64
	for _, v := range RRdb.BizTypes[bid].GLAccounts {
		if strings.Contains(strings.ToLower(v.AcctType), s) || strings.Contains(strings.ToLower(v.Name), s) {
			m = append(m, v.LID)
		}
	}
	return m
}

// GetReceivableAccounts goes throughout the GLAccounts and returns
// an array of LIDs which are of type Receivables
func GetReceivableAccounts(bid int64) []int64 {
	return getAccounts(bid, acctsRcv)
}

// GetSecurityDepositsAccounts goes throughout the GLAccounts and returns
// an array of LIDs which are of type SecurityDeposits
func GetSecurityDepositsAccounts(bid int64) []int64 {
	return getAccounts(bid, secDep)
}
