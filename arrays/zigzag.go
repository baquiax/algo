package arrays

// Given an array of DISTINCT elements, rearrange the elements of array in zig-zag fashion in O(n) time. The converted array should be in form a < b > c < d > e < f.

// Example:

// Input: arr[] = {4, 3, 7, 8, 6, 2, 1}
// if arr[0] > arr[1]: true
//   swap(arr[0], arra[1])

// arr[] = {3, 4, 7, 8, 6, 2, 1}
// if arr[1] < arry[2]: true
//   swap(arry[1], arry[2])
//
// arr[] = {3, 7, 4, 8, 6, 2, 1}
// if arr[2] > arry[3]: true
//   swap(arry[1], arry[2])

// Output: arr[] = {3, 7, 4, 8, 2, 6, 1}

// Input: arr[] = {1, 4, 3, 2}
// Output: arr[] = {1, 4, 2, 3}
//
// arr[] = {1, 4, 3, 2}
// if arr[0] > arry[1]: false
//   swap(arry[1], arry[2])
//
// arr[] = {1, 4, 3, 2}
// if arr[1] < arry[2]: false
//   swap(arry[1], arry[2])
//
// arr[] = {1, 4, 3, 2}
// if arr[2] < arry[3]: false
//   swap(arry[1], arry[2])

func doZigzag(input []int) []int {
	// O(n) in space... This because a good practice in Go, to no
	// have side-effects functions
	numbers := append([]int{}, input...)

	// O(1) in space to keept the flag
	shouldBeLess := true
	// O(n) in time,
	for i := 0; i <= len(numbers)-2; i++ {
		nextI := i + 1
		if shouldBeLess && numbers[i] > numbers[nextI] {
			// do swap
			tmp := numbers[i]
			numbers[i] = numbers[nextI]
			numbers[nextI] = tmp

		}
		if !shouldBeLess && numbers[i] < numbers[nextI] {
			// do swap
			tmp := numbers[i]
			numbers[i] = numbers[nextI]
			numbers[nextI] = tmp

		}

		shouldBeLess = !shouldBeLess
	}

	return numbers
}
