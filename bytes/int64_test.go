package convbytes

import (
	"testing"
)

func TestInt64(t *testing.T) {

	tests := []struct{
		Text   []byte
		Expected int64
	}{
		{
			Text: []byte("0"),
			Expected:     0,
		},
		{
			Text: []byte("-0"), // HYPHEN-MINUS (U+002D)
			Expected:      0,
		},
		{
			Text: []byte("⁻0"), // SUPERSCRIPT MINUS (U+207B)
			Expected:      0,
		},
		{
			Text: []byte("−0"), // MINUS SIGN (U+2212)
			Expected:      0,
		},
		{
			Text: []byte("﹣0"), // SMALL HYPHEN-MINUS (U+FE63)
			Expected:       0,
		},
		{
			Text: []byte("－0"), // FULLWIDTH HYPHEN-MINUS (U+FF0D)
			Expected:       0,
		},
		{
			Text: []byte("+0"), // PLUS SIGN (U+002B)
			Expected:      0,
		},
		{
			Text: []byte("⁺0"), // SUPERSCRIPT PLUS SIGN (U+207A)
			Expected:      0,
		},
		{
			Text: []byte("＋0"), // FULLWIDTH PLUS SIGN (U+FF0B)
			Expected:      0,
		},



		{
			Text: []byte("1"),
			Expected:     1,
		},
		{
			Text: []byte("2"),
			Expected:     2,
		},
		{
			Text: []byte("3"),
			Expected:     3,
		},
		{
			Text: []byte("4"),
			Expected:     4,
		},
		{
			Text: []byte("5"),
			Expected:     5,
		},
		{
			Text: []byte("6"),
			Expected:     6,
		},
		{
			Text: []byte("7"),
			Expected:     7,
		},
		{
			Text: []byte("8"),
			Expected:     8,
		},
		{
			Text: []byte("9"),
			Expected:     9,
		},



		{
			Text: []byte("-1"), // HYPHEN-MINUS (U+002D)
			Expected:     -1,
		},
		{
			Text: []byte("-2"), // HYPHEN-MINUS (U+002D)
			Expected:     -2,
		},
		{
			Text: []byte("-3"), // HYPHEN-MINUS (U+002D)
			Expected:     -3,
		},
		{
			Text: []byte("-4"), // HYPHEN-MINUS (U+002D)
			Expected:     -4,
		},
		{
			Text: []byte("-5"), // HYPHEN-MINUS (U+002D)
			Expected:     -5,
		},
		{
			Text: []byte("-6"), // HYPHEN-MINUS (U+002D)
			Expected:     -6,
		},
		{
			Text: []byte("-7"), // HYPHEN-MINUS (U+002D)
			Expected:     -7,
		},
		{
			Text: []byte("-8"), // HYPHEN-MINUS (U+002D)
			Expected:     -8,
		},
		{
			Text: []byte("-9"), // HYPHEN-MINUS (U+002D)
			Expected:     -9,
		},



		{
			Text: []byte("⁻1"), // SUPERSCRIPT MINUS (U+207B)
			Expected:     -1,
		},
		{
			Text: []byte("⁻2"), // SUPERSCRIPT MINUS (U+207B)
			Expected:     -2,
		},
		{
			Text: []byte("⁻3"), // SUPERSCRIPT MINUS (U+207B)
			Expected:     -3,
		},
		{
			Text: []byte("⁻4"), // SUPERSCRIPT MINUS (U+207B)
			Expected:     -4,
		},
		{
			Text: []byte("⁻5"), // SUPERSCRIPT MINUS (U+207B)
			Expected:     -5,
		},
		{
			Text: []byte("⁻6"), // SUPERSCRIPT MINUS (U+207B)
			Expected:     -6,
		},
		{
			Text: []byte("⁻7"), // SUPERSCRIPT MINUS (U+207B)
			Expected:     -7,
		},
		{
			Text: []byte("⁻8"), // SUPERSCRIPT MINUS (U+207B)
			Expected:     -8,
		},
		{
			Text: []byte("⁻9"), // SUPERSCRIPT MINUS (U+207B)
			Expected:     -9,
		},



		{
			Text: []byte("−1"), // MINUS SIGN (U+2212)
			Expected:     -1,
		},
		{
			Text: []byte("−2"), // MINUS SIGN (U+2212)
			Expected:     -2,
		},
		{
			Text: []byte("−3"), // MINUS SIGN (U+2212)
			Expected:     -3,
		},
		{
			Text: []byte("−4"), // MINUS SIGN (U+2212)
			Expected:     -4,
		},
		{
			Text: []byte("−5"), // MINUS SIGN (U+2212)
			Expected:     -5,
		},
		{
			Text: []byte("−6"), // MINUS SIGN (U+2212)
			Expected:     -6,
		},
		{
			Text: []byte("−7"), // MINUS SIGN (U+2212)
			Expected:     -7,
		},
		{
			Text: []byte("−8"), // MINUS SIGN (U+2212)
			Expected:     -8,
		},
		{
			Text: []byte("−9"), // MINUS SIGN (U+2212)
			Expected:     -9,
		},



		{
			Text: []byte("﹣1"), // SMALL HYPHEN-MINUS (U+FE63)
			Expected:      -1,
		},
		{
			Text: []byte("﹣2"), // SMALL HYPHEN-MINUS (U+FE63)
			Expected:      -2,
		},
		{
			Text: []byte("﹣3"), // SMALL HYPHEN-MINUS (U+FE63)
			Expected:      -3,
		},
		{
			Text: []byte("﹣4"), // SMALL HYPHEN-MINUS (U+FE63)
			Expected:      -4,
		},
		{
			Text: []byte("﹣5"), // SMALL HYPHEN-MINUS (U+FE63)
			Expected:      -5,
		},
		{
			Text: []byte("﹣6"), // SMALL HYPHEN-MINUS (U+FE63)
			Expected:      -6,
		},
		{
			Text: []byte("﹣7"), // SMALL HYPHEN-MINUS (U+FE63)
			Expected:      -7,
		},
		{
			Text: []byte("﹣8"), // SMALL HYPHEN-MINUS (U+FE63)
			Expected:      -8,
		},
		{
			Text: []byte("﹣9"), // SMALL HYPHEN-MINUS (U+FE63)
			Expected:      -9,
		},



		{
			Text: []byte("－1"), // FULLWIDTH HYPHEN-MINUS (U+FF0D)
			Expected:      -1,
		},
		{
			Text: []byte("－2"), // FULLWIDTH HYPHEN-MINUS (U+FF0D)
			Expected:      -2,
		},
		{
			Text: []byte("－3"), // FULLWIDTH HYPHEN-MINUS (U+FF0D)
			Expected:      -3,
		},
		{
			Text: []byte("－4"), // FULLWIDTH HYPHEN-MINUS (U+FF0D)
			Expected:      -4,
		},
		{
			Text: []byte("－5"), // FULLWIDTH HYPHEN-MINUS (U+FF0D)
			Expected:      -5,
		},
		{
			Text: []byte("－6"), // FULLWIDTH HYPHEN-MINUS (U+FF0D)
			Expected:      -6,
		},
		{
			Text: []byte("－7"), // FULLWIDTH HYPHEN-MINUS (U+FF0D)
			Expected:      -7,
		},
		{
			Text: []byte("－8"), // FULLWIDTH HYPHEN-MINUS (U+FF0D)
			Expected:      -8,
		},
		{
			Text: []byte("－9"), // FULLWIDTH HYPHEN-MINUS (U+FF0D)
			Expected:      -9,
		},



		{
			Text: []byte("+1"), // PLUS SIGN (U+002B)
			Expected:      1,
		},
		{
			Text: []byte("+2"), // PLUS SIGN (U+002B)
			Expected:      2,
		},
		{
			Text: []byte("+3"), // PLUS SIGN (U+002B)
			Expected:      3,
		},
		{
			Text: []byte("+4"), // PLUS SIGN (U+002B)
			Expected:      4,
		},
		{
			Text: []byte("+5"), // PLUS SIGN (U+002B)
			Expected:      5,
		},
		{
			Text: []byte("+6"), // PLUS SIGN (U+002B)
			Expected:      6,
		},
		{
			Text: []byte("+7"), // PLUS SIGN (U+002B)
			Expected:      7,
		},
		{
			Text: []byte("+8"), // PLUS SIGN (U+002B)
			Expected:      8,
		},
		{
			Text: []byte("+9"), // PLUS SIGN (U+002B)
			Expected:      9,
		},



		{
			Text: []byte("⁺1"), // SUPERSCRIPT PLUS SIGN (U+207A)
			Expected:      1,
		},
		{
			Text: []byte("⁺2"), // SUPERSCRIPT PLUS SIGN (U+207A)
			Expected:      2,
		},
		{
			Text: []byte("⁺3"), // SUPERSCRIPT PLUS SIGN (U+207A)
			Expected:      3,
		},
		{
			Text: []byte("⁺4"), // SUPERSCRIPT PLUS SIGN (U+207A)
			Expected:      4,
		},
		{
			Text: []byte("⁺5"), // SUPERSCRIPT PLUS SIGN (U+207A)
			Expected:      5,
		},
		{
			Text: []byte("⁺6"), // SUPERSCRIPT PLUS SIGN (U+207A)
			Expected:      6,
		},
		{
			Text: []byte("⁺7"), // SUPERSCRIPT PLUS SIGN (U+207A)
			Expected:      7,
		},
		{
			Text: []byte("⁺8"), // SUPERSCRIPT PLUS SIGN (U+207A)
			Expected:      8,
		},
		{
			Text: []byte("⁺9"), // SUPERSCRIPT PLUS SIGN (U+207A)
			Expected:      9,
		},



		{
			Text: []byte("＋1"), // FULLWIDTH PLUS SIGN (U+FF0B)
			Expected:       1,
		},
		{
			Text: []byte("＋2"), // FULLWIDTH PLUS SIGN (U+FF0B)
			Expected:       2,
		},
		{
			Text: []byte("＋3"), // FULLWIDTH PLUS SIGN (U+FF0B)
			Expected:       3,
		},
		{
			Text: []byte("＋4"), // FULLWIDTH PLUS SIGN (U+FF0B)
			Expected:       4,
		},
		{
			Text: []byte("＋5"), // FULLWIDTH PLUS SIGN (U+FF0B)
			Expected:       5,
		},
		{
			Text: []byte("＋6"), // FULLWIDTH PLUS SIGN (U+FF0B)
			Expected:       6,
		},
		{
			Text: []byte("＋7"), // FULLWIDTH PLUS SIGN (U+FF0B)
			Expected:       7,
		},
		{
			Text: []byte("＋8"), // FULLWIDTH PLUS SIGN (U+FF0B)
			Expected:       8,
		},
		{
			Text: []byte("＋9"), // FULLWIDTH PLUS SIGN (U+FF0B)
			Expected:       9,
		},



		{
			Text:  []byte("10"),
			Expected:      10,
		},
		{
			Text: []byte("-10"), // HYPHEN-MINUS (U+002D)
			Expected:     -10,
		},
		{
			Text: []byte("⁻10"), // SUPERSCRIPT MINUS (U+207B)
			Expected:     -10,
		},
		{
			Text: []byte("−10"), // MINUS SIGN (U+2212)
			Expected:     -10,
		},
		{
			Text: []byte("﹣10"), // SMALL HYPHEN-MINUS (U+FE63)
			Expected:      -10,
		},
		{
			Text: []byte("－10"), // FULLWIDTH HYPHEN-MINUS (U+FF0D)
			Expected:      -10,
		},
		{
			Text: []byte("+10"),
			Expected:      10,
		},
		{
			Text: []byte("⁺10"), // SUPERSCRIPT PLUS SIGN (U+207A)
			Expected:      10,
		},
		{
			Text: []byte("＋10"), // FULLWIDTH PLUS SIGN (U+FF0B)
			Expected:       10,
		},



		{
			Text:  []byte("21"),
			Expected:      21,
		},
		{
			Text: []byte("-21"), // HYPHEN-MINUS (U+002D)
			Expected:     -21,
		},
		{
			Text: []byte("⁻21"), // SUPERSCRIPT MINUS (U+207B)
			Expected:     -21,
		},
		{
			Text: []byte("−21"), // MINUS SIGN (U+2212)
			Expected:     -21,
		},
		{
			Text: []byte("﹣21"), // SMALL HYPHEN-MINUS (U+FE63)
			Expected:      -21,
		},
		{
			Text: []byte("－21"), // FULLWIDTH HYPHEN-MINUS (U+FF0D)
			Expected:      -21,
		},
		{
			Text: []byte("+21"), // PLUS SIGN (U+002B)
			Expected:      21,
		},
		{
			Text: []byte("⁺21"), // SUPERSCRIPT PLUS SIGN (U+207A)
			Expected:      21,
		},
		{
			Text: []byte("＋21"), // FULLWIDTH PLUS SIGN (U+FF0B)
			Expected:       21,
		},



		{
			Text:  []byte("32"),
			Expected:      32,
		},
		{
			Text: []byte("-32"), // HYPHEN-MINUS (U+002D)
			Expected:     -32,
		},
		{
			Text: []byte("⁻32"), // SUPERSCRIPT MINUS (U+207B)
			Expected:     -32,
		},
		{
			Text: []byte("−32"), // MINUS SIGN (U+2212)
			Expected:     -32,
		},
		{
			Text: []byte("﹣32"), // SMALL HYPHEN-MINUS (U+FE63)
			Expected:      -32,
		},
		{
			Text: []byte("－32"), // FULLWIDTH HYPHEN-MINUS (U+FF0D)
			Expected:      -32,
		},
		{
			Text: []byte("+32"), // PLUS SIGN (U+002B)
			Expected:      32,
		},
		{
			Text: []byte("⁺32"), // SUPERSCRIPT PLUS SIGN (U+207A)
			Expected:      32,
		},
		{
			Text: []byte("＋32"), // FULLWIDTH PLUS SIGN (U+FF0B)
			Expected:       32,
		},



		{
			Text:  []byte("12345"),
			Expected:      12345,
		},
		{
			Text: []byte("-12345"), // HYPHEN-MINUS (U+002D)
			Expected:     -12345,
		},
		{
			Text: []byte("⁻12345"), // SUPERSCRIPT MINUS (U+207B)
			Expected:     -12345,
		},
		{
			Text: []byte("−12345"), // MINUS SIGN (U+2212)
			Expected:     -12345,
		},
		{
			Text: []byte("﹣12345"), // SMALL HYPHEN-MINUS (U+FE63)
			Expected:      -12345,
		},
		{
			Text: []byte("－12345"), // FULLWIDTH HYPHEN-MINUS (U+FF0D)
			Expected:      -12345,
		},
		{
			Text: []byte("+12345"), // PLUS SIGN (U+002B)
			Expected:      12345,
		},
		{
			Text: []byte("⁺12345"), // SUPERSCRIPT PLUS SIGN (U+207A)
			Expected:      12345,
		},
		{
			Text: []byte("＋12345"), // FULLWIDTH PLUS SIGN (U+FF0B)
			Expected:       12345,
		},



		{
			Text:  []byte("9876543210"),
			Expected:      9876543210,
		},
		{
			Text: []byte("-9876543210"), // HYPHEN-MINUS (U+002D)
			Expected:     -9876543210,
		},
		{
			Text: []byte("⁻9876543210"), // SUPERSCRIPT MINUS (U+207B)
			Expected:     -9876543210,
		},
		{
			Text: []byte("−9876543210"), // MINUS SIGN (U+2212)
			Expected:      9876543210,
		},
		{
			Text: []byte("﹣9876543210"), // SMALL HYPHEN-MINUS (U+FE63)
			Expected:      -9876543210,
		},
		{
			Text: []byte("－9876543210"), // FULLWIDTH HYPHEN-MINUS (U+FF0D)
			Expected:      -9876543210,
		},
		{
			Text: []byte("+9876543210"), // PLUS SIGN (U+002B)
			Expected:      9876543210,
		},
		{
			Text: []byte("⁺9876543210"), // SUPERSCRIPT PLUS SIGN (U+207A)
			Expected:      9876543210,
		},
		{
			Text: []byte("＋9876543210"), // FULLWIDTH PLUS SIGN (U+FF0B)
			Expected:       9876543210,
		},



		{
			Text:  []byte("9223372036854775807"),
			Expected:      9223372036854775807,
		},
		{
			Text: []byte("-9223372036854775807"), // HYPHEN-MINUS (U+002D)
			Expected:     -9223372036854775807,
		},
		{
			Text: []byte("⁻9223372036854775807"), // SUPERSCRIPT MINUS (U+207B)
			Expected:     -9223372036854775807,
		},
		{
			Text: []byte("−9223372036854775807"), // MINUS SIGN (U+2212)
			Expected:     -9223372036854775807,
		},
		{
			Text: []byte("﹣9223372036854775807"), // SMALL HYPHEN-MINUS (U+FE63)
			Expected:      -9223372036854775807,
		},
		{
			Text: []byte("－9223372036854775807"), // FULLWIDTH HYPHEN-MINUS (U+FF0D)
			Expected:      -9223372036854775807,
		},
		{
			Text: []byte("+9223372036854775807"), // PLUS SIGN (U+002B)
			Expected:      9223372036854775807,
		},
		{
			Text: []byte("⁺9223372036854775807"), // SUPERSCRIPT PLUS SIGN (U+207A)
			Expected:      9223372036854775807,
		},
		{
			Text: []byte("＋9223372036854775807"), // FULLWIDTH PLUS SIGN (U+FF0B)
			Expected:       9223372036854775807,
		},



		{
			Text: []byte("-9223372036854775808"), // HYPHEN-MINUS (U+002D)
			Expected:     -9223372036854775808,
		},
		{
			Text: []byte("⁻9223372036854775808"), // SUPERSCRIPT MINUS (U+207B)
			Expected:     -9223372036854775808,
		},
		{
			Text: []byte("⁻9223372036854775808"), // SUPERSCRIPT MINUS (U+207B)
			Expected:     -9223372036854775808,
		},
		{
			Text: []byte("−9223372036854775808"), // MINUS SIGN (U+2212)
			Expected:     -9223372036854775808,
		},
		{
			Text: []byte("﹣9223372036854775808"), // SMALL HYPHEN-MINUS (U+FE63)
			Expected:      -9223372036854775808,
		},
		{
			Text: []byte("－9223372036854775808"), // FULLWIDTH HYPHEN-MINUS (U+FF0D)
			Expected:      -9223372036854775808,
		},












		//
		// Arabic-Indic Digits
		//



		{
			Text: []byte("٠"),
			Expected:     0,
		},
		{
			Text: []byte("-٠"), // HYPHEN-MINUS (U+002D)
			Expected:      0,
		},
		{
			Text: []byte("⁻٠"), // SUPERSCRIPT MINUS (U+207B)
			Expected:      0,
		},
		{
			Text: []byte("−٠"), // MINUS SIGN (U+2212)
			Expected:      0,
		},
		{
			Text: []byte("﹣٠"), // SMALL HYPHEN-MINUS (U+FE63)
			Expected:       0,
		},
		{
			Text: []byte("－٠"), // FULLWIDTH HYPHEN-MINUS (U+FF0D)
			Expected:       0,
		},
		{
			Text: []byte("+٠"), // PLUS SIGN (U+002B)
			Expected:      0,
		},
		{
			Text: []byte("⁺٠"), // SUPERSCRIPT PLUS SIGN (U+207A)
			Expected:      0,
		},
		{
			Text: []byte("＋٠"), // FULLWIDTH PLUS SIGN (U+FF0B)
			Expected:      0,
		},



		{
			Text: []byte("١"),
			Expected:     1,
		},
		{
			Text: []byte("٢"),
			Expected:     2,
		},
		{
			Text: []byte("٣"),
			Expected:     3,
		},
		{
			Text: []byte("٤"),
			Expected:     4,
		},
		{
			Text: []byte("٥"),
			Expected:     5,
		},
		{
			Text: []byte("٦"),
			Expected:     6,
		},
		{
			Text: []byte("٧"),
			Expected:     7,
		},
		{
			Text: []byte("٨"),
			Expected:     8,
		},
		{
			Text: []byte("٩"),
			Expected:     9,
		},



		{
			Text: []byte("-١"), // HYPHEN-MINUS (U+002D)
			Expected:     -1,
		},
		{
			Text: []byte("-٢"), // HYPHEN-MINUS (U+002D)
			Expected:     -2,
		},
		{
			Text: []byte("-٣"), // HYPHEN-MINUS (U+002D)
			Expected:     -3,
		},
		{
			Text: []byte("-٤"), // HYPHEN-MINUS (U+002D)
			Expected:     -4,
		},
		{
			Text: []byte("-٥"), // HYPHEN-MINUS (U+002D)
			Expected:     -5,
		},
		{
			Text: []byte("-٦"), // HYPHEN-MINUS (U+002D)
			Expected:     -6,
		},
		{
			Text: []byte("-٧"), // HYPHEN-MINUS (U+002D)
			Expected:     -7,
		},
		{
			Text: []byte("-٨"), // HYPHEN-MINUS (U+002D)
			Expected:     -8,
		},
		{
			Text: []byte("-٩"), // HYPHEN-MINUS (U+002D)
			Expected:     -9,
		},



		{
			Text: []byte("⁻١"), // SUPERSCRIPT MINUS (U+207B)
			Expected:     -1,
		},
		{
			Text: []byte("⁻٢"), // SUPERSCRIPT MINUS (U+207B)
			Expected:     -2,
		},
		{
			Text: []byte("⁻٣"), // SUPERSCRIPT MINUS (U+207B)
			Expected:     -3,
		},
		{
			Text: []byte("⁻٤"), // SUPERSCRIPT MINUS (U+207B)
			Expected:     -4,
		},
		{
			Text: []byte("⁻٥"), // SUPERSCRIPT MINUS (U+207B)
			Expected:     -5,
		},
		{
			Text: []byte("⁻٦"), // SUPERSCRIPT MINUS (U+207B)
			Expected:     -6,
		},
		{
			Text: []byte("⁻٧"), // SUPERSCRIPT MINUS (U+207B)
			Expected:     -7,
		},
		{
			Text: []byte("⁻٨"), // SUPERSCRIPT MINUS (U+207B)
			Expected:     -8,
		},
		{
			Text: []byte("⁻٩"), // SUPERSCRIPT MINUS (U+207B)
			Expected:     -9,
		},



		{
			Text: []byte("−١"), // MINUS SIGN (U+2212)
			Expected:     -1,
		},
		{
			Text: []byte("−٢"), // MINUS SIGN (U+2212)
			Expected:     -2,
		},
		{
			Text: []byte("−٣"), // MINUS SIGN (U+2212)
			Expected:     -3,
		},
		{
			Text: []byte("−٤"), // MINUS SIGN (U+2212)
			Expected:     -4,
		},
		{
			Text: []byte("−٥"), // MINUS SIGN (U+2212)
			Expected:     -5,
		},
		{
			Text: []byte("−٦"), // MINUS SIGN (U+2212)
			Expected:     -6,
		},
		{
			Text: []byte("−٧"), // MINUS SIGN (U+2212)
			Expected:     -7,
		},
		{
			Text: []byte("−٨"), // MINUS SIGN (U+2212)
			Expected:     -8,
		},
		{
			Text: []byte("−٩"), // MINUS SIGN (U+2212)
			Expected:     -9,
		},



		{
			Text: []byte("﹣١"), // SMALL HYPHEN-MINUS (U+FE63)
			Expected:      -1,
		},
		{
			Text: []byte("﹣٢"), // SMALL HYPHEN-MINUS (U+FE63)
			Expected:      -2,
		},
		{
			Text: []byte("﹣٣"), // SMALL HYPHEN-MINUS (U+FE63)
			Expected:      -3,
		},
		{
			Text: []byte("﹣٤"), // SMALL HYPHEN-MINUS (U+FE63)
			Expected:      -4,
		},
		{
			Text: []byte("﹣٥"), // SMALL HYPHEN-MINUS (U+FE63)
			Expected:      -5,
		},
		{
			Text: []byte("﹣٦"), // SMALL HYPHEN-MINUS (U+FE63)
			Expected:      -6,
		},
		{
			Text: []byte("﹣٧"), // SMALL HYPHEN-MINUS (U+FE63)
			Expected:      -7,
		},
		{
			Text: []byte("﹣٨"), // SMALL HYPHEN-MINUS (U+FE63)
			Expected:      -8,
		},
		{
			Text: []byte("﹣٩"), // SMALL HYPHEN-MINUS (U+FE63)
			Expected:      -9,
		},



		{
			Text: []byte("－١"), // FULLWIDTH HYPHEN-MINUS (U+FF0D)
			Expected:      -1,
		},
		{
			Text: []byte("－٢"), // FULLWIDTH HYPHEN-MINUS (U+FF0D)
			Expected:      -2,
		},
		{
			Text: []byte("－٣"), // FULLWIDTH HYPHEN-MINUS (U+FF0D)
			Expected:      -3,
		},
		{
			Text: []byte("－٤"), // FULLWIDTH HYPHEN-MINUS (U+FF0D)
			Expected:      -4,
		},
		{
			Text: []byte("－٥"), // FULLWIDTH HYPHEN-MINUS (U+FF0D)
			Expected:      -5,
		},
		{
			Text: []byte("－٦"), // FULLWIDTH HYPHEN-MINUS (U+FF0D)
			Expected:      -6,
		},
		{
			Text: []byte("－٧"), // FULLWIDTH HYPHEN-MINUS (U+FF0D)
			Expected:      -7,
		},
		{
			Text: []byte("－٨"), // FULLWIDTH HYPHEN-MINUS (U+FF0D)
			Expected:      -8,
		},
		{
			Text: []byte("－٩"), // FULLWIDTH HYPHEN-MINUS (U+FF0D)
			Expected:      -9,
		},



		{
			Text: []byte("+١"), // PLUS SIGN (U+002B)
			Expected:      1,
		},
		{
			Text: []byte("+٢"), // PLUS SIGN (U+002B)
			Expected:      2,
		},
		{
			Text: []byte("+٣"), // PLUS SIGN (U+002B)
			Expected:      3,
		},
		{
			Text: []byte("+٤"), // PLUS SIGN (U+002B)
			Expected:      4,
		},
		{
			Text: []byte("+٥"), // PLUS SIGN (U+002B)
			Expected:      5,
		},
		{
			Text: []byte("+٦"), // PLUS SIGN (U+002B)
			Expected:      6,
		},
		{
			Text: []byte("+٧"), // PLUS SIGN (U+002B)
			Expected:      7,
		},
		{
			Text: []byte("+٨"), // PLUS SIGN (U+002B)
			Expected:      8,
		},
		{
			Text: []byte("+٩"), // PLUS SIGN (U+002B)
			Expected:      9,
		},



		{
			Text: []byte("⁺١"), // SUPERSCRIPT PLUS SIGN (U+207A)
			Expected:      1,
		},
		{
			Text: []byte("⁺٢"), // SUPERSCRIPT PLUS SIGN (U+207A)
			Expected:      2,
		},
		{
			Text: []byte("⁺٣"), // SUPERSCRIPT PLUS SIGN (U+207A)
			Expected:      3,
		},
		{
			Text: []byte("⁺٤"), // SUPERSCRIPT PLUS SIGN (U+207A)
			Expected:      4,
		},
		{
			Text: []byte("⁺٥"), // SUPERSCRIPT PLUS SIGN (U+207A)
			Expected:      5,
		},
		{
			Text: []byte("⁺٦"), // SUPERSCRIPT PLUS SIGN (U+207A)
			Expected:      6,
		},
		{
			Text: []byte("⁺٧"), // SUPERSCRIPT PLUS SIGN (U+207A)
			Expected:      7,
		},
		{
			Text: []byte("⁺٨"), // SUPERSCRIPT PLUS SIGN (U+207A)
			Expected:      8,
		},
		{
			Text: []byte("⁺٩"), // SUPERSCRIPT PLUS SIGN (U+207A)
			Expected:      9,
		},



		{
			Text: []byte("＋١"), // FULLWIDTH PLUS SIGN (U+FF0B)
			Expected:       1,
		},
		{
			Text: []byte("＋٢"), // FULLWIDTH PLUS SIGN (U+FF0B)
			Expected:       2,
		},
		{
			Text: []byte("＋٣"), // FULLWIDTH PLUS SIGN (U+FF0B)
			Expected:       3,
		},
		{
			Text: []byte("＋٤"), // FULLWIDTH PLUS SIGN (U+FF0B)
			Expected:       4,
		},
		{
			Text: []byte("＋٥"), // FULLWIDTH PLUS SIGN (U+FF0B)
			Expected:       5,
		},
		{
			Text: []byte("＋٦"), // FULLWIDTH PLUS SIGN (U+FF0B)
			Expected:       6,
		},
		{
			Text: []byte("＋٧"), // FULLWIDTH PLUS SIGN (U+FF0B)
			Expected:       7,
		},
		{
			Text: []byte("＋٨"), // FULLWIDTH PLUS SIGN (U+FF0B)
			Expected:       8,
		},
		{
			Text: []byte("＋٩"), // FULLWIDTH PLUS SIGN (U+FF0B)
			Expected:       9,
		},



		{
			Text:  []byte("١٠"),
			Expected:      10,
		},
		{
			Text: []byte("-١٠"), // HYPHEN-MINUS (U+002D)
			Expected:     -10,
		},
		{
			Text: []byte("⁻١٠"), // SUPERSCRIPT MINUS (U+207B)
			Expected:     -10,
		},
		{
			Text: []byte("−١٠"), // MINUS SIGN (U+2212)
			Expected:     -10,
		},
		{
			Text: []byte("﹣١٠"), // SMALL HYPHEN-MINUS (U+FE63)
			Expected:      -10,
		},
		{
			Text: []byte("－١٠"), // FULLWIDTH HYPHEN-MINUS (U+FF0D)
			Expected:      -10,
		},
		{
			Text: []byte("+١٠"),
			Expected:      10,
		},
		{
			Text: []byte("⁺١٠"), // SUPERSCRIPT PLUS SIGN (U+207A)
			Expected:      10,
		},
		{
			Text: []byte("＋١٠"), // FULLWIDTH PLUS SIGN (U+FF0B)
			Expected:       10,
		},



		{
			Text:  []byte("٢١"),
			Expected:      21,
		},
		{
			Text: []byte("-٢١"), // HYPHEN-MINUS (U+002D)
			Expected:     -21,
		},
		{
			Text: []byte("⁻٢١"), // SUPERSCRIPT MINUS (U+207B)
			Expected:     -21,
		},
		{
			Text: []byte("−٢١"), // MINUS SIGN (U+2212)
			Expected:     -21,
		},
		{
			Text: []byte("﹣٢١"), // SMALL HYPHEN-MINUS (U+FE63)
			Expected:      -21,
		},
		{
			Text: []byte("－٢١"), // FULLWIDTH HYPHEN-MINUS (U+FF0D)
			Expected:      -21,
		},
		{
			Text: []byte("+٢١"), // PLUS SIGN (U+002B)
			Expected:      21,
		},
		{
			Text: []byte("⁺٢١"), // SUPERSCRIPT PLUS SIGN (U+207A)
			Expected:      21,
		},
		{
			Text: []byte("＋٢١"), // FULLWIDTH PLUS SIGN (U+FF0B)
			Expected:       21,
		},



		{
			Text:  []byte("٣٢"),
			Expected:      32,
		},
		{
			Text: []byte("-٣٢"), // HYPHEN-MINUS (U+002D)
			Expected:     -32,
		},
		{
			Text: []byte("⁻٣٢"), // SUPERSCRIPT MINUS (U+207B)
			Expected:     -32,
		},
		{
			Text: []byte("−٣٢"), // MINUS SIGN (U+2212)
			Expected:     -32,
		},
		{
			Text: []byte("﹣٣٢"), // SMALL HYPHEN-MINUS (U+FE63)
			Expected:      -32,
		},
		{
			Text: []byte("－٣٢"), // FULLWIDTH HYPHEN-MINUS (U+FF0D)
			Expected:      -32,
		},
		{
			Text: []byte("+٣٢"), // PLUS SIGN (U+002B)
			Expected:      32,
		},
		{
			Text: []byte("⁺٣٢"), // SUPERSCRIPT PLUS SIGN (U+207A)
			Expected:      32,
		},
		{
			Text: []byte("＋٣٢"), // FULLWIDTH PLUS SIGN (U+FF0B)
			Expected:       32,
		},



		{
			Text:  []byte("١٢٣٤٥"),
			Expected:      12345,
		},
		{
			Text: []byte("-١٢٣٤٥"), // HYPHEN-MINUS (U+002D)
			Expected:     -12345,
		},
		{
			Text: []byte("⁻١٢٣٤٥"), // SUPERSCRIPT MINUS (U+207B)
			Expected:     -12345,
		},
		{
			Text: []byte("−١٢٣٤٥"), // MINUS SIGN (U+2212)
			Expected:     -12345,
		},
		{
			Text: []byte("﹣١٢٣٤٥"), // SMALL HYPHEN-MINUS (U+FE63)
			Expected:      -12345,
		},
		{
			Text: []byte("－١٢٣٤٥"), // FULLWIDTH HYPHEN-MINUS (U+FF0D)
			Expected:      -12345,
		},
		{
			Text: []byte("+١٢٣٤٥"), // PLUS SIGN (U+002B)
			Expected:      12345,
		},
		{
			Text: []byte("⁺١٢٣٤٥"), // SUPERSCRIPT PLUS SIGN (U+207A)
			Expected:      12345,
		},
		{
			Text: []byte("＋١٢٣٤٥"), // FULLWIDTH PLUS SIGN (U+FF0B)
			Expected:       12345,
		},



		{
			Text:  []byte("٩٨٧٦٥٤٣٢١٠"),
			Expected:      9876543210,
		},
		{
			Text: []byte("-٩٨٧٦٥٤٣٢١٠"), // HYPHEN-MINUS (U+002D)
			Expected:     -9876543210,
		},
		{
			Text: []byte("⁻٩٨٧٦٥٤٣٢١٠"), // SUPERSCRIPT MINUS (U+207B)
			Expected:     -9876543210,
		},
		{
			Text: []byte("−٩٨٧٦٥٤٣٢١٠"), // MINUS SIGN (U+2212)
			Expected:      9876543210,
		},
		{
			Text: []byte("﹣٩٨٧٦٥٤٣٢١٠"), // SMALL HYPHEN-MINUS (U+FE63)
			Expected:      -9876543210,
		},
		{
			Text: []byte("－٩٨٧٦٥٤٣٢١٠"), // FULLWIDTH HYPHEN-MINUS (U+FF0D)
			Expected:      -9876543210,
		},
		{
			Text: []byte("+٩٨٧٦٥٤٣٢١٠"), // PLUS SIGN (U+002B)
			Expected:      9876543210,
		},
		{
			Text: []byte("⁺٩٨٧٦٥٤٣٢١٠"), // SUPERSCRIPT PLUS SIGN (U+207A)
			Expected:      9876543210,
		},
		{
			Text: []byte("＋٩٨٧٦٥٤٣٢١٠"), // FULLWIDTH PLUS SIGN (U+FF0B)
			Expected:       9876543210,
		},



		{
			Text:  []byte("٩٢٢٣٣٧٢٠٣٦٨٥٤٧٧٥٨٠٧"),
			Expected:      9223372036854775807,
		},
		{
			Text: []byte("-٩٢٢٣٣٧٢٠٣٦٨٥٤٧٧٥٨٠٧"), // HYPHEN-MINUS (U+002D)
			Expected:     -9223372036854775807,
		},
		{
			Text: []byte("⁻٩٢٢٣٣٧٢٠٣٦٨٥٤٧٧٥٨٠٧"), // SUPERSCRIPT MINUS (U+207B)
			Expected:     -9223372036854775807,
		},
		{
			Text: []byte("−٩٢٢٣٣٧٢٠٣٦٨٥٤٧٧٥٨٠٧"), // MINUS SIGN (U+2212)
			Expected:     -9223372036854775807,
		},
		{
			Text: []byte("﹣٩٢٢٣٣٧٢٠٣٦٨٥٤٧٧٥٨٠٧"), // SMALL HYPHEN-MINUS (U+FE63)
			Expected:      -9223372036854775807,
		},
		{
			Text: []byte("－٩٢٢٣٣٧٢٠٣٦٨٥٤٧٧٥٨٠٧"), // FULLWIDTH HYPHEN-MINUS (U+FF0D)
			Expected:      -9223372036854775807,
		},
		{
			Text: []byte("+٩٢٢٣٣٧٢٠٣٦٨٥٤٧٧٥٨٠٧"), // PLUS SIGN (U+002B)
			Expected:      9223372036854775807,
		},
		{
			Text: []byte("⁺٩٢٢٣٣٧٢٠٣٦٨٥٤٧٧٥٨٠٧"), // SUPERSCRIPT PLUS SIGN (U+207A)
			Expected:      9223372036854775807,
		},
		{
			Text: []byte("＋٩٢٢٣٣٧٢٠٣٦٨٥٤٧٧٥٨٠٧"), // FULLWIDTH PLUS SIGN (U+FF0B)
			Expected:       9223372036854775807,
		},



		{
			Text: []byte("-٩٢٢٣٣٧٢٠٣٦٨٥٤٧٧٥٨٠٨"), // HYPHEN-MINUS (U+002D)
			Expected:     -9223372036854775808,
		},
		{
			Text: []byte("⁻٩٢٢٣٣٧٢٠٣٦٨٥٤٧٧٥٨٠٨"), // SUPERSCRIPT MINUS (U+207B)
			Expected:     -9223372036854775808,
		},
		{
			Text: []byte("⁻٩٢٢٣٣٧٢٠٣٦٨٥٤٧٧٥٨٠٨"), // SUPERSCRIPT MINUS (U+207B)
			Expected:     -9223372036854775808,
		},
		{
			Text: []byte("−٩٢٢٣٣٧٢٠٣٦٨٥٤٧٧٥٨٠٨"), // MINUS SIGN (U+2212)
			Expected:     -9223372036854775808,
		},
		{
			Text: []byte("﹣٩٢٢٣٣٧٢٠٣٦٨٥٤٧٧٥٨٠٨"), // SMALL HYPHEN-MINUS (U+FE63)
			Expected:      -9223372036854775808,
		},
		{
			Text: []byte("－٩٢٢٣٣٧٢٠٣٦٨٥٤٧٧٥٨٠٨"), // FULLWIDTH HYPHEN-MINUS (U+FF0D)
			Expected:      -9223372036854775808,
		},












		//
		// Extended Arabic-Indic Digit
		//



		{
			Text: []byte("۰"),
			Expected:     0,
		},
		{
			Text: []byte("-۰"), // HYPHEN-MINUS (U+002D)
			Expected:      0,
		},
		{
			Text: []byte("⁻۰"), // SUPERSCRIPT MINUS (U+207B)
			Expected:      0,
		},
		{
			Text: []byte("−۰"), // MINUS SIGN (U+2212)
			Expected:      0,
		},
		{
			Text: []byte("﹣۰"), // SMALL HYPHEN-MINUS (U+FE63)
			Expected:       0,
		},
		{
			Text: []byte("－۰"), // FULLWIDTH HYPHEN-MINUS (U+FF0D)
			Expected:       0,
		},
		{
			Text: []byte("+۰"), // PLUS SIGN (U+002B)
			Expected:      0,
		},
		{
			Text: []byte("⁺۰"), // SUPERSCRIPT PLUS SIGN (U+207A)
			Expected:      0,
		},
		{
			Text: []byte("＋۰"), // FULLWIDTH PLUS SIGN (U+FF0B)
			Expected:      0,
		},



		{
			Text: []byte("۱"),
			Expected:     1,
		},
		{
			Text: []byte("۲"),
			Expected:     2,
		},
		{
			Text: []byte("۳"),
			Expected:     3,
		},
		{
			Text: []byte("۴"),
			Expected:     4,
		},
		{
			Text: []byte("۵"),
			Expected:     5,
		},
		{
			Text: []byte("۶"),
			Expected:     6,
		},
		{
			Text: []byte("۷"),
			Expected:     7,
		},
		{
			Text: []byte("۸"),
			Expected:     8,
		},
		{
			Text: []byte("۹"),
			Expected:     9,
		},



		{
			Text: []byte("-۱"), // HYPHEN-MINUS (U+002D)
			Expected:     -1,
		},
		{
			Text: []byte("-۲"), // HYPHEN-MINUS (U+002D)
			Expected:     -2,
		},
		{
			Text: []byte("-۳"), // HYPHEN-MINUS (U+002D)
			Expected:     -3,
		},
		{
			Text: []byte("-۴"), // HYPHEN-MINUS (U+002D)
			Expected:     -4,
		},
		{
			Text: []byte("-۵"), // HYPHEN-MINUS (U+002D)
			Expected:     -5,
		},
		{
			Text: []byte("-۶"), // HYPHEN-MINUS (U+002D)
			Expected:     -6,
		},
		{
			Text: []byte("-۷"), // HYPHEN-MINUS (U+002D)
			Expected:     -7,
		},
		{
			Text: []byte("-۸"), // HYPHEN-MINUS (U+002D)
			Expected:     -8,
		},
		{
			Text: []byte("-۹"), // HYPHEN-MINUS (U+002D)
			Expected:     -9,
		},



		{
			Text: []byte("⁻۱"), // SUPERSCRIPT MINUS (U+207B)
			Expected:     -1,
		},
		{
			Text: []byte("⁻۲"), // SUPERSCRIPT MINUS (U+207B)
			Expected:     -2,
		},
		{
			Text: []byte("⁻۳"), // SUPERSCRIPT MINUS (U+207B)
			Expected:     -3,
		},
		{
			Text: []byte("⁻۴"), // SUPERSCRIPT MINUS (U+207B)
			Expected:     -4,
		},
		{
			Text: []byte("⁻۵"), // SUPERSCRIPT MINUS (U+207B)
			Expected:     -5,
		},
		{
			Text: []byte("⁻۶"), // SUPERSCRIPT MINUS (U+207B)
			Expected:     -6,
		},
		{
			Text: []byte("⁻۷"), // SUPERSCRIPT MINUS (U+207B)
			Expected:     -7,
		},
		{
			Text: []byte("⁻۸"), // SUPERSCRIPT MINUS (U+207B)
			Expected:     -8,
		},
		{
			Text: []byte("⁻۹"), // SUPERSCRIPT MINUS (U+207B)
			Expected:     -9,
		},



		{
			Text: []byte("−۱"), // MINUS SIGN (U+2212)
			Expected:     -1,
		},
		{
			Text: []byte("−۲"), // MINUS SIGN (U+2212)
			Expected:     -2,
		},
		{
			Text: []byte("−۳"), // MINUS SIGN (U+2212)
			Expected:     -3,
		},
		{
			Text: []byte("−۴"), // MINUS SIGN (U+2212)
			Expected:     -4,
		},
		{
			Text: []byte("−۵"), // MINUS SIGN (U+2212)
			Expected:     -5,
		},
		{
			Text: []byte("−۶"), // MINUS SIGN (U+2212)
			Expected:     -6,
		},
		{
			Text: []byte("−۷"), // MINUS SIGN (U+2212)
			Expected:     -7,
		},
		{
			Text: []byte("−۸"), // MINUS SIGN (U+2212)
			Expected:     -8,
		},
		{
			Text: []byte("−۹"), // MINUS SIGN (U+2212)
			Expected:     -9,
		},



		{
			Text: []byte("﹣۱"), // SMALL HYPHEN-MINUS (U+FE63)
			Expected:      -1,
		},
		{
			Text: []byte("﹣۲"), // SMALL HYPHEN-MINUS (U+FE63)
			Expected:      -2,
		},
		{
			Text: []byte("﹣۳"), // SMALL HYPHEN-MINUS (U+FE63)
			Expected:      -3,
		},
		{
			Text: []byte("﹣۴"), // SMALL HYPHEN-MINUS (U+FE63)
			Expected:      -4,
		},
		{
			Text: []byte("﹣۵"), // SMALL HYPHEN-MINUS (U+FE63)
			Expected:      -5,
		},
		{
			Text: []byte("﹣۶"), // SMALL HYPHEN-MINUS (U+FE63)
			Expected:      -6,
		},
		{
			Text: []byte("﹣۷"), // SMALL HYPHEN-MINUS (U+FE63)
			Expected:      -7,
		},
		{
			Text: []byte("﹣۸"), // SMALL HYPHEN-MINUS (U+FE63)
			Expected:      -8,
		},
		{
			Text: []byte("﹣۹"), // SMALL HYPHEN-MINUS (U+FE63)
			Expected:      -9,
		},



		{
			Text: []byte("－۱"), // FULLWIDTH HYPHEN-MINUS (U+FF0D)
			Expected:      -1,
		},
		{
			Text: []byte("－۲"), // FULLWIDTH HYPHEN-MINUS (U+FF0D)
			Expected:      -2,
		},
		{
			Text: []byte("－۳"), // FULLWIDTH HYPHEN-MINUS (U+FF0D)
			Expected:      -3,
		},
		{
			Text: []byte("－۴"), // FULLWIDTH HYPHEN-MINUS (U+FF0D)
			Expected:      -4,
		},
		{
			Text: []byte("－۵"), // FULLWIDTH HYPHEN-MINUS (U+FF0D)
			Expected:      -5,
		},
		{
			Text: []byte("－۶"), // FULLWIDTH HYPHEN-MINUS (U+FF0D)
			Expected:      -6,
		},
		{
			Text: []byte("－۷"), // FULLWIDTH HYPHEN-MINUS (U+FF0D)
			Expected:      -7,
		},
		{
			Text: []byte("－۸"), // FULLWIDTH HYPHEN-MINUS (U+FF0D)
			Expected:      -8,
		},
		{
			Text: []byte("－۹"), // FULLWIDTH HYPHEN-MINUS (U+FF0D)
			Expected:      -9,
		},



		{
			Text: []byte("+۱"), // PLUS SIGN (U+002B)
			Expected:      1,
		},
		{
			Text: []byte("+۲"), // PLUS SIGN (U+002B)
			Expected:      2,
		},
		{
			Text: []byte("+۳"), // PLUS SIGN (U+002B)
			Expected:      3,
		},
		{
			Text: []byte("+۴"), // PLUS SIGN (U+002B)
			Expected:      4,
		},
		{
			Text: []byte("+۵"), // PLUS SIGN (U+002B)
			Expected:      5,
		},
		{
			Text: []byte("+۶"), // PLUS SIGN (U+002B)
			Expected:      6,
		},
		{
			Text: []byte("+۷"), // PLUS SIGN (U+002B)
			Expected:      7,
		},
		{
			Text: []byte("+۸"), // PLUS SIGN (U+002B)
			Expected:      8,
		},
		{
			Text: []byte("+۹"), // PLUS SIGN (U+002B)
			Expected:      9,
		},



		{
			Text: []byte("⁺۱"), // SUPERSCRIPT PLUS SIGN (U+207A)
			Expected:      1,
		},
		{
			Text: []byte("⁺۲"), // SUPERSCRIPT PLUS SIGN (U+207A)
			Expected:      2,
		},
		{
			Text: []byte("⁺۳"), // SUPERSCRIPT PLUS SIGN (U+207A)
			Expected:      3,
		},
		{
			Text: []byte("⁺۴"), // SUPERSCRIPT PLUS SIGN (U+207A)
			Expected:      4,
		},
		{
			Text: []byte("⁺۵"), // SUPERSCRIPT PLUS SIGN (U+207A)
			Expected:      5,
		},
		{
			Text: []byte("⁺۶"), // SUPERSCRIPT PLUS SIGN (U+207A)
			Expected:      6,
		},
		{
			Text: []byte("⁺۷"), // SUPERSCRIPT PLUS SIGN (U+207A)
			Expected:      7,
		},
		{
			Text: []byte("⁺۸"), // SUPERSCRIPT PLUS SIGN (U+207A)
			Expected:      8,
		},
		{
			Text: []byte("⁺۹"), // SUPERSCRIPT PLUS SIGN (U+207A)
			Expected:      9,
		},



		{
			Text: []byte("＋۱"), // FULLWIDTH PLUS SIGN (U+FF0B)
			Expected:       1,
		},
		{
			Text: []byte("＋۲"), // FULLWIDTH PLUS SIGN (U+FF0B)
			Expected:       2,
		},
		{
			Text: []byte("＋۳"), // FULLWIDTH PLUS SIGN (U+FF0B)
			Expected:       3,
		},
		{
			Text: []byte("＋۴"), // FULLWIDTH PLUS SIGN (U+FF0B)
			Expected:       4,
		},
		{
			Text: []byte("＋۵"), // FULLWIDTH PLUS SIGN (U+FF0B)
			Expected:       5,
		},
		{
			Text: []byte("＋۶"), // FULLWIDTH PLUS SIGN (U+FF0B)
			Expected:       6,
		},
		{
			Text: []byte("＋۷"), // FULLWIDTH PLUS SIGN (U+FF0B)
			Expected:       7,
		},
		{
			Text: []byte("＋۸"), // FULLWIDTH PLUS SIGN (U+FF0B)
			Expected:       8,
		},
		{
			Text: []byte("＋۹"), // FULLWIDTH PLUS SIGN (U+FF0B)
			Expected:       9,
		},



		{
			Text:  []byte("۱۰"),
			Expected:      10,
		},
		{
			Text: []byte("-۱۰"), // HYPHEN-MINUS (U+002D)
			Expected:     -10,
		},
		{
			Text: []byte("⁻۱۰"), // SUPERSCRIPT MINUS (U+207B)
			Expected:     -10,
		},
		{
			Text: []byte("−۱۰"), // MINUS SIGN (U+2212)
			Expected:     -10,
		},
		{
			Text: []byte("﹣۱۰"), // SMALL HYPHEN-MINUS (U+FE63)
			Expected:      -10,
		},
		{
			Text: []byte("－۱۰"), // FULLWIDTH HYPHEN-MINUS (U+FF0D)
			Expected:      -10,
		},
		{
			Text: []byte("+۱۰"),
			Expected:      10,
		},
		{
			Text: []byte("⁺۱۰"), // SUPERSCRIPT PLUS SIGN (U+207A)
			Expected:      10,
		},
		{
			Text: []byte("＋۱۰"), // FULLWIDTH PLUS SIGN (U+FF0B)
			Expected:       10,
		},



		{
			Text:  []byte("۲۱"),
			Expected:      21,
		},
		{
			Text: []byte("-۲۱"), // HYPHEN-MINUS (U+002D)
			Expected:     -21,
		},
		{
			Text: []byte("⁻۲۱"), // SUPERSCRIPT MINUS (U+207B)
			Expected:     -21,
		},
		{
			Text: []byte("−۲۱"), // MINUS SIGN (U+2212)
			Expected:     -21,
		},
		{
			Text: []byte("﹣۲۱"), // SMALL HYPHEN-MINUS (U+FE63)
			Expected:      -21,
		},
		{
			Text: []byte("－۲۱"), // FULLWIDTH HYPHEN-MINUS (U+FF0D)
			Expected:      -21,
		},
		{
			Text: []byte("+۲۱"), // PLUS SIGN (U+002B)
			Expected:      21,
		},
		{
			Text: []byte("⁺۲۱"), // SUPERSCRIPT PLUS SIGN (U+207A)
			Expected:      21,
		},
		{
			Text: []byte("＋۲۱"), // FULLWIDTH PLUS SIGN (U+FF0B)
			Expected:       21,
		},



		{
			Text:  []byte("۳۲"),
			Expected:      32,
		},
		{
			Text: []byte("-۳۲"), // HYPHEN-MINUS (U+002D)
			Expected:     -32,
		},
		{
			Text: []byte("⁻۳۲"), // SUPERSCRIPT MINUS (U+207B)
			Expected:     -32,
		},
		{
			Text: []byte("−۳۲"), // MINUS SIGN (U+2212)
			Expected:     -32,
		},
		{
			Text: []byte("﹣۳۲"), // SMALL HYPHEN-MINUS (U+FE63)
			Expected:      -32,
		},
		{
			Text: []byte("－۳۲"), // FULLWIDTH HYPHEN-MINUS (U+FF0D)
			Expected:      -32,
		},
		{
			Text: []byte("+۳۲"), // PLUS SIGN (U+002B)
			Expected:      32,
		},
		{
			Text: []byte("⁺۳۲"), // SUPERSCRIPT PLUS SIGN (U+207A)
			Expected:      32,
		},
		{
			Text: []byte("＋۳۲"), // FULLWIDTH PLUS SIGN (U+FF0B)
			Expected:       32,
		},



		{
			Text:  []byte("۱۲۳۴۵"),
			Expected:      12345,
		},
		{
			Text: []byte("-۱۲۳۴۵"), // HYPHEN-MINUS (U+002D)
			Expected:     -12345,
		},
		{
			Text: []byte("⁻۱۲۳۴۵"), // SUPERSCRIPT MINUS (U+207B)
			Expected:     -12345,
		},
		{
			Text: []byte("−۱۲۳۴۵"), // MINUS SIGN (U+2212)
			Expected:     -12345,
		},
		{
			Text: []byte("﹣۱۲۳۴۵"), // SMALL HYPHEN-MINUS (U+FE63)
			Expected:      -12345,
		},
		{
			Text: []byte("－۱۲۳۴۵"), // FULLWIDTH HYPHEN-MINUS (U+FF0D)
			Expected:      -12345,
		},
		{
			Text: []byte("+۱۲۳۴۵"), // PLUS SIGN (U+002B)
			Expected:      12345,
		},
		{
			Text: []byte("⁺۱۲۳۴۵"), // SUPERSCRIPT PLUS SIGN (U+207A)
			Expected:      12345,
		},
		{
			Text: []byte("＋۱۲۳۴۵"), // FULLWIDTH PLUS SIGN (U+FF0B)
			Expected:       12345,
		},



		{
			Text:  []byte("۹۸۷۶۵۴۳۲۱۰"),
			Expected:      9876543210,
		},
		{
			Text: []byte("-۹۸۷۶۵۴۳۲۱۰"), // HYPHEN-MINUS (U+002D)
			Expected:     -9876543210,
		},
		{
			Text: []byte("⁻۹۸۷۶۵۴۳۲۱۰"), // SUPERSCRIPT MINUS (U+207B)
			Expected:     -9876543210,
		},
		{
			Text: []byte("−۹۸۷۶۵۴۳۲۱۰"), // MINUS SIGN (U+2212)
			Expected:      9876543210,
		},
		{
			Text: []byte("﹣۹۸۷۶۵۴۳۲۱۰"), // SMALL HYPHEN-MINUS (U+FE63)
			Expected:      -9876543210,
		},
		{
			Text: []byte("－۹۸۷۶۵۴۳۲۱۰"), // FULLWIDTH HYPHEN-MINUS (U+FF0D)
			Expected:      -9876543210,
		},
		{
			Text: []byte("+۹۸۷۶۵۴۳۲۱۰"), // PLUS SIGN (U+002B)
			Expected:      9876543210,
		},
		{
			Text: []byte("⁺۹۸۷۶۵۴۳۲۱۰"), // SUPERSCRIPT PLUS SIGN (U+207A)
			Expected:      9876543210,
		},
		{
			Text: []byte("＋۹۸۷۶۵۴۳۲۱۰"), // FULLWIDTH PLUS SIGN (U+FF0B)
			Expected:       9876543210,
		},



		{
			Text:  []byte("۹۲۲۳۳۷۲۰۳۶۸۵۴۷۷۵۸۰۷"),
			Expected:      9223372036854775807,
		},
		{
			Text: []byte("-۹۲۲۳۳۷۲۰۳۶۸۵۴۷۷۵۸۰۷"), // HYPHEN-MINUS (U+002D)
			Expected:     -9223372036854775807,
		},
		{
			Text: []byte("⁻۹۲۲۳۳۷۲۰۳۶۸۵۴۷۷۵۸۰۷"), // SUPERSCRIPT MINUS (U+207B)
			Expected:     -9223372036854775807,
		},
		{
			Text: []byte("−۹۲۲۳۳۷۲۰۳۶۸۵۴۷۷۵۸۰۷"), // MINUS SIGN (U+2212)
			Expected:     -9223372036854775807,
		},
		{
			Text: []byte("﹣۹۲۲۳۳۷۲۰۳۶۸۵۴۷۷۵۸۰۷"), // SMALL HYPHEN-MINUS (U+FE63)
			Expected:      -9223372036854775807,
		},
		{
			Text: []byte("－۹۲۲۳۳۷۲۰۳۶۸۵۴۷۷۵۸۰۷"), // FULLWIDTH HYPHEN-MINUS (U+FF0D)
			Expected:      -9223372036854775807,
		},
		{
			Text: []byte("+۹۲۲۳۳۷۲۰۳۶۸۵۴۷۷۵۸۰۷"), // PLUS SIGN (U+002B)
			Expected:      9223372036854775807,
		},
		{
			Text: []byte("⁺۹۲۲۳۳۷۲۰۳۶۸۵۴۷۷۵۸۰۷"), // SUPERSCRIPT PLUS SIGN (U+207A)
			Expected:      9223372036854775807,
		},
		{
			Text: []byte("＋۹۲۲۳۳۷۲۰۳۶۸۵۴۷۷۵۸۰۷"), // FULLWIDTH PLUS SIGN (U+FF0B)
			Expected:       9223372036854775807,
		},



		{
			Text: []byte("-۹۲۲۳۳۷۲۰۳۶۸۵۴۷۷۵۸۰۸"), // HYPHEN-MINUS (U+002D)
			Expected:     -9223372036854775808,
		},
		{
			Text: []byte("⁻۹۲۲۳۳۷۲۰۳۶۸۵۴۷۷۵۸۰۸"), // SUPERSCRIPT MINUS (U+207B)
			Expected:     -9223372036854775808,
		},
		{
			Text: []byte("⁻۹۲۲۳۳۷۲۰۳۶۸۵۴۷۷۵۸۰۸"), // SUPERSCRIPT MINUS (U+207B)
			Expected:     -9223372036854775808,
		},
		{
			Text: []byte("−۹۲۲۳۳۷۲۰۳۶۸۵۴۷۷۵۸۰۸"), // MINUS SIGN (U+2212)
			Expected:     -9223372036854775808,
		},
		{
			Text: []byte("﹣۹۲۲۳۳۷۲۰۳۶۸۵۴۷۷۵۸۰۸"), // SMALL HYPHEN-MINUS (U+FE63)
			Expected:      -9223372036854775808,
		},
		{
			Text: []byte("－۹۲۲۳۳۷۲۰۳۶۸۵۴۷۷۵۸۰۸"), // FULLWIDTH HYPHEN-MINUS (U+FF0D)
			Expected:      -9223372036854775808,
		},















		{
			Text: []byte("0x0"),
			Expected:     0x0,
		},
		{
			Text: []byte("0x1"),
			Expected:     0x1,
		},
		{
			Text: []byte("0x2"),
			Expected:     0x2,
		},
		{
			Text: []byte("0x3"),
			Expected:     0x3,
		},
		{
			Text: []byte("0x4"),
			Expected:     0x4,
		},
		{
			Text: []byte("0x5"),
			Expected:     0x5,
		},
		{
			Text: []byte("0x6"),
			Expected:     0x6,
		},
		{
			Text: []byte("0x7"),
			Expected:     0x7,
		},
		{
			Text: []byte("0x8"),
			Expected:     0x8,
		},
		{
			Text: []byte("0x9"),
			Expected:     0x9,
		},

		{
			Text: []byte("0xa"),
			Expected:     0xa,
		},
		{
			Text: []byte("0xA"),
			Expected:     0xA,
		},

		{
			Text: []byte("0xb"),
			Expected:     0xb,
		},
		{
			Text: []byte("0xB"),
			Expected:     0xB,
		},

		{
			Text: []byte("0xc"),
			Expected:     0xc,
		},
		{
			Text: []byte("0xC"),
			Expected:     0xC,
		},

		{
			Text: []byte("0xd"),
			Expected:     0xd,
		},
		{
			Text: []byte("0xD"),
			Expected:     0xD,
		},

		{
			Text: []byte("0xe"),
			Expected:     0xe,
		},
		{
			Text: []byte("0xE"),
			Expected:     0xE,
		},

		{
			Text: []byte("0xf"),
			Expected:     0xf,
		},
		{
			Text: []byte("0xF"),
			Expected:     0xF,
		},



		{
			Text: []byte("0x00"),
			Expected:     0x00,
		},
		{
			Text: []byte("0x01"),
			Expected:     0x01,
		},
		{
			Text: []byte("0x02"),
			Expected:     0x02,
		},
		{
			Text: []byte("0x03"),
			Expected:     0x03,
		},
		{
			Text: []byte("0x04"),
			Expected:     0x04,
		},
		{
			Text: []byte("0x05"),
			Expected:     0x05,
		},
		{
			Text: []byte("0x06"),
			Expected:     0x06,
		},
		{
			Text: []byte("0x07"),
			Expected:     0x07,
		},
		{
			Text: []byte("0x08"),
			Expected:     0x08,
		},
		{
			Text: []byte("0x09"),
			Expected:     0x09,
		},

		{
			Text: []byte("0x0a"),
			Expected:     0x0a,
		},
		{
			Text: []byte("0x0A"),
			Expected:     0x0A,
		},

		{
			Text: []byte("0x0b"),
			Expected:     0x0b,
		},
		{
			Text: []byte("0x0B"),
			Expected:     0x0B,
		},

		{
			Text: []byte("0x0c"),
			Expected:     0x0c,
		},
		{
			Text: []byte("0x0C"),
			Expected:     0x0C,
		},

		{
			Text: []byte("0x0d"),
			Expected:     0x0d,
		},
		{
			Text: []byte("0x0D"),
			Expected:     0x0D,
		},

		{
			Text: []byte("0x0e"),
			Expected:     0x0e,
		},
		{
			Text: []byte("0x0E"),
			Expected:     0x0E,
		},

		{
			Text: []byte("0x0f"),
			Expected:     0x0f,
		},
		{
			Text: []byte("0x0F"),
			Expected:     0x0F,
		},



		{
			Text: []byte("0x2f"),
			Expected:     0x2f,
		},
		{
			Text: []byte("0x2F"),
			Expected:     0x2F,
		},



		{
			Text: []byte("0x000"),
			Expected:     0x000,
		},
		{
			Text: []byte("0x001"),
			Expected:     0x001,
		},
		{
			Text: []byte("0x002"),
			Expected:     0x002,
		},
		{
			Text: []byte("0x003"),
			Expected:     0x003,
		},
		{
			Text: []byte("0x004"),
			Expected:     0x004,
		},
		{
			Text: []byte("0x005"),
			Expected:     0x005,
		},
		{
			Text: []byte("0x006"),
			Expected:     0x006,
		},
		{
			Text: []byte("0x007"),
			Expected:     0x007,
		},
		{
			Text: []byte("0x008"),
			Expected:     0x008,
		},
		{
			Text: []byte("0x009"),
			Expected:     0x009,
		},

		{
			Text: []byte("0x00a"),
			Expected:     0x00a,
		},
		{
			Text: []byte("0x00A"),
			Expected:     0x00A,
		},

		{
			Text: []byte("0x00b"),
			Expected:     0x00b,
		},
		{
			Text: []byte("0x00B"),
			Expected:     0x00B,
		},

		{
			Text: []byte("0x00c"),
			Expected:     0x00c,
		},
		{
			Text: []byte("0x00C"),
			Expected:     0x00C,
		},

		{
			Text: []byte("0x00d"),
			Expected:     0x00d,
		},
		{
			Text: []byte("0x00D"),
			Expected:     0x00D,
		},

		{
			Text: []byte("0x00e"),
			Expected:     0x00e,
		},
		{
			Text: []byte("0x00E"),
			Expected:     0x00E,
		},

		{
			Text: []byte("0x00f"),
			Expected:     0x00f,
		},
		{
			Text: []byte("0x00F"),
			Expected:     0x00F,
		},



		{
			Text: []byte("0x3ea"),
			Expected:     0x3ea,
		},
		{
			Text: []byte("0x3EA"),
			Expected:     0x3EA,
		},
		{
			Text: []byte("0x3eA"),
			Expected:     0x3eA,
		},
		{
			Text: []byte("0x3Ea"),
			Expected:     0x3Ea,
		},



		{
			Text: []byte("0x000000000000000"),
			Expected:     0x000000000000000,
		},
		{
			Text: []byte("0x000000000000001"),
			Expected:     0x000000000000001,
		},
		{
			Text: []byte("0x000000000000002"),
			Expected:     0x000000000000002,
		},
		{
			Text: []byte("0x000000000000003"),
			Expected:     0x000000000000003,
		},
		{
			Text: []byte("0x000000000000004"),
			Expected:     0x000000000000004,
		},
		{
			Text: []byte("0x000000000000005"),
			Expected:     0x000000000000005,
		},
		{
			Text: []byte("0x000000000000006"),
			Expected:     0x000000000000006,
		},
		{
			Text: []byte("0x000000000000007"),
			Expected:     0x000000000000007,
		},
		{
			Text: []byte("0x000000000000008"),
			Expected:     0x000000000000008,
		},
		{
			Text: []byte("0x000000000000009"),
			Expected:     0x000000000000009,
		},

		{
			Text: []byte("0x00000000000000a"),
			Expected:     0x00000000000000a,
		},
		{
			Text: []byte("0x00000000000000A"),
			Expected:     0x00000000000000A,
		},

		{
			Text: []byte("0x00000000000000b"),
			Expected:     0x00000000000000b,
		},
		{
			Text: []byte("0x00000000000000B"),
			Expected:     0x00000000000000B,
		},

		{
			Text: []byte("0x00000000000000c"),
			Expected:     0x00000000000000c,
		},
		{
			Text: []byte("0x00000000000000C"),
			Expected:     0x00000000000000C,
		},

		{
			Text: []byte("0x00000000000000d"),
			Expected:     0x00000000000000d,
		},
		{
			Text: []byte("0x00000000000000D"),
			Expected:     0x00000000000000D,
		},

		{
			Text: []byte("0x00000000000000e"),
			Expected:     0x00000000000000e,
		},
		{
			Text: []byte("0x00000000000000E"),
			Expected:     0x00000000000000E,
		},

		{
			Text: []byte("0x00000000000000f"),
			Expected:     0x00000000000000f,
		},
		{
			Text: []byte("0x00000000000000F"),
			Expected:     0x00000000000000F,
		},



		{
			Text: []byte("0xfedcbaABCDEF012"),
			Expected:     0xfedcbaABCDEF012,
		},



		{
			Text: []byte("0x0000000000000000"),
			Expected:     0x0000000000000000,
		},
		{
			Text: []byte("0x0000000000000001"),
			Expected:     0x0000000000000001,
		},
		{
			Text: []byte("0x0000000000000002"),
			Expected:     0x0000000000000002,
		},
		{
			Text: []byte("0x0000000000000003"),
			Expected:     0x0000000000000003,
		},
		{
			Text: []byte("0x0000000000000004"),
			Expected:     0x0000000000000004,
		},
		{
			Text: []byte("0x0000000000000005"),
			Expected:     0x0000000000000005,
		},
		{
			Text: []byte("0x0000000000000006"),
			Expected:     0x0000000000000006,
		},
		{
			Text: []byte("0x0000000000000007"),
			Expected:     0x0000000000000007,
		},
		{
			Text: []byte("0x0000000000000008"),
			Expected:     0x0000000000000008,
		},
		{
			Text: []byte("0x0000000000000009"),
			Expected:     0x0000000000000009,
		},

		{
			Text: []byte("0x000000000000000a"),
			Expected:     0x000000000000000a,
		},
		{
			Text: []byte("0x000000000000000A"),
			Expected:     0x000000000000000A,
		},

		{
			Text: []byte("0x000000000000000b"),
			Expected:     0x000000000000000b,
		},
		{
			Text: []byte("0x000000000000000B"),
			Expected:     0x000000000000000B,
		},

		{
			Text: []byte("0x000000000000000c"),
			Expected:     0x000000000000000c,
		},
		{
			Text: []byte("0x000000000000000C"),
			Expected:     0x000000000000000C,
		},

		{
			Text: []byte("0x000000000000000d"),
			Expected:     0x000000000000000d,
		},
		{
			Text: []byte("0x000000000000000D"),
			Expected:     0x000000000000000D,
		},

		{
			Text: []byte("0x000000000000000e"),
			Expected:     0x000000000000000e,
		},
		{
			Text: []byte("0x000000000000000E"),
			Expected:     0x000000000000000E,
		},

		{
			Text: []byte("0x000000000000000f"),
			Expected:     0x000000000000000f,
		},
		{
			Text: []byte("0x000000000000000F"),
			Expected:     0x000000000000000F,
		},



		{
			Text: []byte("0x1abcdefFEDCBA032"),
			Expected:     0x1abcdefFEDCBA032,
		},



		{
			Text: []byte("0b0"),
			Expected:       0,
		},
		{
			Text: []byte("0b1"),
			Expected:       1,
		},



		{
			Text: []byte("0b00"),
			Expected:        0,
		},
		{
			Text: []byte("0b01"),
			Expected:        1,
		},
		{
			Text: []byte("0b10"),
			Expected:        2,
		},
		{
			Text: []byte("0b11"),
			Expected:        3,
		},



		{
			Text: []byte("0b000"),
			Expected:         0,
		},
		{
			Text: []byte("0b001"),
			Expected:         1,
		},
		{
			Text: []byte("0b010"),
			Expected:         2,
		},
		{
			Text: []byte("0b011"),
			Expected:         3,
		},
		{
			Text: []byte("0b100"),
			Expected:         4,
		},
		{
			Text: []byte("0b101"),
			Expected:         5,
		},
		{
			Text: []byte("0b110"),
			Expected:         6,
		},
		{
			Text: []byte("0b111"),
			Expected:         7,
		},



		{
			Text: []byte("0b0000"),
			Expected:          0,
		},
		{
			Text: []byte("0b0001"),
			Expected:          1,
		},
		{
			Text: []byte("0b0010"),
			Expected:          2,
		},
		{
			Text: []byte("0b0011"),
			Expected:          3,
		},
		{
			Text: []byte("0b0100"),
			Expected:          4,
		},
		{
			Text: []byte("0b0101"),
			Expected:          5,
		},
		{
			Text: []byte("0b0110"),
			Expected:          6,
		},
		{
			Text: []byte("0b0111"),
			Expected:          7,
		},
		{
			Text: []byte("0b1000"),
			Expected:          8,
		},
		{
			Text: []byte("0b1001"),
			Expected:          9,
		},
		{
			Text: []byte("0b1010"),
			Expected:         10,
		},
		{
			Text: []byte("0b1011"),
			Expected:         11,
		},
		{
			Text: []byte("0b1100"),
			Expected:         12,
		},
		{
			Text: []byte("0b1101"),
			Expected:         13,
		},
		{
			Text: []byte("0b1110"),
			Expected:         14,
		},
		{
			Text: []byte("0b1111"),
			Expected:         15,
		},



		{
			Text: []byte("0b000000000000000000000000000000000000000000000000000000000000000"),
			Expected:                                                                     0,
		},
		{
			Text: []byte("0b000000000000000000000000000000000000000000000000000000000000001"),
			Expected:                                                                     1,
		},
		{
			Text: []byte("0b000000000000000000000000000000000000000000000000000000000000010"),
			Expected:                                                                     2,
		},
		{
			Text: []byte("0b000000000000000000000000000000000000000000000000000000000000011"),
			Expected:                                                                     3,
		},
		{
			Text: []byte("0b000000000000000000000000000000000000000000000000000000000000100"),
			Expected:                                                                     4,
		},
		{
			Text: []byte("0b000000000000000000000000000000000000000000000000000000000000101"),
			Expected:                                                                     5,
		},
		{
			Text: []byte("0b000000000000000000000000000000000000000000000000000000000000110"),
			Expected:                                                                     6,
		},
		{
			Text: []byte("0b000000000000000000000000000000000000000000000000000000000000111"),
			Expected:                                                                     7,
		},
		{
			Text: []byte("0b000000000000000000000000000000000000000000000000000000000001000"),
			Expected:                                                                     8,
		},
		{
			Text: []byte("0b000000000000000000000000000000000000000000000000000000000001001"),
			Expected:                                                                     9,
		},
		{
			Text: []byte("0b000000000000000000000000000000000000000000000000000000000001010"),
			Expected:                                                                    10,
		},
		{
			Text: []byte("0b000000000000000000000000000000000000000000000000000000000001011"),
			Expected:                                                                    11,
		},
		{
			Text: []byte("0b000000000000000000000000000000000000000000000000000000000001100"),
			Expected:                                                                    12,
		},
		{
			Text: []byte("0b000000000000000000000000000000000000000000000000000000000001101"),
			Expected:                                                                    13,
		},
		{
			Text: []byte("0b000000000000000000000000000000000000000000000000000000000001110"),
			Expected:                                                                    14,
		},
		{
			Text: []byte("0b000000000000000000000000000000000000000000000000000000000001111"),
			Expected:                                                                    15,
		},
		{
			Text: []byte("0b000000000000000000000000000000000000000000000000000000000010000"),
			Expected:                                                                    16,
		},
		{
			Text: []byte("0b000000000000000000000000000000000000000000000000000000000010001"),
			Expected:                                                                    17,
		},

		{
			Text: []byte("0b010010001000010000010000001000000010000000010000000001000000000"),
			Expected:                                                   2612659598909768192,
		},

		{
			Text: []byte("0b010101010101010101010101010101010101010101010101010101010101010"),
			Expected:                                                   3074457345618258602,
		},
	}


	for testNumber, test := range tests {

		actual, err := Int64(test.Text)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		if expected := test.Expected; expected != actual {
			t.Errorf("For test #%d, expected %d, but actually got %d.", testNumber, expected, actual)
			continue
		}
	}
}
