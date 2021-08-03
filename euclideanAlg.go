package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(euclideanAlgorithm(16, 4))
}

func euclideanAlgorithm(firstNumber, secondNumber int) int {
	a := math.Abs(float64(firstNumber))
	b := math.Abs(float64(secondNumber))
	if b == 0 {
		return int(a)
	} else {
		return euclideanAlgorithm(int(b), int(a)%int(b))
	}
}
