package multipleinterfaces

import "fmt"

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

func Init() {

	var blah Blah1 = Bash { widget: "Bish" }

	

	result := 0

	switch test := blah.(type) {
	case Blah1:
		result = test.Foo(1) + 2
	case Blah2:
		result = test.Foo(1) - 2
	}

	fmt.Printf("%d\n", result)
}