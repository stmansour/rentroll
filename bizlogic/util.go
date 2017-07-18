package bizlogic

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
		Message: "Error inserting assessment = " + (*err).Error(),
	}
	errlist = append(errlist, berr)
	return errlist
}
