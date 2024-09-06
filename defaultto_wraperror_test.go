package fallback_test

import (
	"testing"

	"errors"

	"github.com/reiver/go-fallback"
)

func TestDefaultTo_WrapError_string(t *testing.T) {

	tests := []struct{
		Value string
		Err error
		Contingency string
		Expected string
	}{
		{
			Value: "banana",
			Err: errors.New("oh no!"),
			Contingency: "cherry",
			Expected: "cherry",
		},
		{
			Value: "banana",
			Err: nil,
			Contingency: "cherry",
			Expected: "banana",
		},
	}

	for testNumber, test := range tests {

		fn := func() (string, error) {
			return test.Value, test.Err
		}

		actual := fallback.DefaultTo[string]{test.Contingency}.WrapError(fn())

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual defaulted-value is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			continue
		}
	}
}
