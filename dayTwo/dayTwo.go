package dayTwo

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func Register(registrar map[string]func(io.Writer)) map[string]func(io.Writer) {
	registrar["d2p1"] = SolvePart1
	registrar["d2p2"] = SolvePart2
	return registrar
}

func SolvePart1(writer io.Writer) {
	bytes, err := os.ReadFile("./dayTwo/dayTwoInput.txt")
	if err != nil {
		fmt.Fprintln(writer, err.Error())
		os.Exit(1)
	}

	solution := 0
	line := string(bytes)
	line = strings.Trim(line, "\r\n")

	for numRange := range strings.SplitSeq(line, ",") {

		bounds := strings.Split(numRange, "-")
		lowerBound, _ := strconv.Atoi(bounds[0])
		upperBound, _ := strconv.Atoi(bounds[1])
		currentRange := CreateRange(lowerBound, upperBound, 1)

		for _, id := range currentRange {
			if IsInvalid(id) {
				newId, _ := strconv.Atoi(id)
				solution += newId
			}
		}
	}

	fmt.Fprintln(writer, solution)
}

func SolvePart2(writer io.Writer) {
	bytes, err := os.ReadFile("./dayTwo/dayTwoInput.txt")
	if err != nil {
		fmt.Fprintln(writer, err.Error())
		os.Exit(1)
	}

	solution := 0
	line := string(bytes)
	line = strings.Trim(line, "\r\n")

	for numRange := range strings.SplitSeq(line, ",") {

		bounds := strings.Split(numRange, "-")
		lowerBound, _ := strconv.Atoi(bounds[0])
		upperBound, _ := strconv.Atoi(bounds[1])
		currentRange := CreateRange(lowerBound, upperBound, 1)

		for _, id := range currentRange {
			if IsInvalidP2(id) {
				newId, _ := strconv.Atoi(id)
				solution += newId
			}
		}
	}
}

func IsInvalid(str string) bool {
	if (len(str) % 2) == 0 {
		midPoint := len(str) / 2
		return str[0:midPoint] == str[midPoint:]
	}
	return false
}

func IsInvalidP2(str string) bool {
	if (len(str) % 2) == 0 {
	}
	return false
}

func CreateRange(start, end, step int) []string {
	s := make([]string, 0, (1+(end-start))/step)
	for start <= end {
		s = append(s, strconv.Itoa(start))
		start += step
	}
	return s
}
