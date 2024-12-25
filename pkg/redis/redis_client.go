package redis

import (
	"context"
	"messageOK/internal/entity"

	"github.com/go-redis/redis/v8"
)

type RedisClient interface {
	Set(key string, value interface{}) error
	GetSentMessages() ([]entity.SentMessage, error)
}

type redisClient struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedisClient(addr string) RedisClient {
	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	return &redisClient{
		client: rdb,
		ctx:    context.Background(),
	}
}

func (rc *redisClient) Set(key string, value interface{}) error {
	return rc.client.Set(rc.ctx, key, value, 0).Err()
}

func (rc *redisClient) GetSentMessages() ([]entity.SentMessage, error) {
	keys, err := rc.client.Keys(rc.ctx, "*").Result()
	if err != nil {
		return nil, err
	}

	var messages []entity.SentMessage
	for _, key := range keys {
		val, err := rc.client.Get(rc.ctx, key).Result()
		if err != nil {
			continue
		}

		messages = append(messages, entity.SentMessage{
			MessageID: key,
			SentTime:  val,
		})
	}

	return messages, nil
}
