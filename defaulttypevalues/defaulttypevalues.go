package defaulttypevalues

import "fmt"

func Register(registrar map[string]func()) map[string]func() {
	registrar["defaulttypevalues"] = Run
	return registrar
}

func Run() {
	var blah uint
	fmt.Println(blah)

	var test int
	fmt.Println(test)

	var fakestring string
	fmt.Println(fakestring)
}
