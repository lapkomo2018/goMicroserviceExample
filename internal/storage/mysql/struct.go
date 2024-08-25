package mysql

import (
	"github.com/lapkomo2018/goServiceExample/pkg/model"
	"gorm.io/gorm"
)

type StructStorage struct {
	db *gorm.DB
}

func NewStructStorage(db *gorm.DB) *StructStorage {
	return &StructStorage{
		db: db,
	}
}

func (ss *StructStorage) ExistsByMessage(msg string) (bool, error) {
	var count int64
	err := ss.db.Model(&model.Struct{}).Where("message = ?", msg).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (ss *StructStorage) FindByID(id uint) (*model.Struct, error) {
	str := &model.Struct{}
	err := ss.db.Where("id = ?", id).First(str).Error
	if err != nil {
		return nil, err
	}
	return str, nil
}

func (ss *StructStorage) FindByMessage(msg string) (*model.Struct, error) {
	str := &model.Struct{}
	err := ss.db.Where("message = ?", msg).First(str).Error
	if err != nil {
		return nil, err
	}
	return str, nil
}

func (ss *StructStorage) Add(str *model.Struct) error {
	return ss.db.Create(str).Error
}

func (ss *StructStorage) Save(str *model.Struct) error {
	return ss.db.Save(str).Error
}

func (ss *StructStorage) Delete(str *model.Struct) error {
	return ss.db.Delete(str).Error
}
