package main

import "fmt"

func main() {
	testScores := []float64{
		87.3,
		105,
		63.5,
		27,
	}

	c := clone(testScores)

	fmt.Println(&testScores[0], &c[0], c)

}

func clone(s []float64) []float64 {
	result := make([]float64, len(s))
	for i, v := range s {
		result[i] = v
	}
	return result
}