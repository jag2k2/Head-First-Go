package prose

import (
	"fmt"
	"proseproject/src/prose"
	"testing"
)

func TestTwoElements(t *testing.T) {
	list := []string{"apple", "orange"}
	got := prose.JoinWithCommas(list)
	want := "apple and orange"
	if got != want {
		t.Errorf(errorString(list, got, want))
	}
}

func TestThreeElements(t *testing.T) {
	list := []string{"apple", "orange", "pear"}
	got := prose.JoinWithCommas(list)
	want := "apple, orange, and pear"
	if got != want {
		t.Errorf(errorString(list, got, want))
	}
}

func errorString(list []string, got string, want string) string {
	return fmt.Sprintf("JoinWithCommas(%#v) =  \"%s\", want \"%s\"", list, got, want)
}
