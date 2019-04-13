package rlib

import (
	"context"
	"time"
)

// CloseInfo contains relevant close period information for a business
type CloseInfo struct {
	BID       int64     // Business ID
	LastClose time.Time // last closed period
	CPID      int64     // id of last close
	BKDTRACP  bool      // backdate rental agreements in closed period allowed?
}

// GetCloseInfo returns a struct of information about closed periods and
// related business prefs
//
// INPUTS
//     ctx  = db transaction context
//     BID  = Business ID
//
// RETURNS
//     RAPetsFlowData structure
//     any error encountered
//-----------------------------------------------------------------------------
func GetCloseInfo(ctx context.Context, bid int64) (CloseInfo, error) {
	var ci CloseInfo
	var err error
	var b Business
	var lc ClosePeriod

	ci.BID = bid
	lc, err = GetLastClosePeriod(ctx, bid)
	if err != nil {
		return ci, err
	}
	ci.LastClose = lc.Dt
	ci.CPID = lc.CPID
	if err = getBiz(bid, &b); err != nil {
		return ci, err
	}

	ci.BKDTRACP = b.FLAGS&(1<<1) != 0
	return ci, err
}

// GetAllBizCloseInfo returns a map of CloseInfo structs keyed by the designator
//
// INPUTS
//     ctx  = db transaction context
//     BID  = Business ID
//
// RETURNS
//     RAPetsFlowData structure
//     any error encountered
//-----------------------------------------------------------------------------
func GetAllBizCloseInfo(ctx context.Context) (map[string]CloseInfo, error) {
	var m = map[string]CloseInfo{}
	n, err := GetAllBiz(ctx)
	if err != nil {
		return m, err
	}
	for i := 0; i < len(n); i++ {
		var ci CloseInfo
		var lc ClosePeriod
		ci.BID = n[i].BID
		lc, err = GetLastClosePeriod(ctx, ci.BID)
		if err != nil {
			return m, err
		}
		ci.LastClose = lc.Dt
		ci.CPID = lc.CPID
		ci.BKDTRACP = n[i].FLAGS&(1<<1) != 0
		m[n[i].Designation] = ci
	}
	return m, nil
}
