package main

import (
	"fmt"
	"os"
)

func main() {
	// 依存関係を手動で注入
	// message := model.NewMessage()
	// greeter := model.NewGreeter(message)
	// e := model.NewEvent(greeter)

	// wireのInjectorによる依存性注入
	e, err := InitializeEvent()
	if err != nil {
		fmt.Printf("failed to create event: %s\n", err)
		os.Exit(2)
	}

	e.Start()
}
