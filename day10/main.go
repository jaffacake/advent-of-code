package main

import (
	"day10/adapter"
	"fmt"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("./input.txt")

	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}

	fmt.Println(adapter.CalculateJolts(data))
}
