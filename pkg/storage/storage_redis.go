package storage

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type redisStorage struct {
	client    *redis.Client
	keyPrefix string
	expiry    time.Duration
}

func NewRedisStorage(host string, port int, password string, db int) *redisStorage {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", host, port),
		Password: password,
		DB:       db,
	})

	return &redisStorage{
		client:    redisClient,
		keyPrefix: "secret",
		expiry:    time.Hour * 24,
	}
}

func (s *redisStorage) getKey(uuid string) string {
	if len(s.keyPrefix) < 1 {
		return uuid
	}

	return fmt.Sprintf("%s.%s", s.keyPrefix, uuid)
}

func (s *redisStorage) Get(ctx context.Context, uuid string) (string, error) {
	val, err := s.client.Get(ctx, s.getKey(uuid)).Result()
	if err != nil {
		return "", err
	}

	return val, nil
}

func (s *redisStorage) Set(ctx context.Context, uuid string, content string) error {
	return s.client.Set(ctx, s.getKey(uuid), content, s.expiry).Err()
}

func (s *redisStorage) Delete(ctx context.Context, uuid string) error {
	return s.client.Del(ctx, s.getKey(uuid)).Err()
}
