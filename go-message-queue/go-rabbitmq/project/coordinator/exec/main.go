package main

import (
	"fmt"

	"github.com/dindasigma/go-rabbitmq/project/coordinator"
)

func main() {
	ql := coordinator.NewQueueListener()
	go ql.ListenForNewSource()

	var a string
	fmt.Scanln(&a)
}
