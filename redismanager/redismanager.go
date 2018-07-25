package redismanager

//实现redis的分布式锁

import (
	"github.com/garyburd/redigo/redis"
	"log"
	"time"
	"strconv"
	"sync"
	"fmt"
)

type RedisManager struct {
	Conn redis.Conn
}

var (
	ConnType = "tcp"
	ConnIp	 = "localhost"
	ConnPort = "6380"
	//ConnPassWord = "123456"
)
var mutex sync.Mutex

const LockPrefix = "_Lock_Pre_fix_"

var redisMgrInstance = RedisManager{Conn:nil}

//获取redis实例
func GetInstance() *RedisManager {
	mutex.Lock()

	defer func() {
		mutex.Unlock()
	}()

	if redisMgrInstance.Conn == nil {
		var err error
		password := redis.DialPassword("123456")
		redisMgrInstance.Conn, err = redis.Dial(ConnType, ConnIp + ":" + ConnPort, password)
		if err != nil {
			panic(err)
		}

		//初始化
		redisMgrInstance.onInit()
	}

	return &redisMgrInstance
}

//获取redis实例
func (rMgr *RedisManager) onInit() bool {
	if redisMgrInstance.Conn != nil {
		log.Println("RedisManager onInit")

		//初始化
		replay, err := redis.Strings(redisMgrInstance.Conn.Do("keys", LockPrefix + "*"))
		if err != nil {
			panic(err)
		}

		if replay != nil {
			for _, v := range replay {
				if _, err := redis.Int(redisMgrInstance.Conn.Do("del", v)); err != nil {
					log.Println("RedisManager init, del lock key: " + v + " failed!")
					return false
				}
				log.Println("RedisManager init, del lock key: " + v + " success!")
			}
		}
	}

	return true
}

//根据key获取value
func (rMgr *RedisManager) Get(key string) interface{} {
	if rMgr == nil {
		return nil
	}

	replay, err := redis.String(rMgr.Conn.Do("get", key))
	if err != nil {
		log.Println(err)
	}

	log.Println(replay)

	return replay
}

//加锁
func (rMgr *RedisManager) DistributedLock(lockName string, expireTime uint32) (bool, error) {
	mutex.Lock()
	defer func() {
		mutex.Unlock()
	}()

	if rMgr == nil {
		return false, redis.Error("RedisManager is nil")
	}
	lockName = LockPrefix + lockName
	lockValue := lockName + strconv.FormatInt(time.Now().Unix(), 10)
	replay1, err := redis.Int(rMgr.Conn.Do("setnx", lockName, lockValue))
	if err != nil {
		return false, err
	}

	if replay1 == 1 {
		replay2, err := redis.Int(rMgr.Conn.Do("expire", lockName, expireTime))
		if err != nil {
			if _, err := redis.Int(rMgr.Conn.Do("del", lockName)); err != nil {
				return false, redis.Error("del keys error: " + err.Error())
			}
			return false, redis.Error("expire keys error: " + err.Error())
		}

		if replay2 == 1 {
			return true, nil
		}
		return false, redis.Error("expire return replay is:" + strconv.Itoa(replay2))
	} else {
		return false, nil
	}
}

//加锁
func (rMgr *RedisManager) UnDistributedLock(lockName string) (bool, error) {
	mutex.Lock()

	defer func() {
		mutex.Unlock()
	}()

	if rMgr == nil {
		return false, redis.Error("RedisManager is nil")
	}

	lockName = LockPrefix + lockName
	value, err := redis.String(rMgr.Conn.Do("get", lockName))
	if err != nil {
		return false, redis.Error("UnDistributedLock get key error: " + err.Error())
	}

	if value == "" {
		return false, redis.Error("UnDistributedLock has no LockName:" + lockName)
	}

	if _, err := redis.Int(rMgr.Conn.Do("del", lockName)); err != nil {
		return false, redis.Error("UnDistributedLock del key error: " + err.Error())
	}

	return true, nil
}


func TestNoDistributedLock() {

	var ticket = 300
	for i := 1; i <= 3; i++ {
		go func(id int) {
			for {
				if ticket <= 0 {
					break
				}
				ticket--
				log.Println(fmt.Sprintf("thread %d get ticket, left ticket is:%d", id, ticket))
			}
		}(i)
	}
}

func TestDistributedLock() {

	var ticket = 300
	for i := 1; i <= 3; i++ {
		go func(id int) {
			rMgr := GetInstance()
			for {
				if ok, _ := rMgr.DistributedLock("ticketLock", 10); ok {
					if ticket <= 0 {
						rMgr.UnDistributedLock("ticketLock")
						break
					}

					ticket--
					log.Println(fmt.Sprintf("thread %d get ticket, left ticket is:%d", id, ticket))
					rMgr.UnDistributedLock("ticketLock")
				}
			}
		}(i)
	}
}