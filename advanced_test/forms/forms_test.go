package forms

import "testing"

func TestArea(t *testing.T) {

	t.Run("Rectangle", func(t *testing.T) {
		ret := Rectangle{ 10, 12 }
		expectedArea := 120.0
		receivedArea := ret.Area()

		if receivedArea != expectedArea {
			t.Fatalf("Expected %f but received %f", expectedArea, receivedArea)
		}
	})

	t.Run("Circle", func(t *testing.T) {
		cir := Circle{ 10 }
		expectedArea := 314.16
		receivedArea := cir.Area()

		if receivedArea != expectedArea {
			t.Fatalf("Expected %f but received %f", expectedArea, receivedArea)
		}
	})
}
