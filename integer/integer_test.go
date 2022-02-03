package integer

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {

	assertCorrectRes := func(t testing.TB, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("expected '%d' but got '%d'", want, got)
		}
	}

	t.Run("2 + 2", func(t *testing.T) {
		assertCorrectRes(t, Add(2, 2), 4)
	})

	t.Run("2 + 3", func(t *testing.T) {
		assertCorrectRes(t, Add(2, 3), 5)
	})

}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
