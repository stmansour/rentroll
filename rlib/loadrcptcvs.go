package rlib

// type Receipt struct {
// 	RCPTID   int64
// 	BID      int64
// 	RAID     int64
// 	PMTID    int64
// 	Dt       time.Time
// 	Amount   float64
// 	AcctRule string
// 	Comment  string
// 	RA       []ReceiptAllocation
// }

// 0    1     2      3   4       5
// BID, RAID, PMTID, Dt, Amount, AcctRule
// REH,RA00000001,2,"2004-01-01", 1000.00, "d ${DFLTCASH} _, c 11002 _"
// REH,RA00000001,1,"2015-11-21",  294.66, "c ${DFLTGENRCV} 266.67, d ${DFLTCASH} 266.67, c ${DFLTGENRCV} 13.33, d ${DFLTCASH} 13.33, c ${DFLTGENRCV} 5.33, d ${DFLTCASH} 5.33, c ${DFLTGENRCV} 9.33,d ${DFLTCASH} 9.33"
