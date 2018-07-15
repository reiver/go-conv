package convbytes

import (
	"fmt"
	"unicode/utf8"
)

func Uint64(data []byte) (uint64, error) {

	if 0 >= len(data) {
		return 0, InterpretationError{
			context: "uint64",
			value:   string(data),
			reason:  "no data",
		}
	}


	var r0  rune
	var p []byte
	{
		var size int

		r0,size = utf8.DecodeRune(data)
		if utf8.RuneError == r0 {
			return 0, InterpretationError{
				context: "uint64",
				value:   string(data),
				reason: "invalid UTF-8 data",
			}
		}

		p = data[size:]
	}

	switch r0 {
	case '0':
		if 1 == len(data) {
			return 0, nil
		}

		r1, size := utf8.DecodeRune(p)
		if utf8.RuneError == r1 {
			return 0, InterpretationError{
				context: "uint64",
				value: string(data),
				reason: "invalid UTF-8 data",
			}
		}

		switch r1 {
		case 'x':
			return uint64Hexadecimal(p[size:])

		case 'b':
			return uint64Binary(p[size:])

		default:
			return 0, InterpretationError{
				context: "uint64",
				value:   string(data),
				reason:  fmt.Sprintf("invalid or unrecognized literal"),
			}
		}

	case '1','2','3','4','5','6','7','8','9': // Digits
		return uint64Digits(data)

	case '٠','١','٢','٣','٤','٥','٦','٧','٨','٩': // Arabic-Indic Digits
		return uint64ArabicIndicDigits(data)

	case '۰','۱','۲','۳','۴','۵','۶','۷','۸','۹': // Extended Arabic-Indic Digits
		return uint64ExtendedArabicIndicDigits(data)

	case 'Ⅰ','Ⅱ','Ⅲ','Ⅳ','Ⅴ','Ⅵ','Ⅶ','Ⅷ','Ⅸ','Ⅹ','Ⅺ','Ⅻ','Ⅼ','Ⅽ','Ⅾ','Ⅿ': // Roman Numerals
		return uint64RomanNumerals(data)

	case 'ⅰ','ⅱ','ⅲ','ⅳ','ⅴ','ⅵ','ⅶ','ⅷ','ⅸ','ⅹ','ⅺ','ⅻ','ⅼ','ⅽ','ⅾ','ⅿ': // Roman Numerals
		return uint64SmallRomanNumerals(data)

	default:
		return 0, InterpretationError{
			context: "uint64",
			value:   string(data),
			reason:  "invalid or unrecognized literal",
		}
	}
}

func uint64Digits(data []byte) (uint64, error) {

	if 0 >= len(data) {
		return 0, InterpretationError{
			context: "uint64",
			value:   string(data),
			reason:  "no data",
		}
	}

	var result uint64

	for _, symbol := range data {
		switch {
		case  '0' <= symbol && symbol <= '9':

			x := uint64(symbol - '0')

			result *= 10
			result += x

		default:
			return 0, InterpretationError{
				context: "uint64",
				value:   string(data),
				reason:  fmt.Sprintf("expected only digits (i.e., “0”, “1”, “2”, “3”, “4”, “5”, “6”, “7”, “8”, “9”), but found %q", symbol),
			}
		}
	}

	return result, nil
}

func uint64ArabicIndicDigits(data []byte) (uint64, error) {

	if 0 >= len(data) {
		return 0, InterpretationError{
			context: "uint64",
			value:   string(data),
			reason:  "no data",
		}
	}

	var result uint64
	{
		var p []byte = data
		for 0 < len(p) {
			symbol, size := utf8.DecodeRune(p)
			if utf8.RuneError == symbol {
				return 0, InterpretationError{
					context: "uint64",
					value:   string(data),
					reason:  "invalid UTF-8 data",
				}
			}
			p = p[size:]

			switch {
			case  '٠' <= symbol && symbol <= '٩':

				x := uint64(symbol - '٠')

				result *= 10
				result += x

			default:
				return 0, InterpretationError{
					context: "uint64",
					value:   string(data),
					reason:  fmt.Sprintf("expected only arabic-indic digits (i.e., “٠”, “١”, “٢”, “٣”, “٤”, “٥”, “٦”, “٧”, “٨”, “٩”), but found %q", symbol),
				}
			}
		}
	}

	return result, nil
}

func uint64ExtendedArabicIndicDigits(data []byte) (uint64, error) {

	if 0 >= len(data) {
		return 0, InterpretationError{
			context: "uint64",
			value:   string(data),
			reason:  "no data",
		}
	}

	var result uint64
	{
		var p []byte = data
		for 0 < len(p) {
			symbol, size := utf8.DecodeRune(p)
			if utf8.RuneError == symbol {
				return 0, InterpretationError{
					context: "uint64",
					value:   string(data),
					reason:  "invalid UTF-8 data",
				}
			}
			p = p[size:]

			switch {
			case  '۰' <= symbol && symbol <= '۹':

				x := uint64(symbol - '۰')

				result *= 10
				result += x

			default:
				return 0, InterpretationError{
					context: "uint64",
					value:   string(data),
					reason:  fmt.Sprintf("expected extended arabic-indic digits (i.e., “۰”, “۱”, “۲”, “۳”, “۴”, “۵”, “۶”, “۷”, “۸”, “۹”), but found %q", symbol),
				}
			}
		}
	}

	return result, nil
}

