package model

import "fmt"

type Message string
type Greeter struct {
	Message Message
}

// Greetメソッド
func (g Greeter) Greet() Message {
	return g.Message
}

type Event struct {
	Greeter Greeter
}

func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}

// Messageを返すもの　依存関係なし
func NewMessage() Message {
	return Message("Hi There!")
}

// Greeterを返すもの。 依存関係としてMessageを持っている
func NewGreeter(m Message) Greeter {
	return Greeter{Message: m}
}

// Eventの生成。依存関係としてGreeterを持っている
func NewEvent(g Greeter) Event {
	return Event{Greeter: g}
}
