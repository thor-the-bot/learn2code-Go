package n4

import "testing"

func TestMax(t *testing.T) {
	if got := Max(2, 3); got != 3 {
		t.Fatalf("Max(2,3) = %d; want 3", got)
	}
	if got := Max(5, -1); got != 5 {
		t.Fatalf("Max(5,-1) = %d; want 5", got)
	}
	if got := Max(7,7); got != 7 {
		t.Fatalf("Max(7,7) = %d; want 7", got)
	}
}
