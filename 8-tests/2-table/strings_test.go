package strings

import (
	"strings"
	"testing"
)

const defaultError = "%s (part: %s) - indexes: wanted (%d) <> got (%d)"

func TestIndex(t *testing.T) {
	tests := []struct {
		text   string
		part   string
		wanted int
	}{
		{"Linkin Park is nice", "Linkin", 0},
		{"", "", 0},
		{"Hello", "hello", -1},
		{"Goodbye", "o", 1},
	}

	for _, test := range tests {
		t.Logf("Text %v", test)
		current := strings.Index(test.text, test.part)
		if current != test.wanted {
			t.Errorf(defaultError, test.text, test.part, test.wanted, current)
		}
	}
}
