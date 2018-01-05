package bizlogic

import "fmt"

// bizErrSys just encapsulates returning an error in a []BizError.  The Errno
// is set to 0.
//
// INPUTS
//  err = pointer to an error
//
// RETURNS
//  a slize of BizError containing the error message
//-------------------------------------------------------------------------------------
func bizErrSys(err *error) []BizError {
	var errlist []BizError
	berr := BizError{
		Errno:   0, // system error
		Message: (*err).Error(),
	}
	errlist = append(errlist, berr)
	return errlist
}

// BizErrorListToError returns a single error with each BizErr on a separate line.
// if the supplied errlist is nil or len(errlist) is 0 then the return value is nil
//
// INPUTS
//  errlist = array of BizError
//
// RETURNS
//  a single error with error messages aggregated or nil if errlist was empty or nil
//-------------------------------------------------------------------------------------
func BizErrorListToError(errlist []BizError) error {
	if errlist == nil || 0 == len(errlist) {
		return nil
	}
	errmsg := ""
	for i := 0; i < len(errlist); i++ {
		errmsg += errlist[i].Message + "\n"
	}
	return fmt.Errorf("%s", errmsg)
}
