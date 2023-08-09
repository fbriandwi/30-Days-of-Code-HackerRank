package main

import (
	"fmt"
	"math"
)

func main() {
	var mealCost float64
    var tipPercent, taxPercent int
    fmt.Scan(&mealCost, &tipPercent, &taxPercent)

    totalCost := mealCost + (mealCost * float64(tipPercent) / 100) + (mealCost * float64(taxPercent) / 100)
    roundedTotalCost := math.Round(totalCost)
    fmt.Printf("%.0f\n", roundedTotalCost)}
