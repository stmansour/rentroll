package core

import (
	"fmt"
)

var (
	// ErrInternal :Internal program error
	ErrInternal = fmt.Errorf("Internal Error, please try again later")
	// ErrFileNotFound :File not found error
	ErrFileNotFound = fmt.Errorf("File could not be found")
)
