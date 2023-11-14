package config

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"strconv"
	"sync"
	"time"
)

var once sync.Once

// type global
type singleton struct {
	rdb *redis.Client
}

var (
	connect singleton
)

func GetRedis() *singleton {
	once.Do(func() {
		// 生成配置
		configKV := GetConfig("")
		dbHost, err := strconv.Atoi(configKV["CACHE_DATABASE"])
		if err != nil {
			panic(fmt.Sprintf("error by instance DB: %v", err))
		}
		params := RedisParams{
			Db:       dbHost,
			Port:     configKV["CACHE_PORT"],
			Host:     configKV["CACHE_HOST"],
			Password: configKV["CACHE_PASSWORD"],
		}

		// 生成对象
		connect = singleton{
			rdb: redis.NewClient(&redis.Options{
				Addr:     params.Host + ":" + params.Port,
				Password: params.Password, // no password set
				DB:       params.Db,       // use default DB
			}),
		}
	})

	return &connect
}

var ctx = context.Background()

type RedisParams struct {
	Host     string
	Port     string
	Db       int
	Password string
}

// Get 获取 Key 值
func (conn *singleton) Get(key string) string {
	val, err := conn.rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		return ""
	} else if err != nil {
		panic(err)
	}

	return val
}

// Set 设定 Key
func (conn *singleton) Set(key string, val string) string {
	err := conn.rdb.Set(ctx, key, val, 0).Err()
	if err != nil {
		panic(err)
	}

	return val
}

// Exists 检查 Key 存在
func (conn *singleton) Exists(key string) int {
	val, err := conn.rdb.Exists(ctx, key).Result()
	if err == redis.Nil {
		return 0
	} else if err != nil {
		panic(err)
	}

	return int(val)
}

// SetExpiration 设置 Key 和过期时间
func (conn *singleton) SetExpiration(key string, val string, expiration time.Duration) string {
	err := conn.rdb.Set(ctx, key, val, expiration).Err()
	if err != nil {
		panic(err)
	}

	return val
}

// Expire 重设 Key 过期时间
func (conn *singleton) Expire(key string, expiration time.Duration) string {
	err := conn.rdb.Expire(ctx, key, expiration).Err()
	if err != nil {
		panic(err)
	}

	return key
}

// Delete 移除 Key
func (conn *singleton) Delete(key string) bool {
	err := conn.rdb.Del(ctx, key).Err()
	if err != nil {
		panic(err)
	}
	return true
}

// ExpireTime 获取 Key 过期时间
func (conn *singleton) ExpireTime(key string) time.Duration {
	return conn.rdb.ExpireTime(ctx, key).Val()
}

// ZAdd 有序集合：新增
func (conn *singleton) ZAdd(key string, score float64, value string) bool {
	err := conn.rdb.ZAdd(ctx, key, redis.Z{
		Score:  score,
		Member: value,
	}).Err()
	if err != nil {
		panic(err)
	}
	return true
}

// ZRangeByScore 有序集合：范围查询
func (conn *singleton) ZRangeByScore(key string, scoreMin float64, scoreMax float64) []string {
	result := conn.rdb.ZRangeByScore(ctx, key, &redis.ZRangeBy{
		Min: strconv.FormatFloat(scoreMin, 'f', -1, 64),
		Max: strconv.FormatFloat(scoreMax, 'f', -1, 64),
	})

	if result.Err() != nil {
		panic(result.Err())
	}

	return result.Val()
}

// ZRemRangeByScore 有序集合：范围移除
func (conn *singleton) ZRemRangeByScore(key string, scoreMin float64, scoreMax float64) bool {
	err := conn.rdb.ZRemRangeByScore(ctx, key,
		strconv.FormatFloat(scoreMin, 'f', -1, 64),
		strconv.FormatFloat(scoreMax, 'f', -1, 64),
	).Err()
	if err != nil {
		panic(err)
	}

	return true
}

func (conn *singleton) FlushDB() {
	conn.rdb.FlushDB(ctx)
}
