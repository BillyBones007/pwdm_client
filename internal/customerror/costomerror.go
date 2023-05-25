package customerror

import "errors"

var (
	ErrDataTypeIncorrect error = errors.New("accepts incorrect data type (supported types: Login/Password pair - 1, Cards data - 2, Text data - 3, Binary data - 4)")
	ErrEmptyListRecords  error = errors.New("list new records is empty")
	ErrFileSize          error = errors.New("file size is larger than the limit")

	// ErrLogPwdIncorrect - displayed to the user
	ErrLogPwdIncorrect error = errors.New("login or password incorrect")
	// ErrUnknown - displayed to the user if error do not ErrLogPwdIncorrect
	ErrUnknown error = errors.New("unknown error")
	// ErrGRPCLowPwdIncorrect - return from server
	ErrGRPCLogPwdIncorrect error = errors.New("rpc error: code = Unknown desc = login or password incorrect")
	// ErrUserIsExists - displayed to the user
	ErrUserIsExists error = errors.New("user is exists, try using a different username")
	// ErrGRPCUserIsExists - return from server
	ErrGRPCUserIsExists error = errors.New("rpc error: code = Unknown desc = user is exists")

	ErrInvalildCommand error = errors.New("invalid command")
	ErrReadTextFile    error = errors.New("error reading text file")
	ErrEmptyFields     error = errors.New("required fields is empty")
)
