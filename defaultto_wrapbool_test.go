package fallback_test

import (
	"testing"

	"github.com/reiver/go-fallback"
)

func TestDefaultTo_WrapBool_string(t *testing.T) {

	tests := []struct{
		Value string
		Found bool
		Contingency string
		Expected string
	}{
		{
			Value: "banana",
			Found: false,
			Contingency: "cherry",
			Expected: "cherry",
		},
		{
			Value: "banana",
			Found: true,
			Contingency: "cherry",
			Expected: "banana",
		},
	}

	for testNumber, test := range tests {

		fn := func() (string, bool) {
			return test.Value, test.Found
		}

		actual := fallback.DefaultTo[string]{test.Contingency}.WrapBool(fn())

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual defaulted-value is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			continue
		}
	}
}
