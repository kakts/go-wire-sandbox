package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/google/wire"
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

	// mock関連のコード
	fmt.Printf("Real time greeting: %s [current time elided]\n", initApp().Greet()[0:15])

	// Approach A モックを手動で作成し、そのモックをインジェクターに渡す。
	// このアプローチは、事前にモックを準備する必要がある場合に有効です
	fmt.Println("Mock Approach A")
	mt := newMockTimer()
	// モック用のアプリ生成。引数に依存関係パッケージのモックを渡す
	mockedApp := initMockedAppFromArgs(mt)
	fmt.Println(mockedApp.Greet())

	mt.T = mt.T.AddDate(1999, 0, 0)
	fmt.Println(mockedApp.Greet())

	// Approach B インジェクターにモックを生成させ、依存関係をモックしたアプリの構造体を返す
	fmt.Println("Approach B")
	appWithMocks := initMockedApp()
	fmt.Println(appWithMocks.app.Greet())
	appWithMocks.mt.T = appWithMocks.mt.T.AddDate(1999, 0, 0)
	fmt.Println(appWithMocks.app.Greet())
}

// appSetは実際のアプリを作成するためのプロバイダーセットでsう
var appSet = wire.NewSet(
	wire.Struct(new(app), "*"),
	wire.Struct(new(greeter), "*"),
	wire.InterfaceValue(new(timer), realTime{}), // timerインタフェースを実装している構造体はrealTimeを扱う
)

// appSetWithoutMocksはモック化された依存関係を利用しアプリを作成するプロバイダーセットです。
// モック化された依存関係は除外され、インジェクターへの引数として提供されなければなりません。
var appSetWithoutMocks = wire.NewSet(
	wire.Struct(new(app), "*"),
	wire.Struct(new(greeter), "*"),
)

// mockAppSetはモック化されたアプリと、モック化された依存関係も含んで生成するためのプロバイダーセットです。
// これはアプローチBに利用されます。
var mockAppSet = wire.NewSet(
	wire.Struct(new(app), "*"),
	wire.Struct(new(greeter), "*"),
	wire.Struct(new(appWithMocks), "*"),
	// モック化された各依存パッケージにおいて、プロバイダーを追加し、関連するインターフェースに具体的な型をバインドするためにwire.Bindを使ってください
	newMockTimer,
	wire.Bind(new(timer), new(*mockTimer)),
)

type timer interface {
	Now() time.Time
}

// realTime implements timer with the real time
type realTime struct{}

func (realTime) Now() time.Time {
	return time.Now()
}

// mockTimer implements timer using a mocked time.
type mockTimer struct {
	T time.Time
}

// mockTimerのポインタを返すプロバイダ
func newMockTimer() *mockTimer {
	return &mockTimer{}
}

func (m *mockTimer) Now() time.Time {
	return m.T
}

// greeter issues greetings with the time provided by T.
type greeter struct {
	T timer
}

func (g greeter) Greet() string {
	return fmt.Sprintf("Good day! It is %v", g.T.Now())
}

type app struct {
	g greeter
}

func (a app) Greet() string {
	return a.g.Greet()
}

// appWithMocks はアプローチBのために使われます。これはアプリと、モックかされた依存関係を返します。
type appWithMocks struct {
	app app
	mt  *mockTimer
}
