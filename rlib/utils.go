package rlib

// IsBIDExists checks that BID is available or not
// It checks it through db cache, not actualy by hitting to DB
func IsBIDExists(BID int64) bool {
	for _, bid := range RRdb.BUDlist {
		if bid == BID {
			return true
		}
	}
	return false
}

// IsBUDExists checks that BUD is available or not
// It checks it through db cache, not actualy by hitting to DB
func IsBUDExists(BUD string) bool {
	for bud := range RRdb.BUDlist {
		if bud == BUD {
			return true
		}
	}
	return false
}
