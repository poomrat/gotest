package fizzbuzz_test

import (
	"testing"

	fizzbuzz "github.com/poomrat/gotest"
)

func TestShouldSayOne(t *testing.T) {
	result := fizzbuzz.Say(1)
	expected := "1"
	if result != expected {
		t.Errorf("it should say %q but get %q", expected, result)
	}
}

