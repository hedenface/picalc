package main

import (
	"fmt"
	"math"
)

const (
	maxNum = 99999.0
)

func main() {
	// we start with the closest known approximation
	numerator := 22.0
	denominator := 7.0

	// then we calculate the difference from pi
	approx := numerator / denominator
	piDiff := math.Abs(math.Pi - approx)

	fmt.Printf("pi = %f\n", math.Pi)

	fmt.Printf("(%f / %f) = (%f), difference from pi = %f\n", numerator, denominator, approx, piDiff)

	// now we need to loop the numerator/denominators up to maxNum
	// and keep track of the numerators and denominators that have
	// a closer proximity to pi
	var n, d float64
	for d = 1.0; d <= maxNum; d++ {

        if ((n / d) > math.Pi) {
            continue;
        }

		for n = 1.0; n <= maxNum; n++ {
			approx = (n / d)

            if approx > math.Pi {
                break
            }

			thisDiff := math.Abs(math.Pi - approx)

			if thisDiff <= piDiff {
				numerator = n
				denominator = d
				piDiff = thisDiff
			}
		}
	}

	fmt.Printf("(%f / %f) = (%f), difference from pi = %f\n", numerator, denominator, approx, piDiff)
}
