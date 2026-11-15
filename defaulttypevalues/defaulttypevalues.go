package defaulttypevalues

import (
	"fmt"
	"io"
)

func Register(registrar map[string]func(io.Writer)) map[string]func(io.Writer) {
	registrar["defaulttypevalues"] = Run
	return registrar
}

type Blah struct{}

func Run(writer io.Writer) {
	var blah uint
	fmt.Fprintln(writer, blah)

	var test int
	fmt.Fprintln(writer, test)

	var fakestring string
	fmt.Fprintln(writer, fakestring)

	var foo Blah
	foo = Blah{}
	fmt.Fprintln(writer, foo)

	var foo2 *Blah
	foo2 = &Blah{}
	fmt.Fprintln(writer, foo2)
}
