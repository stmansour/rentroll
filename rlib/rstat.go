package rlib

import "time"

// RSLeaseStatus is a slice of the string meaning of each LeaseStatus
var RSLeaseStatus = []string{
	"Unknown	",
	"Vacant Preleased",
	"Vacant Unleased",
	"On Notice Preleased",
	"On Notice Unleased",
	"Leased",
	"Unavailable",
}

// RSUseStatus is a slice of the string meaning of each UseStatus
var RSUseStatus = []string{
	"Unknown",             // 0
	"In Service",          // 1
	"Administrative",      // 2
	"Employee",            // 3
	"Owner Occupied",      // 4
	"Offline Rennovation", // 5
	"Offline Maintenance", // 6
	"Model",               // 7
}

// RSMakeReadyStatus is a slice of the string meaning of each MakeReadyStatus
var RSMakeReadyStatus = []string{
	"Unknown",
	"In Progress Housekeeping",
	"In Progress Maintenance",
	"Pending Inspection",
	"Ready",
}

// RStatInfo encapsulates a RentableStatus structure along with t
// the associated amount for DtStart - DtStop
type RStatInfo struct {
	Amount float64
	RS     RentableStatus
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
func RStat(xbiz *XBusiness, rid int64, gaps []Period) []RStatInfo {
	var m []RStatInfo
	for i := 0; i < len(gaps); i++ {
		//-------------------------------------------------------------
		// Check for any special rentable status during the gap.
		// Reflect what's happening if we find anything
		//-------------------------------------------------------------
		rsa := GetRentableStatusByRange(rid, &gaps[i].D1, &gaps[i].D2)

		if len(rsa) > 0 {
			for j := 0; j < len(rsa); j++ {
				var rs RStatInfo
				rs.Amount = float64(177)
				rsa[j].DtStart = gaps[i].D1
				rsa[j].DtStop = gaps[i].D2
				rs.RS = rsa[j]
				m = append(m, rs)
			}
		} else {
			var rs = RStatInfo{
				RS: RentableStatus{
					BID:         xbiz.P.BID,
					RID:         rid,
					DtStart:     gaps[i].D1,
					DtStop:      gaps[i].D2,
					UseStatus:   USESTATUSinService,
					LeaseStatus: LEASESTATUSvacantNotRented,
				},
			}
			//----------------------------------------------------------------
			// If there is a RentalAgreement in the future, modify the status
			//----------------------------------------------------------------
			r := GetRentableStatusOnOrAfter(rid, &gaps[i].D1)
			if r.RSID > 0 {
				rs.RS.LeaseStatus = LEASESTATUSvacantRented
			}
			rs.Amount = float64(177)
			m = append(m, rs)
		}
	}
	return m
}

// VacancyGSR returns the GSR amount for the rentable during the supplied
// period
//
// INPUT
//  xbiz  - the business
//  rid   - which rentable
//  d1-d2 - the time period
//
// RETURN
//  the amount of GSR for the period
//-----------------------------------------------------------------------------
func VacancyGSR(xbiz *XBusiness, rid int64, d1, d2 *time.Time) float64 {
	amt := float64(0)
	m := VacancyDetect(xbiz, d1, d2, rid)
	for i := 0; i < len(m); i++ {
		amt += m[i].Amount
	}
	return amt
}

// LeaseStatusStringer returns the string associated with the LeaseStatus
// in struct t.
//-----------------------------------------------------------------------------
func (t *RentableStatus) LeaseStatusStringer() string {
	return LeaseStatusString(t.LeaseStatus)
}

// LeaseStatusString returns the string associated with LeaseStatus ls
//-----------------------------------------------------------------------------
func LeaseStatusString(ls int64) string {
	i := int(ls)
	if i > len(RSLeaseStatus) {
		i = 0
	}
	return RSLeaseStatus[i]
}

// UseStatusStringer returns the string associated with the UseStatus
// in struct t.
//-----------------------------------------------------------------------------
func (t *RentableStatus) UseStatusStringer() string {
	return UseStatusString(t.UseStatus)
}

// UseStatusString returns the string associated with UseStatus us
//-----------------------------------------------------------------------------
func UseStatusString(us int64) string {
	i := int(us)
	if i > len(RSUseStatus) {
		i = 0
	}
	return RSUseStatus[i]
}

// MakeReadyStatusStringer returns the string associated with the MakeReadyStatus
// in struct t.
//-----------------------------------------------------------------------------
func (t *Rentable) MakeReadyStatusStringer() string {
	return MakeReadyStatusString(t.MRStatus)
}

// MakeReadyStatusString returns the string associated with MakeReadyStatus mr
//-----------------------------------------------------------------------------
func MakeReadyStatusString(mr int64) string {
	i := int(mr)
	if i > len(RSMakeReadyStatus) {
		i = 0
	}
	return RSMakeReadyStatus[i]
}
