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
	fmt.Println(threeDiffs)

	return 0
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