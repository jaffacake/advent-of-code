package execution

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type execution struct {
	action string
	value  int
}

func RunExecution(data []byte) int {
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

	for i, value := range executions {
		originalExecution := executions[i]
		switch value.action {
		case "nop":
			executions[i].action = "jmp"
		case "jmp":
			executions[i].action = "nop"
		default:
			continue
		}

		result, err := run(executions)

		if err != nil {
			executions[i] = originalExecution
			continue
		}

		return result
	}

	return 0
}

func run(executions []execution) (int, error) {
	executionHistory := []int{}
	accumulatorValue := 0

	i := 0
	for i < len(executions) {
		e := executions[i]

		executionHistory = append(executionHistory, i)

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
			return 0, errors.New("loop")
		}
	}

	fmt.Println(executionHistory)

	return accumulatorValue, nil
}

func Find(slice []int, val int) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}
