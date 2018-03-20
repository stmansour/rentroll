package rlib

import (
	"context"
	"database/sql"
	"extres"
	"sort"
)

// RAFlow etc.. all are list of all flows exist in the system
const (
	RAFlow string = "RA"
)

// RAFlowPartType is type of rental agreement flow part
type RAFlowPartType int

// DatesRAFlowPart etc. all are constants for rental agreement flow part
const (
	DatesRAFlowPart RAFlowPartType = 1 + iota // must start from 1
	PeopleRAFlowPart
	PetsRAFlowPart
	VehiclesRAFlowPart
	BackGroundInfoRAFlowPart
	RentablesRAFlowPart
	FeesTermsRAFlowPart
)

// IsValid checks the validity of RAFlowPartType raftp
func (raftp RAFlowPartType) IsValid() bool {
	if raftp < DatesRAFlowPart || raftp > FeesTermsRAFlowPart {
		return false
	}

	return true
}

// String representation of RAFlowPartType
func (raftp RAFlowPartType) String() string {
	names := [...]string{
		"Agreement Dates",
		"Payors-Users-Guarantors",
		"Pets",
		"Vehicles",
		"Background-Info",
		"Rentables",
		"Fess-Terms",
	}

	// if not valid then return unknown
	if !(raftp.IsValid()) {
		return "Unknown-RAFlowPart"
	}

	return names[raftp-1]
}

// InsertInitialRAFlow writes a bunch of flow's sections record for a particular RA
// This should be run under atomic transaction mode as per DB design of flow
// This is very special case that we're not returning primary key generated from database
// instead we're generating in form of string which we return if tx will be succeed.
func InsertInitialRAFlow(ctx context.Context, BID int64) (string, error) {

	var (
		flowID string
		err    error
		sess   *Session
		ok     bool
		UID    = int64(0)
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		sess, ok = SessionFromContext(ctx)
		if !ok {
			return flowID, ErrSessionRequired
		}
		UID = sess.UID
	}

	// ------------
	// SPECIAL CASE
	// ------------
	var (
		newTx bool
		tx    *sql.Tx
	)

	if tx, ok = DBTxFromContext(ctx); !ok { // if transaction is NOT supplied
		newTx = true
		tx, err = RRdb.Dbrr.Begin()
		if err != nil {
			return flowID, err
		}
		ctx = SetDBTxContextKey(ctx, tx)
	}

	// getFlowID first
	flowID = getFlowID(UID)

	// initRAFlowPart
	initRAFlowPart := FlowPart{
		BID:       BID,
		Flow:      RAFlow,
		FlowID:    flowID,
		PartType:  0,
		Data:      []byte("null"), // JSON "null" primitive type
		CreateBy:  UID,
		LastModBy: UID,
	}

	// Rental agreement flow parts map init
	// maybe we can just override the above pre-defined initFlowPart struct
	initRAFlowMap := map[RAFlowPartType]FlowPart{
		DatesRAFlowPart:          FlowPart{},
		PeopleRAFlowPart:         FlowPart{},
		PetsRAFlowPart:           FlowPart{},
		VehiclesRAFlowPart:       FlowPart{},
		BackGroundInfoRAFlowPart: FlowPart{},
		RentablesRAFlowPart:      FlowPart{},
		FeesTermsRAFlowPart:      FlowPart{},
	}

	// insert in order to ease
	var keys Int64Range
	for k := range initRAFlowMap {
		keys = append(keys, int64(k))
	}
	sort.Sort(keys)

	// assign part type
	for _, v := range keys {
		partTypeID := RAFlowPartType(v)
		// fmt.Printf("partTypeID: %s: %d\n", partTypeID, partTypeID)

		// get blank flow part
		a := initRAFlowMap[RAFlowPartType(partTypeID)]

		// assign pre-defined init flow data
		a = initRAFlowPart

		// modify part type
		a.PartType = int(partTypeID)

		// insert each flowpart of RA flow
		_, err = InsertFlowPart(ctx, &a)
		if err != nil {
			Ulog("Error while inserting FlowPart BULK-WRITE: %s\n", err.Error())
		}
	}

	if newTx { // if new transaction then commit it
		// if error then rollback
		if err = tx.Commit(); err != nil {
			tx.Rollback()
			Ulog("Error while Committing transaction | inserting FlowPart BULK-WRITE: %s\n", err.Error())
			err = insertError(err, "InitialRAFlow", nil)
			return flowID, err
		}
	}

	return flowID, err
}
