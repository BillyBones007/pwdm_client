package encrypttools

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
)

// Encrypter - main encrypt/decrypt type.
type Encrypter struct {
	cryptKey []byte
}

// NewEncrypter - will set the key to be used for encryption and decryption.
// Automatically generated from the user's password.
func NewEncrypter(password string) *Encrypter {
	k := sha256.Sum256([]byte(password))
	hasher := md5.New()
	hasher.Write(k[:8])
	hash := hasher.Sum(k[16:24])
	key := make([]byte, hex.EncodedLen(len(hash)))
	hex.Encode(key, hash)
	return &Encrypter{cryptKey: key[:16]}
}

// EncryptString - will AES-encrypt and base64 encode the given string.
func (e *Encrypter) EncryptString(text string) (string, error) {
	ciphertext, err := e.EncryptBytes([]byte(text), false)
	if err != nil {
		return "", err
	}
	return ciphertext.(string), nil
}

// EncryptBytes - will AES-encrypt the given byte slice.
// bytemode bool - true returns []byte
// bytemode bool - false returns string
func (e *Encrypter) EncryptBytes(data []byte, bytemode bool) (any, error) {
	// TODO: реализовать условие, если bytemode true - функция возвратит []byte
	// иначе string
	block, err := aes.NewCipher(e.cryptKey)
	if err != nil {
		return nil, err
	}

	// generate IV
	ciphertext := make([]byte, aes.BlockSize+len(data))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	// if bytemode true - returns []byte
	if bytemode {
		cipher.NewCFBEncrypter(block, iv).XORKeyStream(ciphertext[aes.BlockSize:], data)
		return ciphertext, nil
	}
	// else string
	cipher.NewCFBEncrypter(block, iv).XORKeyStream(ciphertext[aes.BlockSize:], data)
	return base64.URLEncoding.EncodeToString(ciphertext), nil
}

// DecryptString - will base64 decode and then AES-decrypt the given string.
// bytemode bool - true returns []byte
// bytemode bool - false returns string
func (e *Encrypter) DecryptString(encrypted string, bytemode bool) (any, error) {
	decodedBytes, err := base64.URLEncoding.DecodeString(encrypted)
	if err != nil {
		return nil, err
	}

	// if bytemode true - returns []byte
	if bytemode {
		ciphertext, err := e.DecryptBytes(decodedBytes)
		if err != nil {
			return nil, err
		}
		return ciphertext, nil
	}
	// else string
	ciphertext, err := e.DecryptBytes(decodedBytes)
	if err != nil {
		return nil, err
	}
	return string(ciphertext), nil
}

// DecryptBytes - will AES-decrypt the given byte slise.
func (e *Encrypter) DecryptBytes(encrypted []byte) ([]byte, error) {
	block, err := aes.NewCipher(e.cryptKey)
	if err != nil {
		return nil, err
	}

	if byteLen := len(encrypted); byteLen < aes.BlockSize {
		return nil, fmt.Errorf("invalid cipher size %d, expicted at least %d", byteLen, aes.BlockSize)
	}

	iv := encrypted[:aes.BlockSize]
	encrypted = encrypted[aes.BlockSize:]

	cipher.NewCFBDecrypter(block, iv).XORKeyStream(encrypted, encrypted)

	return encrypted, nil
}
