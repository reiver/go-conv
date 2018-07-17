package convwriter

import (
	"io"
)

var (
	boolStyle_0_1_false = [...]byte{'0'}
	boolStyle_0_1_true  = [...]byte{'1'}

	boolStyle_F_T_false = [...]byte{'F'}
	boolStyle_F_T_true  = [...]byte{'T'}

	boolStyle_f_t_false = [...]byte{'f'}
	boolStyle_f_t_true  = [...]byte{'t'}

	boolStyle_up_down_false = []byte("⊥")
	boolStyle_up_down_true  = []byte("⊤")

	boolStyle_OFF_ON_false = [...]byte{'O','F','F'}
	boolStyle_OFF_ON_true  = [...]byte{'O','N'}

	boolStyle_off_on_false = [...]byte{'o','f','f'}
	boolStyle_off_on_true  = [...]byte{'o','n'}

	boolStyle_Off_On_false = [...]byte{'O','f','f'}
	boolStyle_Off_On_true  = [...]byte{'O','n'}

	boolStyle_NO_YES_false = [...]byte{'N','O'}
	boolStyle_NO_YES_true  = [...]byte{'Y','E','S'}

	boolStyle_no_yes_false = [...]byte{'n','o'}
	boolStyle_no_yes_true  = [...]byte{'y','e','s'}

	boolStyle_No_Yes_false = [...]byte{'N','o'}
	boolStyle_No_Yes_true  = [...]byte{'Y','e','s'}

	boolStyle_FALSE_TRUE_false = [...]byte{'F','A','L','S','E'}
	boolStyle_FALSE_TRUE_true  = [...]byte{'T','R','U','E'}

	boolStyle_default_false = [...]byte{'f','a','l','s','e'}
	boolStyle_default_true  = [...]byte{'t','r','u','e'}

	boolStyle_False_True_false = [...]byte{'F','a','l','s','e'}
	boolStyle_False_True_true  = [...]byte{'T','r','u','e'}
)

// Bool converts a bool into a human readable form.
//
// For example:
//
//	var writer io.Writer
//
//	//...
//
//	var value bool = true
//	
//	//...
//	
//	convwrite.Bool(writer, value)
//	
//	// Output:
//	// true
//
// Another example:
//
//	var writer io.Writer
//
//	//...
//
//	var value bool = true
//	
//	//...
//	
//	convwrite.Bool(writer, value, "Off")
//	
//	// Output:
//	// On
func Bool(writer io.Writer, value bool, params ...interface{}) (int64, error) {

	if nil == writer {
		return 0, errNilWriter
	}

	var f []byte
	var t []byte
	{
		var style interface{}

		if 0 < len(params) {
			param0 := params[0]

			style = param0
		}


		switch style {
		case '0','1',
		     "0","1":
			f = boolStyle_0_1_false[:]
			t = boolStyle_0_1_true[:]

		case 'F','T',
		     "F","T":
			f = boolStyle_F_T_false[:]
			t = boolStyle_F_T_true[:]


		case 'f','t',
		     "f","t":
			f = boolStyle_f_t_false[:]
			t = boolStyle_f_t_true[:]

		case '⊥','⊤',
		     "⊥","⊤":
			f = boolStyle_up_down_false
			t = boolStyle_up_down_true

		case "OFF","ON":
			f = boolStyle_OFF_ON_false[:]
			t = boolStyle_OFF_ON_true[:]

		case "off","on":
			f = boolStyle_off_on_false[:]
			t = boolStyle_off_on_true[:]

		case "Off","On":
			f = boolStyle_Off_On_false[:]
			t = boolStyle_Off_On_true[:]

		case "NO","YES":
			f = boolStyle_NO_YES_false[:]
			t = boolStyle_NO_YES_true[:]

		case "no","yes":
			f = boolStyle_no_yes_false[:]
			t = boolStyle_no_yes_true[:]

		case "No","Yes":
			f = boolStyle_No_Yes_false[:]
			t = boolStyle_No_Yes_true[:]

		case "FALSE","TRUE":
			f = boolStyle_FALSE_TRUE_false[:]
			t = boolStyle_FALSE_TRUE_true[:]

		case "False","True":
			f = boolStyle_False_True_false[:]
			t = boolStyle_False_True_true[:]

		default:
			f = boolStyle_default_false[:]
			t = boolStyle_default_true[:]
		}
	}



	var n64 int64
	var err error

	switch value {
	case true:
		var n int
		n, err = writer.Write(t)
		n64 = int64(n)
	default: // case false
		var n int
		n, err = writer.Write(f)
		n64 = int64(n)
	}

	return n64, err
}
