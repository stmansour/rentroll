package rlib

import (
	"context"
	"strings"
)

// RollerStrings is a slice of strings needed by Roller for
// processing the states of a Rental Agreement, etc.
var RollerStrings = []string{
	"Application was declined",
}

// CreateRollerStringList creates a stringlist that Roller must have in order
// to process RentalAgreement state changes, etc.
//
//
func CreateRollerStringList(bid int64) error {
	// funcname := "CreateRollerStringList"

	ctx := context.Background()
	//-------------------------------------------------------------------
	// If we already have it, then don't create...
	//-------------------------------------------------------------------
	var t StringList
	err := GetStringListByName(ctx, bid, ROLLERSL, &t) // do we already have a stringlist by this name?
	if err != nil {
		return err
	}
	if t.SLID > 0 {
		return nil
	}

	//-------------------------------------------------------------------
	// We need to create it...
	//-------------------------------------------------------------------
	var sl = StringList{
		BID:  bid,
		Name: ROLLERSL,
	}

	for i := 0; i < len(RollerStrings); i++ {
		var sls = SLString{
			BID:   bid,
			SLID:  sl.SLID,
			Value: strings.TrimSpace(RollerStrings[i]),
		}
		sl.S = append(sl.S, sls)
	}

	_, err = InsertStringList(ctx, &sl)
	if err != nil {
		return err
	}
	return nil
}
