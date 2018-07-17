package convwriter

import (
	"bytes"

	"testing"
)

func TestBool(t *testing.T) {

	tests := []struct{
		Params   []interface{}
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









		{
			Params: []interface{}{"FALSE"},
			Value:                 false,
			Expected:             "FALSE",
		},
		{
			Params: []interface{}{"FALSE"},
			Value:                 true,
			Expected:             "TRUE",
		},



		{
			Params: []interface{}{"TRUE"},
			Value:                 false,
			Expected:             "FALSE",
		},
		{
			Params: []interface{}{"TRUE"},
			Value:                 true,
			Expected:             "TRUE",
		},









		{
			Params: []interface{}{"false"},
			Value:                 false,
			Expected:             "false",
		},
		{
			Params: []interface{}{"false"},
			Value:                 true,
			Expected:             "true",
		},



		{
			Params: []interface{}{"true"},
			Value:                 false,
			Expected:             "false",
		},
		{
			Params: []interface{}{"true"},
			Value:                 true,
			Expected:             "true",
		},









		{
			Params: []interface{}{"False"},
			Value:                 false,
			Expected:             "False",
		},
		{
			Params: []interface{}{"False"},
			Value:                 true,
			Expected:             "True",
		},



		{
			Params: []interface{}{"True"},
			Value:                 false,
			Expected:             "False",
		},
		{
			Params: []interface{}{"True"},
			Value:                 true,
			Expected:             "True",
		},









		{
			Params: []interface{}{'0'},
			Value:                 false,
			Expected:             "0",
		},
		{
			Params: []interface{}{'0'},
			Value:                 true,
			Expected:             "1",
		},



		{
			Params: []interface{}{'1'},
			Value:                 false,
			Expected:             "0",
		},
		{
			Params: []interface{}{'1'},
			Value:                 true,
			Expected:             "1",
		},









		{
			Params: []interface{}{'٠'}, // Arabic-Indic Digit Zero  (٠)
			Value:                 false,
			Expected:             "٠",  // Arabic-Indic Digit Zero  (٠)
		},
		{
			Params: []interface{}{'٠'}, // Arabic-Indic Digit Zero  (٠)
			Value:                 true,
			Expected:             "١",  // Arabic-Indic Digit One   (١)
		},



		{
			Params: []interface{}{'١'}, // Arabic-Indic Digit One   (١)
			Value:                 false,
			Expected:             "٠",  // Arabic-Indic Digit Zero  (٠)
		},
		{
			Params: []interface{}{'١'}, // Arabic-Indic Digit One   (١)
			Value:                 true,
			Expected:             "١",  // Arabic-Indic Digit One   (١)
		},









		{
			Params: []interface{}{'۰'}, // Extended Arabic-Indic Digit Zero  (۰)
			Value:                 false,
			Expected:             "۰",  // Extended Arabic-Indic Digit Zero  (۰)
		},
		{
			Params: []interface{}{'۰'}, // Extended Arabic-Indic Digit Zero  (۰)
			Value:                 true,
			Expected:             "۱",  // Extended Arabic-Indic Digit One   (۱)
		},



		{
			Params: []interface{}{'۱'}, // Extended Arabic-Indic Digit One   (۱)
			Value:                 false,
			Expected:             "۰",  // Extended Arabic-Indic Digit Zero  (۰)
		},
		{
			Params: []interface{}{'۱'}, // Extended Arabic-Indic Digit One   (۱)
			Value:                 true,
			Expected:             "۱",  // Extended Arabic-Indic Digit One   (۱)
		},









		{
			Params: []interface{}{'F'},
			Value:                 false,
			Expected:             "F",
		},
		{
			Params: []interface{}{'F'},
			Value:                 true,
			Expected:             "T",
		},



		{
			Params: []interface{}{'T'},
			Value:                 false,
			Expected:             "F",
		},
		{
			Params: []interface{}{'T'},
			Value:                 true,
			Expected:             "T",
		},









		{
			Params: []interface{}{'f'},
			Value:                 false,
			Expected:             "f",
		},
		{
			Params: []interface{}{'f'},
			Value:                 true,
			Expected:             "t",
		},



		{
			Params: []interface{}{'t'},
			Value:                 false,
			Expected:             "f",
		},
		{
			Params: []interface{}{'t'},
			Value:                 true,
			Expected:             "t",
		},









		{
			Params: []interface{}{'⊥'},
			Value:                 false,
			Expected:             "⊥",
		},
		{
			Params: []interface{}{'⊥'},
			Value:                 true,
			Expected:             "⊤",
		},



		{
			Params: []interface{}{'⊤'},
			Value:                 false,
			Expected:             "⊥",
		},
		{
			Params: []interface{}{'⊤'},
			Value:                 true,
			Expected:             "⊤",
		},









		{
			Params: []interface{}{"OFF"},
			Value:                 false,
			Expected:             "OFF",
		},
		{
			Params: []interface{}{"OFF"},
			Value:                 true,
			Expected:             "ON",
		},



		{
			Params: []interface{}{"ON"},
			Value:                 false,
			Expected:             "OFF",
		},
		{
			Params: []interface{}{"ON"},
			Value:                 true,
			Expected:             "ON",
		},









		{
			Params: []interface{}{"off"},
			Value:                 false,
			Expected:             "off",
		},
		{
			Params: []interface{}{"off"},
			Value:                 true,
			Expected:             "on",
		},



		{
			Params: []interface{}{"on"},
			Value:                 false,
			Expected:             "off",
		},
		{
			Params: []interface{}{"on"},
			Value:                 true,
			Expected:             "on",
		},









		{
			Params: []interface{}{"Off"},
			Value:                 false,
			Expected:             "Off",
		},
		{
			Params: []interface{}{"Off"},
			Value:                 true,
			Expected:             "On",
		},



		{
			Params: []interface{}{"On"},
			Value:                 false,
			Expected:             "Off",
		},
		{
			Params: []interface{}{"On"},
			Value:                 true,
			Expected:             "On",
		},









		{
			Params: []interface{}{"NO"},
			Value:                 false,
			Expected:             "NO",
		},
		{
			Params: []interface{}{"NO"},
			Value:                 true,
			Expected:             "YES",
		},



		{
			Params: []interface{}{"YES"},
			Value:                 false,
			Expected:             "NO",
		},
		{
			Params: []interface{}{"YES"},
			Value:                 true,
			Expected:             "YES",
		},









		{
			Params: []interface{}{"no"},
			Value:                 false,
			Expected:             "no",
		},
		{
			Params: []interface{}{"no"},
			Value:                 true,
			Expected:             "yes",
		},



		{
			Params: []interface{}{"yes"},
			Value:                 false,
			Expected:             "no",
		},
		{
			Params: []interface{}{"yes"},
			Value:                 true,
			Expected:             "yes",
		},









		{
			Params: []interface{}{"No"},
			Value:                 false,
			Expected:             "No",
		},
		{
			Params: []interface{}{"No"},
			Value:                 true,
			Expected:             "Yes",
		},



		{
			Params: []interface{}{"Yes"},
			Value:                 false,
			Expected:             "No",
		},
		{
			Params: []interface{}{"Yes"},
			Value:                 true,
			Expected:             "Yes",
		},
	}


	for testNumber, test := range tests {

		var buffer bytes.Buffer

		n64, err := Bool(&buffer, test.Value, test.Params...)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		if expected, actual := test.Expected, buffer.String(); expected != actual {
			t.Errorf("For test #%d, expected %q, but actually got %q.", testNumber, expected, actual)
			continue
		}

		if expected, actual := int64(len(test.Expected)), n64; expected != actual {
			t.Errorf("For test #%d, expected %d, but actually got %d.", testNumber, expected, actual)
			continue
		}
	}
}
