package convbytes

import (
	"fmt"
	"unicode/utf8"
)

func Int64(data []byte) (int64, error) {

	if 0 >= len(data) {
		return 0, InterpretationError{
			value: string(data),
			reason: "invalid int64 literal",
		}
	}

	var negative bool
	var p      []byte = data

	{
		r0,size := utf8.DecodeRune(data)
		if utf8.RuneError == r0 {
			return 0, InterpretationError{
				value: string(data),
				reason: "invalid int64 literal",
			}
		}

		switch r0 {
		case '+',  // PLUS SIGN (U+002B)
		     '⁺',  // SUPERSCRIPT PLUS SIGN (U+207A)
		     '＋': // FULLWIDTH PLUS SIGN (U+FF0B)
			p = p[size:]
		case '-',  // HYPHEN-MINUS (U+002D)
		     '⁻',  // SUPERSCRIPT MINUS (U+207B)
		     '−',  // MINUS SIGN (U+2212)
		     '﹣', // SMALL HYPHEN-MINUS (U+FE63)
		     '－': // FULLWIDTH HYPHEN-MINUS (U+FF0D)
			p = p[size:]
			negative = true
		default:
		}
	}

	var i64 int64
	{
		ui64, err := Uint64(p)
		if nil != err {
			switch err.(type) {
			case InterpretationError:
				return 0, InterpretationError{
					value: string(data),
					reason: "invalid int64 literal",
				}

			default:
				return 0, err
			}
		}

		switch negative {
		case false:

			const maxInt64 uint64 = 1<<63 - 1

			if maxInt64 < ui64 {
				return 0, Overflow{
					value: string(data),
					reason: fmt.Sprintf("maximum int64 is %d, %d is outside of that range", maxInt64, ui64),
				}
			}


			i64 = int64(ui64)

		default: // case true

			const minInt64Abs uint64 = -1*(-1 << 63)

			if minInt64Abs < ui64 {
				return 0, Overflow{
					value: string(data),
					reason: fmt.Sprintf("minimum int64 is −%d, −%d is outside of that range", minInt64Abs, ui64),
				}
			}

			i64 = -1 * int64(ui64)
		}
	}

	return i64, nil
}
