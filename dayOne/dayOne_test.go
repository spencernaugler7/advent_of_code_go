package dayone

import (
	"strings"
	"testing"
)

func TestMod(t *testing.T) {

}

func TestSampleInput(t *testing.T) {
	var input = `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`

	hitsZero := FindPass(strings.Lines(input))

	if hitsZero != 3 {
		t.Errorf("actual was %d answer should be 3", hitsZero)
	}
}

func TestSampleInputPartTwo(t *testing.T) {
	var input = `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`

	hitsZero := FindPassPartTwo(strings.Lines(input))

	if hitsZero != 6 {
		t.Errorf("actual was %d answer should be 6", hitsZero)
	}
}
