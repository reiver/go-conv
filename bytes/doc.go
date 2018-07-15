/*
Package convbytes provides tools for converting text into builtin and custom Go data types.


Example

Here is an example of a common usage of this package:

	import "github.com/reiver/go-conv/bytes"
	
	// ...
	
	var text []bytes = []byte("-25")
	
	// ...
	
	value, err := convbytes.Int64(text)


Converting A String To An Boolean

In boolean algebra, there are conceptually 2 literals:

• false

• true

Humans write these 2 boolean literals, in text, in a number of different ways.

For example, common ways that “false” gets written are:

• false

• f

• no

• off

• 0

• ٠

• ۰

• ⊥

(These are all case-insensitive. So this “false”, “FALSE”, “False”, “fAlSe”, etc are all the same.)

Also, for example, common ways that “true” gets written are:

• true

• t

• yes

• on

• 1

• ١

• ۱

• ⊤

(These are all case-insensitive. So this “true”, “TRUE”, “True”, “tRuE”, etc are all the same.)

This package provides a tool that can help you convert text (stored in a []byte) into a Go bool

For example:

	import "github.com/reiver/go-conv/bytes"
	
	// ...
	
	var text []bytes = []byte("True")
	
	// ...
	
	value, err := convbytes.Bool(text)

*/
package convbytes
