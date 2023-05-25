package filetools

import (
	"errors"
	"io/ioutil"
	"os"

	"github.com/BillyBones007/pwdm_client/internal/customerror"
)

// Limit file size.
const fileSizeLimit int64 = 1024 * 1024 * 3 // 3Mb

// ValidFileSize - checks file size(3Мb limit).
func validFileSize(file string) (bool, error) {
	fileInfo, err := os.Stat(file)
	if err != nil {
		return false, err
	}
	if fileInfo.Size() > fileSizeLimit {
		return false, customerror.ErrFileSize
	}
	return true, nil
}

// ReadTextFile - read text file and returns string.
func ReadTextFile(path string) (string, error) {
	_, err := validFileSize(path)
	if err != nil {
		if errors.Is(err, customerror.ErrFileSize) {
			return "", customerror.ErrFileSize
		} else {
			return "", customerror.ErrReadTextFile
		}
	}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", customerror.ErrReadTextFile
	}
	return string(data), nil
}

// ReadBinaryFile - read binary file and returns []byte.
func ReadBinaryFile(path string) ([]byte, error) {
	_, err := validFileSize(path)
	if err != nil {
		if errors.Is(err, customerror.ErrFileSize) {
			return nil, customerror.ErrFileSize
		} else {
			return nil, customerror.ErrReadTextFile
		}
	}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, customerror.ErrReadTextFile
	}
	return data, nil
}

// WriteTextFile - creates and writes a text file with the specified name.
func WriteTextFile(data string, name string) error {
	err := ioutil.WriteFile(name, []byte(data), 0644)
	if err != nil {
		return err
	}
	return nil
}

// WriteBinaryFile - creates and writes a binary file with the specified name.
func WriteBinaryFile(data []byte, name string) error {
	err := ioutil.WriteFile(name, data, 0644)
	if err != nil {
		return err
	}
	return nil
}
