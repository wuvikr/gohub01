package verifycode

import (
	"time"

	"github.com/wuvikr/gohub01/pkg/config"
	"github.com/wuvikr/gohub01/pkg/redis"
)

type RedisStore struct {
	RedisClient *redis.RedisClient
	KeyPrefix   string
}

func (s *RedisStore) Set(key string, value string) bool {
	key = s.KeyPrefix + key
	expireTime := time.Duration(config.MustGet[int]("verifycode.expire_time")) * time.Minute
	if config.MustGet[string]("app.env") != "production" {
		expireTime = time.Duration(config.MustGet[int]("verifycode.debug_expire_time")) * time.Minute
	}

	return s.RedisClient.Set(key, value, expireTime)
}

func (s *RedisStore) Get(key string, clear bool) string {
	key = s.KeyPrefix + key
	val := s.RedisClient.Get(key)
	if clear {
		s.RedisClient.Del(key)
	}
	return val
}

func (s *RedisStore) Verify(key, answer string, clear bool) bool {
	v := s.Get(key, clear)
	return v == answer
}
