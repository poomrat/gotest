package fizzbuzz_test

import (
	"testing"

	fizzbuzz "github.com/poomrat/gotest"
)

func TestShouldSayOne(t *testing.T) {
	// Arrange

	// Action
	result := fizzbuzz.Say(1)
	expected := "1"

	// Assertion
	if result != expected {
		t.Errorf("it should say %q but get %q", expected, result)
	}
}

func TestShouldSayTwo(t *testing.T) {
	// Arrange

	//Action
	result := fizzbuzz.Say(2)

	// Assertion
	expected := "2"
	if result != expected {
		t.Errorf("it should say %q but get %q", expected, result)
	}
}
