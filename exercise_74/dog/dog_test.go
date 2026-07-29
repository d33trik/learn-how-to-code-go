package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	tests := map[string]struct {
		input int
		want  int
	}{
		"positive numbers": {
			input: 10,
			want:  70,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := Years(tc.input)

			if tc.want != got {
				t.Fatalf("want: %d, got: %d", tc.want, got)
			}
		})
	}
}

func TestYearsTwo(t *testing.T) {
	tests := map[string]struct {
		input int
		want  int
	}{
		"positive numbers": {
			input: 10,
			want:  70,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := YearsTwo(tc.input)

			if tc.want != got {
				t.Fatalf("want: %d, got: %d", tc.want, got)
			}
		})
	}
}

func BenchmarkYears(b *testing.B) {
	for b.Loop() {
		Years(10)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for b.Loop() {
		YearsTwo(10)
	}
}

func ExampleYears() {
	years := Years(10)
	fmt.Println(years)
	// Output: 70
}

func ExampleYearsTwo() {
	years := YearsTwo(10)
	fmt.Println(years)
	// Output: 70
}
