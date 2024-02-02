package arrays

import (
	"fmt"
	"reflect"
	"testing"
)

func ExampleSum() {
	got := Sum([]int{1, 2, 3, 4, 5})
	fmt.Println(got)
	// Output: 15
}

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	got := Sum(numbers)
	expected := 15
	if got != expected {
		t.Errorf("expected %d got %d", expected, got)
	}
	t.Run("adding slices", func(t *testing.T) {
		got := SumSlices([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	got := SumAllTails([]int{1, 2}, []int{0, 9}, []int{})
	want := []int{2, 9, 0}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
