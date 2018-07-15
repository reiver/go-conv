# go-conv

Package **conv** provides tools for converting text into Go data types.

This includes text stored in `[]byte`, `[]rune`, and `string`.

For example, converting a string to an (Golang) integer.
Or converting a string to a (Golang) boolean.
Etc.


## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-conv

[![GoDoc](https://godoc.org/github.com/reiver/go-conv?status.svg)](https://godoc.org/github.com/reiver/go-conv)


## Example
```go
import "github.com/reiver/go-conv/bytes"

// ...

var text []bytes = []byte("-25")

// ...

value, err := convbytes.Int64(text)
```

```go
import "github.com/reiver/go-conv/bytes"

// ...

var text []bytes = []byte("true")

// ...

value, err := convbytes.Bool(text)
```
