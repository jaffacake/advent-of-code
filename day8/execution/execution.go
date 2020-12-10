package execution

import (
	"fmt"
	"strconv"
	"strings"
)

type execution struct {
	action string
	value  int
}

func RunExecution(data []byte) int {
	// acc
	accumulatorValue := 0
	// nop - ignore
	// jmp - move index/pointer

	values := strings.Split(string(data), "\n")
	var executions []execution

	for _, value := range values {

		e := execution{
			action: value[:3],
		}
		e.value, _ = strconv.Atoi(value[4:])
		executions = append(executions, e)
	}
	fmt.Println(executions)

	executionHistory := []int{}

	// run executions
	i := 0
	for i < len(executions) {
		executionHistory = append(executionHistory, i)
		e := executions[i]
		if e.action == "acc" {
			accumulatorValue += e.value
			i++
		} else if e.action == "nop" {
			i++
		} else {
			// must be jump
			i += e.value
		}

		if Find(executionHistory, i) {
			break;
		}
	}

	fmt.Println(executionHistory)

	return accumulatorValue
}

func Find(slice []int, val int) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}