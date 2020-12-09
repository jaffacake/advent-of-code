package customs

import (
	"fmt"
	"strings"
)

func CustomsQuestionnaire(data []byte) int {
	values := strings.Split(string(data), "\n\n")

	answersTotal := 0
	var answers map[string]int
	for _, questionnaireGroup := range values {
		answers = make(map[string]int)
		groupAnswers := strings.Split(questionnaireGroup, "\n")

		for _, groupAnswer := range groupAnswers {
			for _, groupIndividualAnswers := range []byte(groupAnswer) {
				answers[string(groupIndividualAnswers)] += 1
			}
		}

		fmt.Println(answers)
		fmt.Println(len(groupAnswers))
		answersForGroup := 0
		for _, answersCheck := range answers {
			if answersCheck == len(groupAnswers) {
				answersForGroup += 1
			}
		}
		fmt.Println(answersForGroup)
		answersTotal += answersForGroup
	}

	return answersTotal
}
