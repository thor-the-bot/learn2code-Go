package helloworld

import "testing"

func TestHelloWorld(t *testing.T) {
	want := "Hello, world!"
	if got := HelloWorld(); got != want {
		t.Fatalf("HelloWorld() = %q; want %q", got, want)
	}
}
