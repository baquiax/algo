package arrays

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDoZigzag(t *testing.T) {
	tests := []struct {
		given []int
		want  []int
	}{
		{
			given: []int{1, 4, 3, 2},
			want:  []int{1, 4, 2, 3},
		},
		{
			given: []int{4, 3, 7, 8, 6, 2, 1},
			want:  []int{3, 7, 4, 8, 2, 6, 1},
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Should return %v When %v is given", test.want, test.given), func(t *testing.T) {
			result := doZigzag(test.given)

			if !reflect.DeepEqual(result, test.want) {
				t.Fatalf("want %v but got %v", test.want, result)
			}
		})
	}

}
