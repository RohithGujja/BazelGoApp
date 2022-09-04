package hello

import "testing"

func TestHello(t *testing.T) {
    expected := "Hello World!!"
    actual := Hello()
	if actual != expected {
	    t.Errorf("expected %s, but got %s", expected, actual)
	}
}