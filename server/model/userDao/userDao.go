package userDao

//userDao 数据访问对象Dao   data access object
//对user对象的各种方法

import (
	"MassUserComm/common/user"
	"MassUserComm/server/model/error"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
)

// UserDao 用户数据访问对象
type UserDao struct {
	RedisDB *redis.Client
}

var UDao = UserDao{}

func (u *UserDao) Ping() (err error) {
	_, err = u.RedisDB.Ping().Result()
	return
}

// GetUserById 根据id查询用户是否存在
func (this *UserDao) GetUserById(id int) (user user.User, err error) {

	//err = this.rdb.Set("100", "{\"userId\":100,\"userPwd\":\"123456\",\"userName\":\"Tzh\"}", 0).Err()
	//if err != nil {
	//	fmt.Print(err)
	//	return
	//}
	////get方法获取string
	//userString, err := this.rdb.Get("100").Result()
	//fmt.Println(userString)

	//do方法获取string,传入参数组成redis语句： hget users 100
	userString, err := this.RedisDB.Do("hget", "users", id).Result()
	if err != nil {
		err = model.ERROR_USER_NOTEXISTS
		return
	}

	fmt.Println(userString)
	//将string反序列化放入user结构体中
	err = json.Unmarshal([]byte(userString.(string)), &user)
	if err != nil {
		fmt.Println("json.Unmarshal err", err)
		err = model.ERROR_SERVER
		return
	}
	return
}

// AddUser 添加用户（注册）
func (this UserDao) AddUser(user user.User) (err error) {
	//先判断用户id是否已存在
	_, err = this.GetUserById(user.UserId)
	if err == nil { //说明已存在
		err = model.ERROR_USER_EXISTS
		return
	}

	//将user序列化后写入redis
	jsonUser, err := json.Marshal(user)

	_, err = this.RedisDB.Do("hset", "users", user.UserId, string(jsonUser)).Result()
	if err != nil {
		err = model.ERROR_SERVER
		return
	}
	return
}
