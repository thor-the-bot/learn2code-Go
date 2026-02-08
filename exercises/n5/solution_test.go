package n5

import "testing"

func TestIsEven(t *testing.T) {
	if !IsEven(2) {
		t.Fatalf("IsEven(2) = false; want true")
	}
	if IsEven(3) {
		t.Fatalf("IsEven(3) = true; want false")
	}
	if !IsEven(0) {
		t.Fatalf("IsEven(0) = false; want true")
	}
}
