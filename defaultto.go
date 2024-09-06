package fallback

// DefaultTo is used to make a something default to a contingency-value.
type DefaultTo[T any] struct {
	Contingency T
}

func (receiver DefaultTo[T]) WrapBool(value T, found bool) T {
	if !found {
		return receiver.Contingency
	}

	return value
}

func (receiver DefaultTo[T]) WrapError(value T, err error) T {
	if nil != err {
		return receiver.Contingency
	}

	return value
}
