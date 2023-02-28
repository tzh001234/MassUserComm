package userProcess

import (
	"MassUserComm/common/message"
	"encoding/json"
	"fmt"
	"net"
)

// ProcessOnlineUsersMessage  处理OnlineUsersMessage类型的消息
func ProcessOnlineUsersMessage(conn net.Conn, mes message.Message) (err error) {
	//将mes.Data转换成[]byte,并反序列化
	var onlineUsersMessage message.OnlineUsersMessage
	err = json.Unmarshal([]byte(mes.Data), &onlineUsersMessage)
	if err != nil {
		fmt.Println("json.Unmarshal err:", err)
		return
	}
	//准备返回 请求结果消息
	var resultMes message.Message
	resultMes.Type = message.OnlineUsersResultMessageType
	var onlineUsersResultMessage message.OnlineUsersResultMessage //将其序列化后放入resultMes.Data中
	//业务逻辑(给要返回的结构体赋值)
	onlineUsersResultMessage.UsersIdAndName = UsersIdAndName

	//将返回消息序列化后放入resultMes.Data中
	data, err := json.Marshal(onlineUsersResultMessage)
	if err != nil {
		fmt.Println("json.Marshal err", err)
		return
	}
	resultMes.Data = string(data)
	//发送resultMes
	err = ServerSend(resultMes, conn)
	if err != nil {
		fmt.Println("ServerSend err", err)
		return
	}
	return
}
