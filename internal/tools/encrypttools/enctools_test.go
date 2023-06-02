package encrypttools

import (
	"crypto/rand"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetCryptKey(t *testing.T) {
	testsEqual := []struct {
		name      string
		password1 string
		password2 string
	}{
		{
			name:      "Equal_test1",
			password1: "my password",
			password2: "my password",
		},
		{
			name:      "Equal_test2",
			password1: "equal password",
			password2: "equal password",
		},
	}

	for _, tt := range testsEqual {
		t.Run(tt.name, func(t *testing.T) {
			enc1 := NewEncrypter(tt.password1)
			enc2 := NewEncrypter(tt.password2)
			fmt.Printf("Key 1: %v\n", enc1.cryptKey)
			fmt.Printf("Key 2: %v\n", enc2.cryptKey)
			assert.Equal(t, enc1.cryptKey, enc2.cryptKey)
		})
	}
	testsNotEqual := []struct {
		name      string
		password1 string
		password2 string
	}{
		{
			name:      "Not_equal_test1",
			password1: "my password",
			password2: "password",
		},
		{
			name:      "Not_equal_test2",
			password1: "equal password",
			password2: "not equal password",
		},
	}

	for _, tt := range testsNotEqual {
		t.Run(tt.name, func(t *testing.T) {
			enc1 := NewEncrypter(tt.password1)
			enc2 := NewEncrypter(tt.password2)

			assert.NotEqual(t, enc1.cryptKey, enc2.cryptKey)
		})
	}
}

func TestEncDecBytes(t *testing.T) {
	data1, _ := helperRandBytes(32)
	data2, _ := helperRandBytes(10)
	enc := NewEncrypter("my password")
	tests := []struct {
		name  string
		data  []byte
		wrong []byte
	}{
		{
			name:  "one test",
			data:  data1,
			wrong: data1,
		},
		{
			name:  "two test",
			data:  data2,
			wrong: data2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encripted, err := enc.EncryptBytes(tt.data, true)
			if err != nil {
				fmt.Println(err)
				t.Fail()
			}
			encrByte := encripted.([]byte)
			fmt.Printf("Test data %v\nEncrypted data %v\n", tt.data, encripted)
			decripted, _ := enc.DecryptBytes(encrByte)
			fmt.Printf("Decrypted data %v\n", decripted)

			assert.Equal(t, decripted, tt.wrong)

		})
	}

}

func TestEncDecString(t *testing.T) {
	data1 := "First string"
	data2 := "Second string"
	enc := NewEncrypter("my password")
	tests := []struct {
		name  string
		data  string
		wrong string
	}{
		{
			name:  "one test",
			data:  data1,
			wrong: data1,
		},
		{
			name:  "two test",
			data:  data2,
			wrong: data2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encripted, err := enc.EncryptString(tt.data)
			if err != nil {
				fmt.Println(err)
				t.Fail()
			}
			fmt.Printf("Test data %s\nEncrypted data %s\n", tt.data, encripted)
			decripted, _ := enc.DecryptString(encripted, false)
			decriptData := decripted.(string)
			fmt.Printf("Decrypted data %v\n", decriptData)

			assert.Equal(t, decriptData, tt.wrong)

		})
	}

}

func BenchmarkEncryptBytes(b *testing.B) {
	enc := NewEncrypter("my password")
	data, _ := helperRandBytes(1024 * 1024 * 10)
	var encripted []byte
	b.ResetTimer()

	b.Run("encrypt", func(b *testing.B) {
		encData, err := enc.EncryptBytes(data, true)
		if err != nil {
			fmt.Println(err)
			b.Fail()
		}
		encripted = encData.([]byte)
	})

	b.Run("decrypt", func(b *testing.B) {
		decData, err := enc.DecryptBytes(encripted)
		if err != nil {
			fmt.Println(err)
			b.Fail()
		}
		decData = []byte("OK")
		fmt.Println(string(decData))
	})
}

// helperRandBytes - helper function.
// Uses in tests.
func helperRandBytes(size int) ([]byte, error) {
	b := make([]byte, size)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}
