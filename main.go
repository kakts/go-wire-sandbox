package main

import (
	"context"
	"fmt"
	"os"
)

type Bar string

func main() {
	// 依存関係を手動で注入
	// message := model.NewMessage()
	// greeter := model.NewGreeter(message)
	// e := model.NewEvent(greeter)

	// wireのInjectorによる依存性注入
	e, err := InitializeEvent("test phrase")
	if err != nil {
		fmt.Printf("failed to create event: %s\n", err)
		os.Exit(2)
	}

	e.Start()

	baz, err := InitializeBaz(context.Background())
	if err != nil {
		fmt.Printf("failed to create Baz: %s\n", err)
		os.Exit(2)
	}
	fmt.Println("----")
	fmt.Println(baz.X)

	bar := InitializeBar()
	fmt.Println(bar)

	foobar := InitializeFooBar()
	fmt.Println(foobar.MyBar)
	fmt.Println(foobar.MyFoo)

	valueFoo := InitializeValueFoo()
	fmt.Println(valueFoo)

	fieldFoo1 := InitializeFieldFoo1()
	fmt.Println(fieldFoo1)

	fieldFoo2 := InitializeFieldFoo2()
	fmt.Println(fieldFoo2)
}
