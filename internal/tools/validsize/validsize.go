package validsize

import (
	"os"

	"github.com/BillyBones007/pwdm_client/internal/customerror"
)

// Limit file size.
const fileSizeLimit int64 = 1024 * 1024 * 10 // 10Mb

// ValidFileSize - checks file size(10Мb limit).
func ValidFileSize(file string) (bool, error) {
	fileInfo, err := os.Stat(file)
	if err != nil {
		return false, err
	}
	if !(fileInfo.Size() <= fileSizeLimit) {
		return false, customerror.ErrFileSize
	}
	return true, nil
}
