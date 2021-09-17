package storage

import (
	"mfaspike/internal/domain"
	"time"

	"gorm.io/gorm"
)

type userEntity struct {
	gorm.Model
	Id       string `gorm:"primaryKey"`
	Name     string
	Timezone string
}

type UserStore struct {
	client *gorm.DB
}

func NewUserStore(client *gorm.DB) (UserStore, error) {
	store := UserStore{
		client,
	}

	if err := store.migrate(); err != nil {
		return UserStore{}, err
	}

	return store, nil
}

func (store *UserStore) migrate() error {
	return store.client.AutoMigrate(&userEntity{})
}

func (store *UserStore) Read(contact string) (domain.User, error) {
	entity := userEntity{}

	err := store.client.First(&entity, contact).Error

	if err != nil {
		return domain.User{}, err
	}

	tz, _ := time.LoadLocation(entity.Timezone)

	return domain.User{
		Contact:  entity.Id,
		Name:     entity.Name,
		Timezone: *tz,
	}, nil
}

func (store *UserStore) Write(user *domain.User) error {
	err := store.client.Create(&userEntity{
		Id:       user.Contact,
		Name:     user.Name,
		Timezone: user.Timezone.String(),
	}).Error

	return err
}

func (store *UserStore) Delete(contact string) error {
	err := store.client.Delete(&userEntity{
		Id: contact,
	}).Error

	return err
}
