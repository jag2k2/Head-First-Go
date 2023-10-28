package greetings

import "testing"
import "regexp"

func TestHelloName(t *testing.T) {
	name := "Gladys"
	expected := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Gladys2")
	if !expected.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, expected)
	}
}