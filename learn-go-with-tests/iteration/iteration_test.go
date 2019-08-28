package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	got := Repeat("a", 4)
	want := "aaaa"

	if got != want {
		t.Errorf("Expected %q, got %q", want, got)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 4)
	}
}

func ExampleRepeat() {
	repeat := Repeat("x", 3)
	fmt.Printf(repeat)
	// Output: xxx
}