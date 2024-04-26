// Package mapper provides functionality for mapping data between different types.
//
// This package defines a generic MapFunc type, which represents a function for mapping values from type T to type U.
// It also includes convenience methods for mapping individual values and slices of values.
package mapper

type MapFunc[T any, U any] func(T) U

// Map is a aliased call to the underlying function, it is optional
// to define this method, but it can be useful for readability.
func (a MapFunc[T, U]) Map(v T) U {
	return a(v)
}

// MapEach is a convenience method for mapping a slice of items to a slice of
// the same length of a different type.
func (a MapFunc[T, U]) MapEach(v []T) []U {
	result := make([]U, len(v))
	for i, item := range v {
		result[i] = a(item)
	}
	return result
}

// MapErr is a convenience method for mapping a value to a different type
// and returning an error if one is provided.
func (a MapFunc[T, U]) MapErr(v T, err error) (U, error) {
	if err != nil {
		var zero U
		return zero, err
	}

	return a(v), nil
}

// MapEachErr is a convenience method for mapping a slice of items to a slice of
// the same length of a different type, and returning an error if one is provided.
func (a MapFunc[T, U]) MapEachErr(v []T, err error) ([]U, error) {
	if err != nil {
		return nil, err
	}

	return a.MapEach(v), nil
}
