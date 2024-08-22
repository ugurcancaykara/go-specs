package main

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("rectangeles", func(t *testing.T) {
		rectangle := Rectangle{Width: 6.0, Height: 4.0}
		got := rectangle.Perimeter()
		want := 20.0

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})
	t.Run("circles", func(t *testing.T) {
		circle := Circle{Radius: 6.0}
		got := circle.Perimeter()
		want := 37.69911184307752

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})

}

func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{Width: 6.0, Height: 4.0}
		checkArea(t, rectangle, 24.0)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)

	})
}
