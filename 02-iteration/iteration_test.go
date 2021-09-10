package iteration

import (
	"fmt"
	"testing"
)

func ExampleRepeat() {
	repeated := Repeat("z", 5)
	fmt.Println(repeated)
	// Output: zzzzz
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func TestRepeat(t *testing.T) {
	t.Run("shuld repeat 5 times a letter 'a'", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})

	t.Run("shuld repeat 8 times a letter 'z'", func(t *testing.T) {
		repeated := Repeat("z", 8)
		expected := "zzzzzzzz"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})
}
