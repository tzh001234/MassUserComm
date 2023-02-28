package userProcess

import (
	"MassUserComm/common/message"
	"MassUserComm/common/utils"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
)

//Login 登录函数
func Login(userId int, userPwd string) (conn net.Conn, err error) {
	//连接到服务器
	conn, err = net.Dial("tcp", "127.0.0.1:8889")
	if err != nil {
		fmt.Println("net.Dial连接到服务器失败")
		return
	}

	//定义消息实例
	var mes message.Message
	mes.Type = message.LoginMessageType

	//定义LoginMessage存放实际信息
	var loginMessage = message.LoginMessage{
		UserId:  userId,
		UserPwd: userPwd,
	}
	//将loginMessage序列化,赋给mes
	data, err := json.Marshal(loginMessage)
	if err != nil {
		fmt.Println("json.Marshal err:", err)
		return
	}
	mes.Data = string(data)
	//至此，mes赋值完成，接下来将其发送

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err:", err)
		return
	}
	//发送数据
	err = utils.WritePKG(conn, data)

	//等待读取请求结果
	mes, err = utils.ReadPKG(conn)
	if err != nil {
		if err == io.EOF {
			fmt.Println("服务端退出，我客户端也退出")
			return
		}
		fmt.Println("readPKG err", err)
		return
	}

	//识别请求结果
	err = ClientProcessMes(mes)

	return
}

func ProcessLoginResultMessage(mes message.Message) (err error) {
	//将mes.Data转换成[]byte,并反序列化
	var loginResultMessage message.LoginResultMessage
	err = json.Unmarshal([]byte(mes.Data), &loginResultMessage)
	if err != nil {
		fmt.Println("json.Unmarshal err:", err)
		return
	}
	fmt.Println("状态码：", loginResultMessage.Code)
	if loginResultMessage.Code != 200 {
		err = errors.New(loginResultMessage.Error)
	}
	return
}
