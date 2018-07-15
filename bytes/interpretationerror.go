package convbytes

import (
	"fmt"
)

type InterpretationError struct {
	value  string
	reason string
}

func (receiver InterpretationError) Error() string {
	return fmt.Sprintf("conv: interpretation error of %q: %s", receiver.value, receiver.reason)
}
