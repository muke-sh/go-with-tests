package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}

	got := Perimeter(rectangle)
	want := 40.0

	assert(t, got, want)
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{shape: Rectangle{Width: 12.0, Height: 6.0}, want: 72.0},
		{shape: Circle{Radius: 10.0}, want: 314.1592653589793},
		{shape: Triangle{Base: 12, Height: 6}, want: 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()

		assert(t, got, tt.want)
	}
}

func assert(t testing.TB, got float64, want float64) {

	t.Helper()
	if got != want {
		t.Errorf("got %g want %g", got, want)
	}
}
