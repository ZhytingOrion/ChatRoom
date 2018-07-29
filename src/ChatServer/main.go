package main

import (
	"msgdef/gensvr"
)

func main() {
	svr, err := gensvr.New("tcp", ":5678", 10000)

	if err != nil {
		panic(err)
	}

	svr.Run()
}