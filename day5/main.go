package main

import (
	"day5/boardingPass"
	"fmt"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("./input.txt")

	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}

	fmt.Println(boardingPass.CalculateSeat(data))
}