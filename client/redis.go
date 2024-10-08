package client

import (
	"github.com/Mohamadreza-shad/ucl-draw/config"
	"github.com/redis/go-redis/v9"
)

func NewRedisClient() (redis.UniversalClient, error) {
	redisURI := config.RedisURI()
	masterName := config.RedisMasterName()
	opts, err := redis.ParseURL(redisURI)
	if err != nil {
		return nil, err
	}
	return redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs: []string{
			opts.Addr,
		},
		DB:                    opts.DB,
		Username:              opts.Username,
		Password:              opts.Password,
		SentinelUsername:      opts.Username,
		SentinelPassword:      opts.Password,
		MaxRetries:            opts.MaxRetries,
		MinRetryBackoff:       opts.MinRetryBackoff,
		MaxRetryBackoff:       opts.MaxRetryBackoff,
		DialTimeout:           opts.DialTimeout,
		ReadTimeout:           opts.ReadTimeout,
		WriteTimeout:          opts.WriteTimeout,
		ContextTimeoutEnabled: opts.ContextTimeoutEnabled,
		PoolFIFO:              opts.PoolFIFO,
		PoolSize:              opts.PoolSize,
		PoolTimeout:           opts.PoolTimeout,
		MinIdleConns:          opts.MinIdleConns,
		MaxIdleConns:          opts.MaxIdleConns,
		ConnMaxIdleTime:       opts.ConnMaxIdleTime,
		ConnMaxLifetime:       opts.ConnMaxLifetime,
		MaxRedirects:          opts.MaxRetries,
		ReadOnly:              false,
		RouteByLatency:        false,
		RouteRandomly:         false,
		MasterName:            masterName,
	}), nil
}
