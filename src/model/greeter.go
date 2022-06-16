package model

import (
	"errors"
	"fmt"
	"time"
)

type Message string
type Greeter struct {
	Message Message
	Grumpy  bool
}

// Greetメソッド
func (g Greeter) Greet() Message {
	if g.Grumpy {
		return Message("Go away!")
	}
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
func NewMessage(phrasa string) Message {
	return Message(phrasa)
}

// Greeterを返すもの。 依存関係としてMessageを持っている
func NewGreeter(m Message) Greeter {
	var grumpy bool
	if time.Now().Unix()%2 == 0 {
		grumpy = true
	}
	return Greeter{Message: m, Grumpy: grumpy}
}

// Eventの生成。依存関係としてGreeterを持っている
func NewEvent(g Greeter) (Event, error) {
	if g.Grumpy {
		// return Event{}, errors.New("could not create event: event greeter is grumpy")
		return Event{}, errors.New("could not create event: event greeter is grumpy")
	}
	return Event{Greeter: g}, nil
}
