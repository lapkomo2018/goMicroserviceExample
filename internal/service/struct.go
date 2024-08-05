package service

import (
	"github.com/lapkomo2018/goServiceExample/internal/core"
)

type (
	Hasher interface {
		Hash(password string) string
		Verify(hash string, password string) bool
	}

	StructStorage interface {
		ExistsByMessage(msg string) (bool, error)
		FindByID(id uint) (*core.Struct, error)
		FindByMessage(msg string) (*core.Struct, error)
		Add(str *core.Struct) error
		Save(str *core.Struct) error
		Delete(str *core.Struct) error
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

func (ss *StructService) Add(str *core.Struct) error {
	str.Message = ss.hasher.Hash(str.Message)
	return ss.storage.Add(str)
}

func (ss *StructService) FindByID(id uint) (*core.Struct, error) {
	str, err := ss.storage.FindByID(id)
	if err != nil {
		return nil, err
	}
	return str, nil
}
