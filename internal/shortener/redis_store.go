package shortener

import (
	"context"
	"encoding/json"

	"github.com/redis/go-redis/v9"
)

type redisStore struct {
	client *redis.Client
}

func NewRedisStore(addr string, password string) *redisStore {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
	})

	return &redisStore{client: client}
}

func (s *redisStore) Save(ctx context.Context, redirect Redirect) error {
	redirectJson, err := json.Marshal(redirect)
	if err != nil {
		return err
	}

	status := s.client.Set(ctx, redirect.Alias, redirectJson, 0)

	if status.Err() != nil {
		return status.Err()
	}
	return nil
}

func (s *redisStore) Get(ctx context.Context, alias string) (Redirect, error) {
	value, err := s.client.Get(ctx, alias).Result()
	if err != nil {
		return Redirect{}, err
	}

	var redirect Redirect

	if err := json.Unmarshal([]byte(value), &redirect); err != nil {
		return Redirect{}, err
	}

	return redirect, nil
}

func (s *redisStore) GetAll(ctx context.Context) ([]Redirect, error) {
	return []Redirect{}, nil
}
