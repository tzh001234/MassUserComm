package userProcess

import (
	"MassUserComm/common/message"
	"MassUserComm/server/model/error"
	"MassUserComm/server/model/userDao"
	"encoding/json"
	"fmt"
	"net"
)

// ProcessLoginMessage 处理LoginMessage类型的消息
func ProcessLoginMessage(conn net.Conn, mes message.Message) (err error) {
	//将mes.Data转换成[]byte,并反序列化
	var loginMessage message.LoginMessage
	err = json.Unmarshal([]byte(mes.Data), &loginMessage)
	if err != nil {
		fmt.Println("json.Unmarshal err:", err)
		return
	}
	//准备返回 请求结果消息
	var resultMes message.Message
	resultMes.Type = message.LoginResultMessageType
	var loginResultMessage message.LoginResultMessage //将其序列化后放入resultMes.Data中
	//业务逻辑(给要返回的结构体赋值)
	//根据id拿到数据库中的user信息
	u, err := userDao.UDao.GetUserById(loginMessage.UserId)
	//判断请求是否可以登录成功
	if err != nil {
		if err == model.ERROR_USER_NOTEXISTS {
			loginResultMessage.Code = 500
			fmt.Println("查无此人")
			loginResultMessage.Error = error.Error(model.ERROR_USER_NOTEXISTS)
		} else {
			loginResultMessage.Code = 500
			loginResultMessage.Error = error.Error(model.ERROR_SERVER)
		}
	} else if u.UserPwd != loginMessage.UserPwd {
		fmt.Println("密码错误，登录失败！")
		loginResultMessage.Code = 500
		loginResultMessage.Error = error.Error(model.ERROR_USER_PWD)
	} else {
		fmt.Println("登录成功！")
		loginResultMessage.Code = 200
	}

	//将返回消息序列化后放入resultMes.Data中
	data, err := json.Marshal(loginResultMessage)
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
	AddOnlineUsers(conn, u)
	return
}
