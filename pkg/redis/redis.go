package redis

import (
	"context"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/wuvikr/gohub01/pkg/logger"
	"go.uber.org/zap"
)

type RedisClient struct {
	Client  *redis.Client
	Context context.Context
}

var (
	Redis *RedisClient
	once  sync.Once
)

func ConnectRedis(address, username, password string, db int) {
	once.Do(func() {
		Redis = &RedisClient{
			Client: redis.NewClient(&redis.Options{
				Addr:     address,
				Username: username,
				Password: password,
				DB:       db,
			}),
			Context: context.Background(),
		}

		err := Redis.Client.Ping(Redis.Context).Err()
		if err != nil {
			logger.Error("redis connect error", zap.Error(err))
		}
	})
}

// Set 存储 key 对应的 value，且设置过期时间
func (rds *RedisClient) Set(key string, value interface{}, expiration time.Duration) bool {
	if err := rds.Client.Set(rds.Context, key, value, expiration).Err(); err != nil {
		logger.Error("Redis Set failed", zap.Error(err))
		return false
	}
	return true
}

// Get 获取 key 对应的 value
func (rds *RedisClient) Get(key string) string {
	result, err := rds.Client.Get(rds.Context, key).Result()
	if err != nil {
		if err != redis.Nil {
			logger.Error("Redis Get failed", zap.Error(err))
		}
		return ""
	}
	return result
}

// Has 判断一个 key 是否存在
func (rds *RedisClient) Has(key string) bool {
	result, err := rds.Client.Exists(rds.Context, key).Result()
	if err != nil {
		logger.Error("Redis Exists failed", zap.Error(err))
		return false
	}
	return result > 0
}

// Del 删除存储在 redis 里的数据
func (rds *RedisClient) Del(key string) bool {
	if err := rds.Client.Del(rds.Context, key).Err(); err != nil {
		logger.Error("Redis Del failed", zap.Error(err))
		return false
	}
	return true
}

// FlushDB 清空当前 redis db
func (rds *RedisClient) FlushDB() bool {
	if err := rds.Client.FlushDB(rds.Context).Err(); err != nil {
		logger.Error("Redis FlushDB failed", zap.Error(err))
		return false
	}
	return true
}

// Increment 当前 key 的值加一，并返回加一后的值
func (rds *RedisClient) Increment(key string) int64 {
	result, err := rds.Client.Incr(rds.Context, key).Result()
	if err != nil {
		logger.Error("Redis Increment failed", zap.Error(err))
		return 0
	}
	return result
}

// IncrementBy 当前 key 的值加上 value，并返回加上 value 后的值
func (rds *RedisClient) IncrementBy(key string, value int64) int64 {
	result, err := rds.Client.IncrBy(rds.Context, key, value).Result()
	if err != nil {
		logger.Error("Redis IncrementBy failed", zap.Error(err))
		return 0
	}
	return result
}

// Decrement 当前 key 的值减一，并返回减一后的值
func (rds *RedisClient) Decrement(key string) int64 {
	result, err := rds.Client.Decr(rds.Context, key).Result()
	if err != nil {
		logger.Error("Redis Decrement failed", zap.Error(err))
		return 0
	}
	return result
}

// DecrementBy 当前 key 的值减去 value，并返回减去 value 后的值
func (rds *RedisClient) DecrementBy(key string, value int64) int64 {
	result, err := rds.Client.DecrBy(rds.Context, key, value).Result()
	if err != nil {
		logger.Error("Redis DecrementBy failed", zap.Error(err))
		return 0
	}
	return result
}
