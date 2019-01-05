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

	// Action
	result := fizzbuzz.Say(2)

	// Assertion
	expected := "2"
	if result != expected {
		t.Errorf("it should say %q but get %q", expected, result)
	}
}

func TestShouldSayFizzForThree(t *testing.T) {
	// Arrange

	// Action
	result := fizzbuzz.Say(3)

	// Assertion
	expected := "fizz"
	if result != expected {
		t.Errorf("it should say %q but get %q", expected, result)
	}
}

func TestShouldSayFizzFor99(t *testing.T) {
	// Arrange

	// Action
	result := fizzbuzz.Say(99)

	// Assertion
	expected := "fizz"
	if result != expected {
		t.Errorf("it should say %q but get %q", expected, result)
	}
}

func TestShouldSayBuzzForFive(t *testing.T) {
	// Arrange

	// Action
	result := fizzbuzz.Say(5)

	// Assertion
	expected := "buzz"
	if result != expected {
		t.Errorf("it should say %q but get %q", expected, result)
	}
}

func TestShouldSayBuzzFor100(t *testing.T) {
	// Arrange

	// Action
	result := fizzbuzz.Say(100)

	// Assertion
	expected := "buzz"
	if result != expected {
		t.Errorf("it should say %q but get %q", expected, result)
	}
}

func TestShouldSayFizzBuzzForFifteen(t *testing.T) {
	// Arrange

	// Action
	result := fizzbuzz.Say(15)

	// Assertion
	expected := "fizzbuzz"
	if result != expected {
		t.Errorf("it should say %q but get %q", expected, result)
	}
}

func TestShouldSayFizzBuzzFor150(t *testing.T) {
	// Arrange

	// Action
	result := fizzbuzz.Say(150)

	// Assertion
	expected := "fizzbuzz"
	if result != expected {
		t.Errorf("it should say %q but get %q", expected, result)
	}
}
