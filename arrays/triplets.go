package arrays

import (
	"sort"
)

// Given an array of distinct integers and a sum value.
// Find count of triplets with sum smaller than given sum value.
// Examples:

// Input : arr[] = {-2, 0, 1, 3}
//         sum = 2.
// Output : 2
// Explanation :  Below are triplets with sum less than 2
//                (-2, 0, 1) and (-2, 0, 3)

// Input : arr[] = {5, 1, 3, 4, 7}
//         sum = 12.
// Output : 4
// Explanation :  Below are triplets with sum less than 12
//                (1, 3, 4), (1, 3, 5), (1, 3, 7) and
//                (1, 4, 5)
// Seems to be a O(n^2)
func findTriplets(numbers []int, targetSum int) int {
	// Brainstorm
	// 1. We need to generate all the triplets and then sum up the values
	// and the discarging those that sum up >= targetSum
	//

	if len(numbers) < 3 {
		return 0
	}

	// step# 1: array: [5, 1, 3, 4, 7], target: 12
	sort.Ints(numbers)

	// step# 2: [1, 3, 4, 5, 7]
	// firstIndex: 0, secondIndex: 1, thirdIndex: 2
	// pair sum: 1 + 3 < 12 = true
	// pair + third: 4 + 4 < 12 = true
	// firstIndex: 0, secondIndex: 2, thirdIndex: 3
	// triplets: 1
	// pair sum: 1 + 4 < 12 = true
	// pair + third: 5 + 5 < 12 true
	// triplets: 2
	// pair sum: 1 + 5 < 12 = true
	// pair + third: 6 + 7 < 12 false
	// paur sum: 3

	lastIndex := len(numbers) - 1
	tripletsCounter := 0
	firstIndex := 0
	secondIndex := 1
	thirdIndex := 2

	for {
		if thirdIndex > lastIndex {
			secondIndex++
			thirdIndex = secondIndex + 1
		}

		if secondIndex > lastIndex-1 {
			firstIndex++
			secondIndex = firstIndex + 1
			thirdIndex = secondIndex + 1
		}

		if firstIndex > lastIndex-2 {
			break
		}

		firstNumber := numbers[firstIndex]
		secondNumber := numbers[secondIndex]
		thirdNumber := numbers[thirdIndex]

		// base case: if in the ordered list we found in the left hand side
		// a number greater or equals than the target number, it is not
		// needed to continue iterating.
		if firstNumber >= targetSum {
			break
		}

		pairSum := firstNumber + secondNumber
		if pairSum >= targetSum {
			secondIndex++
			thirdIndex = secondIndex + 1
			continue
		}

		fullSum := pairSum + thirdNumber
		thirdIndex++
		if fullSum >= targetSum {
			continue
		}

		tripletsCounter++
	}

	return tripletsCounter
}

// O(n^2)
func findTriplets2(numbers []int, sumTarget int) int {
	result := 0
	sort.Ints(numbers)

	for i := 0; i < len(numbers)-2; i++ {
		j := i + 1
		k := len(numbers) - 1

		for j < k {
			sum := numbers[i] + numbers[j] + numbers[k]
			if sum >= sumTarget {
				k--
				continue
			}

			result += (k - j)
			j++
		}
	}

	return result
}
