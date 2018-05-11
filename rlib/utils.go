package rlib

// BIDExists checks that BID is available or not
// It checks it through db cache, not actually by hitting to DB
func BIDExists(BID int64) bool {
	for _, bid := range RRdb.BUDlist {
		if bid == BID {
			return true
		}
	}
	return false
}

// BUDExists checks that BUD is available or not
// It checks it through db cache, not actually by hitting to DB
func BUDExists(BUD string) bool {
	for bud := range RRdb.BUDlist {
		if bud == BUD {
			return true
		}
	}
	return false
}

// GetBUDFromBIDList return the BUD for BID from pre-populated
// list of BUD:BID map, i.e, RRdb.BUDlist
func GetBUDFromBIDList(BID int64) XJSONBud {
	var BUD string
	for bud, bid := range RRdb.BUDlist {
		if bid == BID {
			BUD = bud
			break
		}
	}
	return XJSONBud(BUD)
}
