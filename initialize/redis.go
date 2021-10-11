package initialize

import (
	"context"
	"fmt"
	"github.com/Summrry-top/gin-glory-api/global"
	"github.com/go-redis/redis/v8"
	"time"
)

// 连接redis
func ConnectRedis() bool {
	client := redis.NewClient(&redis.Options{
		Addr:     global.ServerConfig.Redis.Addr,
		Password: global.ServerConfig.Redis.Password,
		DB:       global.ServerConfig.Redis.Db,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if _, err := client.Ping(ctx).Result(); err != nil {
		fmt.Println("连接redis数据库失败")
		return false
	}
	global.Redis = client
	return true
}

// func ConnectRedis() {
// 	// redis连接池
// 	global.RedisClient = &redis.Pool{
// 		Dial: func() (redis.Conn, error) {
// 			c, err := redis.Dial("tcp", "127.0.0.1:6379")
// 			if err != nil {
// 				return nil, err
// 			}
// 			return c, nil
// 		},
// 		//最大的空闲连接数，表示即使没有redis连接时依然可以保持N个空闲的连接，而不被清除，随时处于待命状态。
// 		MaxIdle: 1,
// 		//最大的激活连接数，表示同时最多有N个连接
// 		MaxActive: 10,
// 		//最大的空闲连接等待时间，超过此时间后，空闲连接将被关闭
// 		IdleTimeout: 180 * time.Second,
// 	}
// }
