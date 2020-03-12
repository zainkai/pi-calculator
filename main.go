package main

import (
	"fmt"
	"math"
	"math/rand"
)

type Pi struct {
	// randomFloat64 a uniform distributed Number [0.0, 1.0)
	randomFloat64 func() float64
}

func distOrigin(x, y float64) float64 {
	x *= x
	y *= y

	return math.Sqrt(x + y)
}

func (p Pi) Calculate(percision int) float64 {
	countInCircle := 0

	for i := 0; i < percision; i++ {
		x, y := p.randomFloat64(), p.randomFloat64()
		if distOrigin(x, y) <= 1 {
			countInCircle++
		}
	}

	return 4.0 * float64(countInCircle) / float64(percision)
}

func main() {
	pi := Pi{
		// anonymous function
		func() float64 { return rand.Float64() },
	}

	fmt.Println(pi.Calculate(10000))
}
