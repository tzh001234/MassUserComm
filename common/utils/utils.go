// Package utils 常用的工具、结构体等（包的读写功能）
package utils

import (
	"MassUserComm/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

// ReadPKG 读取消息（包）
func ReadPKG(conn net.Conn) (mes message.Message, err error) {
	buffer := make([]byte, 8096) //存放接收到的数据
	//	fmt.Println("等待读取消息。。")
	_, err = conn.Read(buffer[0:4])
	if err != nil {
		return
	}
	//根据buffer[0:4]转成一个uint32类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buffer[0:4])
	fmt.Println("读到的buffer转成的uint32=", pkgLen)
	n, err := conn.Read(buffer[:pkgLen]) //会覆盖之前的长度信息，目前总长度=pkgLen,全为有效信息（结构体信息）
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Read err:", err)
		return
	}
	//将buffer[:pkgLen]的值反序列化 （读取到的实际请求数据，不包含长度信息）
	err = json.Unmarshal(buffer[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarshal err:", err)
		return
	}

	return
}

// WritePKG 发送消息（包）
func WritePKG(conn net.Conn, data []byte) (err error) {
	//用[]byte表示长度
	DataLen := uint32(len(data))
	BytesDataLen := make([]byte, 4)
	binary.BigEndian.PutUint32(BytesDataLen, DataLen)
	//先发送长度（[]byte格式）
	_, err = conn.Write(BytesDataLen)
	if err != nil {
		fmt.Println("conn.Write发送失败")
		return
	}
	//发送数据
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write err:", err)
		return
	}
	fmt.Printf("发送长度为%d的data成功:%s\n", DataLen, string(data))

	return
}
