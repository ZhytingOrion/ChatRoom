package room

import (
	"fmt"
	"proto/chat"
	"sync"
	"zeus/net/server"
)

type Room struct {
	mtx     sync.Mutex
	senders []server.ISession
}

type ISender interface {
	Send(msg server.IMsg)
}

var TheRoom *Room = &Room{}

func (r *Room) Add(s server.ISession) {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	r.senders = append(r.senders, s)
}

func (r *Room) Del(s server.ISession) {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	for i, sndr := range r.senders {
		if sndr != s {
			continue
		}
		fmt.Println("Deleted!")
		r.senders = append(r.senders[:i], r.senders[i+1:]...)
		return
	}
}

func (r *Room) Chat(content string) {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	r.broadcast(content)
}

func (r *Room) broadcast(content string) {
	fmt.Printf("broadcast: %s\n", content)
	msg := &chat.ChatMessage{
		Who:     "Unknown",
		Content: content,
	}
	for _, sndr := range r.senders {
		sndr.Send(msg)
	}
}
