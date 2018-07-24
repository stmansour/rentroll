package ws

// Fees2RA copies fees into permanent table Assessments.
import (
	"context"
	"fmt"
	"rentroll/rlib"
	"time"
)

// Fees2RA handles all the updates necessary to move the fees defined in a flow
// into the permanent tables.
//
// INPUTS
//     ctx    - db context for transactions
//     x - all the contextual info we need for performing this operation
//
// RETURNS
//     Any errors encountered
//-----------------------------------------------------------------------------
func Fees2RA(ctx context.Context, x *WriteHandlerContext) error {
	var err error
	rlib.Console("Entered Fees2RA\n")

	//--------------------------------------------------
	// Handle Rentables first...
	//--------------------------------------------------
	rlib.Console("A\n")
	for i := 0; i < len(x.raf.Rentables); i++ {
		for j := 0; j < len(x.raf.Rentables[i].Fees); j++ {
			if x.raf.Rentables[i].Fees[j].ASMID > 0 {
				rlib.Console("i = %d, j = %d, ASMID = %d\n", i, j, x.raf.Rentables[i].Fees[j].ASMID)
				if err = F2RAUpdateExistingAssessment(ctx, x, &x.raf.Rentables[i].Fees[j], rlib.ELEMRENTABLE, x.raf.Rentables[i].RID, 0); err != nil {
					return err
				}
			}
		}
	}
	rlib.Console("B\n")
	//--------------------------------------------------
	// Handle pet fees...
	//--------------------------------------------------
	for i := 0; i < len(x.raf.Pets); i++ {
		for j := 0; j < len(x.raf.Pets[i].Fees); j++ {
			if 0 < x.raf.Pets[i].Fees[j].ASMID {
				if err = F2RAUpdateExistingAssessment(ctx, x, &x.raf.Pets[i].Fees[j], rlib.ELEMPET, x.raf.Pets[i].PETID, x.raf.Pets[i].TMPTCID); err != nil {
					return err
				}
			}
		}
	}
	//--------------------------------------------------
	// Handle vehicle fees...
	//--------------------------------------------------
	rlib.Console("C\n")
	for i := 0; i < len(x.raf.Vehicles); i++ {
		for j := 0; j < len(x.raf.Vehicles[i].Fees); j++ {
			if 0 < x.raf.Vehicles[i].Fees[j].ASMID {
				if err = F2RAUpdateExistingAssessment(ctx, x, &x.raf.Vehicles[i].Fees[j], rlib.ELEMVEHICLE, x.raf.Vehicles[i].VID, x.raf.Vehicles[i].TMPTCID); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

// F2RAUpdateExistingAssessment handles all the updates necessary to move the
// supplied field into the permanent tables.
//
// INPUTS
//     ctx  - db context for transactions
//     x    - all the contextual info we need for performing this operation
//     elt  - element type if is this is bound to a pet or vehicle
//     id   - RID if elt == rlib.ELEMRENTABLE, or tmpid of the element
//            (TMPPETID, TMPVID), valid if elt > 0
//     tcid - tmptcid of the transactant responsible, valid if elt > 0
//
// RETURNS
//     Any errors encountered
//-----------------------------------------------------------------------------
func F2RAUpdateExistingAssessment(ctx context.Context, x *WriteHandlerContext, fee *rlib.RAFeesData, eltype, id, tmptcid int64) error {
	if fee.ASMID == int64(0) {
		return fmt.Errorf("fee.ASMID must be > 0")
	}
	a, err := rlib.GetAssessment(ctx, fee.ASMID)
	if err != nil {
		return err
	}

	//-------------------------------------------------------------------
	// skip any assessments that finished prior to this Rental Agreement
	//-------------------------------------------------------------------
	dt := time.Time(x.raf.Dates.AgreementStart)
	stop := time.Time(fee.Stop)
	if stop.Before(dt) {
		return nil // don't need to process this one
	}

	//-------------------------------------------------------------------
	// skip any non-recurring assessment that has been paid...
	//-------------------------------------------------------------------
	if a.FLAGS&3 == 2 {
		return nil // don't need to process this
	}

	//-------------------------------------------------------------------
	// If it's recurring we'll just stop it on the start date of the new
	// rental agreement
	//-------------------------------------------------------------------
	a.Stop = dt
	if err = rlib.UpdateAssessment(ctx, &a); err != nil {
		return err
	}

	//-------------------------------------------------------------------
	// Create a new assessment from this day forward...
	//-------------------------------------------------------------------
	var b rlib.Assessment
	b.Comment = fmt.Sprintf("Continuation of ASMID %d from RAID %d", a.ASMID, a.RAID)
	Start := time.Time(fee.Start) // the start time will be either the fee start
	if Start.Before(dt) {         // or the start of the new rental agreement
		Start = dt // whichever is later
	}
	b.Stop = time.Time(fee.Stop)
	b.BID = a.BID

	//-------------------------------------------------------------------
	// Set the Element Type and ID if necessary
	//-------------------------------------------------------------------
	b.AssocElemType = eltype
	b.AssocElemID = id

	//-------------------------------------------------------------------
	// find the RID associated with this pet
	//-------------------------------------------------------------------
	switch eltype {
	case rlib.ELEMRENTABLE:
		b.RID = id
	case rlib.ELEMPET:
		if b.RID = GetRIDForTMPTCID(ctx, x, tmptcid); b.RID <= 0 {
			return fmt.Errorf("No RID associated with TMPTCID = %d", tmptcid)
		}
	case rlib.ELEMVEHICLE:
		if b.RID = GetRIDForTMPTCID(ctx, x, tmptcid); b.RID <= 0 {
			return fmt.Errorf("No RID associated with TMPTCID = %d", tmptcid)
		}
	}
	b.Amount = fee.ContractAmount
	rlib.Console("bid = %d, fee ARID = %d\n", b.BID, fee.ARID)
	b.AcctRule = ""
	b.RentCycle = fee.RentCycle
	b.RAID = x.ra.RAID
	b.Start = time.Time(fee.Start)
	b.Stop = time.Time(fee.Stop)
	b.RentCycle = fee.RentCycle
	b.ProrationCycle = rlib.RRdb.BizTypes[b.BID].AR[fee.ARID].DefaultProrationCycle
	b.InvoiceNo = 0
	b.ARID = fee.ARID
	b.FLAGS = a.FLAGS & (1<<3 | 1<<4) // 1<<3 = PEDID required, 1<<4 = VID required
	b.Comment = fmt.Sprintf("Updated from ASMID=%d", a.ASMID)

	_, err = rlib.InsertAssessment(ctx, &b)
	if err != nil {
		return err
	}

	return nil
}

// GetRIDForTMPTCID finds the RID associated with the supplied tmptcid.
// This routine is called when we have a pet or a vehicle and we need to
// know what RID to associate it with. The RATiePeopleData datastruct is
// organized by tmptcid.  PRID is the Parent RID for that person.
//
// INPUTS
//     ctx     - db context for transactions
//     x       - all the contextual info we need for performing this operation
//     tmptcid - tmptcid for person we want the associated RID
//
// RETURNS
//     RID of associated rentable, or -1 if not found
//-----------------------------------------------------------------------------
func GetRIDForTMPTCID(ctx context.Context, x *WriteHandlerContext, tmptcid int64) int64 {
	for i := 0; i < len(x.raf.Tie.People); i++ {
		if x.raf.Tie.People[i].TMPTCID == tmptcid {
			return x.raf.Tie.People[i].PRID
		}
	}
	return -1
}
