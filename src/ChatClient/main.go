package main

import (
	"fmt"
	"msgdef/genclt"
	"proto/chat"
)

func main() {
	session, err := genclt.Dial("tcp", "127.0.0.1:5678")
	if err != nil {
		panic(err)
	}

	proc := session.Proc_chat
	proc.Done = make(chan bool)

	var data string

	session.Start()

	for {

		fmt.Println("Your message: ")
		fmt.Scanln(&data)

		if data == "quit" {
			break
		}

		req := &chat.ChatRequest{
			Data: data,
		}
		session.Send(req)

	}
	<-proc.Done
}
