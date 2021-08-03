package degToRad

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(degToRad(12))
}

func degToRad(deg float64) float64 {
	return deg * math.Pi / 180
}
