package ws

import (
	"rentroll/rlib"
)

// getBUDFromBIDList return the BUD for BID from pre-populated
// list of BUD:BID map, i.e, rlib.RRdb.BUDlist
func getBUDFromBIDList(BID int64) rlib.XJSONBud {
	var BUD string
	for bud, bid := range rlib.RRdb.BUDlist {
		if bid == BID {
			BUD = bud
			break
		}
	}
	return rlib.XJSONBud(BUD)
}
