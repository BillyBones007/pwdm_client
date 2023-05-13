package customerror

import "errors"

var (
	ErrDataTypeIncorrect error = errors.New("accepts incorrect data type (supported types: Login/Password pair - 1, Cards data - 2, Text data - 3, Binary data - 4)")
	ErrEmptyListRecords  error = errors.New("list new records is empty")
	ErrFileSize          error = errors.New("file size is larger than the limit")
)
