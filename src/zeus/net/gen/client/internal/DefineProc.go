package internal

/*
替换字符
	${SERVICE_NAME}
*/
const kTemplateMsgProc = `
package proc

// Code generated by gen.
// 本文件是对应 ${SERVICE_NAME} 的处理器实现文件。
// 框架代码由 gen 生成，具体实现需要手工填写。
// 再次生成时会合并原有实现。

import (
	${S2C_MSG_IMPORTS}
	"zeus/net/client"
)

// Proc_${SERVICE_NAME} 是消息处理类(Processor).
// 必须实现 MsgProc_*() 接口。
type Proc_${SERVICE_NAME} struct {
	sess client.ISession // 一般都需要包含session对象

	// 可能还应该包含用户和房间对象
	// user *User
	// room *Room
}`