package main

import (
	"day1/expenses"
	"fmt"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("./input.txt")

	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}

	fmt.Println(expenses.Calculate(data))
}