package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("./input.txt")

	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}

	values := strings.Split(string(data), "\n")

	for _, value := range values {
		if intValue, err := strconv.Atoi(value); err == nil {
			for _, value2 := range values {
				if intValue2, err := strconv.Atoi(value2); err == nil {
					if intValue + intValue2 == 2020 {
						fmt.Println("match")
						fmt.Println(intValue)
						fmt.Println(intValue2)
						fmt.Println(intValue * intValue2)

						return
					}
				}
			}
		}
	}
}