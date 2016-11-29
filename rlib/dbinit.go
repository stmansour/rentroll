package rlib

import "time"

// DefaultAccts are the required accounts for all businesses
var NewBizDefaultAccts = []GLAccount{
	{Status: 2, Type: GLCASH, Name: "Bank Account"},
	{Status: 2, Type: GLGENRCV, Name: "General Accounts Receivable"},
	{Status: 2, Type: GLGSRENT, Name: "Gross Scheduled Rent"},
	{Status: 2, Type: GLLTL, Name: "Loss to Lease"},
	{Status: 2, Type: GLVAC, Name: "Vacancy"},
	{Status: 2, Type: 15, Name: "Security Deposit Receivable"},
	{Status: 2, Type: GLSECDEP, Name: "Security Deposit Assessment"},
	{Status: 2, Type: GLOWNREQUITY, Name: "Owner Equity"},
}

// NewBusinessInit should be called immediately after creating a new business in order to
// create the initial database information needed to support it.
func NewBusinessInit(bid int64) error {

	for i := 0; i < len(NewBizDefaultAccts); i++ {
		l := NewBizDefaultAccts[i]
		l.BID = bid
		lid, err := InsertLedger(&l)
		if err != nil {
			return err
		}
		var lm LedgerMarker
		lm.BID = bid
		lm.LID = lid
		lm.State = 3
		lm.Dt = time.Date(2000, time.January, 0, 0, 0, 0, 0, time.UTC)
		err = InsertLedgerMarker(&lm)
		if err != nil {
			return err
		}
	}
	return nil
}
