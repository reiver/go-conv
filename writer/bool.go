package convwriter

import (
	"io"
)

func Bool(writer io.Writer, value bool) (int64, error) {

	if nil == writer {
		return 0, errNilWriter
	}

	var t [4]byte = [4]byte{'t','r','u','e'}
	var f [5]byte = [5]byte{'f','a','l','s','e'}

	var n64 int64
	var err error

	switch value {
	case true:
		var n int
		n, err = writer.Write(t[:])
		n64 = int64(n)
	default: // case false
		var n int
		n, err = writer.Write(f[:])
		n64 = int64(n)
	}

	return n64, err
}
