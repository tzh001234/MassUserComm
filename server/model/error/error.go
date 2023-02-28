// Package model 自定义错误
package model

import "errors"

var (
	ERROR_USER_NOTEXISTS = errors.New("用户不存在")
	ERROR_USER_EXISTS    = errors.New("用户已存在")
	ERROR_USER_PWD       = errors.New("密码不正确")

	ERROR_SERVER          = errors.New("服务端错误")
	ERROR_TYPE_NOTEXISTS  = errors.New("请求类型不存在")
	ERROR_USER_NOT_ONLINE = errors.New("用户未登录")
)
