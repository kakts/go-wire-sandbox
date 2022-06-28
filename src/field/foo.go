package field

// https://github.com/google/wire/blob/main/docs/guide.md#use-fields-of-a-struct-as-providers

type Foo struct {
	S string
	N int
	F float64
}

// wire.FieldsOfを使う場合はこの関数が不要になる
func GetS(foo Foo) string {
	return foo.S
}

func ProvideFoo() Foo {
	return Foo{S: "Hello, World!", N: 1, F: 3.14}
}
