package main

import (
	"day2/password"
	"fmt"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("./input.txt")

	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}

	fmt.Println(password.ValidatePasswords(data))
}