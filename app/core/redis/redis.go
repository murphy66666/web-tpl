package redis

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"sync"
	"web-tpl/app/core/config"
)

var redisClientSyncMap map[string]*redis.Client
var redisClientMutex sync.RWMutex

func Load(flag string, conf *config.Model, key ...string) *redis.Client {
	var k = "default"
	if len(key) != 0 {
		k = key[0]
	}

	sgKey := flag + k
	redisClientMutex.RLocker()
	if c, ok := redisClientSyncMap[sgKey]; ok {
		redisClientMutex.RUnlock()
		return c
	}

	redisClientMutex.RUnlock()

	//2验证
	redisClientMutex.Lock()
	defer redisClientMutex.Unlock()
	if c, ok := redisClientSyncMap[sgKey]; ok {
		return c
	}

	rdsConf, ok := conf.Redis[k]
	if !ok {
		panic(fmt.Sprintf("redis %s config not exist,please check your yaml config!",
			key[0]))
	}

	c := initRedis(flag == "write", &rdsConf)
	redisClientSyncMap[sgKey] = c

	return c
}

// 验证redis 地址是否为空
func validParams(addr string) {
	if addr == "" {
		panic("redis can not allow empty addr")
	}
}

func initRedis(isWrite bool, conf *config.Redis) *redis.Client {
	var rdsConf config.RedisItem
	if isWrite {
		rdsConf = conf.Write
	} else {
		rdsConf = conf.Read
	}
	validParams(rdsConf.Addr)

	return redis.NewClient(&redis.Options{
		Addr:         rdsConf.Addr,
		PoolSize:     rdsConf.PoolSize,
		Password:     rdsConf.Password,
		IdleTimeout:  rdsConf.IdleTimeout,
		ReadTimeout:  rdsConf.ReadTimeout,
		WriteTimeout: rdsConf.WriteTimeout,
		MaxRetries:   rdsConf.Retries,
		MinIdleConns: rdsConf.MinIdleConns,
		DB:           rdsConf.DB,
	})
}
