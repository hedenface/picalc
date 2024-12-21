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
	piDiff := math.Abs(math.Pi - (numerator / denominator))

	fmt.Printf("%f / %f has approximation (%f)\n", numerator, denominator, piDiff)

	// now we need to loop the numerator/denominators up to maxNum
	// and keep track of the numerators and denominators that have
	// a closer proximity to pi
	for d := 1.0; d <= maxNum; d++ {
		for n := 1.0; n <= maxNum; n++ {
			thisDiff := math.Abs(math.Pi - (n / d))

			if thisDiff <= piDiff {
				numerator = n
				denominator = d
				piDiff = thisDiff
			}
		}
	}

	fmt.Printf("Got %f / %f as closest approximation (%f)\n", numerator, denominator, piDiff)
}
