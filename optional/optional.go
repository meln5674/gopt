package optional

// Optional represents a value that may not may not be present.
// Unlike pointers, Optional lives on the stack, can can be copied
type Optional[T any] struct {
	value   T
	present bool
}

// Of returns an optional of a valud that is present
func Of[T any](value T) Optional[T] {
	return Optional[T]{value: value, present: true}
}

// Empty returns an optional that is missing
func Empty[T any]() Optional[T] {
	return Optional[T]{}
}

// OfPointer returns an optional that is present if the pointer non-nil, otherwise one that is missing
func OfPointer[T any](ptr *T) Optional[T] {
	if ptr == nil {
		return Empty[T]()
	}
	return Of[T](*ptr)
}

// Present returns true if the optional value is present, false if it is missing
func (o *Optional[T]) Present() bool {
	return o.present
}

// GetOrPanic returns the underlying value if present, otherwise it panics
func (o *Optional[T]) GetOrPanic() T {
	if !o.present {
		panic("GetOrPanic called on missing Optional")
	}
	return o.value
}

// GetOrDefault returns the underlying value if present, otherwise a default value
func (o *Optional[T]) GetOrDefault(default_ T) T {
	if !o.present {
		return default_
	}
	return o.value
}

// AsPointer returns the address of the underlying value if it is present, otherwise nil
func (o *Optional[T]) AsPointer() *T {
	if !o.present {
		return nil
	}
	return &o.value
}

// AsCopyPointer returns the address of a shallow copy (assignment) of the underlying value if it is present, otherwise nil
func (o *Optional[T]) AsCopyPointer() *T {
	if !o.present {
		return nil
	}
	value := o.value
	return &value
}

// Apply calls a function with the underlying value and wraps it in an optional if present, otherwise returns a missing optional of the return type
func Apply[T any, U any](f func(T) U, o Optional[T]) Optional[U] {
	if !o.present {
		return Empty[U]()
	}
	return Of[U](f(o.value))
}

// ApplyPtr calls a function with the address of the underlying value and wraps it in an optional if present, otherwise returns a missing optional of the return type
func ApplyPtr[T any, U any](f func(*T) U, o *Optional[T]) Optional[U] {
	if !o.present {
		return Empty[U]()
	}
	return Of[U](f(&o.value))
}

// Map calls a function with an optional return value with the underlying value if present, otherwise returns a missing optional of the return type
func Map[T any, U any](f func(T) Optional[U], o Optional[T]) Optional[U] {
	if !o.present {
		return Empty[U]()
	}
	return f(o.value)
}

// Map calls a function with an optional return value with the address of the underlying value if present, otherwise returns a missing optional of the return type
func MapPtr[T any, U any](f func(*T) Optional[U], o *Optional[T]) Optional[U] {
	if !o.present {
		return Empty[U]()
	}
	return f(&o.value)
}
