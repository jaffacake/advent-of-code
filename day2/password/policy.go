package password

import (
	"regexp"
	"strconv"
	"strings"
)

type policy struct {
	min int
	max int
	letter string
	password string
}

func (p policy) testPassword() bool {
	containsCount := strings.Count(p.password, p.letter)

	return containsCount >= p.min && containsCount <= p.max
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
			password: policyData[4],
		}

		p.min, _ = strconv.Atoi(policyData[1])
		p.max, _ = strconv.Atoi(policyData[2])

		if p.testPassword() {
			correctPasswords++
		}
	}

	return correctPasswords
}
