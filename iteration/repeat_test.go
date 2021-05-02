package iteration

import (
	"fmt"
	"testing"
)

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 12)
	}
}

func ExampleRepeat() {
	repeat := Repeat("b", 5)
	fmt.Println(repeat)
	// Output: bbbbb
}
func TestIteration(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, repeated, expected string) {
		t.Helper()
		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	}

	t.Run("repeat the character 6 times by using loop", func(t *testing.T) {
		repeated := Repeat("a", 6)
		expected := "aaaaaa"
		assertCorrectMessage(t, repeated, expected)
	})

	t.Run("repeat the character 6 times by using strings.Repeat", func(t *testing.T) {
		repeated := RepeatWithPackage("a", 6)
		expected := "aaaaaa"
		assertCorrectMessage(t, repeated, expected)
	})

}
