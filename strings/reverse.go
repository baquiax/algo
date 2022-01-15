package strings

/*
Given a string, that contains a special character together
with alphabets (‘a’ to ‘z’ and ‘A’ to ‘Z’), reverse the string
in a way that special characters are not affected.

 Examples:
 	Input:   str = "a,b$c"
	Output:  str = "c,b$a"

Note that $ and , are not moved anywhere.
Only subsequence "abc" is reversed

	Input:   str = "Ab,c,de!$"
	Output:  str = "ed,c,bA!$"
*/

// O(n) + O(m) in time
// O(n) + O(m) in space
//
// This is a mid level algorithm, but it's not too bad. ~10 mins invested
func reverse(input string) string {
	// Brainstorming
	// - We only need to reverse [a-zA-Z] values keeping the rest intact
	// - ASCII ranges are: [a-z] = 97-122, [A-Z] = 65-90
	//
	// How to solve it?
	// 1. Traverse the string and generate: O(n) in time
	// - An array without special characters O(n) in space
	// - An array with a small struct witht he special char and its original position
	// 2. Reverse array without special characters
	// 3. Traverse the array with the special characters and insert them in the correct position
	//

	type specialRune struct {
		rune rune
		pos  int
	}

	// O(n) in space
	simpleString := make([]rune, 0)

	// O(m) in space
	specialChars := make([]specialRune, 0)

	// O(n) in time
	for i, r := range input {
		if r >= 97 && r <= 122 || r >= 65 && r <= 90 {
			simpleString = append([]rune{r}, simpleString...)
			continue
		}

		specialChars = append(specialChars, specialRune{
			rune: r,
			pos:  i,
		})
	}

	// O(m) in time
	for _, v := range specialChars {
		simpleString = append(simpleString[:v.pos], append([]rune{v.rune}, simpleString[v.pos:]...)...)
	}

	return string(simpleString)
}

// O(n) in time
// O(n) in space
//
// Due the imutability in Go of strings, O(n) in space is required
// to mutate the array representation of the string, but in general this
// is better than the prior one
func reverse2(input string) string {
	// 1. Run the usual reverse algorithm (using a temp val) but with a little change
	// 2. The small change will be having an extra index to control the end position

	inputSlice := []byte(input)
	leftPosition := 0
	rightPosition := len(inputSlice) - 1

	for leftPosition < rightPosition {
		if !isAlphabet(inputSlice[leftPosition]) {
			leftPosition++
			continue
		}

		if !isAlphabet(inputSlice[rightPosition]) {
			rightPosition--
			continue
		}

		tmp := inputSlice[leftPosition]
		inputSlice[leftPosition] = inputSlice[rightPosition]
		inputSlice[rightPosition] = tmp

		leftPosition++
		rightPosition--
	}

	return string(inputSlice)
}

func isAlphabet(r byte) bool {
	return r >= 97 && r <= 122 || r >= 65 && r <= 90
}
