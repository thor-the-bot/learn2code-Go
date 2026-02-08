package n3

import "testing"

func TestVariablesDemo(t *testing.T) {
	i, s, b := VariablesDemo()
	// Expectations: some sensible non-zero/example values
	if i != 42 {
		t.Fatalf("expected int 42, got %d", i)
	}
	if s != "gopher" {
		t.Fatalf("expected string 'gopher', got %q", s)
	}
	if b != true {
		t.Fatalf("expected bool true, got %v", b)
	}
}
