package mathematics

import (
	"fmt"
	"strconv"
)

// Avg is responsible to calc the average
func Avg(numbers ...float64) float64 {
	total := 0.0
	for _, num := range numbers {
		total += num
	}

	avg := total / float64(len(numbers))
	roundedAvg, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", avg), 64)

	return roundedAvg
}
