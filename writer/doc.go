/*
Package convwriter provides tools for converting builtin and custom Go data types into text.


Example

Here is an example of a common usage of this package:

	import "github.com/reiver/go-conv/writer"
	
	// ...

	var writer io.Writer

	//...
	
	var value uint64 = 25
	
	// ...
	
	value, err := convwriter.Uint64(writer, value)

Here is another similar example:

	import "github.com/reiver/go-conv/writer"
	
	// ...

	var writer io.Writer

	//...
	
	var value bool = true
	
	// ...
	
	value, err := convwriter.Bool(writer, value)

*/
package convwriter
