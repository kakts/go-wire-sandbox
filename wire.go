//go:build wireinject
// +build wireinject

package main // wire.go

import (
	"context"

	"github.com/google/wire"
	"github.com/kakts/go-wire-sandbox/src/field"
	"github.com/kakts/go-wire-sandbox/src/foobar"
	"github.com/kakts/go-wire-sandbox/src/foobarbaz"
	"github.com/kakts/go-wire-sandbox/src/model"
	"github.com/kakts/go-wire-sandbox/src/value"
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

// model.MyFooerの生成
// TODO
type test struct {
	f model.Fooer
}

func InitializeBar() string {
	wire.Build(model.MyFooerSet)
	// 戻り値の値は型があっていればなんでも良い
	return ""
}

func InitializeFooBar() foobar.FooBar {
	wire.Build(foobar.Set)
	return foobar.FooBar{}
}

func InitializeValueFoo() value.Foo {
	wire.Build(wire.Value(value.Foo{X: 42}))
	return value.Foo{}
}

// Foo.SのインジェクションのプロバイダーにGetSを使う場合
func InitializeFieldFoo1() string {
	wire.Build(
		field.ProvideFoo,
		field.GetS,
	)
	return ""
}

// Foo.SのインジェクションのプロバイダーにGetSを使わず、wire.FieldsOfを使う場合
func InitializeFieldFoo2() string {
	wire.Build(
		field.ProvideFoo,
		wire.FieldsOf(new(field.Foo), "S"),
	)
	return ""
}
