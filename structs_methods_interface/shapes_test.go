package main

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2g want %.2g", got, want)
	}
}

func TestArea(t *testing.T) {
	rectangle := Rectangle{5.0, 5.0}
	got := Area(rectangle)
	want := 25.0

	if got != want {
		t.Errorf("got %.2g want %.2g", got, want)
	}
}
