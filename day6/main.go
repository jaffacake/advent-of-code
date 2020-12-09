package main

import (
	"day6/customs"
	"fmt"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("./input.txt")

	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}

	fmt.Println(customs.CustomsQuestionnaire(data))
}