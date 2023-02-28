package userProcess

import (
	"MassUserComm/common/message"
	"MassUserComm/server/model/error"
	"MassUserComm/server/model/userDao"
	"encoding/json"
	"fmt"
	"net"
)

// ProcessRegisterMessage 处理RegisterMessage类型的消息
func ProcessRegisterMessage(conn net.Conn, mes message.Message) (err error) {
	//将mes.Data转换成[]byte,并反序列化
	var registerMessage message.RegisterMessage
	err = json.Unmarshal([]byte(mes.Data), &registerMessage)
	if err != nil {
		fmt.Println("json.Unmarshal err:", err)
		return
	}
	//准备返回 请求结果消息
	var resultMes message.Message
	resultMes.Type = message.RegisterResultMessageType
	var registerResultMessage message.RegisterResultMessage //将其序列化后放入resultMes.Data中
	//业务逻辑(给要返回的结构体赋值)
	//判断是否注册成功
	err = userDao.UDao.AddUser(registerMessage.User)
	if err != nil {
		if err == model.ERROR_SERVER {
			registerResultMessage.Code = 404
		} else if err == model.ERROR_USER_EXISTS {
			registerResultMessage.Code = 500
		}
		registerResultMessage.Error = error.Error(model.ERROR_SERVER)
	} else {
		registerResultMessage.Code = 200
	}
	//将返回消息序列化后放入resultMes.Data中
	data, err := json.Marshal(registerResultMessage)
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
