package redisApi

import (
	"fmt"
	"github.com/go-redis/redis"
)

var ClientRedis *redis.Client

const (
	REDIS_NETWORK = "tcp"
	REDIS_HOST = ""
	REDIS_PORT = ""
	REDIS_PASSWORD = ""
	REDIS_DB = 0
)

func init() {
	options := redis.Options{
		Network: REDIS_NETWORK,
		Addr: fmt.Sprintf("%s:%s", REDIS_HOST, REDIS_PORT),
		Dialer: nil,
		OnConnect: nil,
		Password: REDIS_PASSWORD,
		DB: REDIS_DB,
		MaxRetries: 0,
		MinRetryBackoff: 0,
		MaxRetryBackoff: 0,
		DialTimeout: 0,
		ReadTimeout: 0,
		WriteTimeout: 0,
		PoolSize: 0,
		MinIdleConns: 0,
		MaxConnAge: 0,
		PoolTimeout: 0,
		IdleTimeout: 0,
		IdleCheckFrequency: 0,
		TLSConfig: nil,
		//Limiter: nil,
	}
	ClientRedis = redis.NewClient(&options)
	
}

func String(){
	ClientRedis.Set("golang_redis", "golang", 0)
}