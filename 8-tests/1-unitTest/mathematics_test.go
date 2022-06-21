package mathematics

import "testing"

const defaultError = "Expected: %v, but received %v"

func TestAvg(t *testing.T) {
	expected := 7.28
	value := Avg(7.2, 9.9, 6.1, 5.9)

	if value != expected {
		t.Errorf(defaultError, expected, value)
	}
}
