package expenses

import (
	"strconv"
	"strings"
)

func Calculate(data []byte) (int, error) {
	values := strings.Split(string(data), "\n")
	intValues := make([]int, len(values) - 1)

	for i, value := range values {
		if intValue, err := strconv.Atoi(value); err == nil && intValue != 0 {
			intValues[i] = intValue
		}
	}

	for _, value := range intValues {
		for _, value2 := range intValues {
			for _, value3 := range intValues {
				if value + value2 + value3 == 2020 {
					return value * value2 * value3, nil
				}
			}
		}
	}

	return 0, nil
}
