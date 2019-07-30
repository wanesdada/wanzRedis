package store

import (
	"fmt"
	"math/rand"
	"redis/cache"
	"redis/utils"
	"time"

	"github.com/garyburd/redigo/redis"
	log "github.com/sirupsen/logrus"
)

var (
	Rc *cache.Cache
	Re error
)

var RetryMinTimeInterval int64 = 5

// Lock请求锁的最大时间间隔(毫秒)
var RetryMaxTimeInterval int64 = 30

//缓存key
//TODO: fix typo
const (

	REDIS_URL  			= ""



	CACHE_KEY_TOKEN                  = "_CACHE_KEY_TOKEN_"                    //token 缓存
)

var Rm Redis

type Redis struct{}

func init() {
	Rc, Re = cache.NewCache(REDIS_URL)
	if Re != nil {
		log.Fatal(Re)
	}
}

//TODO: overwrite cache

func Get(merchant, key string) interface{} {
	return Rc.Get(merchant + key)
}

func RedisBytes(merchant, key string) (data []byte, err error) {
	return Rc.RedisBytes(merchant + key)
}

func RedisString(merchant, key string) (data string, err error) {
	return Rc.RedisString(merchant + key)
}

func RedisInt(merchant, key string) (data int, err error) {
	return Rc.RedisInt(merchant + key)
}

func RedisFloat64(merchant, key string) (data float64, err error) {
	data, err = redis.Float64(Rc.Get(merchant+key), err)
	return
}

func Put(merchant, key string, val interface{}, timeout time.Duration) error {
	return Rc.Put(merchant+key, val, timeout)
}

func SetNX(merchant, key string, val interface{}, timeout time.Duration) bool {
	return Rc.SetNX(merchant+key, val, timeout)
}

func Delete(merchant, key string) error {
	return Rc.Delete(merchant + key)
}

func IsExist(merchant, key string) bool {
	return Rc.IsExist(merchant + key)
}

func LPush(merchant, key string, val interface{}) error {
	return Rc.LPush(merchant+key, val)
}

func Brpop(merchant, key string, callback func([]byte)) {
	Rc.Brpop(merchant+key, callback)
}

func GetRedisTTL(merchant, key string) time.Duration {
	return Rc.GetRedisTTL(merchant + key)
}

func Incrby(merchant, key string, num int) (interface{}, error) {
	return Rc.Incrby(merchant+key, num)
}

func Lock(merchant, key string, timeout time.Duration) bool {
	for true {
		result, err := Rc.Do("SET", merchant+key, 1, "NX", "EX", int64(timeout/time.Second))
		if err == nil && result == "OK" {
			return true
		}
		if err != nil {
			if err.Error() != "redigo: nil returned" {
				fmt.Println("Locker Lock Error", err)
				break
			}
		}
		sleepTimeInterval()
	}
	return false
}

func sleepTimeInterval() {
	var unixNano = time.Now().UnixNano()
	var r = rand.New(rand.NewSource(unixNano))
	var randValue = RetryMinTimeInterval + r.Int63n(RetryMaxTimeInterval-RetryMinTimeInterval)
	time.Sleep(time.Duration(randValue) * time.Millisecond)
}

func CheckPwd5Time(merchant, key string) int {
	key = merchant + key
	count := 5
	if Rc.IsExist(key) {
		count, _ = Rc.RedisInt(key)
	}
	count--
	Rc.Put(key, count, utils.GetTodayLastSecond())
	return count
}

func CheckPwd4Time(merchant, key string) int {
	key = merchant + key
	count := 2
	if Rc.IsExist(key) {
		count, _ = Rc.RedisInt(key)
	}
	count--
	Rc.Put(key, count, utils.GetTodayLastSecond())
	return count
}

func CheckProduct10Time(merchant, key string) int {
	key = merchant + key
	count := 5
	if Rc.IsExist(key) {
		count, _ = Rc.RedisInt(key)
	}
	count--
	Rc.Put(key, count, 10*time.Minute)
	return count
}
