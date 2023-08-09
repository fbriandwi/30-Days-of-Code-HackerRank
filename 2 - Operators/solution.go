package main

import (
	"fmt"
)

func main() {
	var mealCost float64
	var tipPercent, taxPercent int
	fmt.Scan(&mealCost, &tipPercent, &taxPercent)

	totalCost := mealCost + (mealCost * float64(tipPercent) * 0.01) + (mealCost * float64(taxPercent) * 0.01)
	fmt.Printf("The total meal cost is %.0f dollars.\n", totalCost)
}
