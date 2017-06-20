package bizlogic

import (
	"fmt"
	"rentroll/rlib"
)

// ValidateAssessment checks to see whether the assessment violates any
// business logic.
func ValidateAssessment(a *rlib.Assessment) []BizError {
	var e []BizError
	if a.RID > 0 {

		//--------------------------------------------------------------------------
		//  Check for assessment timeframe prior to or after Rentable's type being defined
		//--------------------------------------------------------------------------
		rtl := rlib.GetRentableTypeRefs(a.RID) // these are returned in chronological order
		l := len(rtl)
		if l == 0 {
			e = append(e, BizErrors[RentableTypeUnknown])
		} else {
			if a.Stop.Before(rtl[0].DtStart) || a.Start.After(rtl[l-1].DtStop) {
				e = append(e, BizErrors[RentableTypeUnknown])
			}
		}

		//--------------------------------------------------------------------------
		//  Check for assessment timeframe prior to or after Rentable's status being defined
		//--------------------------------------------------------------------------
		rsl := rlib.GetRentableStatusByRange(a.RID, &a.Start, &a.Stop)
		l = len(rsl)
		if l == 0 {
			fmt.Printf("ValidateAssessment: l=0\n")
			e = append(e, BizErrors[RentableStatusUnknown])
		} else {
			fmt.Printf("ValidateAssessment: a.Start-Stop = %s - %s\n", a.Start.Format(rlib.RRDATEINPFMT), a.Stop.Format(rlib.RRDATEINPFMT))
			fmt.Printf("ValidateAssessment: rtl = %s - %s\n", rtl[0].DtStart.Format(rlib.RRDATEINPFMT), rtl[l-1].DtStop.Format(rlib.RRDATEINPFMT))
			if a.Stop.Before(rtl[0].DtStart) || a.Start.After(rtl[l-1].DtStop) {
				e = append(e, BizErrors[RentableStatusUnknown])
			}
		}

	}
	return e
}
