package daytwo

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Solve() {
	bytes, err := os.ReadFile("./dayTwo/dayTwoInput.txt")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	solution := 0

	line := string(bytes)
	ranges := strings.Split(line, ",")

	for _, numRange := range ranges {

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

	print(solution)
}

func IsInvalid(stringRep string) bool {
	for i := 0; i <= len(stringRep)-2; i++ {
		if stringRep[i] == stringRep[i+1] {
			return true
		}
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
