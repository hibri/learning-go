package checkout

import "strings"

func scan(items string) int {

	var total = 0

	for i := 0; i < len(items); i++ {
		total += getPrice(string(items[i]))
	}
	return total - discount(items)
}

func getPrice(item string) int {
	prices := make(map[string]int)
	prices["A"] = 50
	prices["B"] = 30

	price := prices[item]

	return price

}

func discount(items string) int {

	countA := strings.Count(items, "A")
	countB := strings.Count(items, "B")

	if (countA > 0) && ((countA % 3) == 0) {
		return 20
	} else if (countB > 0) && ((countB % 2) == 0) {
		return 15
	} else {
		return 0
	}

}
