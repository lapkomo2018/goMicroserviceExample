package service

import (
	"github.com/lapkomo2018/goServiceExample/pkg/model"
)

type (
	Hasher interface {
		Hash(password string) string
		Verify(hash string, password string) bool
	}

	StructStorage interface {
		ExistsByMessage(msg string) (bool, error)
		FindByID(id uint) (*model.Struct, error)
		FindByMessage(msg string) (*model.Struct, error)
		Add(str *model.Struct) error
		Save(str *model.Struct) error
		Delete(str *model.Struct) error
	}

	StructService struct {
		storage StructStorage
		hasher  Hasher
	}
)

func NewStructService(structStorage StructStorage, hasher Hasher) *StructService {
	return &StructService{
		storage: structStorage,
		hasher:  hasher,
	}
}

func (ss *StructService) ChangeMessage(id uint, message string) error {
	str, err := ss.storage.FindByID(id)
	if err != nil {
		return err
	}

	str.Message = message
	return ss.storage.Save(str)
}

func (ss *StructService) Add(str *model.Struct) error {
	str.Message = ss.hasher.Hash(str.Message)
	return ss.storage.Add(str)
}

func (ss *StructService) FindByID(id uint) (*model.Struct, error) {
	str, err := ss.storage.FindByID(id)
	if err != nil {
		return nil, err
	}
	return str, nil
}
