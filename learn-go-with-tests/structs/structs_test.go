package structs

import "testing"

func assertError(t *testing.T, shape Shape, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("%#v Got %.2f, want %.2f", shape, got, want)
	}
}

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0
	assertError(t, rectangle, got, want)
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name string
		shape Shape
		hasArea float64
	} {
		{name: "Rectangles", shape: Rectangle{Height: 12.0, Width: 6.0}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10.0}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	checkArea := func(t *testing.T, shape Shape, want float64) {
		got := shape.Area()
		assertError(t, shape, got, want)
	}

	for _, test := range areaTests {
		t.Run(test.name, func(t *testing.T) {
			checkArea(t, test.shape, test.hasArea)
		})
	}
}