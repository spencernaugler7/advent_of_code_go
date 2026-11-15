package dayOne

import (
	"fmt"
	"io"
	"iter"
	"math"
	"os"
	"strconv"
	"strings"
)

func Register(registrar map[string]func(io.Writer)) map[string]func(io.Writer) {
	registrar["d1p1"] = SolvePart1
	// registrar["d1p2"] = SolvePart2
	return registrar
}

// dial goes from 0 - 99
// dial starts at 50
// solution: number of times the dial is at exactly zero by the end of the sequence
func SolvePart1(writer io.Writer) {
	bytes, err := os.ReadFile("./dayOne/partOneInput.txt")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	result := 0

	result = FindPass(strings.Lines(string(bytes)))
	fmt.Fprintln(writer, result)
}

func FindPass(input iter.Seq[string]) int {
	var hitsZero int = 0
	var position int = 50

	for line := range input {
		dir := line[0:1]
		length, err := strconv.Atoi(strings.TrimSpace(line)[1:])
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		if dir == "R" {
			position = int(math.Mod(float64(position+length), 100))
		} else {
			turn := math.Mod(float64(position-length), 100)
			if turn < 0 {
				position = int(turn) + 100
			} else {
				position = int(turn)
			}
		}

		if position == 0 {
			hitsZero++
		}
	}

	return hitsZero
}

func FindPassPartTwo(input iter.Seq[string]) int {
	var hitsZero int = 0
	var position int = 50

	for line := range input {
		dir := line[0:1]
		length, err := strconv.Atoi(strings.TrimSpace(line)[1:])
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		if dir == "R" {
			turns := position + length
			rotations := int(math.Floor(float64(turns / 100)))
			position = int(math.Mod(float64(position+length), 100))

			hitsZero += rotations
		} else {
			rotations := int(math.Floor(math.Abs(float64(position-length)) / 100))

			turn := math.Mod(float64(position-length), 100)
			if turn < 0 {
				position = int(turn) + 100
			} else {
				position = int(turn)
			}
			hitsZero += rotations
		}

		if position == 0 {
			hitsZero++
		}
	}

	return hitsZero
}
