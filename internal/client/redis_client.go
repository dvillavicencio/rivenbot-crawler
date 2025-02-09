package client

import (
	"context"
	"encoding/json"

	"github.com/redis/go-redis/v9"
	"rivenbot/internal/dto"
)

type RedisClient interface {
	GetManifestEntity(ctx context.Context, hash string) (*dto.ManifestObject, error)
}

type RedisService struct {
	Client *redis.Client
}

func NewRedisService(client *redis.Client) *RedisService {
	return &RedisService{Client: client}
}

// Returns a given manifest entity based on a hash
func (r *RedisService) GetManifestEntity(ctx context.Context, hash string) (*dto.ManifestObject, error) {
	result, err := r.Client.Get(ctx, hash).Result()
	if err != nil {
		return nil, err
	}

	var response *dto.ManifestObject
	if err := json.Unmarshal([]byte(result), response); err != nil {
		return nil, err
	}

	return response, nil
}
