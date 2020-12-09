package customs

import (
	"fmt"
	"strings"
)

func CustomsQuestionnaire(data []byte) int {
	values := strings.Split(string(data), "\n\n")

	answersTotal := 0
	var answers map[string]bool
	for _, questionnaireGroup := range values {
		answers = make(map[string]bool)
		groupAnswers := strings.Split(questionnaireGroup, "\n")

		for _, groupAnswer := range groupAnswers {
			for _, groupIndividualAnswers := range []byte(groupAnswer) {
				answers[string(groupIndividualAnswers)] = true
			}
		}

		fmt.Println(len(answers))
		answersTotal += len(answers)
	}

	return answersTotal
}
