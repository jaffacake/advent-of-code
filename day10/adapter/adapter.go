package adapter

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func CalculateJolts(data []byte) int {
	adapters := formatData(data)
	sort.Ints(adapters)

	adapters = append(adapters, adapters[len(adapters)+1] + 3)

	//start from 0

	// looop
	oneDiffs := 0
	twoDiffs := 0
	threeDiffs := 0

	for i, adapter := range adapters {
		diff := 0
		if i == 0 {
			diff = adapter - 0
		} else {
			diff = adapter - adapters[i-1]
		}

		switch diff {
		case 0:
			break
		case 1:
			oneDiffs++
		case 2:
			twoDiffs++
		case 3:
			threeDiffs++
		default:
			fmt.Println("difference too big")
		}
	}

	threeDiffs++

	fmt.Println(oneDiffs)
	fmt.Println(twoDiffs)
	fmt.Println(threeDiffs)

	return 0
}

func CalculatePossibleCombinations(data []byte) int {
	adapters := formatData(data)
	adapters = append(adapters, 0)
	sort.Ints(adapters)
	adapters = append(adapters, adapters[len(adapters)-1] + 3)

	return checkForCombinations(adapters, 0)
}
var adapterCombinationsTried = make(map[int]int)

func checkForCombinations(adapters []int, startAdapter int) int {
	possibleCombinations := 0
	nextAdapter := startAdapter + 1
	furthestPossibleAdapter := startAdapter + 3

	if furthestPossibleAdapter >= len(adapters) {
		return 1
	} else if _, ok := adapterCombinationsTried[adapters[startAdapter]]; ok {
		return adapterCombinationsTried[adapters[startAdapter]]
	}

	for i := nextAdapter; i <= furthestPossibleAdapter; i++ {
		lower := adapters[startAdapter]
		upper := adapters[i]

		if (upper - lower) >= 1 && (upper - lower) <= 3 {
			possibleCombinations += checkForCombinations(adapters, i)
		}
	}

	adapterCombinationsTried[adapters[startAdapter]] = possibleCombinations

	return possibleCombinations
}

func formatData(data []byte) []int {
	values := strings.Split(string(data), "\n")

	var transmission []int
	for _, value := range values {
		intValue, _ := strconv.Atoi(value)
		transmission = append(transmission, intValue)
	}

	return transmission
}