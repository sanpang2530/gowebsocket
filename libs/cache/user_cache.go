/**
 * Created by GoLand.
 * User: link1st
 * Date: 2019-07-25
 * Time: 17:28
 */

package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/link1st/gowebsocket/mod/user/user_model"

	"github.com/link1st/gowebsocket/libs/redis_lib"
	"github.com/redis/go-redis/v9"
)

const (
	userOnlinePrefix = "acc:user:online:" // 用户在线状态

	userOnlineCacheTime = 24 * 60 * 60
)

/*********************  查询用户是否在线  ************************/
func getUserOnlineKey(userKey string) (key string) {
	key = fmt.Sprintf("%s%s", userOnlinePrefix, userKey)
	return
}

func GetUserOnlineInfo(userKey string) (userOnline *user_model.UserOnline, err error) {
	redisClient := redis_lib.GetClient()

	key := getUserOnlineKey(userKey)

	data, err := redisClient.Get(context.Background(), key).Bytes()
	if err != nil {
		if err == redis.Nil {
			fmt.Println("GetUserOnlineInfo", userKey, err)
			return
		}
		fmt.Println("GetUserOnlineInfo", userKey, err)
		return
	}

	userOnline = &user_model.UserOnline{}
	err = json.Unmarshal(data, userOnline)
	if err != nil {
		fmt.Println("获取用户在线数据 json Unmarshal", userKey, err)
		return
	}

	fmt.Println("获取用户在线数据", userKey, "time", userOnline.LoginTime, userOnline.HeartbeatTime, "AccIp", userOnline.AccIp, userOnline.IsLogoff)
	return
}

// SetUserOnlineInfo 设置用户在线数据
func SetUserOnlineInfo(userKey string, userOnline *user_model.UserOnline) (err error) {

	redisClient := redis_lib.GetClient()
	key := getUserOnlineKey(userKey)

	valueByte, err := json.Marshal(userOnline)
	if err != nil {
		fmt.Println("设置用户在线数据 json Marshal", key, err)
		return
	}

	_, err = redisClient.Do(context.Background(), "setEx", key, userOnlineCacheTime, string(valueByte)).Result()
	if err != nil {
		fmt.Println("设置用户在线数据 ", key, err)
		return
	}

	return
}
