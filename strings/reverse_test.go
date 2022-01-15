package strings

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "a,b$c",
			want:  "c,b$a",
		},
		{
			input: "Ab,c,de!$",
			want:  "ed,c,bA!$",
		},
		{
			input: "a!!!b.c.d,e'f,ghi",
			want:  "i!!!h.g.f,e'd,cba",
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Return %s When %s is given", test.want, test.input), func(t *testing.T) {
			result := reverse(test.input)

			if result != test.want {
				t.Errorf("got %s want %s", result, test.want)
			}
		})
	}
}

func TestReverse2(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "a,b$c",
			want:  "c,b$a",
		},
		{
			input: "Ab,c,de!$",
			want:  "ed,c,bA!$",
		},
		{
			input: "a!!!b.c.d,e'f,ghi",
			want:  "i!!!h.g.f,e'd,cba",
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Return %s When %s is given", test.want, test.input), func(t *testing.T) {
			result := reverse2(test.input)

			if result != test.want {
				t.Errorf("got %s want %s", result, test.want)
			}
		})
	}
}

// cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
// BenchmarkReverse-12    	 1308486	       923.2 ns/op	     872 B/op	      31 allocs/op
// PASS
func BenchmarkReverse(b *testing.B) {
	for n := 0; n < b.N; n++ {
		reverse("a!!!b.c.d,e'f,ghi")
	}
}

// cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
// BenchmarkReverse2-12    	27015704	        42.85 ns/op	      24 B/op	       1 allocs/op
// PASS
func BenchmarkReverse2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		reverse2("a!!!b.c.d,e'f,ghi")
	}
}
