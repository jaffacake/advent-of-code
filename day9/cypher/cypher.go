package cypher

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func formatData(data []byte) []int {
	values := strings.Split(string(data), "\n")

	var transmission []int
	for _, value := range values {
		intValue, _ := strconv.Atoi(value)
		transmission = append(transmission, intValue)
	}

	return transmission
}

func FindWeakness(data []byte) int {
	transmission := formatData(data)

	// 5 for test
	// 25 for real run
	preambleLength := 25
	i := preambleLength

	invalidValue := 0

	//10884537

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

func CalculateWeaknessArea(data []byte) int {
	weaknessNumber := FindWeakness(data)
	transmission := formatData(data)

	attackVectorStart := 0
	attackVectorEnd := 0

	for indexStart, number := range transmission {
		calculation := number

		for indexEnd, addNumber := range transmission[indexStart+1:] {
			calculation += addNumber

			if calculation == weaknessNumber {
				attackVectorStart = indexStart
				attackVectorEnd = indexEnd
				fmt.Println("found weakness vector")
				fmt.Println(indexStart)
				fmt.Println(indexEnd)
				break
			}
		}
	}

	attackVector := transmission[attackVectorStart:attackVectorStart + attackVectorEnd + 2]

	fmt.Println(sumSlice(attackVector))
	sort.Ints(attackVector)

	return attackVector[0] + attackVector[len(attackVector) - 1]
}

func sumSlice(vars []int) int {
	result := 0
	for _, numb := range vars {
		result += numb
	}
	return result
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