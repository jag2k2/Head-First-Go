package greetings

import (
	"myproject/src/greetings"
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Gladys"
	expected := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := greetings.Hello("Gladys")
	if !expected.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, expected)
	}
}
