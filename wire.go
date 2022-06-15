//go:build wireinject
// +build wireinject

package main // wire.go

import (
	"github.com/google/wire"
	"github.com/kakts/go-wire-sandbox/src/model"
)

// Eventの生成
func InitilizeEvent() (model.Event, error) {
	wire.Build(model.NewEvent, model.NewGreeter, model.NewMessage)
	return model.Event{}, nil
}
