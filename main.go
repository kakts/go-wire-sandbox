package main

import "github.com/kakts/go-wire-sandbox/src/model"

func main() {
	// 依存関係を手動で注入
	message := model.NewMessage()
	greeter := model.NewGreeter(message)
	event := model.NewEvent(greeter)

	event.Start()
}
