package rlib

import "fmt"

// InsertBusiness writes a new Business record.
// returns the new business ID and any associated error
func InsertBusiness(b *Business) (int64, error) {
	var bid = int64(0)
	res, err := RRdb.Prepstmt.InsertBusiness.Exec(b.Designation, b.Name, b.DefaultOccupancyType, b.ParkingPermitInUse, b.LastModBy)
	if nil == err {
		id, err := res.LastInsertId()
		if err == nil {
			bid = int64(id)
		}
	}
	return bid, err
}

// InsertJournalEntry writes a new journal entry to the database
func InsertJournalEntry(j *Journal) (int64, error) {
	var rid = int64(0)
	res, err := RRdb.Prepstmt.InsertJournal.Exec(j.BID, j.RAID, j.Dt, j.Amount, j.Type, j.ID, j.Comment, j.LastModBy)
	if nil == err {
		id, err := res.LastInsertId()
		if err == nil {
			rid = int64(id)
		}
	}
	return rid, err
}

// InsertJournalAllocationEntry writes a new journalallocation record to the database
func InsertJournalAllocationEntry(ja *JournalAllocation) error {
	_, err := RRdb.Prepstmt.InsertJournalAllocation.Exec(ja.JID, ja.RID, ja.Amount, ja.ASMID, ja.AcctRule)
	return err
}

// InsertJournalMarker writes a new journalmarker record to the database
func InsertJournalMarker(jm *JournalMarker) error {
	_, err := RRdb.Prepstmt.InsertJournalMarker.Exec(jm.BID, jm.State, jm.DtStart, jm.DtStop)
	return err
}

// InsertLedgerMarker writes a new journalmarker record to the database
func InsertLedgerMarker(l *LedgerMarker) error {
	_, err := RRdb.Prepstmt.InsertLedgerMarker.Exec(l.BID, l.PID, l.GLNumber, l.Status, l.State, l.DtStart, l.DtStop, l.Balance, l.Type, l.Name)
	if err != nil {
		fmt.Printf("InsertLedgerMarker: err = %#v\n", err)
	}
	return err
}

// InsertLedgerEntry writes a new journal entry to the database
func InsertLedgerEntry(l *Ledger) (int64, error) {
	var rid = int64(0)
	res, err := RRdb.Prepstmt.InsertLedger.Exec(l.BID, l.JID, l.JAID, l.GLNumber, l.Dt, l.Amount, l.Comment, l.LastModBy)
	if nil == err {
		id, err := res.LastInsertId()
		if err == nil {
			rid = int64(id)
		}
	} else {
		Ulog("Error inserting ledger:  %v\n", err)
	}
	return rid, err
}

// InsertAssessmentType writes a new assessmenttype record to the database
func InsertAssessmentType(a *AssessmentType) error {
	_, err := RRdb.Prepstmt.InsertAssessmentType.Exec(a.Name, a.Description, a.LastModBy)
	return err
}

// InsertRentableSpecialty writes a new RentableSpecialtyType record to the database
func InsertRentableSpecialty(a *RentableSpecialty) error {
	_, err := RRdb.Prepstmt.InsertRentableSpecialtyType.Exec(a.RSPID, a.BID, a.Name, a.Fee, a.Description)
	return err
}

// InsertRentableMarketRates writes a new marketrate record to the database
func InsertRentableMarketRates(r *RentableMarketRate) error {
	_, err := RRdb.Prepstmt.InsertRentableMarketRates.Exec(r.RTID, r.MarketRate, r.DtStart, r.DtStop)
	return err
}

// InsertRentableType writes a new RentableType record to the database
func InsertRentableType(a *RentableType) (int64, error) {
	var rid = int64(0)
	res, err := RRdb.Prepstmt.InsertRentableType.Exec(a.RTID, a.BID, a.Style, a.Name, a.Frequency, a.Proration, a.Report, a.ManageToBudget, a.LastModBy)
	if nil == err {
		id, err := res.LastInsertId()
		if err == nil {
			rid = int64(id)
		}
	} else {
		Ulog("Error inserting RentableType:  %v\n", err)
	}
	return rid, err
}

// InsertBuilding writes a new Building record to the database
func InsertBuilding(a *Building) (int64, error) {
	var rid = int64(0)
	res, err := RRdb.Prepstmt.InsertBuilding.Exec(a.BID, a.Address, a.Address2, a.City, a.State, a.PostalCode, a.Country, a.LastModBy)
	if nil == err {
		id, err := res.LastInsertId()
		if err == nil {
			rid = int64(id)
		}
	} else {
		Ulog("Error inserting Building:  %v\n", err)
	}
	return rid, err
}

