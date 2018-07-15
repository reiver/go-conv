package convbytes

func Uint64(data []byte) (uint64, error) {

	return 0, InterpretationError{
		value: string(data),
		reason: "invalid uint64 literal",
	}
}
