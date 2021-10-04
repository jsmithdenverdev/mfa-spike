package cache

import (
	"context"
	"mfaspike"
	"time"

	"github.com/go-redis/redis/v8"
)

type CodeCache struct {
	client *redis.Client
}

func NewCodeCache() CodeCache {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	return CodeCache{
		client,
	}
}

func (c CodeCache) Read(contact string) (mfaspike.Code, error) {
	result := c.client.Get(context.Background(), contact)

	if err := result.Err(); err != nil {
		return mfaspike.Code{}, err
	}

	return mfaspike.Code{
		Contact: contact,
		Code:    result.Val(),
	}, nil
}

func (c CodeCache) Write(code *mfaspike.Code) error {
	// set the record with a 15 minute ttl - the record will automatically expire
	ttl := time.Duration(15) * time.Minute
	result := c.client.Set(context.Background(), code.Contact, code.Code, ttl)

	if err := result.Err(); err != nil {
		return err
	}

	return nil
}

func (c CodeCache) Delete(contact string) error {
	result := c.client.Del(context.Background(), contact)

	if err := result.Err(); err != nil {
		return err
	}

	return nil
}
