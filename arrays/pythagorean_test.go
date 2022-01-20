package arrays

import (
	"testing"
)

func TestPythagoreanTriplet(t *testing.T) {
	tests := map[string]struct {
		given []int
		want  bool
	}{
		"test 1": {
			given: []int{3, 1, 4, 6, 5},
			want:  true,
		},
		"test 2": {
			given: []int{10, 4, 6, 12, 5},
			want:  false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := hasPythagoreanTriplet(test.given)
			if test.want != result {
				t.Fatalf("wanted %v but got %v", test.want, result)
			}
		})
	}
}
