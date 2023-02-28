package userProcess

import (
	"MassUserComm/common/message"
	"MassUserComm/common/utils"
	"encoding/json"
	"fmt"
	"net"
)

// ClientSend 将mes发送给服务器
func ClientSend(mes message.Message, conn net.Conn, L chan message.Message) (err error) {

	//至此，mes赋值完成，接下来将其序列化
	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err:", err)
		return
	}
	//发送数据
	err = utils.WritePKG(conn, data)
	if err != nil {
		return
	}

	//等待读取请求结果(利用管道阻塞)
	mes = <-L

	//识别请求结果
	err = ClientProcessMes(mes)

	return
}

// ClientProcessMes 识别并处理服务器发来的消息
func ClientProcessMes(mes message.Message) (err error) {
	switch mes.Type {
	case message.LoginResultMessageType:
		err = ProcessLoginResultMessage(mes)
		if err != nil {
			return
		}
	case message.RegisterResultMessageType:
		err = ProcessRegisterResultMessage(mes)
		if err != nil {
			return
		}
	case message.OnlineUsersResultMessageType:
		err = ProcessOnlineUsersResultMessage(mes)
		if err != nil {
			return
		}
	default:
		fmt.Println("未能识别服务器发送的消息类型")
	}

	return
}
