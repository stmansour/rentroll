package rlib

// RSLeaseStatus is a slice of the string meaning of each LeaseStatus
var RSLeaseStatus = []string{
	"Unknown	",
	"Vacant-rented",
	"Vacant Not-Rented",
	"On-Notice-Preleased",
	"On-Notice-Available",
	"Leased",
	"Unavailable",
}

// RSUseStatus is a slice of the string meaning of each UseStatus
var RSUseStatus = []string{
	"Unknown",
	"Administrative",
	"In Service",
	"Employee",
	"Model",
	"Offline Rennovation",
	"Offline Maintenance",
	"Owner Occupied",
}

// RSMakeReadyStatus is a slice of the string meaning of each MakeReadyStatus
var RSMakeReadyStatus = []string{
	"Unknown",
	"In Progress Housekeeping",
	"In Progress Maintenance",
	"Pending Inspection",
	"Ready",
}

// RStat returns the status for the supplied rentable during the Periods
// in gaps.  If the status is defined by RentableStatus entries those
// values will be used.
//
// If there is no database status available for the rentable during the
// requested period, a RentableStatus struct will be returned with the
// following values:
//   RSID:        0 (which further reinforces that it's not in the database)
//   LeaseStatus: VacantNotRented or VacantRented, depending on whether
//                there is a RentableStatus record for the rentable.
//   UseStatus:   InService
//
// INPUTS
//   BID       - which business
//   RID       - the rentable id
//   gaps      - slice of Periods of interest
//
// RETURNS
//   []RentableStatus
//             - slice of RStatus structs that defines the state
//               of the rentable during each Period in gaps
//   error     - any error encountered
//-----------------------------------------------------------------------------
func RStat(bid, rid int64, gaps []Period) []RentableStatus {
	var m []RentableStatus
	for i := 0; i < len(gaps); i++ {
		rsa := GetRentableStatusByRange(rid, &gaps[i].D1, &gaps[i].D2)
		if len(rsa) == 0 {
			var rs = RentableStatus{
				BID:         bid,
				RID:         rid,
				DtStart:     gaps[i].D1,
				DtStop:      gaps[i].D2,
				UseStatus:   USESTATUSinService,
				LeaseStatus: LEASESTATUSvacantNotRented,
			}
			//----------------------------------------------------------------
			// If there is a RentalAgreement in the future, modify the status
			//----------------------------------------------------------------
			r := GetRentableStatusOnOrAfter(rid, &gaps[i].D1)
			if r.RSID > 0 {
				rs.LeaseStatus = LEASESTATUSvacantRented
			}
			// now add to return values
			m = append(m, rs)
		} else {
			m = append(m, rsa...)
		}
	}
	return m
}

// LeaseStatusStringer returns the string associated with the LeaseStatus
// in struct t.
func (t *RentableStatus) LeaseStatusStringer() string {
	return LeaseStatusString(t.LeaseStatus)
}

// LeaseStatusString returns the string associated with LeaseStatus ls
func LeaseStatusString(ls int64) string {
	i := int(ls)
	if i > len(RSLeaseStatus) {
		i = 0
	}
	return RSLeaseStatus[i]
}

// UseStatusStringer returns the string associated with the UseStatus
// in struct t.
func (t *RentableStatus) UseStatusStringer() string {
	return UseStatusString(t.UseStatus)
}

// UseStatusString returns the string associated with UseStatus us
func UseStatusString(us int64) string {
	i := int(us)
	if i > len(RSUseStatus) {
		i = 0
	}
	return RSUseStatus[i]
}

// MakeReadyStatusStringer returns the string associated with the MakeReadyStatus
// in struct t.
func (t *Rentable) MakeReadyStatusStringer() string {
	return MakeReadyStatusString(t.MRStatus)
}

// MakeReadyStatusString returns the string associated with MakeReadyStatus mr
func MakeReadyStatusString(mr int64) string {
	i := int(mr)
	if i > len(RSMakeReadyStatus) {
		i = 0
	}
	return RSMakeReadyStatus[i]
}
