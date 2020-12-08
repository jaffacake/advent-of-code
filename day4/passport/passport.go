package passport

import (
	"fmt"
	"github.com/go-playground/validator"
	"strconv"
	"strings"
)

type Passport struct {
	Byr int `validate:"required,numeric,min=1920,max=2002"`
	Iyr int `validate:"required,numeric,min=2010,max=2020"`
	Eyr int `validate:"required,numeric,min=2020,max=2030"`
	Hgt string `validate:"required,height"`
	Hcl string `validate:"required,hexcolor"`
	Ecl string `validate:"required,alpha,oneof=amb blu brn gry grn hzl oth"`
	Pid string `validate:"required,numeric,len=9"`
	Cid string
}

func heightValidator(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	if strings.Contains(value, "cm") {
		values := strings.Split(value, "cm")
		height, _ := strconv.Atoi(values[0])

		return height >= 150 && height <= 193
	} else if strings.Contains(value, "in") {
		values := strings.Split(value, "in")
		height, _ := strconv.Atoi(values[0])

		return height >= 59 && height <= 76
	}

	return false
}


func Validate(data []byte) int {
	validate := validator.New()
	_ = validate.RegisterValidation("height", heightValidator)

	values := strings.Split(string(data), "\n\n")

	validPassports := 0

	for _, passport := range values {
		passportLines := strings.Split(passport, "\n")
		var passportItems []string
		var passportStruct Passport
		for _, passportDataLine := range passportLines {
			for _, data := range strings.Split(passportDataLine, " ") {
				passportItems = append(passportItems, data)
			}
		}

		for _, passportItem := range passportItems {
			passportParts := strings.Split(passportItem, ":")

			switch passportParts[0] {
			case "byr":
				passportStruct.Byr, _ = strconv.Atoi(passportParts[1])
			case "iyr":
				passportStruct.Iyr, _ = strconv.Atoi(passportParts[1])
			case "eyr":
				passportStruct.Eyr, _ = strconv.Atoi(passportParts[1])
			case "hgt":
				passportStruct.Hgt = passportParts[1]
			case "hcl":
				passportStruct.Hcl = passportParts[1]
			case "ecl":
				passportStruct.Ecl = passportParts[1]
			case "pid":
				passportStruct.Pid = passportParts[1]
			case "cid":
				passportStruct.Cid = passportParts[1]
			}
		}

		fmt.Println(passportStruct)
		err := validate.Struct(passportStruct)
		fmt.Println(err)
		if err == nil {
			validPassports++
		}
	}

	return validPassports
}
