// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/kakts/go-wire-sandbox/src/model"
)

// Injectors from wire.go:

// Eventの生成
func InitializeEvent(phrase string) (model.Event, error) {
	message := model.NewMessage(phrase)
	greeter := model.NewGreeter(message)
	event, err := model.NewEvent(greeter)
	if err != nil {
		return model.Event{}, err
	}
	return event, nil
}
