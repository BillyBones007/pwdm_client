package storage

import (
	"fmt"

	"github.com/BillyBones007/pwdm_client/internal/customerror"
	"github.com/BillyBones007/pwdm_client/internal/datatypes"
	"github.com/BillyBones007/pwdm_client/internal/storage/models"
)

// Main storage.
type Storage struct {
	LogPwdData map[int]models.InfoModel
	CardData   map[int]models.InfoModel
	TextData   map[int]models.InfoModel
	BinaryData map[int]models.InfoModel
}

// NewStorage - returns a pointer to the Storage.
func NewStorage() *Storage {
	storage := &Storage{
		LogPwdData: make(map[int]models.InfoModel),
		CardData:   make(map[int]models.InfoModel),
		TextData:   make(map[int]models.InfoModel),
		BinaryData: make(map[int]models.InfoModel),
	}
	return storage
}

// Clear - clear the data in storage.
func (s *Storage) Clear(dataType ...int32) error {
	// TODO: решить на будущее, может просто создать новую map
	// и передать в поле Data новый указатель, а старую map
	// очистит в последствии GC
	for _, dtype := range dataType {
		switch dtype {
		case datatypes.LoginPasswordDataType:
			for k := range s.LogPwdData {
				delete(s.LogPwdData, k)
			}
			return nil
		case datatypes.CardDataType:
			for k := range s.CardData {
				delete(s.CardData, k)
			}
			return nil
		case datatypes.TextDataType:
			for k := range s.TextData {
				delete(s.TextData, k)
			}
			return nil
		case datatypes.BinaryDataType:
			for k := range s.BinaryData {
				delete(s.BinaryData, k)
			}
			return nil
		default:
			return customerror.ErrDataTypeIncorrect
		}
	}
	return nil
}

// UpdateStorage - updates the storage.
func (s *Storage) UpdateStorage(dataType int32, listRecords []models.InfoModel) error {
	// TODO: обновляет базу данных свежими данными с сервера
	switch dataType {
	case datatypes.LoginPasswordDataType:
		err := s.updateMap(datatypes.LoginPasswordDataType, s.LogPwdData, listRecords)
		if err != nil {
			return err
		}
		return nil
	case datatypes.CardDataType:
		err := s.updateMap(datatypes.CardDataType, s.CardData, listRecords)
		if err != nil {
			return err
		}
		return nil
	case datatypes.TextDataType:
		err := s.updateMap(datatypes.TextDataType, s.TextData, listRecords)
		if err != nil {
			return err
		}
		return nil
	case datatypes.BinaryDataType:
		err := s.updateMap(datatypes.BinaryDataType, s.BinaryData, listRecords)
		if err != nil {
			return err
		}
		return nil
	default:
		return customerror.ErrDataTypeIncorrect
	}
}

// updateMap - helper function for updating storage.
func (s *Storage) updateMap(dataType int32, table map[int]models.InfoModel, listRecords []models.InfoModel) error {
	err := s.Clear(dataType)
	if err != nil {
		return err
	}
	i := 1
	if len(listRecords) > 0 {
		for _, v := range listRecords {
			if v.Type == dataType {
				table[i] = v
				i++
			}
		}
		return nil
	}
	return customerror.ErrEmptyListRecords
}

// GetIdRecord - returns id record from map.
// Accepts map key and data type.
// If key not found - returns -1.
func (s *Storage) GetIdRecord(key int, dataType int32) int32 {
	switch dataType {
	case datatypes.LoginPasswordDataType:
		if v, ok := s.LogPwdData[key]; ok {
			return v.Id
		}
		return -1
	case datatypes.CardDataType:
		if v, ok := s.CardData[key]; ok {
			return v.Id
		}
		return -1
	case datatypes.TextDataType:
		if v, ok := s.TextData[key]; ok {
			fmt.Printf("INFO: ID %d\n", v.Id)
			return v.Id
		}
		return -1
	case datatypes.BinaryDataType:
		if v, ok := s.BinaryData[key]; ok {
			return v.Id
		}
		return -1
	default:
		return -1
	}
}

// GetListRecords - returns the map records models.InfoModel for current data type.
func (s *Storage) GetListRecords(dataType int32) map[int]models.InfoModel {
	switch dataType {
	case datatypes.LoginPasswordDataType:
		return s.LogPwdData
	case datatypes.CardDataType:
		return s.CardData
	case datatypes.TextDataType:
		return s.TextData
	case datatypes.BinaryDataType:
		return s.BinaryData
	default:
		return nil
	}
}
