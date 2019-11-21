package foobarbaz

type Foo struct {
	X int
}

func ProvideFoo() Foo {
	return Foo{X: 42}
}
