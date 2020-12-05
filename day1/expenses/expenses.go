package expenses

import (
	"strconv"
	"strings"
)

func Calculate(data []byte) (int, error) {
	values := strings.Split(string(data), "\n")

	for _, value := range values {
		if intValue, err := strconv.Atoi(value); err == nil {
			for _, value2 := range values {
				if intValue2, err := strconv.Atoi(value2); err == nil {
					if intValue + intValue2 == 2020 {
						return intValue * intValue2, nil
					}
				}
			}
		}
	}

	return 0, nil
}
