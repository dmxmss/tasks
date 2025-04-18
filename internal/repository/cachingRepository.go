package repository

import (
	e "github.com/dmxmss/tasks/error"
	"github.com/redis/go-redis/v9"

	"context"
	"time"
	"encoding/json"
	"log"
)

type CachingRepository interface {
	GetCached(string) ([]byte, error)
	SetCached(string, any, time.Duration) error
}

type cachingRepository struct {
	ctx context.Context
	client *redis.Client 
}

func NewCachingRepository(ctx context.Context, client *redis.Client) CachingRepository {
	return &cachingRepository{
		ctx: ctx,
		client: client,
	}
}

func (c *cachingRepository) GetCached(key string) ([]byte, error) {
	val, err := c.client.Get(c.ctx, key).Bytes()
	if err != nil {
		log.Printf("%s", err)
		return nil, e.ErrRedisKeyNotFound
	}

	return val, nil
}

func (c *cachingRepository) SetCached(key string, val any, exp time.Duration) error {
	valSerialized, err := json.Marshal(val)
	if err != nil {
		return e.ErrJSONError
	}

	if err := c.client.Set(c.ctx, key, valSerialized, exp).Err(); err != nil {
		return e.ErrRedisSettingValue
	}

	return nil
}
