package defaulttypevalues

import "fmt"

func Register(registrar map[string]func()) map[string]func() {
	registrar["defaulttypevalues"] = Run
	return registrar
}

type Blah struct{}

func Run() {
	var blah uint
	fmt.Println(blah)

	var test int
	fmt.Println(test)

	var fakestring string
	fmt.Println(fakestring)

	var foo Blah
	foo = Blah{}
	fmt.Println(foo)

	var foo2 *Blah
	foo2 = &Blah{}
	fmt.Println(foo2)
}
