package rlib

import (
	"context"
	"strings"
)

// RollerStrings is a slice of strings needed by Roller for
// processing the states of a Rental Agreement, etc.
var RollerStrings = []string{
	"Application was declined",     // 0: automatic - when RA is terminated due to applicaion being declined
	"Rental Agreement was updated", // 1: automatic - when RA is updated with a changed version
}

// MSGAPPDECLINED et al are indeces into the the stringlist for ROLLERSL.
const (
	MSGAPPDECLINED = 0
	MSGRAUPDATED   = 1
)

// GetRollerStringList creates a stringlist that Roller must have in order
// to process RentalAgreement state changes, etc.
//
//
func GetRollerStringList(ctx context.Context, bid int64) (StringList, error) {
	// funcname := "GetRollerStringList"

	//-------------------------------------------------------------------
	// If we already have it, then don't create...
	//-------------------------------------------------------------------
	var t StringList
	err := GetStringListByName(ctx, bid, ROLLERSL, &t) // do we already have a stringlist by this name?
	if err != nil {
		return t, err
	}
	if t.SLID > 0 {
		return t, nil
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
		return sl, err
	}
	return sl, nil
}
