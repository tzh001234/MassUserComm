package userProcess

import (
	"MassUserComm/common/message"
	"encoding/json"
	"fmt"
	"net"
)

// OnlineUsers 获取在线用户列表
func OnlineUsers(conn net.Conn, L chan message.Message) (err error) {

	//定义消息实例
	var mes message.Message
	mes.Type = message.OnlineUsersMessageType

	//定义LoginMessage存放实际信息
	var onlineUsersMessage = message.OnlineUsersMessage{}
	//将loginMessage序列化,赋给mes
	data, err := json.Marshal(onlineUsersMessage)
	if err != nil {
		fmt.Println("json.Marshal err:", err)
		return
	}
	mes.Data = string(data)
	//至此，mes赋值完成，接下来将其发送
	err = ClientSend(mes, conn, L)
	return
}

func ProcessOnlineUsersResultMessage(mes message.Message) (err error) {
	//将mes.Data转换成[]byte,并反序列化
	var onlineUsersResultMessage message.OnlineUsersResultMessage
	err = json.Unmarshal([]byte(mes.Data), &onlineUsersResultMessage)
	if err != nil {
		fmt.Println("json.Unmarshal err:", err)
		return
	}
	fmt.Println(onlineUsersResultMessage.UsersIdAndName)
	return
}
