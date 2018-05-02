package bizlogic

import (
	"context"
	"fmt"
	"rentroll/rlib"
)

// ValidateTaskDescriptor checks to see whether the task descriptor violates any
// business logic.
//
// INPUTS
//    a = the task descriptor to validate
//
// RETURNS
//    a slice of BizErrors
//-------------------------------------------------------------------------------------
func ValidateTaskDescriptor(ctx context.Context, a *rlib.TaskDescriptor) []BizError {
	var e []BizError
	var tldid, bid int64

	//------------------------------------------------------------
	// Validate that we have a TaskListDefinition that exists...
	//------------------------------------------------------------
	qry := fmt.Sprintf("SELECT TLDID,BID FROM TaskListDefinition WHERE TLDID=%d", a.TLDID)
	de := rlib.RRdb.Dbrr.QueryRow(qry).Scan(&tldid, &bid)
	if de != nil {
		if rlib.IsSQLNoResultsError(de) {
			s := fmt.Sprintf(BizErrors[UnknownTLDID].Message, a.TLDID)
			b := BizError{Errno: UnknownTLDID, Message: s}
			e = append(e, b)
			return e
		}
		return bizErrSys(&de)
	}

	//----------------------------------------------------
	// Validate that it is part of the same Business...
	//----------------------------------------------------
	if bid != a.BID {
		s := fmt.Sprintf(BizErrors[ImproperTLDID].Message, a.TLDID, a.BID)
		b := BizError{Errno: ImproperTLDID, Message: s}
		e = append(e, b)
		return e
	}

	//----------------------------------------------------------------
	// Ensure that there is a name.
	//----------------------------------------------------------------
	if len(a.Name) == 0 {
		rlib.Console("*** MISSING NAME.  Task Descriptor a.TDID = %d has 0 length name\n", a.TDID)
		s := fmt.Sprintf(BizErrors[TaskDescrMissingName].Message, a.TLDID, a.BID)
		b := BizError{Errno: MissingName, Message: s}
		e = append(e, b)
	}

	return e
}
