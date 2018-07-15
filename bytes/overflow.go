package convbytes

import (
	"fmt"
)

type Overflow struct {
	value   string
	reason  string
	context string
}

func (receiver Overflow) Error() string {
	return fmt.Sprintf("conv: [%s] overflow caused by %q: %s", receiver.context, receiver.value, receiver.reason)
}

func (receiver Overflow) Context() string {
	return receiver.context
}

func (receiver Overflow) Reason() string {
	return receiver.reason
}

func (receiver Overflow) Value() string {
	return receiver.value
}
