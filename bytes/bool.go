package convbytes

import (
	"bytes"
)

var (
	boolLiterals = []struct{
		Literal []byte
		Value     bool
	}{
		{
			Literal: []byte{'f','a','l','s','e'},
			Value:   false,
		},
		{
			Literal: []byte{'t','r','u','e'},
			Value:   true,
		},



		{
			Literal: []byte{'f'},
			Value:   false,
		},
		{
			Literal: []byte{'t'},
			Value:   true,
		},



		{
			Literal: []byte{'n','o'},
			Value:   false,
		},
		{
			Literal: []byte{'y','e','s'},
			Value:   true,
		},



		{
			Literal: []byte{'o','f','f'},
			Value:   false,
		},
		{
			Literal: []byte{'o','n'},
			Value:   true,
		},



		{
			Literal: []byte{'\u0030'}, // Digit Zero (0)
			Value:   false,
		},
		{
			Literal: []byte{'\u0031'}, // Digit One (1)
			Value:   true,
		},



		{
			Literal: []byte("\u0660"), // Arabic-Indic Digit Zero (٠)
			Value:   false,
		},
		{
			Literal: []byte("\u0661"), // Arabic-Indic Digit One (١)
			Value:   true,
		},



		{
			Literal: []byte("\u06F0"), // Extended Arabic-Indic Digit Zero (۰)
			Value:   false,
		},
		{
			Literal: []byte("\u06F1"), // Extended Arabic-Indic Digit One (۱)
			Value:   true,
		},



		{
			Literal: []byte("\u22A5"), // Up Tack (⊥)
			Value:   false,
		},
		{
			Literal: []byte("\u22A4"), // Down Tack (⊤)
			Value:   true,
		},
	}
)

// Bool is parses UTF-8 text, interprets it as a bool literal, and tries to return a bool.
//
// Supported literals include:
//
// • false
//
// • true
//
//
// • f
//
// • t
//
//
// • no
//
// • yes
//
//
// • off
//
// • on
//
//
// • 0
//
// • 1
//
//
// • ٠
//
// • ١
//
//
// • ۰
//
// • ۱
//
//
// • ⊥
//
// • ⊤
//
// (These are all case-insensitive, so, for example, “TRUE”, “true”, “True”, etc, all work.)
func Bool(data []byte) (bool, error) {

	for _, boolLiteral := range boolLiterals {

		if bytes.EqualFold(data, boolLiteral.Literal) {
			return boolLiteral.Value, nil
		}

	}

	return false, InterpretationError{
		value: string(data),
		reason: "invalid bool literal",
	}
}
