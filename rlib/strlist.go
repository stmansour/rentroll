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
	"Voided by an amended version", // 2: automatic - when RA in the future is amended by a different RA
}

// MSGAPPDECLINED et al are indeces into the the stringlist for ROLLERSL.
const (
	MSGAPPDECLINED = 0
	MSGRAUPDATED   = 1
	MSGRACANCELLED = 2
)

// GetRollerStringList creates a stringlist that Roller must have in order
// to process RentalAgreement state changes, etc. When new strings are added
// tho RollerStrings initial values list over time, they will be added to
// the database if they don't exist.
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
		//----------------------------------------------------
		// Are all strings accounted for? If not, update...
		//----------------------------------------------------
		for i := len(t.S); i < len(RollerStrings); i++ {
			var s = SLString{
				BID:   bid,
				SLID:  t.SLID,
				Value: RollerStrings[i],
			}
			_, err = InsertSLString(ctx, &s)
			if err != nil {
				return t, err
			}
			t.S = append(t.S, s)
		}
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
