package main

import "fmt"

func averageMarks(scores []float64) float64 {
	sum := 0.0
	for _, score := range scores {
		sum += score
	}
	average := sum / float64(len(scores))
	if len(scores) == 0 {
		return 1
	}
	return average
}

func main() {
	scores := []float64{78.5, 82.0, 90.5, 85.5, 89.0}
	average := averageMarks(scores)
	fmt.Printf("Average: %.2f\n", average)
}
