package bags

import (
	"fmt"
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
	fmt.Println(bagRules)

	// Calculate which bags can contain shiny gold bags
	result := 0
	for ruleBag, rule := range bagRules {
		count := canContainShinyGoldBags(bagRules, rule)

		if count > 0 {
			fmt.Println(ruleBag)
		}
		result += count
	}
	fmt.Println(result)

	return result
}

func canContainShinyGoldBags(bagRules map[string]map[string]int, rules map[string]int) int {
	count := 0
	for bag, _ := range rules {
		if strings.Contains(bag, "shiny gold bag") {
			return 1
		}
		count += canContainShinyGoldBags(bagRules, bagRules[bag])
	}

	if count > 0 {
		return 1
	}
	return 0
}