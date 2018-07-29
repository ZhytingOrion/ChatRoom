package proc

// Code generated by gen.
// 本文件是对应 chat 的处理器实现文件。
// 框架代码由 gen 生成，具体实现需要手工填写。
// 再次生成时会合并原有实现。

import (
	"proto/chat"

	"zeus/net/client"

	"fmt"
)

// Proc_chat 是消息处理类(Processor).
// 必须实现 MsgProc_*() 接口。
type Proc_chat struct {
	sess client.ISession // 一般都需要包含session对象

	// 可能还应该包含用户和房间对象
	// user *User
	// room *Room

	Done chan bool
}

func (p *Proc_chat) MsgProc_ChatResponse(msg *chat.ChatResponse) {
	fmt.Println("Got chat response: ", msg)
	//close(p.Done)
}

func (p *Proc_chat) MsgProc_ChatMessage(msg *chat.ChatMessage) {
	fmt.Println("Got chat message: ", msg)
	//close(p.Done)
}