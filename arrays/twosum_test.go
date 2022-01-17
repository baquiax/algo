package arrays

import (
	"testing"
)

func TestGetsumPair(t *testing.T) {
	tests := []struct {
		name   string
		array  []int
		target int
		want   bool
	}{
		{
			name:   "test 1",
			array:  []int{1, 2, 3, 4, 5},
			target: 9,
			want:   true,
		},
		{
			name:   "test 2",
			array:  []int{1, 2, 3, 4, 5},
			target: -1,
			want:   false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := getSumPair(test.array, test.target)

			if result != test.want {
				t.Fatalf("expected %v but got %v", test.want, result)
			}
		})
	}
}

func TestGetsumPair2(t *testing.T) {
	tests := []struct {
		name   string
		array  []int
		target int
		want   bool
	}{
		{
			name:   "test 1",
			array:  []int{1, 2, 3, 4, 5},
			target: 9,
			want:   true,
		},
		{
			name:   "test 2",
			array:  []int{1, 2, 3, 4, 5},
			target: -1,
			want:   false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := getSumPair2(test.array, test.target)

			if result != test.want {
				t.Fatalf("expected %v but got %v", test.want, result)
			}
		})
	}
}

func BenchmarkWorstGetPair(b *testing.B) {
	var exists bool
	var array = make([]int, 100)
	for i := 0; i <= b.N; i++ {
		exists = getSumPair(array, -1)
	}

	if exists {
		b.Fatal()
	}
}

func BenchmarkWorstGetPair2(b *testing.B) {
	var exists bool
	var array = make([]int, 1000)
	for i := 0; i <= b.N; i++ {
		exists = getSumPair2(array, -1)
	}

	if exists {
		b.Fatal()
	}
}

func BenchmarkGetPair(b *testing.B) {
	var exists bool
	for i := 0; i <= b.N; i++ {
		exists = getSumPair([]int{1, 2, 4, 6, 5, 7, 8}, 3)
	}

	if !exists {
		b.Fatal()
	}
}

func BenchmarkGetPair2(b *testing.B) {
	var exists bool
	for i := 0; i <= b.N; i++ {
		exists = getSumPair2([]int{1, 2, 4, 6, 5, 7, 8}, 3)
	}

	if !exists {
		b.Fatal()
	}
}
