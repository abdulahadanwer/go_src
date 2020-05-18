package prose

import (
	"fmt"
	"testing"
)

func TestForOneElement(t *testing.T) {
	list := []string{"apple"}
	want := "apple"
	got := JoinWithCommas(list)
	if want != got {
		t.Errorf(errorString(list, got, want))
	}

}

func TestForTwoElements(t *testing.T) {
	list := []string{"apple", "orange"}
	want := "apple and orange"
	got := JoinWithCommas(list)
	if want != got {
		t.Errorf(errorString(list, got, want))
	}

}

func TestForThreeElements(t *testing.T) {
	list := []string{"apple", "orange", "pear"}
	want := "apple, orange, and pear"
	got := JoinWithCommas(list)
	if JoinWithCommas(list) != "apple, orange, and pear" {
		t.Errorf(errorString(list, got, want))
	}
}

func errorString(list []string, got string, want string) string {
	return fmt.Sprintf("JoinWithCommas(%#v) = \"%s\", want \"%s\"", list, got, want)
}
