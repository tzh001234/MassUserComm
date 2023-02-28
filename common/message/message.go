//消息实例
package message

import "MassUserComm/common/user"

const (
	LoginMessageType             = "loginMessage"       //客户端发送的登录消息（类型）
	LoginResultMessageType       = "loginResultMessage" //服务器发送的登录返回消息（类型）
	RegisterMessageType          = "registerMessage"
	RegisterResultMessageType    = "registerResultMessage"
	OnlineUsersMessageType       = "onlineUsersMessage"
	OnlineUsersResultMessageType = "onlineUsersResultMessage"
)

// Message 实际发送的结构体，包含要发送的另一个结构体的类型和序列化后的结果
type Message struct {
	Type string `json:"type"` //消息类型
	Data string `json:"data"` //消息体本身（序列化之后的某个结构体）
}

// LoginMessage 登录请求消息 （客户端发送）
type LoginMessage struct {
	UserId   int    `json:"userId"`
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"`
}

// LoginResultMessage 登录返回结果消息 （服务端返回）
type LoginResultMessage struct {
	Code  int    `json:"code"`  //状态码  500表示未注册  200表示登陆成功   404
	Error string `json:"error"` //返回错误信息
}

// RegisterMessage 注册请求消息
type RegisterMessage struct {
	User user.User `json:"user"`
}

// RegisterResultMessage 注册返回结果消息
type RegisterResultMessage struct {
	Code  int    `json:"code"`  //状态码
	Error string `json:"error"` //返回错误信息
}

// OnlineUsersMessage 查询在线用户表请求消息
type OnlineUsersMessage struct {
}

// OnlineUsersResultMessage 查询在线用户表返回结果消息
type OnlineUsersResultMessage struct {
	UsersIdAndName []string `json:"usersIdAndName"`
}
