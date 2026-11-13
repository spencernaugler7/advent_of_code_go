package envpattern

import "fmt"

type TestParentStruct struct {
	*Env
}

func Init() {

	testEnv := Null()

	test := TestParentStruct{Env: testEnv}
	blah := test.Registrar

	fmt.Println(blah)
}
