package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(radToDeg(12))
}

func radToDeg(deg float64) float64 {
	return deg * 180 / math.Pi
}
