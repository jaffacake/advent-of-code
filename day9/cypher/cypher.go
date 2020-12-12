package cypher

import (
	"strconv"
	"strings"
)

func FindWeakness(data []byte) int {
	values := strings.Split(string(data), "\n")

	var transmission []int
	for _, value := range values {
		intValue, _ := strconv.Atoi(value)
		transmission = append(transmission, intValue)
	}

	// 5 for test
	// 25 for real run
	preambleLength := 25
	i := preambleLength

	invalidValue := 0

loop:
	for i < len(transmission) {
		for _, value1 := range transmission[MaxOf(i - preambleLength - 1, 0):i] {
			for _, value2 := range transmission[MaxOf(i - preambleLength - 1, 0):i] {
				if value1 + value2 == transmission[i] {
					i++
					continue loop
				}
			}
		}

		invalidValue = transmission[i]
		break
	}

	return invalidValue
}

func MaxOf(vars ...int) int {
	min := vars[0]

	for _, i := range vars {
		if min < i {
			min = i
		}
	}

	return min
}