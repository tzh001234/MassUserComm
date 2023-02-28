package main

import (
	"MassUserComm/server/model/userDao"
	"MassUserComm/server/process/userProcess"
	"fmt"
	"github.com/go-redis/redis"
	"net"
)

func main() {
	InitRedis()

	//开启服务器，等待客户端连接
	fmt.Println("服务器将在8889端口监听")
	listener, err := net.Listen("tcp", "127.0.0.1:8889")
	if err != nil {
		fmt.Println("net.Listen err=", err)
		return
	}
	fmt.Println("等待客户端连接中...")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listen.Accept err:", err)
		}
		//连接成功则启动一个协程和客户端保持连接
		go userProcess.Process(conn)
	}

}

func InitRedis() {
	userDao.UDao.RedisDB = redis.NewClient(&redis.Options{ //获取连接
		Addr:     "127.0.0.1:6379", //服务器  ip:port
		Password: "",               //密码
		DB:       0,                //选择数据库（0-15）
	})

	err := userDao.UDao.Ping()
	if err != nil {
		fmt.Println("Redis连接失败")
		return
	}
	fmt.Println("Redis连接成功")
}