func uint64Hexadecimal(data []byte) (uint64, error) {

	if 0 >= len(data) {
		return 0, InterpretationError{
			context: "uint64",
			value:   string(data),
			reason:  "no data",
		}
	}

	var result uint64

	for _, symbol := range data {
		switch {
		case  '0' <= symbol && symbol <= '9':

			x := uint64(symbol - '0')

			result *= 16
			result += x

		case  'A' <= symbol && symbol <= 'F':

			x := uint64(symbol - 'A' + 10)

			result *= 16
			result += x

		case  'a' <= symbol && symbol <= 'f':

			x := uint64(symbol - 'a' + 10)

			result *= 16
			result += x

		default:
			return 0, InterpretationError{
				context: "uint64",
				value:   string(data),
				reason:  fmt.Sprintf("expected hexadecimal digits (i.e., “0”, “1”, “2”, “3”, “4”, “5”, “6”, “7”, “8”, “9”, “A”, “a”, “B”, “b”, “C”, “c”, “D”, “d”, “E”, “e”, “F”, “f”), but found %q", symbol),
			}
		}
	}

	return result, nil
}

func uint64Binary(data []byte) (uint64, error) {

	if 0 >= len(data) {
		return 0, InterpretationError{
			context: "uint64",
			value:   string(data),
			reason:  "no data",
		}
	}

	var result uint64

	for _, symbol := range data {
		switch symbol {
		case  '0','1':

			x := uint64(symbol - '0')

			result *= 2
			result += x

		default:
			return 0, InterpretationError{
				context: "uint64",
				value:   string(data),
				reason:  fmt.Sprintf("expected binary digits (i.e., “0”, “1”), but found %q", symbol),
			}
		}
	}

	return result, nil
}

func uint64RomanNumerals(data []byte) (uint64, error) {

	if 0 >= len(data) {
		return 0, InterpretationError{
			context: "uint64",
			value:   string(data),
			reason:  "no data",
		}
	}

	{
		var p []byte = data

		symbol, size := utf8.DecodeRune(p)
		if utf8.RuneError == symbol {
			return 0, InterpretationError{
				context: "uint64",
				value:   string(data),
				reason:  "invalid UTF-8 data",
			}
		}
		p = p[size:]

		if 0 == len(p) {
			return romanNumeral(symbol)
		}
	}

	var result uint64
	{
		var prev uint64
		var p []byte = data
		for 0 < len(p) {
			symbol, size := utf8.DecodeRune(p)
			if utf8.RuneError == symbol {
				return 0, InterpretationError{
					context: "uint64",
					value:   string(data),
					reason:  "invalid UTF-8 data",
				}
			}
			p = p[size:]

			value, err := romanNumeral(symbol)
			if nil != err {
				return 0, err
			}

			if 0 != prev && prev >= value {
				result += value
			} else {
				result -= prev
				result += value
				result -= prev
			}

			prev = value
		}
	}

	return result, nil
}

func romanNumeral(symbol rune) (uint64, error) {
	switch symbol {
	case 'Ⅰ':
		return 1, nil

	case 'Ⅱ':
		return 2, nil

	case 'Ⅲ':
		return 3, nil

	case 'Ⅳ':
		return 4, nil

	case 'Ⅴ':
		return 5, nil

	case 'Ⅵ':
		return 6, nil

	case 'Ⅶ':
		return 7, nil

	case 'Ⅷ':
		return 8, nil

	case 'Ⅸ':
		return 9, nil

	case 'Ⅹ':
		return 10, nil

	case 'Ⅺ':
		return 11, nil

	case 'Ⅻ':
		return 12, nil

	case 'Ⅼ':
		return 50, nil

	case 'Ⅽ':
		return 100, nil

	case 'Ⅾ':
		return 500, nil

	case 'Ⅿ':
		return 1000, nil

	default:
		return 0, InterpretationError{
			value: string(symbol),
			reason: "invalid roman numerals",
		}
	}
}

func uint64SmallRomanNumerals(data []byte) (uint64, error) {

	if 0 >= len(data) {
		return 0, InterpretationError{
			context: "uint64",
			value:   string(data),
			reason:  "no data",
		}
	}

	{
		var p []byte = data

		symbol, size := utf8.DecodeRune(p)
		if utf8.RuneError == symbol {
			return 0, InterpretationError{
				context: "uint64",
				value:   string(data),
				reason:  "invalid UTF-8 data",
			}
		}
		p = p[size:]

		if 0 == len(p) {
			return smallRomanNumeral(symbol)
		}
	}

	var result uint64
	{
		var prev uint64
		var p []byte = data
		for 0 < len(p) {
			symbol, size := utf8.DecodeRune(p)
			if utf8.RuneError == symbol {
				return 0, InterpretationError{
					context: "uint64",
					value:   string(data),
					reason:  "invalid UTF-8 data",
				}
			}
			p = p[size:]

			value, err := smallRomanNumeral(symbol)
			if nil != err {
				return 0, err
			}

			if 0 != prev && prev >= value {
				result += value
			} else {
				result -= prev
				result += value
				result -= prev
			}

			prev = value
		}
	}

	return result, nil
}

func smallRomanNumeral(symbol rune) (uint64, error) {
	switch symbol {
	case 'ⅰ':
		return 1, nil

	case 'ⅱ':
		return 2, nil

	case 'ⅲ':
		return 3, nil

	case 'ⅳ':
		return 4, nil

	case 'ⅴ':
		return 5, nil

	case 'ⅵ':
		return 6, nil

	case 'ⅶ':
		return 7, nil

	case 'ⅷ':
		return 8, nil

	case 'ⅸ':
		return 9, nil

	case 'ⅹ':
		return 10, nil

	case 'ⅺ':
		return 11, nil

	case 'ⅻ':
		return 12, nil

	case 'ⅼ':
		return 50, nil

	case 'ⅽ':
		return 100, nil

	case 'ⅾ':
		return 500, nil

	case 'ⅿ':
		return 1000, nil

	default:
		return 0, InterpretationError{
			value: string(symbol),
			reason: "invalid small roman numerals",
		}
	}
}
