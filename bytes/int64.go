package convbytes

func Int64(data []byte) (int64, error) {

	return 0, InterpretationError{
		value: string(data),
		reason: "invalid int64 literal",
	}
}
