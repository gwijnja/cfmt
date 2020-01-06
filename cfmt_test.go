package cfmt

import (
	"testing"
)

func TestSprintf(t *testing.T) {
	var got, want string

	got = Sprintf(0, 0, "")
	want = "\x1b[0m\x1b[10m\x1b[0m"
	if got != want {
		t.Errorf("Sprintf() name = Empty, want = %v, got = %v", want, got)
	}

	got = Sprintf(1, 2, "Foo")
	want = "\x1b[1m\x1b[12mFoo\x1b[0m"
	if got != want {
		t.Errorf("Sprintf() name = No args, want = %v, got = %v", want, got)
	}

	got = Sprintf(3, 4, "Foo %d and %d", 5, 6)
	want = "\x1b[3m\x1b[14mFoo 5 and 6\x1b[0m"
	if got != want {
		t.Errorf("Sprintf() name = Two args, want = %v, got = %v", want, got)
	}
}
