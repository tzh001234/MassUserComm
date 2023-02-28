package list

import (
	"MassUserComm/client/process/userProcess"
	"fmt"
)

// RegisterList 2.注册目录
func RegisterList() {

	var userId int
	var userPwd string
	var UserName string

	//注册信息
	fmt.Println("注册中。。")

	fmt.Println("请输入id：")
	fmt.Scanf("%v\n", &userId)

	fmt.Println("请输入密码：")
	fmt.Scanf("%v\n", &userPwd)

	fmt.Println("请输入用户名：")
	fmt.Scanf("%v\n", &UserName)

	fmt.Printf("你输入的：%d, %s, %s", userId, userPwd, UserName)

	conn, err := userProcess.Register(userId, userPwd, UserName)
	if err != nil {
		fmt.Println("注册失败：", err)
		return
	}
	defer conn.Close()
	fmt.Println("注册成功")
}
