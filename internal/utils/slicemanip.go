package utils

import (
	"iter"
	"slices"
)

// ---- Slice-based API ----

// SliceMap transforms each element in a slice according to the provided function
func SliceMap[T, U any](slice []T, f func(T) U) []U {
	result := make([]U, len(slice))
	for i, v := range slice {
		result[i] = f(v)
	}
	return result
}

// SliceFilter returns elements from a slice that satisfy the predicate function
func SliceFilter[T any](slice []T, predicate func(T) bool) []T {
	result := make([]T, 0)
	for _, v := range slice {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

// SliceFind returns the first element that satisfies the predicate function
// Returns the value and a boolean indicating if an element was found
func SliceFind[T any](slice []T, predicate func(T) bool) (T, bool) {
	for _, v := range slice {
		if predicate(v) {
			return v, true
		}
	}
	var zero T
	return zero, false
}

// SliceReduce applies a function against an accumulator and each element in the slice
func SliceReduce[T, U any](slice []T, initialValue U, reducer func(acc U, current T) U) U {
	result := initialValue
	for _, v := range slice {
		result = reducer(result, v)
	}
	return result
}

// SliceForEach executes a provided function once for each slice element
func SliceForEach[T any](slice []T, action func(T)) {
	for _, v := range slice {
		action(v)
	}
}

// SliceSome tests whether at least one element satisfies the provided testing function
func SliceSome[T any](slice []T, predicate func(T) bool) bool {
	return slices.ContainsFunc(slice, predicate)
}

// SliceEvery tests whether all elements satisfy the provided testing function
func SliceEvery[T any](slice []T, predicate func(T) bool) bool {
	for _, v := range slice {
		if !predicate(v) {
			return false
		}
	}
	return true
}

// SliceIncludes determines whether a slice includes a certain value
func SliceIncludes[T comparable](slice []T, valueToFind T) bool {
	return slices.Contains(slice, valueToFind)
}

// SliceFlatMap maps each element using a mapping function, then flattens the result
func SliceFlatMap[T, U any](slice []T, f func(T) []U) []U {
	var result []U
	for _, v := range slice {
		result = append(result, f(v)...)
	}
	return result
}

// SliceChunk splits a slice into chunks of the specified size
func SliceChunk[T any](slice []T, size int) [][]T {
	if size <= 0 {
		panic("chunk size must be greater than 0")
	}

	result := make([][]T, 0, (len(slice)+size-1)/size)
	for i := 0; i < len(slice); i += size {
		end := i + size
		if end > len(slice) {
			end = len(slice)
		}
		result = append(result, slice[i:end])
	}
	return result
}

// ---- Iterator-based API (original) ----

// Map transforms each element in a sequence according to the provided function
func Map[T, U any](seq iter.Seq[T], f func(T) U) iter.Seq[U] {
	return func(yield func(U) bool) {
		for a := range seq {
			if !yield(f(a)) {
				return
			}
		}
	}
}

// Filter returns elements from a sequence that satisfy the predicate function
func Filter[T any](seq iter.Seq[T], predicate func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for a := range seq {
			if predicate(a) {
				if !yield(a) {
					return
				}
			}
		}
	}
}

// Find returns the first element that satisfies the predicate function
func Find[T any](seq iter.Seq[T], predicate func(T) bool) (T, bool) {
	var result T
	found := false

	for a := range seq {
		if predicate(a) {
			result = a
			found = true
			break
		}
	}

	return result, found
}

// Helper functions for iter.Seq conversions
func ToSlice[T any](seq iter.Seq[T]) []T {
	result := []T{}
	for a := range seq {
		result = append(result, a)
	}
	return result
}

func FromSlice[T any](slice []T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, v := range slice {
			if !yield(v) {
				return
			}
		}
	}
}
