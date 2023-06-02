package storage

import (
	"testing"

	"github.com/BillyBones007/pwdm_client/internal/datatypes"
	"github.com/BillyBones007/pwdm_client/internal/storage/models"
	"github.com/stretchr/testify/assert"
)

func TestStorage(t *testing.T) {
	t.Run("Clear storage", func(t *testing.T) {
		storage := NewStorage()
		storage.LogPwdData[1] = models.InfoModel{Id: 1, Title: "Test", Type: datatypes.LoginPasswordDataType}
		err := storage.Clear(datatypes.LoginPasswordDataType)
		assert.NoError(t, err)
		_, ok := storage.LogPwdData[1]
		assert.False(t, ok)

	})

	t.Run("Update storage", func(t *testing.T) {
		storage := NewStorage()
		storage.LogPwdData[1] = models.InfoModel{Id: 1, Title: "One", Type: datatypes.LoginPasswordDataType}
		records := []models.InfoModel{
			{Id: 1, Type: datatypes.LoginPasswordDataType, Title: "Two"},
			{Id: 2, Type: datatypes.LoginPasswordDataType, Title: "Three"},
		}
		err := storage.UpdateStorage(datatypes.LoginPasswordDataType, records)
		assert.NoError(t, err)
		l := len(storage.LogPwdData)
		assert.Equal(t, l, 2)
	})

	t.Run("Update storage (empty struct)", func(t *testing.T) {
		storage := NewStorage()
		storage.LogPwdData[1] = models.InfoModel{Id: 1, Title: "One", Type: datatypes.LoginPasswordDataType}
		records := []models.InfoModel{}
		err := storage.UpdateStorage(datatypes.LoginPasswordDataType, records)
		assert.Error(t, err)
		l := len(storage.LogPwdData)
		assert.Equal(t, l, 0)
	})

	t.Run("Get list records", func(t *testing.T) {
		storage := NewStorage()
		storage.LogPwdData[1] = models.InfoModel{Id: 1, Title: "One", Type: datatypes.LoginPasswordDataType}
		records := []models.InfoModel{
			{Id: 2, Type: datatypes.LoginPasswordDataType, Title: "Two"},
			{Id: 3, Type: datatypes.LoginPasswordDataType, Title: "Three"},
		}
		err := storage.UpdateStorage(datatypes.LoginPasswordDataType, records)
		assert.NoError(t, err)
		list := storage.GetListRecords(datatypes.LoginPasswordDataType)
		assert.Equal(t, len(list), 2)
	})

	t.Run("Get list records (empty storage)", func(t *testing.T) {
		storage := NewStorage()
		list := storage.GetListRecords(datatypes.LoginPasswordDataType)
		assert.Equal(t, len(list), 0)
	})

	t.Run("Get id record", func(t *testing.T) {
		storage := NewStorage()
		storage.LogPwdData[1] = models.InfoModel{Id: 1, Title: "One", Type: datatypes.LoginPasswordDataType}
		records := []models.InfoModel{
			{Id: 2, Type: datatypes.LoginPasswordDataType, Title: "Two"},
			{Id: 3, Type: datatypes.LoginPasswordDataType, Title: "Three"},
		}
		err := storage.UpdateStorage(datatypes.LoginPasswordDataType, records)
		assert.NoError(t, err)
		id := storage.GetIdRecord(1, datatypes.LoginPasswordDataType)
		assert.Equal(t, id, int32(2))
	})

	t.Run("Get id record (bad key)", func(t *testing.T) {
		storage := NewStorage()
		storage.LogPwdData[1] = models.InfoModel{Id: 1, Title: "One", Type: datatypes.LoginPasswordDataType}
		id := storage.GetIdRecord(2, datatypes.LoginPasswordDataType)
		assert.Equal(t, id, int32(-1))
	})
}
