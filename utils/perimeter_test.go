package utils

import "testing"

func TestPerimeter(t *testing.T) {

	reactangle := Rectangle{10.0, 10.0}
	got := Perimeter(reactangle)
	want := 40.0

	if got != want {
		t.Errorf("got %g want: %g ", got, want)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{10.0, 20.0}, want: 200}, {name: "Circle", shape: Circle{10}, want: 314.1592653589793},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("got %g want: %g", got, tt.want)
			}

		})
	}

	//	checkArea := func(t testing.TB, shap Shape, want float64) {
	//
	//		t.Helper()
	//		got := shap.Area()
	//		if got != want {
	//			t.Errorf("got: %g want: %g", got, want)
	//		}
	//
	//	}
	//	t.Run("rectangle", func(t *testing.T) {
	//		reactangle := Rectangle{10.0, 20.0}
	//		checkArea(t, reactangle, 200)
	//	})
	//
	//	t.Run("circle", func(t *testing.T) {
	//		circle := Circle{10}
	//		checkArea(t, circle, 314.1592653589793)
	//
	//	})

}
