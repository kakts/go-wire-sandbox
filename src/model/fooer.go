package model

import "github.com/google/wire"

// interfaceと具体型のバインディング用
// https://github.com/google/wire/blob/main/docs/guide.md#binding-interfaces

type Fooer interface {
	Foo() string
}

type MyFooer string

func (b *MyFooer) Foo() string {
	return string(*b)
}

// provideMyFooer 具体型のMyFooer(Fooerインタフェースを実装している)を返すプロバイダー
func ProvideMyFooer() *MyFooer {
	b := new(MyFooer)
	*b = "Hello, World!"
	return b
}

type Bar string

// provideBar Barを返す。 引数はFooerインタフェースを渡すが実際はMyFooer
func ProvideBar(f Fooer) string {
	// f will be a *MyFooer
	return f.Foo()
}

var MyFooerSet = wire.NewSet(
	ProvideMyFooer,
	wire.Bind(new(Fooer), new(*MyFooer)),
	ProvideBar)