// InsertBuildingWithID writes a new Building record to the database with the supplied bldgid
// the building ID must be set in the supplied building struct ptr (a.BLDGID).
func InsertBuildingWithID(a *Building) (int64, error) {
	var rid = int64(0)
	res, err := RRdb.Prepstmt.InsertBuildingWithID.Exec(a.BLDGID, a.BID, a.Address, a.Address2, a.City, a.State, a.PostalCode, a.Country, a.LastModBy)
	if nil == err {
		id, err := res.LastInsertId()
		if err == nil {
			rid = int64(id)
		}
	} else {
		Ulog("InsertBuildingWithID: error inserting Building:  %v\n", err)
		Ulog("Bldg = %#v\n", *a)
	}
	return rid, err
}

// InsertRentable writes a new Rentable record to the database
func InsertRentable(a *Rentable) (int64, error) {
	var rid = int64(0)
	res, err := RRdb.Prepstmt.InsertRentable.Exec(a.RTID, a.BID, a.Name, a.Assignment, a.Report, a.DefaultOccType, a.OccType, a.LastModBy)
	if nil == err {
		id, err := res.LastInsertId()
		if err == nil {
			rid = int64(id)
		}
	} else {
		Ulog("InsertRentable: error inserting Building:  %v\n", err)
		Ulog("Rentable = %#v\n", *a)
	}
	return rid, err
}

// InsertTransactant writes a new transactant record to the database
func InsertTransactant(a *Transactant) (int64, error) {
	var tid = int64(0)
	res, err := RRdb.Prepstmt.InsertTransactant.Exec(a.TID, a.PID, a.PRSPID, a.FirstName, a.MiddleName, a.LastName, a.PrimaryEmail, a.SecondaryEmail, a.WorkPhone, a.CellPhone, a.Address, a.Address2, a.City, a.State, a.PostalCode, a.Country, a.LastModBy)
	if nil == err {
		id, err := res.LastInsertId()
		if err == nil {
			tid = int64(id)
		}
	} else {
		Ulog("InsertTransactant: error inserting Transactant:  %v\n", err)
		Ulog("Transactant = %#v\n", *a)
	}
	return tid, err
}

// InsertTenant writes a new tenant record to the database
func InsertTenant(a *Tenant) (int64, error) {
	var tid = int64(0)
	res, err := RRdb.Prepstmt.InsertTenant.Exec(a.TCID, a.Points, a.CarMake, a.CarModel, a.CarColor, a.CarYear, a.LicensePlateState, a.LicensePlateNumber, a.ParkingPermitNumber, a.AccountRep, a.DateofBirth, a.EmergencyContactName, a.EmergencyContactAddress, a.EmergencyContactTelephone, a.EmergencyEmail, a.AlternateAddress, a.ElibigleForFutureOccupancy, a.Industry, a.Source, a.InvoicingCustomerNumber)
	if nil == err {
		id, err := res.LastInsertId()
		if err == nil {
			tid = int64(id)
		}
	} else {
		Ulog("InsertTenant: error inserting Tenant:  %v\n", err)
		Ulog("Tenant = %#v\n", *a)
	}
	return tid, err
}

// InsertPayor writes a new tenant record to the database
func InsertPayor(a *Payor) (int64, error) {
	var tid = int64(0)
	res, err := RRdb.Prepstmt.InsertPayor.Exec(a.TCID, a.CreditLimit, a.EmployerName, a.EmployerStreetAddress, a.EmployerCity, a.EmployerState, a.EmployerPostalCode, a.EmployerEmail, a.EmployerPhone, a.Occupation, a.LastModBy)
	if nil == err {
		id, err := res.LastInsertId()
		if err == nil {
			tid = int64(id)
		}
	} else {
		Ulog("InsertPayor: error inserting Payor:  %v\n", err)
		Ulog("Payor = %#v\n", *a)
	}
	return tid, err
}

// InsertProspect writes a new tenant record to the database
func InsertProspect(a *Prospect) (int64, error) {
	var tid = int64(0)
	res, err := RRdb.Prepstmt.InsertProspect.Exec(a.TCID, a.ApplicationFee, a.LastModBy)
	if nil == err {
		id, err := res.LastInsertId()
		if err == nil {
			tid = int64(id)
		}
	} else {
		Ulog("InsertProspect: error inserting Prospect:  %v\n", err)
		Ulog("Prospect = %#v\n", *a)
	}
	return tid, err
}
