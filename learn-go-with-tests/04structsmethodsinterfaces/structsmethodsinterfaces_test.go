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

/* func TestArea(t *testing.T) {
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
} */

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{shape: Rectangle{Width: 12, Height: 6}, want: 72.10},
		{shape: Circle{Radius: 10}, want: 314.1592653589793},
		{shape: Triangle{Base: 12, Height: 6}, want: 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			// We can change our error message into %#v got %g want %g. The %#v format string will print out our struct with the values in its field, so the developer can see at a glance the properties that are being tested.
			t.Errorf("got %v want %g", got, tt.want)
		}
	}

}
