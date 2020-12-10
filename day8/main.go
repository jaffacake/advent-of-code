package main

import (
	"day8/execution"
	"fmt"
	"io/ioutil"
)

func main()  {
	data, err := ioutil.ReadFile("./input.txt")

	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}

	fmt.Println(execution.RunExecution(data))
}
