package cache

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"strings"

	"github.com/redis/go-redis/v9"
)

type Service struct{}

func NewService() *Service {
	GetRedisClient()
	return &Service{}
}

// func (s *Service) GetSession(ctx context.Context, key string) (*Session, error) {
// 	client := GetRedisClient()

// 	data, err := client.Get(ctx, key).Result()
// 	if err != nil {
// 		if err.Error() == "redis: nil" {
// 			return nil, fmt.Errorf("session not found or expired")
// 		}
// 		return nil, fmt.Errorf("failed to get from cache: %w", err)
// 	}

// 	var payload Session
// 	err = json.Unmarshal([]byte(data), &payload)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to unmarshal payload: %w", err)
// 	}

// 	return payload, nil
// }

func unlinkOrDel(ctx context.Context, client *redis.Client, keys []string) (int64, error) {
	if len(keys) == 0 {
		return 0, nil
	}
	pipe := client.Pipeline()
	unlink := pipe.Unlink(ctx, keys...)
	_, err := pipe.Exec(ctx)
	if err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "unknown command") ||
			strings.Contains(strings.ToLower(err.Error()), "unlink") {
			pipe2 := client.Pipeline()
			for _, k := range keys {
				pipe2.Del(ctx, k)
			}
			_, err2 := pipe2.Exec(ctx)
			if err2 != nil {
				return 0, err2
			}
			return int64(len(keys)), nil
		}
		return 0, err
	}
	return unlink.Val(), nil
}

func generateSessionKey() (string, error) {
	bytes := make([]byte, 16)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(bytes), nil
}
