package jusulsa

import "testing"

func Test1(t *testing.T) {
	if 123456789 != removeChar(" 123,456,789", "") {
		t.Error("removeChar(123) error.")
	}
	if 123 != removeChar(" 123", "") {
		t.Error("removeChar(123) error.")
	}
	if 0 != removeChar(" ", "") {
		t.Error("removeChar(' ') error.")
	}
}
