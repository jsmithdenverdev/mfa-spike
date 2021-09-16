package storage

import (
	"errors"
	"mfaspike/internal/domain"

	"gorm.io/gorm"
)

type mfaEntity struct {
	gorm.Model
	Contact string `gorm:"primaryKey"`
	Code    string
}

type MfaStore struct {
	Client *gorm.DB
}

func (store *MfaStore) Migrate() error {
	return store.Client.AutoMigrate(&mfaEntity{})
}

func (store *MfaStore) Read(contact string) (domain.MfaCode, error) {
	entity := mfaEntity{}

	tx := store.Client.First(&entity, contact)

	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return domain.MfaCode{}, domain.ErrNoCode
		}

		return domain.MfaCode{}, tx.Error
	}

	return domain.MfaCode{
		Contact: entity.Contact,
		Code:    entity.Code,
	}, nil
}

func (store *MfaStore) Write(code *domain.MfaCode) error {
	tx := store.Client.Create(&mfaEntity{
		Contact: code.Contact,
		Code:    code.Code,
	})

	return tx.Error
}

func (store *MfaStore) Delete(contact string) error {
	tx := store.Client.Delete(&mfaEntity{
		Contact: contact,
	})

	return tx.Error
}
