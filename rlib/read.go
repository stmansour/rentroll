package rlib

import (
	"database/sql"
	"encoding/hex"
)

// As the database structures change, having the calls that Read from the database into these structures located
// in one place simplifies maintenance

// The routine which is dealing with *sql.Row, it will check for sql.ErrNoRows(no rows in results error)
// if found it will ignore the error and assign nil to original err variable with help of SkipSQLNoRowsError routine.
// It is caller responsibility to check for zero-value for a resource (returned back from Get-* method) if
// it's want to consider "no resource found" as an Error.

// ReadAR reads a full AR structure from the database based on the supplied row object
func ReadAR(row *sql.Row, a *AR) error {
	err := row.Scan(&a.ARID, &a.BID, &a.Name, &a.ARType, &a.DebitLID, &a.CreditLID, &a.Description, &a.RARequired, &a.DtStart, &a.DtStop, &a.FLAGS, &a.DefaultAmount, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
	SkipSQLNoRowsError(&err)
	return err
}

// ReadARs reads a full AR structure from the database based on the supplied rows object
func ReadARs(rows *sql.Rows, a *AR) error {
	return rows.Scan(&a.ARID, &a.BID, &a.Name, &a.ARType, &a.DebitLID, &a.CreditLID, &a.Description, &a.RARequired, &a.DtStart, &a.DtStop, &a.FLAGS, &a.DefaultAmount, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
}

// ReadAssessment reads a full Assessment structure of data from the database based on the supplied Rows pointer.
func ReadAssessment(row *sql.Row, a *Assessment) error {
	err := row.Scan(&a.ASMID, &a.PASMID, &a.RPASMID, &a.AGRCPTID, &a.BID, &a.RID, &a.ATypeLID, &a.RAID, &a.Amount,
		&a.Start, &a.Stop, &a.RentCycle, &a.ProrationCycle, &a.InvoiceNo, &a.AcctRule, &a.ARID, &a.FLAGS, &a.Comment,
		&a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
	SkipSQLNoRowsError(&err)
	return err
}

// ReadAssessments reads a full Assessment structure of data from the database based on the supplied Rows pointer.
func ReadAssessments(rows *sql.Rows, a *Assessment) error {
	return rows.Scan(&a.ASMID, &a.PASMID, &a.RPASMID, &a.AGRCPTID, &a.BID, &a.RID, &a.ATypeLID, &a.RAID, &a.Amount,
		&a.Start, &a.Stop, &a.RentCycle, &a.ProrationCycle, &a.InvoiceNo, &a.AcctRule, &a.ARID, &a.FLAGS, &a.Comment,
		&a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
}

// ReadBuildingData reads the data for a building object from db based on the supplied pointer
func ReadBuildingData(row *sql.Row, a *Building) error {
	err := row.Scan(&a.BLDGID, &a.BID, &a.Address, &a.Address2, &a.City, &a.State, &a.PostalCode, &a.Country, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
	SkipSQLNoRowsError(&err)
	return err
}

// ReadBusiness reads a full Business structure from the database based on the supplied row object
func ReadBusiness(row *sql.Row, a *Business) error {
	err := row.Scan(&a.BID, &a.Designation, &a.Name, &a.DefaultRentCycle, &a.DefaultProrationCycle, &a.DefaultGSRPC, &a.ClosePeriodTLID, &a.FLAGS, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
	SkipSQLNoRowsError(&err)
	return err
}

// ReadBusinesses reads a full Business structure from the database based on the supplied rows object
func ReadBusinesses(rows *sql.Rows, a *Business) error {
	return rows.Scan(&a.BID, &a.Designation, &a.Name, &a.DefaultRentCycle, &a.DefaultProrationCycle, &a.DefaultGSRPC, &a.ClosePeriodTLID, &a.FLAGS, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
}

// ReadBusinessProperties reads a full BusinessProperties structure from the database based on the supplied row object
func ReadBusinessProperties(row *sql.Row, a *BusinessProperties) error {
	err := row.Scan(&a.BPID, &a.BID, &a.Name, &a.Data, &a.FLAGS, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
	SkipSQLNoRowsError(&err)
	return err
}

// ReadBusinessPropertiess reads a full BusinessProperties structure from the database based on the supplied rows object
func ReadBusinessPropertiess(rows *sql.Rows, a *BusinessProperties) error {
	return rows.Scan(&a.BPID, &a.BID, &a.Name, &a.Data, &a.FLAGS, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
}

// ReadClosePeriod reads a full ClosePeriod structure from the database based on the supplied row object
func ReadClosePeriod(row *sql.Row, a *ClosePeriod) error {
	err := row.Scan(&a.CPID, &a.BID, &a.TLID, &a.Dt, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
	SkipSQLNoRowsError(&err)
	return err
}

// ReadClosePeriods reads a full ClosePeriod structure from the database based on the supplied rows object
func ReadClosePeriods(rows *sql.Rows, a *ClosePeriod) error {
	return rows.Scan(&a.CPID, &a.BID, &a.TLID, &a.Dt, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
}

// ReadCustomAttribute reads a full CustomAttribute structure from the database based on the supplied row object
func ReadCustomAttribute(row *sql.Row, a *CustomAttribute) error {
	err := row.Scan(&a.CID, &a.BID, &a.Type, &a.Name, &a.Value, &a.Units, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
	SkipSQLNoRowsError(&err)
	return err
}

// ReadCustomAttributes reads a full CustomAttribute structure from the database based on the supplied rows object
func ReadCustomAttributes(rows *sql.Rows, a *CustomAttribute) error {
	return rows.Scan(&a.CID, &a.BID, &a.Type, &a.Name, &a.Value, &a.Units, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
}

// ReadCustomAttributeRef reads a full CustomAttributeRef structure from the database based on the supplied row object
func ReadCustomAttributeRef(row *sql.Row, a *CustomAttributeRef) error {
	err := row.Scan(&a.ElementType, &a.BID, &a.ID, &a.CID, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
	SkipSQLNoRowsError(&err)
	return err
}

// ReadCustomAttributeRefs reads a full CustomAttributeRef structure from the database based on the supplied rows object
func ReadCustomAttributeRefs(rows *sql.Rows, a *CustomAttributeRef) error {
	return rows.Scan(&a.ElementType, &a.BID, &a.ID, &a.CID, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
}

// ReadDemandSource reads a full DemandSource structure from the database based on the supplied row object
func ReadDemandSource(row *sql.Row, a *DemandSource) error {
	err := row.Scan(&a.SourceSLSID, &a.BID, &a.Name, &a.Industry, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
	SkipSQLNoRowsError(&err)
	return err
}

// ReadDemandSources reads a full DemandSource structure from the database based on the supplied rows object
func ReadDemandSources(rows *sql.Rows, a *DemandSource) error {
	return rows.Scan(&a.SourceSLSID, &a.BID, &a.Name, &a.Industry, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
}

// ReadDeposit reads a full Deposit structure from the database based on the supplied row object
func ReadDeposit(row *sql.Row, a *Deposit) error {
	err := row.Scan(&a.DID, &a.BID, &a.DEPID, &a.DPMID, &a.Dt, &a.Amount, &a.ClearedAmount, &a.FLAGS, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
	SkipSQLNoRowsError(&err)
	return err
}

// ReadDeposits reads a full Deposit structure from the database based on the supplied rows object
func ReadDeposits(rows *sql.Rows, a *Deposit) error {
	return rows.Scan(&a.DID, &a.BID, &a.DEPID, &a.DPMID, &a.Dt, &a.Amount, &a.ClearedAmount, &a.FLAGS, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
}

// ReadDepository reads a full Depository structure from the database based on the supplied row object
func ReadDepository(row *sql.Row, a *Depository) error {
	err := row.Scan(&a.DEPID, &a.BID, &a.LID, &a.Name, &a.AccountNo, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
	SkipSQLNoRowsError(&err)
	return err
}

// ReadDepositories reads a full Depository structure from the database based on the supplied rows object
func ReadDepositories(rows *sql.Rows, a *Depository) error {
	return rows.Scan(&a.DEPID, &a.BID, &a.LID, &a.Name, &a.AccountNo, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
}

// ReadDepositMethod reads a full DepositMethod structure from the database based on the supplied row object
func ReadDepositMethod(row *sql.Row, a *DepositMethod) error {
	err := row.Scan(&a.DPMID, &a.BID, &a.Method, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
	SkipSQLNoRowsError(&err)
	return err
}

// ReadDepositMethods reads a full DepositMethod structure from the database based on the supplied row object
func ReadDepositMethods(rows *sql.Rows, a *DepositMethod) error {
	return rows.Scan(&a.DPMID, &a.BID, &a.Method, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
}

// ReadDepositPart reads a full DepositPart structure from the database based on the supplied row object
func ReadDepositPart(row *sql.Row, a *DepositPart) error {
	err := row.Scan(&a.DPID, &a.DID, &a.BID, &a.RCPTID, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
	SkipSQLNoRowsError(&err)
	return err
}

// ReadDepositParts reads a full DepositPart structure from the database based on the supplied row object
func ReadDepositParts(rows *sql.Rows, a *DepositPart) error {
	return rows.Scan(&a.DPID, &a.DID, &a.BID, &a.RCPTID, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
}

// ReadExpense reads a full Expense structure from the database based on the supplied row object
func ReadExpense(row *sql.Row, a *Expense) error {
	err := row.Scan(&a.EXPID, &a.RPEXPID, &a.BID, &a.RID, &a.RAID, &a.Amount, &a.Dt, &a.AcctRule, &a.ARID, &a.FLAGS, &a.Comment, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
	SkipSQLNoRowsError(&err)
	return err
}

// ReadExpenses reads a full Expense structure from the database based on the supplied row object
func ReadExpenses(rows *sql.Rows, a *Expense) error {
	return rows.Scan(&a.EXPID, &a.RPEXPID, &a.BID, &a.RID, &a.RAID, &a.Amount, &a.Dt, &a.AcctRule, &a.ARID, &a.FLAGS, &a.Comment, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
}

//------------------
// FLOW
//------------------

// ReadFlow reads a full Flow structure from the database based on the supplied row object
func ReadFlow(row *sql.Row, a *Flow) error {
	err := row.Scan(&a.FlowID, &a.BID, &a.UserRefNo, &a.FlowType, &a.ID, &a.Data, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
	SkipSQLNoRowsError(&err)
	return err
}

// ReadFlows reads a full Flow structure from the database based on the supplied rows object
func ReadFlows(rows *sql.Rows, a *Flow) error {
	return rows.Scan(&a.FlowID, &a.BID, &a.UserRefNo, &a.FlowType, &a.ID, &a.Data, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
}

//------------------
// GLAccount
//------------------

// ReadGLAccount reads a full Ledger structure of data from the database based on the supplied Rows pointer.
func ReadGLAccount(row *sql.Row, a *GLAccount) error {
	err := row.Scan(&a.LID, &a.PLID, &a.BID, &a.RAID, &a.TCID, &a.GLNumber,
		/*&a.Status,*/ &a.Name, &a.AcctType, &a.AllowPost,
		&a.FLAGS, &a.Description, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
	SkipSQLNoRowsError(&err)
	return err
}

// ReadGLAccounts reads a full Ledger structure of data from the database based on the supplied Rows pointer.
func ReadGLAccounts(rows *sql.Rows, a *GLAccount) error {
	return rows.Scan(&a.LID, &a.PLID, &a.BID, &a.RAID, &a.TCID, &a.GLNumber,
		/*&a.Status,*/ &a.Name, &a.AcctType, &a.AllowPost,
		&a.FLAGS, &a.Description, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
}

// ReadInvoice reads a full Invoice structure of data from the database based on the supplied Rows pointer.
func ReadInvoice(row *sql.Row, a *Invoice) error {
	err := row.Scan(&a.InvoiceNo, &a.BID, &a.Dt, &a.DtDue, &a.Amount, &a.DeliveredBy, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
	SkipSQLNoRowsError(&err)
	return err
}

// ReadInvoices reads a full Invoice structure of data from the database based on the supplied Rows pointer.
func ReadInvoices(rows *sql.Rows, a *Invoice) error {
	return rows.Scan(&a.InvoiceNo, &a.BID, &a.Dt, &a.DtDue, &a.Amount, &a.DeliveredBy, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
}

// ReadInvoiceAssessments reads a full InvoiceAssessment structure of data from the database based on the supplied Rows pointer.
func ReadInvoiceAssessments(rows *sql.Rows, a *InvoiceAssessment) error {
	return rows.Scan(&a.InvoiceNo, &a.BID, &a.ASMID, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
}

// ReadInvoicePayors reads a full InvoicePayor structure of data from the database based on the supplied Rows pointer.
func ReadInvoicePayors(rows *sql.Rows, a *InvoicePayor) error {
	return rows.Scan(&a.InvoiceNo, &a.BID, &a.PID, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
}

// ReadJournal reads a full Journal structure of data from the database based on the supplied Rows pointer.
func ReadJournal(row *sql.Row, a *Journal) error {
	err := row.Scan(&a.JID, &a.BID, &a.Dt, &a.Amount, &a.Type, &a.ID, &a.Comment, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
	SkipSQLNoRowsError(&err)
	return err
}

// ReadJournals reads a full Journal structure of data from the database based on the supplied Rows pointer.
func ReadJournals(rows *sql.Rows, a *Journal) error {
	return rows.Scan(&a.JID, &a.BID, &a.Dt, &a.Amount, &a.Type, &a.ID, &a.Comment, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
}

// ReadJournalMarker reads a full JournalMarker structure of data from the database based on the supplied Rows pointer.
func ReadJournalMarker(row *sql.Row, a *JournalMarker) error {
	err := row.Scan(&a.JMID, &a.BID, &a.State, &a.DtStart, &a.DtStop, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
	SkipSQLNoRowsError(&err)
	return err
}

// ReadJournalMarkers reads a full JournalMarker structure of data from the database based on the supplied Rows pointer.
func ReadJournalMarkers(rows *sql.Rows, a *JournalMarker) error {
	return rows.Scan(&a.JMID, &a.BID, &a.State, &a.DtStart, &a.DtStop, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
}

// ReadJournalAllocation reads a full JournalAllocation structure of data from the database based on the supplied Rows pointer.
func ReadJournalAllocation(row *sql.Row, a *JournalAllocation) error {
	err := row.Scan(&a.JAID, &a.BID, &a.JID, &a.RID, &a.RAID, &a.TCID, &a.RCPTID, &a.Amount, &a.ASMID, &a.EXPID, &a.AcctRule, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
	SkipSQLNoRowsError(&err)
	return err
}

// ReadJournalAllocations reads a full JournalAllocation structure of data from the database based on the supplied Rows pointer.
func ReadJournalAllocations(rows *sql.Rows, a *JournalAllocation) error {
	return rows.Scan(&a.JAID, &a.BID, &a.JID, &a.RID, &a.RAID, &a.TCID, &a.RCPTID, &a.Amount, &a.ASMID, &a.EXPID, &a.AcctRule, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
}

// ReadLedgerEntry reads a full LedgerEntry structure of data from the database based on the supplied Rows pointer.
func ReadLedgerEntry(row *sql.Row, a *LedgerEntry) error {
	err := row.Scan(&a.LEID, &a.BID, &a.JID, &a.JAID, &a.LID, &a.RAID, &a.RID, &a.TCID, &a.Dt, &a.Amount, &a.Comment, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
	SkipSQLNoRowsError(&err)
	return err
}

// ReadLedgerEntries reads a full LedgerEntry structure of data from the database based on the supplied Rows pointer.
func ReadLedgerEntries(rows *sql.Rows, a *LedgerEntry) error {
	return rows.Scan(&a.LEID, &a.BID, &a.JID, &a.JAID, &a.LID, &a.RAID, &a.RID, &a.TCID, &a.Dt, &a.Amount, &a.Comment, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
}

// ReadRentableSpecialty read a full RentableSpecialty structure of data from db based on sql.Row pointer
func ReadRentableSpecialty(row *sql.Row, a *RentableSpecialty) error {
	err := row.Scan(&a.RSPID, &a.BID, &a.Name, &a.Fee, &a.Description, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
	SkipSQLNoRowsError(&err)
	return err
}

// ReadRentableSpecialties reads a full RentableSpecialty structure of data from db based on sql.Rows pointer
func ReadRentableSpecialties(rows *sql.Rows, a *RentableSpecialty) error {
	return rows.Scan(&a.RSPID, &a.BID, &a.Name, &a.Fee, &a.Description, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
}

// ReadRentableSpecialtyRef read a full ReadRentableSpecialtyRef structure of data from db based on sql.Row pointer
func ReadRentableSpecialtyRef(row *sql.Row, a *RentableSpecialtyRef) error {
	err := row.Scan(&a.BID, &a.RID, &a.RSPID, &a.DtStart, &a.DtStop, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
	SkipSQLNoRowsError(&err)
	return err
}

// ReadRentableSpecialtyRefs reads a full ReadRentableSpecialtyRef structure of data from db based on sql.Rows pointer
func ReadRentableSpecialtyRefs(rows *sql.Rows, a *RentableSpecialtyRef) error {
	return rows.Scan(&a.BID, &a.RID, &a.RSPID, &a.DtStart, &a.DtStop, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
}

// ReadRentableMarketRate reads a full RentableMarketRate structure of data from the database based on the supplied Rows pointer.
func ReadRentableMarketRate(row *sql.Row, a *RentableMarketRate) error {
	err := row.Scan(&a.RMRID, &a.RTID, &a.BID, &a.MarketRate, &a.DtStart, &a.DtStop, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
	SkipSQLNoRowsError(&err)
	return err
}

// ReadRentableMarketRates reads a full RentableMarketRate structure of data from the database based on the supplied Rows pointer.
func ReadRentableMarketRates(rows *sql.Rows, a *RentableMarketRate) error {
	return rows.Scan(&a.RMRID, &a.RTID, &a.BID, &a.MarketRate, &a.DtStart, &a.DtStop, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
}

// ReadLedgerMarker reads a full LedgerMarker structure of data from the database based on the supplied Rows pointer.
func ReadLedgerMarker(row *sql.Row, a *LedgerMarker) error {
	err := row.Scan(&a.LMID, &a.LID, &a.BID, &a.RAID, &a.RID, &a.TCID, &a.Dt, &a.Balance, &a.State, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
	SkipSQLNoRowsError(&err)
	return err
}

// ReadLedgerMarkers reads a full LedgerMarker structure of data from the database based on the supplied Rows pointer.
func ReadLedgerMarkers(rows *sql.Rows, a *LedgerMarker) error {
	return rows.Scan(&a.LMID, &a.LID, &a.BID, &a.RAID, &a.RID, &a.TCID, &a.Dt, &a.Balance, &a.State, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
}

// ReadNote reads a full Note structure from the database based on the supplied row object
func ReadNote(row *sql.Row, a *Note) error {
	err := row.Scan(&a.NID, &a.BID, &a.NLID, &a.PNID, &a.NTID, &a.RID, &a.RAID, &a.TCID, &a.Comment, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
	SkipSQLNoRowsError(&err)
	return err
}

// ReadNotes reads a full Note structure from the database based on the supplied row object
func ReadNotes(rows *sql.Rows, a *Note) error {
	return rows.Scan(&a.NID, &a.BID, &a.NLID, &a.PNID, &a.NTID, &a.RID, &a.RAID, &a.TCID, &a.Comment, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
}

// ReadNoteList reads a full NoteList structure from the database based on the supplied row object
func ReadNoteList(row *sql.Row, a *NoteList) error {
	err := row.Scan(&a.NLID, &a.BID, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
	SkipSQLNoRowsError(&err)
	return err
}

// ReadNoteLists reads a full Note structure from the database based on the supplied row object
func ReadNoteLists(rows *sql.Rows, a *NoteList) error {
	return rows.Scan(&a.NLID, &a.BID, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
}

// ReadNoteType reads a full NoteType structure from the database based on the supplied row object
func ReadNoteType(row *sql.Row, a *NoteType) error {
	err := row.Scan(&a.NTID, &a.BID, &a.Name, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
	SkipSQLNoRowsError(&err)
	return err
}

// ReadNoteTypes reads a full NoteType structure from the database based on the supplied row object
func ReadNoteTypes(rows *sql.Rows, a *NoteType) error {
	return rows.Scan(&a.NTID, &a.BID, &a.Name, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
}

// ReadPaymentType reads a full PaymentType structure from the database based on the supplied row object
func ReadPaymentType(row *sql.Row, a *PaymentType) error {
	err := row.Scan(&a.PMTID, &a.BID, &a.Name, &a.Description, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
	SkipSQLNoRowsError(&err)
	return err
}

// ReadPaymentTypes reads a full PaymentType structure from the database based on the supplied rows object
func ReadPaymentTypes(rows *sql.Rows, a *PaymentType) error {
	return rows.Scan(&a.PMTID, &a.BID, &a.Name, &a.Description, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
}

// ReadProspect reads a full Prospect structure from the database based on the supplied row object
func ReadProspect(row *sql.Row, a *Prospect) error {
	err := row.Scan(&a.TCID, &a.BID, &a.CompanyAddress,
		&a.CompanyCity, &a.CompanyState, &a.CompanyPostalCode, &a.CompanyEmail, &a.CompanyPhone, &a.Occupation,
		&a.DesiredUsageStartDate, &a.RentableTypePreference, &a.FLAGS,
		&a.EvictedDes, &a.ConvictedDes, &a.BankruptcyDes, &a.Approver, &a.DeclineReasonSLSID,
		&a.OtherPreferences, &a.FollowUpDate, &a.CSAgent, &a.OutcomeSLSID,
		&a.CurrentAddress, &a.CurrentLandLordName, &a.CurrentLandLordPhoneNo, &a.CurrentReasonForMoving,
		&a.CurrentLengthOfResidency, &a.PriorAddress, &a.PriorLandLordName, &a.PriorLandLordPhoneNo,
		&a.PriorReasonForMoving, &a.PriorLengthOfResidency, &a.CommissionableThirdParty,
		&a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
	SkipSQLNoRowsError(&err)
	return err
}

// ReadProspects reads a full Prospect structure from the database based on the supplied rows object
func ReadProspects(rows *sql.Rows, a *Prospect) error {
	return rows.Scan(&a.TCID, &a.BID, &a.CompanyAddress,
		&a.CompanyCity, &a.CompanyState, &a.CompanyPostalCode, &a.CompanyEmail, &a.CompanyPhone, &a.Occupation,
		&a.DesiredUsageStartDate, &a.RentableTypePreference, &a.FLAGS,
		&a.EvictedDes, &a.ConvictedDes, &a.BankruptcyDes, &a.Approver, &a.DeclineReasonSLSID,
		&a.OtherPreferences, &a.FollowUpDate, &a.CSAgent, &a.OutcomeSLSID,
		&a.CurrentAddress, &a.CurrentLandLordName, &a.CurrentLandLordPhoneNo, &a.CurrentReasonForMoving,
		&a.CurrentLengthOfResidency, &a.PriorAddress, &a.PriorLandLordName, &a.PriorLandLordPhoneNo,
		&a.PriorReasonForMoving, &a.PriorLengthOfResidency, &a.CommissionableThirdParty,
		&a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
}

// ReadRatePlan reads a full RatePlan structure from the database based on the supplied row object
func ReadRatePlan(row *sql.Row, a *RatePlan) error {
	err := row.Scan(&a.RPID, &a.BID, &a.Name, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
	SkipSQLNoRowsError(&err)
	return err
}

// ReadRatePlans reads a full RatePlan structure from the database based on the supplied row object
func ReadRatePlans(rows *sql.Rows, a *RatePlan) error {
	return rows.Scan(&a.RPID, &a.BID, &a.Name, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
}

// ReadRatePlanRef reads a full RatePlanRef structure from the database based on the supplied row object
func ReadRatePlanRef(row *sql.Row, a *RatePlanRef) error {
	err := row.Scan(&a.RPRID, &a.BID, &a.RPID, &a.DtStart, &a.DtStop, &a.FeeAppliesAge, &a.MaxNoFeeUsers,
		&a.AdditionalUserFee, &a.PromoCode, &a.CancellationFee, &a.FLAGS, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
	SkipSQLNoRowsError(&err)
	return err
}

// ReadRatePlanRefs reads a full RatePlanRef structure from the database based on the supplied row object
func ReadRatePlanRefs(rows *sql.Rows, a *RatePlanRef) error {
	return rows.Scan(&a.RPRID, &a.BID, &a.RPID, &a.DtStart, &a.DtStop, &a.FeeAppliesAge, &a.MaxNoFeeUsers,
		&a.AdditionalUserFee, &a.PromoCode, &a.CancellationFee, &a.FLAGS, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
}

// ReadRatePlanRefRTRate reads a full RatePlanRefRTRate structure from the database based on the supplied row object
func ReadRatePlanRefRTRate(row *sql.Row, a *RatePlanRefRTRate) error {
	err := row.Scan(&a.RPRID, &a.BID, &a.RTID, &a.FLAGS, &a.Val, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
	SkipSQLNoRowsError(&err)
	return err
}

// ReadRatePlanRefRTRates reads a full RatePlanRefRTRate structure from the database based on the supplied row object
func ReadRatePlanRefRTRates(rows *sql.Rows, a *RatePlanRefRTRate) error {
	return rows.Scan(&a.RPRID, &a.BID, &a.RTID, &a.FLAGS, &a.Val, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
}

// ReadRatePlanRefSPRate reads a full RatePlanRefSPRate structure from the database based on the supplied row object
func ReadRatePlanRefSPRate(row *sql.Row, a *RatePlanRefSPRate) error {
	err := row.Scan(&a.RPRID, &a.BID, &a.RTID, &a.RSPID, &a.FLAGS, &a.Val, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
	SkipSQLNoRowsError(&err)
	return err
}

// ReadRatePlanRefSPRates reads a full RatePlanRefSPRate structure from the database based on the supplied row object
func ReadRatePlanRefSPRates(rows *sql.Rows, a *RatePlanRefSPRate) error {
	return rows.Scan(&a.RPRID, &a.BID, &a.RTID, &a.RSPID, &a.FLAGS, &a.Val, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
}

// ReadReceipt reads a full Receipt structure of data from the database based on the supplied Rows pointer.
func ReadReceipt(row *sql.Row, a *Receipt) error {
	err := row.Scan(&a.RCPTID, &a.PRCPTID, &a.BID, &a.TCID, &a.PMTID, &a.DEPID, &a.DID, &a.RAID, &a.Dt, &a.DocNo, &a.Amount, &a.AcctRuleReceive, &a.ARID, &a.AcctRuleApply, &a.FLAGS, &a.Comment,
		&a.OtherPayorName, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
	SkipSQLNoRowsError(&err)
	return err
}

// ReadReceipts reads a full Receipt structure of data from the database based on the supplied Rows pointer.
func ReadReceipts(rows *sql.Rows, a *Receipt) error {
	return rows.Scan(&a.RCPTID, &a.PRCPTID, &a.BID, &a.TCID, &a.PMTID, &a.DEPID, &a.DID, &a.RAID, &a.Dt, &a.DocNo, &a.Amount, &a.AcctRuleReceive, &a.ARID, &a.AcctRuleApply, &a.FLAGS, &a.Comment,
		&a.OtherPayorName, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
}

// ReadReceiptAllocation reads a full ReceiptAllocation structure of data from the database based on the supplied Rows pointer.
func ReadReceiptAllocation(row *sql.Row, a *ReceiptAllocation) error {
	err := row.Scan(&a.RCPAID, &a.RCPTID, &a.BID, &a.RAID, &a.Dt, &a.Amount, &a.ASMID, &a.FLAGS, &a.AcctRule, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
	SkipSQLNoRowsError(&err)
	return err
}

// ReadReceiptAllocations reads a full ReceiptAllocation structure of data from the database based on the supplied Rows pointer.
func ReadReceiptAllocations(rows *sql.Rows, a *ReceiptAllocation) error {
	return rows.Scan(&a.RCPAID, &a.RCPTID, &a.BID, &a.RAID, &a.Dt, &a.Amount, &a.ASMID, &a.FLAGS, &a.AcctRule, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
}

// ReadRentableTypeDown reads a full RentableTypeDown structure of data from the database based on the supplied Row pointer.
func ReadRentableTypeDown(rows *sql.Rows, a *RentableTypeDown) error {
	return rows.Scan(&a.RID, &a.RentableName)
}

// ReadRentable reads a full Rentable structure of data from the database based on the supplied Row pointer.
func ReadRentable(row *sql.Row, a *Rentable) error {
	err := row.Scan(&a.RID, &a.BID, &a.RentableName, &a.AssignmentTime, &a.MRStatus, &a.DtMRStart, &a.Comment, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
	SkipSQLNoRowsError(&err)
	return err
}

// ReadRentables reads a full Rentable structure of data from the database based on the supplied Rows pointer.
func ReadRentables(rows *sql.Rows, a *Rentable) error {
	return rows.Scan(&a.RID, &a.BID, &a.RentableName, &a.AssignmentTime, &a.MRStatus, &a.DtMRStart, &a.Comment, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
}

// ReadRentableType reads a full RentableType structure of data from the database based on the supplied Row pointer.
func ReadRentableType(row *sql.Row, a *RentableType) error {
	err := row.Scan(&a.RTID, &a.BID, &a.Style, &a.Name, &a.RentCycle, &a.Proration, &a.GSRPC, &a.ARID, &a.FLAGS, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
	SkipSQLNoRowsError(&err)
	return err
}

// ReadRentableTypes reads a full RentableType structure of data from the database based on the supplied Rows pointer.
func ReadRentableTypes(rows *sql.Rows, a *RentableType) error {
	return rows.Scan(&a.RTID, &a.BID, &a.Style, &a.Name, &a.RentCycle, &a.Proration, &a.GSRPC, &a.ARID, &a.FLAGS, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
}

// ReadRentableTypeRef reads a full RentableTypeRef structure of data from the database based on the supplied Row pointer.
func ReadRentableTypeRef(row *sql.Row, a *RentableTypeRef) error {
	err := row.Scan(&a.RTRID, &a.RID, &a.BID, &a.RTID, &a.OverrideRentCycle, &a.OverrideProrationCycle, &a.DtStart, &a.DtStop,
		&a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
	SkipSQLNoRowsError(&err)
	return err
}

// ReadRentableTypeRefs reads a full RentableTypeRef structure of data from the database based on the supplied Rows pointer.
func ReadRentableTypeRefs(rows *sql.Rows, a *RentableTypeRef) error {
	return rows.Scan(&a.RTRID, &a.RID, &a.BID, &a.RTID, &a.OverrideRentCycle, &a.OverrideProrationCycle, &a.DtStart, &a.DtStop,
		&a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
}

// ReadRentableStatus reads a full RentableStatus structure of data from the database based on the supplied Row pointer.
func ReadRentableStatus(row *sql.Row, a *RentableStatus) error {
	err := row.Scan(&a.RSID, &a.RID, &a.BID, &a.DtStart, &a.DtStop, &a.DtNoticeToVacate, &a.UseStatus, &a.LeaseStatus,
		&a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
	SkipSQLNoRowsError(&err)
	return err
}

// ReadRentableStatuss reads a full RentableStatus structure of data from the database based on the supplied Rows pointer.
func ReadRentableStatuss(rows *sql.Rows, a *RentableStatus) error {
	return rows.Scan(&a.RSID, &a.RID, &a.BID, &a.DtStart, &a.DtStop, &a.DtNoticeToVacate, &a.UseStatus, &a.LeaseStatus,
		&a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
}

// ReadRentalAgreement reads a full RentalAgreement structure of data from the database based on the supplied Row pointer.
func ReadRentalAgreement(row *sql.Row, a *RentalAgreement) error {
	err := row.Scan(&a.RAID, &a.RATID, &a.BID, &a.NLID, &a.AgreementStart, &a.AgreementStop, &a.PossessionStart,
		&a.PossessionStop, &a.RentStart, &a.RentStop, &a.RentCycleEpoch, &a.UnspecifiedAdults, &a.UnspecifiedChildren,
		&a.Renewal, &a.SpecialProvisions,
		&a.LeaseType, &a.ExpenseAdjustmentType, &a.ExpensesStop, &a.ExpenseStopCalculation, &a.BaseYearEnd,
		&a.ExpenseAdjustment, &a.EstimatedCharges, &a.RateChange, &a.NextRateChange, &a.PermittedUses, &a.ExclusiveUses,
		&a.ExtensionOption, &a.ExtensionOptionNotice, &a.ExpansionOption, &a.ExpansionOptionNotice, &a.RightOfFirstRefusal,
		&a.FLAGS, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
	SkipSQLNoRowsError(&err)
	return err
}

// ReadRentalAgreements reads a full RentalAgreement structure of data from the database based on the supplied Rows pointer.
func ReadRentalAgreements(rows *sql.Rows, a *RentalAgreement) error {
	return rows.Scan(&a.RAID, &a.RATID, &a.BID, &a.NLID, &a.AgreementStart, &a.AgreementStop, &a.PossessionStart,
		&a.PossessionStop, &a.RentStart, &a.RentStop, &a.RentCycleEpoch, &a.UnspecifiedAdults, &a.UnspecifiedChildren,
		&a.Renewal, &a.SpecialProvisions,
		&a.LeaseType, &a.ExpenseAdjustmentType, &a.ExpensesStop, &a.ExpenseStopCalculation, &a.BaseYearEnd,
		&a.ExpenseAdjustment, &a.EstimatedCharges, &a.RateChange, &a.NextRateChange, &a.PermittedUses, &a.ExclusiveUses,
		&a.ExtensionOption, &a.ExtensionOptionNotice, &a.ExpansionOption, &a.ExpansionOptionNotice, &a.RightOfFirstRefusal,
		&a.FLAGS, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
}

/*// ReadRentalAgreementGrids reads a full RentalAgreementGrid structure of data from the database based on the supplied Rows pointer.
func ReadRentalAgreementGrids(rows *sql.Rows, a *RentalAgreementGrid) error {
	return rows.Scan(&a.RAID, &a.TCIDPayor, &a.AgreementStart, &a.AgreementStop)
}*/

// ReadRentalAgreementPayor reads a full RentalAgreementPayor structure of data from the database based on the supplied Row pointer.
func ReadRentalAgreementPayor(row *sql.Row, a *RentalAgreementPayor) error {
	err := row.Scan(&a.RAPID, &a.RAID, &a.BID, &a.TCID, &a.DtStart, &a.DtStop, &a.FLAGS, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
	SkipSQLNoRowsError(&err)
	return err
}

// ReadRentalAgreementPayors reads a full RentalAgreementPayor structure of data from the database based on the supplied Rows pointer.
func ReadRentalAgreementPayors(rows *sql.Rows, a *RentalAgreementPayor) error {
	return rows.Scan(&a.RAPID, &a.RAID, &a.BID, &a.TCID, &a.DtStart, &a.DtStop, &a.FLAGS, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
}

// ReadRentalAgreementPet reads a full RentalAgreementPet structure of data from the database based on the supplied Row pointer.
func ReadRentalAgreementPet(row *sql.Row, a *RentalAgreementPet) error {
	err := row.Scan(&a.PETID, &a.BID, &a.RAID, &a.TCID, &a.Type, &a.Breed, &a.Color, &a.Weight, &a.Name, &a.DtStart, &a.DtStop,
		&a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
	SkipSQLNoRowsError(&err)
	return err
}

// ReadRentalAgreementPets reads a full RentalAgreementPet structure of data from the database based on the supplied Rows pointer.
func ReadRentalAgreementPets(rows *sql.Rows, a *RentalAgreementPet) error {
	return rows.Scan(&a.PETID, &a.BID, &a.RAID, &a.TCID, &a.Type, &a.Breed, &a.Color, &a.Weight, &a.Name, &a.DtStart, &a.DtStop,
		&a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
}

// ReadRentalAgreementRentable reads a full RentalAgreementRentable structure of data from the database based on the supplied Row pointer.
func ReadRentalAgreementRentable(row *sql.Row, a *RentalAgreementRentable) error {
	err := row.Scan(&a.RARID, &a.RAID, &a.BID, &a.RID, &a.CLID, &a.ContractRent, &a.RARDtStart, &a.RARDtStop, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
	SkipSQLNoRowsError(&err)
	return err
}

// ReadRentalAgreementRentables reads a full RentalAgreementRentable structure of data from the database based on the supplied Rows pointer.
func ReadRentalAgreementRentables(rows *sql.Rows, a *RentalAgreementRentable) error {
	return rows.Scan(&a.RARID, &a.RAID, &a.BID, &a.RID, &a.CLID, &a.ContractRent, &a.RARDtStart, &a.RARDtStop, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
}

// ReadRentalAgreementTemplate reads a full RentalAgreementTemplate structure of data from the database based on the supplied Row pointer.
func ReadRentalAgreementTemplate(row *sql.Row, a *RentalAgreementTemplate) error {
	err := row.Scan(&a.RATID, &a.BID, &a.RATemplateName, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
	SkipSQLNoRowsError(&err)
	return err
}

// ReadRentalAgreementTemplates reads a full RentalAgreementTemplate structure of data from the database based on the supplied Rows pointer.
func ReadRentalAgreementTemplates(rows *sql.Rows, a *RentalAgreementTemplate) error {
	return rows.Scan(&a.RATID, &a.BID, &a.RATemplateName, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
}

// ReadRentableUser reads a full RentableUser structure of data from the database based on the supplied Row pointer.
func ReadRentableUser(row *sql.Row, a *RentableUser) error {
	err := row.Scan(&a.RUID, &a.RID, &a.BID, &a.TCID, &a.DtStart, &a.DtStop, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
	SkipSQLNoRowsError(&err)
	return err
}

// ReadRentableUsers reads a full RentableUser structure of data from the database based on the supplied Rows pointer.
func ReadRentableUsers(rows *sql.Rows, a *RentableUser) error {
	return rows.Scan(&a.RUID, &a.RID, &a.BID, &a.TCID, &a.DtStart, &a.DtStop, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
}

// ReadStringList reads a full StringList structure from the database based on the supplied row object
func ReadStringList(row *sql.Row, a *StringList) error {
	err := row.Scan(&a.SLID, &a.BID, &a.Name, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
	SkipSQLNoRowsError(&err)
	return err
}

// ReadStringLists reads a full StringList structure from the database based on the supplied rows object
func ReadStringLists(rows *sql.Rows, a *StringList) error {
	return rows.Scan(&a.SLID, &a.BID, &a.Name, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
}

// ReadSubAR reads a full SubAR structure from the database based on the supplied row object
func ReadSubAR(row *sql.Row, a *SubAR) error {
	err := row.Scan(&a.SARID, &a.ARID, &a.SubARID, &a.BID, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
	SkipSQLNoRowsError(&err)
	return err
}

// ReadSubARs reads a full SubAR structure from the database based on the supplied row object
func ReadSubARs(row *sql.Rows, a *SubAR) error {
	err := row.Scan(&a.SARID, &a.ARID, &a.SubARID, &a.BID, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
	SkipSQLNoRowsError(&err)
	return err
}

// ReadSLString reads a full SLString structure from the database based on the supplied row object
func ReadSLString(row *sql.Row, a *SLString) error {
	err := row.Scan(&a.SLSID, &a.BID, &a.SLID, &a.Value, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
	SkipSQLNoRowsError(&err)
	return err
}

// ReadSLStrings reads a full SLString structure from the database based on the supplied rows
func ReadSLStrings(rows *sql.Rows, a *SLString) error {
	return rows.Scan(&a.SLSID, &a.BID, &a.SLID, &a.Value, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
}

//---------------------
//  TASKS
//---------------------

// ReadTask reads a full Task structure from the database based on the supplied row object
func ReadTask(row *sql.Row, a *Task) error {
	err := row.Scan(&a.TID, &a.BID, &a.TLID, &a.Name, &a.Worker, &a.DtDue, &a.DtPreDue, &a.DtDone, &a.DtPreDone, &a.FLAGS, &a.DoneUID, &a.PreDoneUID, &a.Comment, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
	SkipSQLNoRowsError(&err)
	return err
}

// ReadTasks reads a full Task structure from the database based on the supplied rows
func ReadTasks(rows *sql.Rows, a *Task) error {
	return rows.Scan(&a.TID, &a.BID, &a.TLID, &a.Name, &a.Worker, &a.DtDue, &a.DtPreDue, &a.DtDone, &a.DtPreDone, &a.FLAGS, &a.DoneUID, &a.PreDoneUID, &a.Comment, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
}

// ReadTaskList reads a full TaskList structure from the database based on the supplied row object
//            flds = "TLID,     BID, TLDID,    Name,    Cycle,    DtDue,    DtPreDue,    DtDone,    DtPreDone,    FLAGS,    DoneUID,    PreDoneUID,    Comment,    CreateTS,    CreateBy,    LastModTime,    LastModBy"
func ReadTaskList(row *sql.Row, a *TaskList) error {
	err := row.Scan(&a.TLID, &a.BID, &a.PTLID, &a.TLDID, &a.Name, &a.Cycle, &a.DtDue, &a.DtPreDue, &a.DtDone, &a.DtPreDone, &a.FLAGS, &a.DoneUID, &a.PreDoneUID, &a.EmailList, &a.DtLastNotify, &a.DurWait, &a.Comment, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
	SkipSQLNoRowsError(&err)
	return err
}

// ReadTaskLists reads a full TaskList structure from the database based on the supplied rows
func ReadTaskLists(rows *sql.Rows, a *TaskList) error {
	return rows.Scan(&a.TLID, &a.BID, &a.PTLID, &a.TLDID, &a.Name, &a.Cycle, &a.DtDue, &a.DtPreDue, &a.DtDone, &a.DtPreDone, &a.FLAGS, &a.DoneUID, &a.PreDoneUID, &a.EmailList, &a.DtLastNotify, &a.DurWait, &a.Comment, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
}

// ReadTaskDescriptor reads a full TaskDescriptor structure from the database based on the supplied row object
func ReadTaskDescriptor(row *sql.Row, a *TaskDescriptor) error {
	err := row.Scan(&a.TDID, &a.BID, &a.TLDID, &a.Name, &a.Worker, &a.EpochDue, &a.EpochPreDue, &a.FLAGS, &a.Comment, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
	SkipSQLNoRowsError(&err)
	return err
}

// ReadTaskDescriptors reads a full TaskDescriptor structure from the database based on the supplied rows
func ReadTaskDescriptors(rows *sql.Rows, a *TaskDescriptor) error {
	return rows.Scan(&a.TDID, &a.BID, &a.TLDID, &a.Name, &a.Worker, &a.EpochDue, &a.EpochPreDue, &a.FLAGS, &a.Comment, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
}

// ReadTaskListDefinition reads a full TaskListDefinition structure from the database based on the supplied row object
func ReadTaskListDefinition(row *sql.Row, a *TaskListDefinition) error {
	err := row.Scan(&a.TLDID, &a.BID, &a.Name, &a.Cycle, &a.Epoch, &a.EpochDue, &a.EpochPreDue, &a.FLAGS, &a.EmailList, &a.DurWait, &a.Comment, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
	SkipSQLNoRowsError(&err)
	return err
}

// ReadTaskListDefinitions reads a full TaskListDefinition structure from the database based on the supplied rows
func ReadTaskListDefinitions(rows *sql.Rows, a *TaskListDefinition) error {
	return rows.Scan(&a.TLDID, &a.BID, &a.Name, &a.Cycle, &a.Epoch, &a.EpochDue, &a.EpochPreDue, &a.FLAGS, &a.EmailList, &a.DurWait, &a.Comment, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
}

//---------------------
//  TRANSACTANT
//---------------------

// ReadTransactant reads a full Transactant structure from the database based on the supplied row object
func ReadTransactant(row *sql.Row, a *Transactant) error {
	err := row.Scan(&a.TCID, &a.BID, &a.NLID, &a.FirstName, &a.MiddleName, &a.LastName, &a.PreferredName,
		&a.CompanyName, &a.IsCompany, &a.PrimaryEmail, &a.SecondaryEmail, &a.WorkPhone, &a.CellPhone,
		&a.Address, &a.Address2, &a.City, &a.State, &a.PostalCode, &a.Country, &a.Website, &a.FLAGS,
		&a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
	SkipSQLNoRowsError(&err)
	return err
}

// ReadTransactants reads a full Transactant structure from the database based on the supplied rows object
func ReadTransactants(rows *sql.Rows, a *Transactant) error {
	return rows.Scan(&a.TCID, &a.BID, &a.NLID, &a.FirstName, &a.MiddleName, &a.LastName, &a.PreferredName,
		&a.CompanyName, &a.IsCompany, &a.PrimaryEmail, &a.SecondaryEmail, &a.WorkPhone, &a.CellPhone,
		&a.Address, &a.Address2, &a.City, &a.State, &a.PostalCode, &a.Country, &a.Website, &a.FLAGS,
		&a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
}

// ReadTransactantTypeDowns reads the TCID and full name of Transactants based on the supplied rows object
func ReadTransactantTypeDowns(rows *sql.Rows, a *TransactantTypeDown) error {
	return rows.Scan(&a.TCID, &a.FirstName, &a.MiddleName, &a.LastName, &a.CompanyName, &a.IsCompany)
}

// ReadPayor reads a full Payor structure from the database based on the supplied row object
func ReadPayor(row *sql.Row, a *Payor) error {
	var b1, d1 string
	err := row.Scan(&a.TCID, &a.BID, &a.CreditLimit, &a.TaxpayorID, &a.ThirdPartySource, &a.EligibleFuturePayor,
		&a.FLAGS, &b1, &d1, &a.GrossIncome, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
	SkipSQLNoRowsError(&err)
	if err != nil {
		return err
	}
	b, err := hex.DecodeString(b1)
	if err != nil {
		return err
	}
	a.SSN, err = DecryptOrEmpty(b)
	if err != nil {
		return err
	}
	d, err := hex.DecodeString(d1)
	a.DriversLicense, err = DecryptOrEmpty(d)
	if err != nil {
		return err
	}
	return nil
}

// ReadPayors reads a full Payor structure from the database based on the supplied rows object
func ReadPayors(rows *sql.Rows, a *Payor) error {
	return rows.Scan(&a.TCID, &a.BID, &a.CreditLimit, &a.TaxpayorID, &a.ThirdPartySource, &a.EligibleFuturePayor,
		&a.FLAGS, &a.SSN, &a.DriversLicense, &a.GrossIncome, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
}

// ReadUser reads a full User structure from the database based on the supplied row object
func ReadUser(row *sql.Row, a *User) error {
	err := row.Scan(&a.TCID, &a.BID, &a.Points, &a.DateofBirth, &a.EmergencyContactName, &a.EmergencyContactAddress,
		&a.EmergencyContactTelephone, &a.EmergencyContactEmail, &a.AlternateAddress, &a.EligibleFutureUser, &a.FLAGS, &a.Industry, &a.SourceSLSID,
		&a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
	SkipSQLNoRowsError(&err)
	return err
}

// ReadUsers reads a full User structure from the database based on the supplied rows object
func ReadUsers(rows *sql.Rows, a *User) error {
	return rows.Scan(&a.TCID, &a.BID, &a.Points, &a.DateofBirth, &a.EmergencyContactName, &a.EmergencyContactAddress,
		&a.EmergencyContactTelephone, &a.EmergencyContactEmail, &a.AlternateAddress, &a.EligibleFutureUser, &a.FLAGS, &a.Industry, &a.SourceSLSID,
		&a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
}

// ReadVehicle reads a full Vehicle structure from the database based on the supplied row object
func ReadVehicle(row *sql.Row, a *Vehicle) error {
	err := row.Scan(&a.VID, &a.TCID, &a.BID, &a.VehicleType, &a.VehicleMake, &a.VehicleModel, &a.VehicleColor, &a.VehicleYear,
		&a.VIN, &a.LicensePlateState, &a.LicensePlateNumber, &a.ParkingPermitNumber, &a.DtStart, &a.DtStop,
		&a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
	SkipSQLNoRowsError(&err)
	return err
}

// ReadVehicles reads a full Vehicle structure from the database based on the supplied rows object
func ReadVehicles(rows *sql.Rows, a *Vehicle) error {
	return rows.Scan(&a.VID, &a.TCID, &a.BID, &a.VehicleType, &a.VehicleMake, &a.VehicleModel, &a.VehicleColor, &a.VehicleYear,
		&a.VIN, &a.LicensePlateState, &a.LicensePlateNumber, &a.ParkingPermitNumber, &a.DtStart, &a.DtStop,
		&a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
}
