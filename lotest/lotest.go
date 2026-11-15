package lotest

import (
	"fmt"
	"io"

	"github.com/samber/lo"
)

func Register(registrar map[string]func(io.Writer)) map[string]func(io.Writer) {
	registrar["loTest"] = Run
	return registrar
}

func Run(writer io.Writer) {
	numbers := lo.Range(40)
	squared := lo.Map(numbers, func(blah, _ int) int { return blah * blah })
	sum := lo.Sum(squared)
	fmt.Fprintf(writer, "%d", sum)
}
