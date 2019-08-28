package arrays

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("Testing with 5 elements", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15
		if got != want {
			t.Errorf("Got %d, want %d, using %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("Testing with {1, 2} and {3, 6}", func(t *testing.T) {
		slice1 := []int{1, 2}
		slice2 := []int{3, 6}

		got := SumAll(slice1, slice2)
		want := []int{3, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %v, want %v, using [%v, %v]", got, want, slice1, slice2)
		}
	})
}

func BenchmarkSum(b *testing.B) {
	numbers := []int{1, 2, 3, 4}
	for i := 0; i < b.N; i++ {
		Sum(numbers)
	}
}

func BenchmarkSumAll(b *testing.B) {
	slice1 := []int{1, 2}
	slice2 := []int{3, 6}
	for i := 0; i < b.N; i++ {
		SumAll(slice1, slice2)
	}
}

func ExampleSum() {
	numbers := []int{1, 2, 3, 4}
	res := Sum(numbers)
	fmt.Printf("%d", res)
	// Output: 10
}

func ExampleSumAll() {
	slice1 := []int{1, 2}
	slice2 := []int{3, 6}
	res := SumAll(slice1, slice2)
	fmt.Printf("%v", res)
	// Output: [3 9]
}