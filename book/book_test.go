package book

import (
	"testing"
)

func TestToJSON(t *testing.T) {
	actual := Book{"One", "Two", "Three", "Six"}.ToJSON()
	expected := `{"title":"One","author":"Two","isbn":"Three","description":"Six"}`
	if expected != string(actual) {
		t.Errorf("Error occured while testing ToJSON: '%s' != '%s'", expected, actual)
	}
}

func TestFromJSON(t *testing.T) {
	actual := FromJSON([]byte(`{"title":"One","author":"Two","isbn":"Three","description":"Six"}`))
	expected := Book{"One", "Two", "Three", "Six"}

	if actual != expected {
		t.Errorf("Error occured while testing FromJSON: '%s' != '%s'", expected, actual)
	}
}
