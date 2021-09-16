package storage

import (
	"mfaspike/internal/domain"
	"time"

	"gorm.io/gorm"
)

type userEntity struct {
	gorm.Model
	Contact  string `gorm:"primaryKey"`
	Name     string
	Timezone string
}

type UserStore struct {
	Client *gorm.DB
}

func (store *UserStore) Migrate() error {
	return store.Client.AutoMigrate(&userEntity{})
}

func (store *UserStore) Read(contact string) (domain.User, error) {
	entity := userEntity{}

	tx := store.Client.First(&entity, contact)

	if tx.Error != nil {
		return domain.User{}, tx.Error
	}

	tz, _ := time.LoadLocation(entity.Timezone)

	return domain.User{
		Contact:  entity.Contact,
		Name:     entity.Name,
		Timezone: *tz,
	}, nil
}

func (store *UserStore) Write(user *domain.User) error {
	tx := store.Client.Create(&userEntity{
		Contact:  user.Contact,
		Name:     user.Name,
		Timezone: user.Timezone.String(),
	})

	return tx.Error
}

func (store *UserStore) Delete(contact string) error {
	tx := store.Client.Delete(&userEntity{
		Contact: contact,
	})

	return tx.Error
}
