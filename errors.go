package nxprd

import "fmt"

// ExtendedErrorCode defines errors of the wrapper. The C library has its own
// error codes.
type ExtendedErrorCode int

// All extended error codes of the wrapper.
const (
	// TimeoutErr is used if a timeout occurred when using Discover().
	TimeoutErr ExtendedErrorCode = 1 + iota
	// NoErr is used if no error occurred.
	NoErr
)

var extCodes = [...]string{
	"TimeoutErr",
	"NoErr",
}

// String returns the name of an ExtendedErrorCode (emulating an enum).
func (ec ExtendedErrorCode) String() string {
	return extCodes[ec-1]
}

// NxpError defines the error struct that is used to indicate an error of the
// C library or the GO wrapper.
type NxpError struct {
	// Comp defines the component of the C library where the error occurred.
	// 0 stands for "no error".
	Comp int
	// Code defines the actual error that occurred in the C library.
	// 0 stands for "no error".
	Code int
	// ExtCode defines the extended error code of the wrapper. ExtCode will have
	// the NoErr value, if no extended error occurred.
	ExtCode ExtendedErrorCode
	// Msg describes the error.
	Msg string
}

// Error returns the error information as a human readable string.
func (e *NxpError) Error() string {
	return fmt.Sprintf("component code: %x, error code: %x, ext error code: %s msg: %s", e.Comp, e.Code, e.ExtCode, e.Msg)
}

func createLibErr(status int) *NxpError {
	return &NxpError{(status >> 8), (status & 0x00FF), NoErr, ""}
}

func createExtErr(extCode ExtendedErrorCode, msg string) *NxpError {
	return &NxpError{0, 0, extCode, msg}
}
