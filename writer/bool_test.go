package convwriter

import (
	"bytes"

	"testing"
)

func TestBool(t *testing.T) {

	tests := []struct{
		Value    bool
		Expected string
	}{
		{
			Value:     false,
			Expected: "false",
		},
		{
			Value:     true,
			Expected: "true",
		},
	}


	for testNumber, test := range tests {

		var buffer bytes.Buffer

		n64, err := Bool(&buffer, test.Value)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		if expected, actual := int64(len(test.Expected)), n64; expected != actual {
			t.Errorf("For test #%d, expected %d, but actually got %d.", testNumber, expected, actual)
			continue
		}

		if expected, actual := test.Expected, buffer.String(); expected != actual {
			t.Errorf("For test #%d, expected %q, but actually got %q.", testNumber, expected, actual)
			continue
		}
	}
}
