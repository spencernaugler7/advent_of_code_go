package multipleinterfaces

import (
	"fmt"
	"io"
)

// want to see what happens when I define two method that are exactly the same

type Blah1 interface {
	Foo(tisk int) int
}

type Blah2 interface {
	Foo(tisk int) int
}

type Bash struct {
	widget string
}

func (b Bash) Foo(bish int) int {
	return bish * 2
}

func Register(registrar map[string]func(io.Writer)) map[string]func(io.Writer) {
	registrar["multipleInterfaces"] = Init
	return registrar
}

func Init(writer io.Writer) {
	var blah Blah1 = Bash{widget: "Bish"}

	result := 0

	switch test := blah.(type) {
	case Blah1:
		result = test.Foo(1) + 2
	case Blah2:
		result = test.Foo(1) - 2
	}

	fmt.Fprintf(writer, "%d\n", result)
}
