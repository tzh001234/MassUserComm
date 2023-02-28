package list

import (
	"MassUserComm/client/process/serverProcess"
	"MassUserComm/client/process/userProcess"
	"MassUserComm/common/message"
	"fmt"
	"net"
	"os"
)

// LoginList 登录目录
func LoginList() {

	var userId int
	var userPwd string
	//登录信息
	fmt.Println("登录中。。")

	fmt.Println("请输入id：")
	fmt.Scanf("%v\n", &userId)

	fmt.Println("请输入密码：")
	fmt.Scanf("%v\n", &userPwd)
	fmt.Printf("您输入的userId：%d  userPwd:%v\n", userId, userPwd)
	conn, err := userProcess.Login(userId, userPwd)
	if err != nil {
		fmt.Println("登录失败：", err)
		return
	}
	defer conn.Close()

	fmt.Println("登录成功")
	L := make(chan message.Message, 10)
	go serverProcess.ProcessServerMessage(conn, L)

	//登录成功后的页面
	for loop {
		var key int
		fmt.Println("----恭喜登录成功------")
		fmt.Println("\t1、显示用户列表")
		fmt.Println("\t2、发送消息")
		fmt.Println("\t3、信息列表")
		fmt.Println("\t4、退出登录")
		fmt.Println("\t5、退出系统")
		fmt.Println("\t请选择（1-4）：")
		fmt.Scanf("%d\n", &key)
		//分配下级菜单
		switch key {
		case 1:
			OnlineUsers(conn, L)
		case 2:
			fmt.Println("发送消息函数")

			loop = false //会退出此菜单（循环），返回上一级菜单
		case 3:
			fmt.Println("信息列表函数")
			loop = false
		case 4:
			fmt.Println("退出登录。。")
			conn.Close()
			loop = false
		case 5:
			fmt.Println("退出系统。。")
			os.Exit(0)
		default:
			fmt.Println("你的输入有误")
		}

	}
	//退出登录之后让首页菜单继续循环打印
	loop = true
}

func OnlineUsers(conn net.Conn, L chan message.Message) {
	err := userProcess.OnlineUsers(conn, L)
	if err != nil {
		fmt.Println("获取在线用户列表失败：", err)
	}
}
