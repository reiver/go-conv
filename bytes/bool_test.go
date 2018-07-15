package convbytes

import (
	"testing"
)

func TestBool(t *testing.T) {

	tests := []struct{
		Text   []byte
		Expected bool
	}{
		{
			Text: []byte("f"),
			Expected: false,
		},
		{
			Text: []byte("t"),
			Expected: true,
		},



		{
			Text: []byte("F"),
			Expected: false,
		},
		{
			Text: []byte("T"),
			Expected: true,
		},



		{
			Text: []byte("0"),
			Expected: false,
		},
		{
			Text: []byte("1"),
			Expected: true,
		},



		{
			Text: []byte("٠"), // Arabic-Indic Digit Zero (٠) (\u0660)
			Expected: false,
		},
		{
			Text: []byte("١"), // Arabic-Indic Digit One (١) (\u0661)
			Expected: true,
		},



		{
			Text: []byte("۰"), // Extended Arabic-Indic Digit Zero (۰) (\u06F0)
			Expected: false,
		},
		{
			Text: []byte("۱"), // Extended Arabic-Indic Digit One (۱) (\u06F1)
			Expected: true,
		},



		{
			Text: []byte("⊥"), // Up Tack (⊥) (\u22A5)
			Expected: false,
		},
		{
			Text: []byte("⊤"), // Down Tack (⊤) (\u22A4)
			Expected: true,
		},



		{
			Text: []byte("off"),
			Expected: false,
		},
		{
			Text: []byte("on"),
			Expected: true,
		},



		{
			Text: []byte("OFF"),
			Expected: false,
		},
		{
			Text: []byte("ON"),
			Expected: true,
		},



		{
			Text: []byte("Off"),
			Expected: false,
		},
		{
			Text: []byte("On"),
			Expected: true,
		},



		{
			Text: []byte("oFf"),
			Expected: false,
		},
		{
			Text: []byte("oN"),
			Expected: true,
		},



		{
			Text: []byte("no"),
			Expected: false,
		},
		{
			Text: []byte("yes"),
			Expected: true,
		},



		{
			Text: []byte("NO"),
			Expected: false,
		},
		{
			Text: []byte("YES"),
			Expected: true,
		},



		{
			Text: []byte("No"),
			Expected: false,
		},
		{
			Text: []byte("Yes"),
			Expected: true,
		},



		{
			Text: []byte("nO"),
			Expected: false,
		},
		{
			Text: []byte("yEs"),
			Expected: true,
		},



		{
			Text: []byte("false"),
			Expected: false,
		},
		{
			Text: []byte("true"),
			Expected: true,
		},



		{
			Text: []byte("FALSE"),
			Expected: false,
		},
		{
			Text: []byte("TRUE"),
			Expected: true,
		},



		{
			Text: []byte("False"),
			Expected: false,
		},
		{
			Text: []byte("True"),
			Expected: true,
		},



		{
			Text: []byte("fAlSe"),
			Expected: false,
		},
		{
			Text: []byte("tRuE"),
			Expected: true,
		},
	}


	for testNumber, test := range tests {

		actual, err := Bool(test.Text)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		if expected := test.Expected; expected != actual {
			t.Errorf("For test #%d, expected %t, but actually got %t.", testNumber, expected, actual)
			continue
		}

	}
}

func TestBoolError(t *testing.T) {

	tests := []struct{
		Text []byte
	}{
		{
			Text: nil,
		},



		{
			Text: []byte(""),
		},



		{
			Text: []byte("z"),
		},



		{
			Text: []byte("apple"),
		},
		{
			Text: []byte("banana"),
		},
		{
			Text: []byte("cherry"),
		},



		{
			Text: []byte("fals"),
		},
		{
			Text: []byte("tru"),
		},



		{
			Text: []byte("2"),
		},
		{
			Text: []byte("٢"),
		},
		{
			Text: []byte("۲"),
		},



		{
			Text: []byte("ff"),
		},
		{
			Text: []byte("tt"),
		},



		{
			Text: []byte("false0"),
		},
		{
			Text: []byte("true1"),
		},
	}


	for testNumber, test := range tests {

		_, err := Bool(test.Text)
		if nil == err {
			t.Errorf("For test #%d, expected an error, but did not actually got one: (%T) %v", testNumber, err, err)
			continue
		}
	}
}
