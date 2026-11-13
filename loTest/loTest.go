package lotest

import (
	"fmt"

	"github.com/samber/lo"
)

func Register(registrar map[string]func()) map[string]func() {
	registrar["fragments"] = Run
	return registrar
}

func Run() {
	numbers := lo.Range(40)
	squared := lo.Map(numbers, func(blah, _ int) int { return blah * blah })
	sum := lo.Sum(squared)
	fmt.Printf("%d", sum)
}
