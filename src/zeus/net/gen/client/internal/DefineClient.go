package internal

const kTemplateClient = `
package genclt

// Code generated by gen. Do not edit!
// 代码由 gen 生成。不要手工编辑！

import (
	"${CURRENT_IMPORT_DIR}genclt/internal"
	"${CURRENT_IMPORT_DIR}genclt/proc"
	"zeus/net/client"
)

type IMsg = client.IMsg

// 封装 client.Session, 并添加MsgProc。
type Session struct {
	*client.Session
	
	${MSG_PROC_MEMBERS}
}

// Dial 创建一个连接.
// 使用 genclt.Dial() 创建连接，可在 Start() 之前更改消息处理器。
func Dial(protocal string, addr string) (*Session, error) {
	sess, err := client.Dial(protocal, addr)
	if err != nil {
		return nil, err
	}

	result := &Session{
		Session: sess,
		
		${MSG_PROC_INIT_MEMBERS}
	}
	
	return result, nil
}

func (s *Session) Start() {
	${REG_MSG_PROC_METHODS}
	s.Session.Start()
}
`
