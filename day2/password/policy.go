package password

import (
	"regexp"
	"strconv"
	"strings"
)

type policy struct {
	int1     int
	int2     int
	letter   string
	password []byte
}

func (p policy) testPassword() bool {
	containsCount := strings.Count(string(p.password), p.letter)

	return containsCount >= p.int1 && containsCount <= p.int2
}

func (p policy) testSecondPolicy() bool {
	matches := 0
	if string(p.password[p.int1 - 1]) == p.letter {
		matches++
	}

	if string(p.password[p.int2 - 1]) == p.letter {
		matches++
	}

	return matches == 1
}

func ValidatePasswords(data []byte) int {
	values := strings.Split(string(data), "\n")
	r := regexp.MustCompile(`(\d*)-(\d*) (.): (\S*)`)
	correctPasswords := 0

	for _, value := range values {
		policyData := r.FindStringSubmatch(value)

		if len(policyData) == 0 {
			break
		}

		p := policy{
			letter:   policyData[3],
			password: []byte(policyData[4]),
		}

		p.int1, _ = strconv.Atoi(policyData[1])
		p.int2, _ = strconv.Atoi(policyData[2])

		if p.testSecondPolicy() {
			correctPasswords++
		}
	}

	return correctPasswords
}
