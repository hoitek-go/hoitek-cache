package drivers

import (
	"context"
	"errors"
	redisV8 "github.com/go-redis/redis/v8"
	redisV9 "github.com/go-redis/redis/v9"
	"github.com/hoitek-go/hoitek-cache/config"
	"github.com/hoitek-go/hoitek-cache/versions"
	"time"
)

type RedisVersionType interface {
	*redisV8.Client | *redisV9.Client | any
}

type Redis struct{}

var redisConf *config.Redis
var redisClient RedisVersionType

func initRedisConf() {
	if redisConf == nil {
		globalConfig := config.Global.(config.Redis)
		redisConf = &globalConfig
	}
}

func GetRedisClient() RedisVersionType {
	if redisClient == nil {
		ctx = context.Background()
		if redisConf.Version == versions.REDIS_V6 {
			redisClient = redisV8.NewClient(&redisV8.Options{
				Addr:     redisConf.GetAddress(),
				Password: redisConf.GetPassword(),
				DB:       redisConf.GetDatabase(),
			})
		}
		if redisConf.Version == versions.REDIS_V7 {
			redisClient = redisV9.NewClient(&redisV9.Options{
				Addr:     redisConf.GetAddress(),
				Password: redisConf.GetPassword(),
				DB:       redisConf.GetDatabase(),
			})
		}
	}
	return redisClient
}

func (r Redis) Set(key string, value interface{}) error {
	initConf()
	if redisConf.Version == versions.REDIS_V6 {
		redisClient := GetRedisClient().(*redisV8.Client)
		return redisClient.Set(ctx, key, value, 0).Err()
	}
	if redisConf.Version == versions.REDIS_V7 {
		redisClient := GetRedisClient().(*redisV9.Client)
		return redisClient.Set(ctx, key, value, 0).Err()
	}
	return errors.New("version not found")
}

func (r Redis) Get(key string) (interface{}, error) {
	initConf()
	if redisConf.Version == versions.REDIS_V6 {
		redisClient := GetRedisClient().(*redisV8.Client)
		cmd := redisClient.Get(ctx, key)
		if err := cmd.Err(); err != nil {
			return nil, err
		}
		return cmd.Val(), nil
	}
	if redisConf.Version == versions.REDIS_V7 {
		redisClient := GetRedisClient().(*redisV9.Client)
		cmd := redisClient.Get(ctx, key)
		if err := cmd.Err(); err != nil {
			return nil, err
		}
		return cmd.Val(), nil
	}
	return nil, errors.New("version not found")
}

func (r Redis) Delete(key string) error {
	initConf()
	if redisConf.Version == versions.REDIS_V6 {
		redisClient := GetRedisClient().(*redisV8.Client)
		return redisClient.Del(ctx, key).Err()
	}
	if redisConf.Version == versions.REDIS_V7 {
		redisClient := GetRedisClient().(*redisV9.Client)
		return redisClient.Del(ctx, key).Err()
	}
	return errors.New("version not found")
}

func (r Redis) Exists(key string) (bool, error) {
	initConf()
	if redisConf.Version == versions.REDIS_V6 {
		redisClient := GetRedisClient().(*redisV8.Client)
		cmd := redisClient.Exists(ctx, key)
		if err := cmd.Err(); err != nil {
			return false, err
		}
		return cmd.Val() > 0, nil
	}
	if redisConf.Version == versions.REDIS_V7 {
		redisClient := GetRedisClient().(*redisV9.Client)
		cmd := redisClient.Exists(ctx, key)
		if err := cmd.Err(); err != nil {
			return false, err
		}
		return cmd.Val() > 0, nil
	}
	return false, errors.New("version not found")
}

func (r Redis) Expire(key string, seconds int) error {
	initConf()
	if redisConf.Version == versions.REDIS_V6 {
		redisClient := GetRedisClient().(*redisV8.Client)
		return redisClient.Expire(ctx, key, time.Duration(seconds)).Err()
	}
	if redisConf.Version == versions.REDIS_V7 {
		redisClient := GetRedisClient().(*redisV9.Client)
		return redisClient.Expire(ctx, key, time.Duration(seconds)).Err()
	}
	return errors.New("version not found")
}

func (r Redis) TTL(key string) (int, error) {
	initConf()
	if redisConf.Version == versions.REDIS_V6 {
		redisClient := GetRedisClient().(*redisV8.Client)
		cmd := redisClient.TTL(ctx, key)
		if err := cmd.Err(); err != nil {
			return 0, err
		}
		return int(cmd.Val().Seconds()), nil
	}
	if redisConf.Version == versions.REDIS_V7 {
		redisClient := GetRedisClient().(*redisV9.Client)
		cmd := redisClient.TTL(ctx, key)
		if err := cmd.Err(); err != nil {
			return 0, err
		}
		return int(cmd.Val().Seconds()), nil
	}
	return 0, errors.New("version not found")
}

func (r Redis) Flush() error {
	initConf()
	if redisConf.Version == versions.REDIS_V6 {
		redisClient := GetRedisClient().(*redisV8.Client)
		return redisClient.FlushAll(ctx).Err()
	}
	if redisConf.Version == versions.REDIS_V7 {
		redisClient := GetRedisClient().(*redisV9.Client)
		return redisClient.FlushAll(ctx).Err()
	}
	return errors.New("version not found")
}

func (r Redis) Close() error {
	initConf()
	if redisConf.Version == versions.REDIS_V6 {
		redisClient := GetRedisClient().(*redisV8.Client)
		return redisClient.Close()
	}
	if redisConf.Version == versions.REDIS_V7 {
		redisClient := GetRedisClient().(*redisV9.Client)
		return redisClient.Close()
	}
	return errors.New("version not found")
}

func (r Redis) Ping() error {
	initConf()
	if redisConf.Version == versions.REDIS_V6 {
		redisClient := GetRedisClient().(*redisV8.Client)
		return redisClient.Ping(ctx).Err()
	}
	if redisConf.Version == versions.REDIS_V7 {
		redisClient := GetRedisClient().(*redisV9.Client)
		return redisClient.Ping(ctx).Err()
	}
	return errors.New("version not found")
}
