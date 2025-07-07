package main

import "fmt"

func assignGrade(score int) string {
	switch {
	case score >= 90 && score <= 100:
		return "A"
	case score >= 80 && score < 90:
		return "B"
	case score >= 70 && score < 80:
		return "C"
	case score >= 60 && score < 70:
		return "D"
	case score < 60 && score >= 0:
		return "F"
	}
	return "Invalid"
}

func main() {
	scores := []int{95, 85, 73, 65, 59, -5, 101}

	for _, score := range scores {
		grade := assignGrade(score)
		fmt.Printf("Score: %d -> Grade: %s\n", score, grade)
	}
}
