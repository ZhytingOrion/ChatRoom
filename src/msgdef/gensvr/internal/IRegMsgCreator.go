package internal

// Code generated by gen. Do not edit!
// 代码由 gen 生成。不要手工编辑！

import (
	"zeus/net/server"
)

type _IRegMsgCreator interface {
	RegMsgCreator(msgID server.MsgID, msgCreator func() server.IMsg)
}
