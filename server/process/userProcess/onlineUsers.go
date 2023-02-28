package userProcess

import (
	"MassUserComm/common/user"
	model "MassUserComm/server/model/error"
	"fmt"
	"net"
)

// OnlineUsersMap 在线用户表
var OnlineUsersMap map[int]*UserProcess

var UsersIdAndName []string

func init() {
	OnlineUsersMap = make(map[int]*UserProcess, 10)
	UsersIdAndName = make([]string, 10)
	UsersIdAndName[0] = "\tUserId\tUserName\n"
}

func AddOnlineUsers(conn net.Conn, u user.User) {
	var userProcess = UserProcess{
		Conn:     conn,
		UserName: u.UserName,
		UserId:   u.UserId,
	}
	OnlineUsersMap[u.UserId] = &userProcess
	fmt.Println(OnlineUsersMap)

	UsersIdAndName = append(UsersIdAndName, fmt.Sprintf("\n\tid:%v \tname:%v\t", u.UserId, u.UserName))

}

//删除
func DeleteOnlineUsers(userId int) {
	delete(OnlineUsersMap, userId)
}

//查询所有在线用户
func GetAllOnlineUsers() (s string) {

	for k, v := range OnlineUsersMap {

		s += fmt.Sprintf("id:%d,name:%s\n", k, v.UserName)

	}
	return
}

//根据id返回对应用户
func GetOnlineUserById(UserId int) (u *UserProcess, err error) {

	u, ok := OnlineUsersMap[UserId]
	if !ok {
		err = model.ERROR_USER_NOT_ONLINE
	}
	return
}
