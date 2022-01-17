package arrays

import "sort"

// Write a program that, given an array A[] of n numbers and another number x,
// determines whether or not there exist two elements in A[] whose sum is exactly x.

// Examples:

// Input: arr[] = {0, -1, 2, -3, 1}
//         sum = -2
// Output: -3, 1
//          Valid pair exists.

// If we calculate the sum of the output,
// 1 + (-3) = -2

// Input: arr[] = {1, -2, 1, 0, 5}
//        sum = 0
// Output:
//         No valid pair exists.

// 1. sort the numbers O(n)
// 2. have two indexes, left and right
// 3. start sum up the numbers forn A[left] and A[right]
// 4. if the result is greater than the target, right--
// 4. if the sum result is less than the target, left++

// O(n log n) + O(n)
func getSumPair(array []int, target int) bool {
	sort.Ints(array)

	i := 0
	j := len(array) - 1
	for i < j {
		sum := array[i] + array[j]
		if sum > target {
			j--
			continue
		}

		if sum < target {
			i++
			continue
		}

		return true

	}

	return false
}

// O(n) time complexity but O(n) in the worst in space
//
func getSumPair2(array []int, target int) bool {
	waitingPair := make(map[int]bool)

	for i := 0; i < len(array); i++ {
		diff := target - array[i]
		if waitingPair[diff] {
			return true
		}

		waitingPair[array[i]] = true
	}
	return false
}
