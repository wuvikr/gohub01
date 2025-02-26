package captcha

import (
	"errors"
	"time"

	"github.com/wuvikr/gohub01/pkg/config"
	"github.com/wuvikr/gohub01/pkg/redis"
)

type RedisStore struct {
	RedisClient *redis.RedisClient
	KeyPrefix   string
}

func (s *RedisStore) Set(key string, value string) error {
	expireTime := config.MustGet[int]("captcha.expire_time")

	if config.MustGet[string]("app.env") != "production" {
		expireTime = config.MustGet[int]("captcha.debug_expire_time")
	}

	if ok := s.RedisClient.Set(s.KeyPrefix+key, value, time.Duration(expireTime)*time.Minute); !ok {
		return errors.New("无法存储验证码答案")
	}
	return nil
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
