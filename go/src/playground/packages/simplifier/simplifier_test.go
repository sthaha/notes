package simplifier

import (
	"regexp"
	"testing"
)

func TestSimplify(t *testing.T) {
	onlyLower := regexp.MustCompile("[a-z ]")

	simplified := Simplify("Hello world!!")
	if !onlyLower.MatchString(simplified) {
		t.Fatalf("Found an error")
	}

	if Simplify("Hello world!!") != "hello world" {
		t.Fatalf("Found an error")
	}
}
