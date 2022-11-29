package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}

	got := Perimeter(rectangle)
	want := 40.0

	assert(t, got, want)
}

func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()

		got := shape.Area()
		assert(t, got, want)
	}

	t.Run("rectangle", func(t *testing.T) {

		rectangle := Rectangle{12.0, 6.0}
		checkArea(t, rectangle, 72.0)
	})

	t.Run("circle", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589792)
	})

}

func assert(t testing.TB, got float64, want float64) {

	t.Helper()
	if got != want {
		t.Errorf("got %g want %g", got, want)
	}
}
