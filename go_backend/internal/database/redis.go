package database

import (
    "context"
    "github.com/redis/go-redis/v9"
)

func NewRedisClient(addr, password string) *redis.Client {
    return redis.NewClient(&redis.Options{
        Addr:     addr,
        Password: password,
        DB:       0,
    })
}

func TestRedisConnection(client *redis.Client) error {
    ctx := context.Background()
    return client.Ping(ctx).Err()
}