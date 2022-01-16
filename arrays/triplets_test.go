package arrays

import (
	"fmt"
	"testing"
)

func TestFindTriples(t *testing.T) {
	tests := []struct {
		numbers   []int
		sumTarget int
		want      int
	}{
		{
			numbers:   []int{-2, 0, 1, 3},
			sumTarget: 2,
			want:      2,
		},
		{
			numbers:   []int{5, 1, 3, 4, 7},
			sumTarget: 12,
			want:      4,
		},
		{
			numbers:   []int{5, 6, 4, 3, 3},
			sumTarget: 4,
			want:      0,
		},
		{
			numbers:   []int{},
			sumTarget: 4,
			want:      0,
		},
		{
			numbers:   []int{1, 50, 13, 23, 40, 33, 3, 1},
			sumTarget: 6,
			want:      1,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Should get %d When %v is given", test.want, test.numbers), func(t *testing.T) {
			result := findTriplets(test.numbers, test.sumTarget)

			if result != test.want {
				t.Fatalf("wanted %d but got %d", test.want, result)
			}
		})
	}

}

func TestFindTriples2(t *testing.T) {
	tests := []struct {
		numbers   []int
		sumTarget int
		want      int
	}{
		{
			numbers:   []int{-2, 0, 1, 3},
			sumTarget: 2,
			want:      2,
		},
		{
			numbers:   []int{5, 1, 3, 4, 7},
			sumTarget: 12,
			want:      4,
		},
		{
			numbers:   []int{5, 6, 4, 3, 3},
			sumTarget: 4,
			want:      0,
		},
		{
			numbers:   []int{},
			sumTarget: 4,
			want:      0,
		},
		{
			numbers:   []int{1, 50, 13, 23, 40, 33, 3, 1},
			sumTarget: 6,
			want:      1,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Should get %d When %v is given", test.want, test.numbers), func(t *testing.T) {
			result := findTriplets2(test.numbers, test.sumTarget)

			if result != test.want {
				t.Fatalf("wanted %d but got %d", test.want, result)
			}
		})
	}

}
func BenchmarkFindTriples(b *testing.B) {
	var result int
	for i := 0; i <= b.N; i++ {
		zeros := make([]int, 1000)
		result = findTriplets(zeros, 1000)
	}

	b.Logf("Last triplets found: %d", result)
}

func BenchmarkFindTriples2(b *testing.B) {
	var result int
	for i := 0; i <= b.N; i++ {
		zeros := make([]int, 1000)
		result = findTriplets2(zeros, 1000)
	}

	b.Logf("Last triplets found: %d", result)
}
