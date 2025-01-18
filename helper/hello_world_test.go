package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Rizqi")

	if result != "Hello, Rizqi" {
		// failed
		panic("Result is not Hello, Rizqi")
	}
}
