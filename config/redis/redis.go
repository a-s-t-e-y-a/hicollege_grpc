package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type RedisConnectionOptions struct {
	Address  string
	Password string
	DB       int
	Protocol int
}

func (redisOption RedisConnectionOptions) RedisConnection() *redis.Client {
	connection := redis.NewClient(&redis.Options{
		Addr:     redisOption.Address,
		Password: redisOption.Password,
		DB:       redisOption.DB,
		Protocol: redisOption.Protocol,
	})

	ctx := context.Background()
	err := connection.Set(ctx, "foo", "bar", 0).Err()
	if err != nil {
		panic(err)
	}
	return connection

}
