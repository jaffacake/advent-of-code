package bags

import (
	"strconv"
	"strings"
)

func CalculateBags(data []byte) int {
	values := strings.Split(string(data), "\n")

	bagRules := map[string]map[string]int{}
	for _, rule := range values {
		rules := strings.Split(rule, " contain ")

		//3 wavy bronze bags, 5 clear tomato bags.
		rules[1] = strings.Replace(rules[1], ".", "", 1)
		splitRules := strings.Split(rules[1], ", ")

		if !strings.Contains(rules[1], "no other bags") {
			rulesMap := map[string]int{}
			for _, rule := range splitRules {
				numberOfBags, _ := strconv.Atoi(string(rule[0]))
				bagName := strings.Replace(rule[2:], "bags", "replace", 1)
				bagName = strings.Replace(bagName, "bag", "replace", 1)
				bagName = strings.Replace(bagName, "replace", "bags", 1)
				rulesMap[bagName] = numberOfBags
			}
			bagRules[rules[0]] = rulesMap
		}
	}

	// Calculate which bags can contain shiny gold bags
	result := 0

	for ruleBag, rule := range bagRules {
		if strings.Contains(ruleBag, "shiny gold bag") {
			result += canContainShinyGoldBags(bagRules, rule)
		}
	}

	return result
}

func canContainShinyGoldBags(bagRules map[string]map[string]int, rules map[string]int) int {
	bagCount := 0
	for bag, numberOfBags := range rules {
		if strings.Contains(bag, "contain no other bags") {
			return bagCount
		}

		bagCount += numberOfBags
		bagCount += numberOfBags * canContainShinyGoldBags(bagRules, bagRules[bag])
	}

	return bagCount
}