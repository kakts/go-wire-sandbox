//go:build wireinject
// +build wireinject

package main // wire.go

import (
	"context"

	"github.com/google/wire"
	"github.com/kakts/go-wire-sandbox/src/foobarbaz"
	"github.com/kakts/go-wire-sandbox/src/model"
)

// Eventの生成
func InitializeEvent(phrase string) (model.Event, error) {

	wire.Build(model.NewEvent, model.NewGreeter, model.NewMessage)
	return model.Event{}, nil
}

// Bazの生成
func InitializeBaz(ctx context.Context) (foobarbaz.Baz, error) {
	wire.Build(foobarbaz.FooBarBazSet)
	return foobarbaz.Baz{}, nil
}
