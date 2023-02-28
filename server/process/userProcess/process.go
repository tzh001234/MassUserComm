package userProcess

//总处理器

import (
	"MassUserComm/common/message"
	"MassUserComm/common/utils"
	model "MassUserComm/server/model/error"
	"encoding/json"
	"fmt"
	"io"
	"net"
)

type UserProcess struct {
	Conn     net.Conn
	UserId   int
	UserName string
}

// Process 处理和某个客户端的通讯(此处维护单个登录用户的连接)
func Process(conn net.Conn) {
	defer conn.Close()

	for {
		//读取请求
		mes, err := utils.ReadPKG(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Println("收到EOF,连接断开")
				return
			} else {
				fmt.Println("此用户连接断开")
				return
			}
		}
		//识别并处理请求
		err = ServerProcessMes(conn, mes)
		if err != nil {
			fmt.Println("serverProcessMes err", err)
		}

	}
}

// ServerProcessMes 识别请求并处理
func ServerProcessMes(conn net.Conn, mes message.Message) (err error) {

	switch mes.Type { //判断接收的请求类型
	case message.LoginMessageType: //此处处理登录
		err = ProcessLoginMessage(conn, mes)
	case message.RegisterMessageType: //此处处理注册
		err = ProcessRegisterMessage(conn, mes)
	case message.OnlineUsersMessageType: //此处处理查询在线用户
		err = ProcessOnlineUsersMessage(conn, mes)

	default:
		err = model.ERROR_TYPE_NOTEXISTS
	}
	return
}

func ServerSend(resultMes message.Message, conn net.Conn) (err error) {
	data, err := json.Marshal(resultMes)
	if err != nil {
		fmt.Println("json.Marshal err", err)
		return
	}
	//发送序列化后的结果（data）
	err = utils.WritePKG(conn, data)

	return
}
