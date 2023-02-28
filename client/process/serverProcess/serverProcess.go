// Package serverProcess 与服务器保持连接
package serverProcess

import (
	"MassUserComm/common/message"
	"MassUserComm/common/utils"
	"fmt"
	"net"
)

// ProcessServerMessage 保持通讯，处理服务端的消息
func ProcessServerMessage(conn net.Conn, L chan message.Message) {

	for {
		//fmt.Println("客户端正在保持通讯，等待读取服务器发送的消息（忽略此消息）")
		//接收消息（到mes）
		mes, err := utils.ReadPKG(conn)
		if err != nil {
			fmt.Println("与服务器的连接已断开，系统退出")
			//os.Exit(0)
			return
		}
		//识别并分配服务器发来的消息
		L <- mes

	}

}
