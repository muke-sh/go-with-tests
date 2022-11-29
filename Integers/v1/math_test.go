package integers

import "testing"

func TestMath(t *testing.T) {

	t.Run("add", func(t *testing.T) {
		got := Add(2, 2)
		want := 4

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("add", func(t *testing.T) {
		got := Multiply(2, 3)
		want := 6

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

}
