package arrays

import (
	"math"
	"sort"
)

// Given an array of integers, write a function that returns true if
// there is a triplet (a, b, c) that satisfies a^2 + b^2 = c^2. => a^2+b^2-c^2=0

// Example:

// Input: arr[] = {3, 1, 4, 6, 5}
// Output: True
// There is a Pythagorean triplet (3, 4, 5).

// Input: arr[] = {10, 4, 6, 12, 5}
// Output: False
// There is no Pythagorean triplet.

// 1. Looking the first triplet (one is enough)
// Naive approach is doing a triple nested for to build all the triplets possible,
// checking if the three elements in every iteration fulfill the criteria. O(n^3)
//
// 2. I can order the array.  O(n log n)
// 2.1 Iterate over the ordered array O(n)
// 3.2 Iterate in the array with two indexes trying to find those two numbers that gives as
// a result the element in the first loop O(n)
// O(n log n) + O (n) * O(n) => O(n lon g) + O(n^2)

// NOTE: A B and C should be on that order in the array.
func hasPythagoreanTriplet(array []int) bool {
	sort.Ints(array) // n log n

	// a^2 = c^2 - b^2
	for c := len(array) - 1; c >= 0; c-- { // n
		a := 0
		b := len(array) - 1
		cSquare := math.Pow(float64(array[c]), 2)
		for a < b {
			aSquare := math.Pow(float64(array[a]), 2)
			bSquare := math.Pow(float64(array[b]), 2)
			currentC := aSquare + bSquare
			if currentC < cSquare {
				a++
				continue
			}

			if currentC > cSquare {
				b--
				continue
			}

			return true
		}
	}

	return false
}
