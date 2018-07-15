package convbytes

import (
	"fmt"
)

type InterpretationError struct {
	value   string
	reason  string
	context string
}

func (receiver InterpretationError) Error() string {
	return fmt.Sprintf("conv: [%s] interpretation error of %q: %s", receiver.context, receiver.value, receiver.reason)
}

func (receiver InterpretationError) Context() string {
	return receiver.context
}

func (receiver InterpretationError) Reason() string {
	return receiver.reason
}

func (receiver InterpretationError) Value() string {
	return receiver.value
}
