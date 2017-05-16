package bizlogic

import (
	"fmt"
	"rentroll/rlib"
	"time"
)

// If an overpayment is made, then the balance is stored with the Rental Agreement.
// If there are multiple rental agreements for which the payor is responsible,
// then the balance is stored with the oldest Rental Agreement.

// GetAllUnpaidReceiptsForPayor determines all the Rental Agreements for which
// the supplied Transactant is Payor at time dt, then returns a list of all unpaid
// assessments associated with these Rental Agreements.
func GetAllUnpaidReceiptsForPayor(bid, tcid int64, dt *time.Time) []rlib.Assessment {
	var a []rlib.Assessment
	m := rlib.GetRentalAgreementsByPayor(bid, tcid, dt) // Determine which Rental Agreements the Payor is responsible for...
	for i := 0; i < len(m); i++ {                       // build the list of unpaid assessments
		n := rlib.GetUnpaidAssessmentsByRAID(m[i].RAID) // the list is presorted by Start date ascending
		a = append(a, n...)
	}
	return a
}

// AddReceiptToBooks adds the receipt and amount to the journal and the ledgers
func AddReceiptToBooks() {

}

// AutoProcessReceipt applies the amount of the supplied receipt to allocate
// payments to all unpaid based assessments for which the payor is
// responsible.
func AutoProcessReceipt(r *rlib.Receipt) {
	remaining := r.Amount // this is the total amount of assessments we can pay off

	// Push the full amount of this check into the supplied GL Account...

	m := GetAllUnpaidReceiptsForPayor(r.BID, r.TCID, &r.Dt)

	//----------------------------------------------------------------------
	// Use the remaining balance to pay off this assessment (or as much
	// of the assessment as possible)
	//----------------------------------------------------------------------
	for i := 0; i < len(m); i++ {
		fmt.Printf("Assessment %d, Amount = %8.2f, AR = %d\n", m[i].ASMID, m[i].Amount, m[i].ARID)
		if remaining > m[i].Amount {

		}

	}

	//----------------------------------------------------------------------
	// After all payments have been made, if there is any balance, apply it
	// to the oldest Rental Agreement
	//----------------------------------------------------------------------

}
