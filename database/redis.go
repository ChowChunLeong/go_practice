package database

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var RedisDb *RedisClient

// RedisClient represents a Redis client object
type RedisClient struct {
	Client *redis.Client
	ctx    context.Context
}

// NewRedisClient creates a new RedisClient object
func NewRedisClient() *RedisClient {
	RedisDb = &RedisClient{
		Client: redis.NewClient(&redis.Options{
			Addr:     "localhost:6379", // Redis server address
			Password: "",               // No password set
			DB:       0,                // Use default DB
		}),
		ctx: context.Background(),
	}
	return RedisDb
}

// GetKeyValue retrieves the value for a given key from Redis
func (r *RedisClient) GetKeyValue(key string) (string, error) {
	return r.Client.Get(r.ctx, key).Result()
}

// SetKeyValue sets a key-value pair in Redis
func (r *RedisClient) SetKeyValue(key, value string) error {
	return r.Client.Set(r.ctx, key, value, 0).Err()
}

// PushToList pushes elements onto a Redis list
func (r *RedisClient) PushToList(key string, values ...interface{}) error {
	return r.Client.LPush(r.ctx, key, values...).Err()
}

// PopFromList pops an element from a Redis list
func (r *RedisClient) PopFromList(key string) (string, error) {
	return r.Client.RPop(r.ctx, key).Result()
}

// GetListLength returns the length of a Redis list
func (r *RedisClient) GetListLength(key string) (int64, error) {
	return r.Client.LLen(r.ctx, key).Result()
}

// GetListRange returns elements from a Redis list within the specified range
func (r *RedisClient) GetListRange(key string, start, stop int64) ([]string, error) {
	return r.Client.LRange(r.ctx, key, start, stop).Result()
}

func SetUpRedisDatabaseConnection() {
	RedisDb = NewRedisClient()
}
