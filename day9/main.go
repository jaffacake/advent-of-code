package main

import (
	"day9/cypher"
	"fmt"
	"io/ioutil"
)

func main()  {
	data, err := ioutil.ReadFile("./input.txt")

	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}

	fmt.Println(cypher.CalculateWeaknessArea(data))
}
