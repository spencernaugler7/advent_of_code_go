package daytwo

import (
	"fmt"
	"os"
)


func Solve() {
	bytes, err := os.ReadFile("./dayTwo/dayTwoInput.txt")

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}


}
