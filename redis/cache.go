package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

var (
	logger = logrus.New()
	client *redis.Client
	ctx    = context.Background()
)

func init() {
	client = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetLevel(logrus.InfoLevel)
}

// GetCache retrieves data from Redis
func GetCache(key string) (string, error) {
	val, err := client.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", nil // Cache miss
		}
		logger.WithField("key", key).Error("Error fetching from Redis: ", err)
		return "", err
	}
	logger.WithField("key", key).Info("Cache hit")
	return val, nil
}

// SetCache sets data in Redis
func SetCache(key string, value string, ttl time.Duration) error {
	err := client.Set(ctx, key, value, ttl).Err()
	if err != nil {
		logger.WithField("key", key).Error("Error setting Redis cache: ", err)
		return err
	}
	logger.WithField("key", key).Info("Cache set successfully")
	return nil
}
