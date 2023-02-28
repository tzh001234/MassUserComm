// Package list 目录
package list

import (
	"fmt"
)

//判断是否还要继续显示菜单
var loop = true

func List() {
	//一级菜单
	for loop {
		var key int
		//接收用户的选择
		fmt.Println("------欢迎------")
		fmt.Println("\t1、登录聊天室")
		fmt.Println("\t2、注册用户")
		fmt.Println("\t3、退出系统")
		fmt.Println("\t请选择（1-3）：")
		fmt.Scanf("%d\n", &key)
		//分配二级菜单
		switch key {
		case 1:
			LoginList()
		case 2:
			RegisterList()
		case 3:
			fmt.Println("正在退出")
			loop = false
		default:
			fmt.Println("你的输入有误,请重新输入")
		}

	}
}
