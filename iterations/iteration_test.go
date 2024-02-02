package iteration

import (
	"fmt"
	"testing"
)

func ExampleRepeat() {
	repeated := Repeat("a", 5)
	fmt.Println(repeated)
	//Output: aaaaa
}
func TestRepeat(t *testing.T) {
	t.Run("repeat using cycles", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"

		if repeated != expected {
			t.Errorf("expected %q got %q", expected, repeated)
		}
	})

	t.Run("repeat using built-in method", func(t *testing.T) {
		repeated := RepeatWithFunction("a", 5)
		expected := "aaaaa"

		if repeated != expected {
			t.Errorf("expected %q got %q", expected, repeated)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}
