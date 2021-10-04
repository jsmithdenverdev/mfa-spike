package storage

import (
	"errors"
	"mfaspike"

	"gorm.io/gorm"
)

type mfaEntity struct {
	gorm.Model
	Id   string `gorm:"primaryKey"`
	Code string
}

type MfaStore struct {
	client *gorm.DB
}

func NewMfaStore(client *gorm.DB) (MfaStore, error) {
	store := MfaStore{
		client,
	}

	if err := store.migrate(); err != nil {
		return MfaStore{}, err
	}

	return store, nil
}

func (store *MfaStore) migrate() error {
	return store.client.AutoMigrate(&mfaEntity{})
}

func (store *MfaStore) Read(contact string) (mfaspike.Code, error) {
	entity := mfaEntity{}

	err := store.client.First(&entity, contact).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return mfaspike.Code{}, mfaspike.ErrNoCode
		}

		return mfaspike.Code{}, err
	}

	return mfaspike.Code{
		Contact: entity.Id,
		Code:    entity.Code,
	}, nil
}

func (store *MfaStore) Write(code *mfaspike.Code) error {
	err := store.client.Create(&mfaEntity{
		Id:   code.Contact,
		Code: code.Code,
	}).Error

	return err
}

func (store *MfaStore) Delete(contact string) error {
	err := store.client.Delete(&mfaEntity{
		Id: contact,
	}).Error

	return err
}
