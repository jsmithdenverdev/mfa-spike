package storage

import (
	"errors"
	"mfaspike/internal/domain"

	"gorm.io/gorm"
)

type mfaEntity struct {
	gorm.Model
	contact string `gorm:"primaryKey"`
	code    string
}

type MfaStore struct {
	client *gorm.DB
}

func NewMfaStore(client *gorm.DB) MfaStore {
	return MfaStore{
		client,
	}
}

func (store *MfaStore) Migrate() error {
	return store.client.AutoMigrate(&mfaEntity{})
}

func (store *MfaStore) Read(contact string) (domain.MfaCode, error) {
	entity := mfaEntity{}

	tx := store.client.First(&entity, contact)

	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return domain.MfaCode{}, domain.ErrNoCode
		}

		return domain.MfaCode{}, tx.Error
	}

	return domain.MfaCode{
		Contact: entity.contact,
		Code:    entity.code,
	}, nil
}

func (store *MfaStore) Write(code *domain.MfaCode) error {
	tx := store.client.Create(&mfaEntity{
		contact: code.Contact,
		code:    code.Code,
	})

	return tx.Error
}

func (store *MfaStore) Delete(contact string) error {
	tx := store.client.Delete(&mfaEntity{
		contact: contact,
	})

	return tx.Error
}
