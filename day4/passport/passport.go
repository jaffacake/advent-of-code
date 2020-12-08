package passport

import (
	"fmt"
	"strings"
)

var requiredFields = []string{
	"byr",
	"iyr",
	"eyr",
	"hgt",
	"hcl",
	"ecl",
	"pid",
	"cid",
}

func Validate(data []byte) int {
	values := strings.Split(string(data), "\n\n")

	validPassports := 0

	for _, passport := range values {
		validatedFields := 0
		passportLines := strings.Split(passport, "\n")

		fmt.Println(passport)
		for _, field := range requiredFields {
			valid := false
			for _, passportDataLine := range passportLines {
				if strings.Contains(passportDataLine, field) {
					valid = true
				}
			}

			if valid || (!valid && field == "cid") {
				validatedFields++
			}
		}

		if validatedFields >= 8 {
			fmt.Println("valid")
			validPassports++
		} else {
			fmt.Println("invalid")
		}
	}

	return validPassports
}
