package foobar

import "github.com/google/wire"

type Foo int
type Bar int

func ProvideFoo() Foo {
	return Foo(1)
}

func ProvideBar() Bar {
	return Bar(5)
}

type FooBar struct {
	MyFoo Foo
	MyBar Bar
}

// FooBarのMyFoo, MyBarフィールドの値をProvideFoo,ProvideBarの値に代入するように指定する。
// wire.Struct(new(FooBar), "*")でも同じ
var Set = wire.NewSet(
	ProvideFoo,
	ProvideBar,
	wire.Struct(new(FooBar), "MyFoo", "MyBar"))
