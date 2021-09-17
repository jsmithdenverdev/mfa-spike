package storage

import (
	"mfaspike/internal/domain"
	"time"

	"gorm.io/gorm"
)

type userEntity struct {
	gorm.Model
	contact  string `gorm:"primaryKey"`
	name     string
	timezone string
}

type UserStore struct {
	client *gorm.DB
}

func NewUserStore(client *gorm.DB) UserStore {
	return UserStore{
		client,
	}
}

func (store *UserStore) Migrate() error {
	return store.client.AutoMigrate(&userEntity{})
}

func (store *UserStore) Read(contact string) (domain.User, error) {
	entity := userEntity{}

	tx := store.client.First(&entity, contact)

	if tx.Error != nil {
		return domain.User{}, tx.Error
	}

	tz, _ := time.LoadLocation(entity.timezone)

	return domain.User{
		Contact:  entity.contact,
		Name:     entity.name,
		Timezone: *tz,
	}, nil
}

func (store *UserStore) Write(user *domain.User) error {
	tx := store.client.Create(&userEntity{
		contact:  user.Contact,
		name:     user.Name,
		timezone: user.Timezone.String(),
	})

	return tx.Error
}

func (store *UserStore) Delete(contact string) error {
	tx := store.client.Delete(&userEntity{
		contact: contact,
	})

	return tx.Error
}
