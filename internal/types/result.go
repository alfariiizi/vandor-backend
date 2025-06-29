package types

import (
	"encoding/json"
	"fmt"
	"runtime"
	"strings"
)

// Result represents either a success value or an error, similar to Rust's Result<T, E>
type Result[T any] struct {
	value T
	err   *Error
}

// Error represents an error with stack trace information
type Error struct {
	message string
	cause   error
	stack   []StackFrame
}

// StackFrame represents a single frame in the stack trace
type StackFrame struct {
	Function string `json:"function"`
	File     string `json:"file"`
	Line     int    `json:"line"`
}

// Constructor functions

// Ok creates a successful Result with the given value
func Ok[T any](value T) Result[T] {
	return Result[T]{value: value, err: nil}
}

// Err creates a failed Result with the given error
func Err[T any](err error) Result[T] {
	return Result[T]{err: newError(err)}
}

// Errf creates a failed Result with a formatted error message
func Errf[T any](format string, args ...any) Result[T] {
	return Result[T]{err: newError(fmt.Errorf(format, args...))}
}

// WrapErr creates a failed Result by wrapping an existing error with additional context
func WrapErr[T any](err error, message string) Result[T] {
	if err == nil {
		return Ok(*new(T))
	}
	return Result[T]{err: wrapError(err, message)}
}

// Core methods

// IsOk returns true if the Result contains a value
func (r Result[T]) IsOk() bool {
	return r.err == nil
}

// IsErr returns true if the Result contains an error
func (r Result[T]) IsErr() bool {
	return r.err != nil
}

// Unwrap returns the value if Ok, zero value if Err
func (r Result[T]) Unwrap() T {
	if r.IsErr() {
		var zero T
		return zero
	}
	return r.value
}

// MustUnwrap returns the value if Ok, panics if Err (for cases where panic is desired)
func (r Result[T]) MustUnwrap() T {
	if r.IsErr() {
		panic(fmt.Sprintf("called MustUnwrap on Err value: %v", r.err))
	}
	return r.value
}

// UnwrapOr returns the value if Ok, or the default value if Err
func (r Result[T]) UnwrapOr(defaultValue T) T {
	if r.IsErr() {
		return defaultValue
	}
	return r.value
}

// UnwrapOrElse returns the value if Ok, or calls the function if Err
func (r Result[T]) UnwrapOrElse(fn func(*Error) T) T {
	if r.IsErr() {
		return fn(r.err)
	}
	return r.value
}

// Expect returns the value if Ok, panics with custom message if Err
func (r Result[T]) Expect(message string) T {
	if r.IsErr() {
		panic(fmt.Sprintf("%s: %v", message, r.err))
	}
	return r.value
}

// UnwrapPtr returns a pointer to the value if Ok, nil if Err
// This is especially useful for pointer types to avoid double pointers
func (r Result[T]) UnwrapPtr() *T {
	if r.IsErr() {
		return nil
	}
	return &r.value
}

// UnwrapOk returns the value and true if Ok, zero value and false if Err
func (r Result[T]) UnwrapOk() (T, bool) {
	if r.IsErr() {
		var zero T
		return zero, false
	}
	return r.value, true
}

// Error returns the error if Err, nil if Ok
func (r Result[T]) Error() *Error {
	return r.err
}

// Transformation methods

// Map transforms the value if Ok, preserves error if Err
func Map[T, U any](r Result[T], fn func(T) U) Result[U] {
	if r.IsErr() {
		return Result[U]{err: r.err}
	}
	return Ok(fn(r.value))
}

// MapErr transforms the error if Err, preserves value if Ok
func (r Result[T]) MapErr(fn func(*Error) error) Result[T] {
	if r.IsOk() {
		return r
	}
	return Result[T]{err: newError(fn(r.err))}
}

// AndThen chains Results (flatMap equivalent)
func AndThen[T, U any](r Result[T], fn func(T) Result[U]) Result[U] {
	if r.IsErr() {
		return Result[U]{err: r.err}
	}
	return fn(r.value)
}

// OrElse provides an alternative Result if Err
func (r Result[T]) OrElse(fn func(*Error) Result[T]) Result[T] {
	if r.IsOk() {
		return r
	}
	return fn(r.err)
}

// Utility methods

// Match provides pattern matching-like functionality
func (r Result[T]) Match(onOk func(T), onErr func(*Error)) {
	if r.IsOk() {
		onOk(r.value)
	} else {
		onErr(r.err)
	}
}

// String implements the Stringer interface
func (r Result[T]) String() string {
	if r.IsOk() {
		return fmt.Sprintf("Ok(%v)", r.value)
	}
	return fmt.Sprintf("Err(%v)", r.err)
}

// MarshalJSON implements json.Marshaler
func (r Result[T]) MarshalJSON() ([]byte, error) {
	if r.IsOk() {
		return json.Marshal(map[string]any{
			"status": "ok",
			"value":  r.value,
		})
	}
	return json.Marshal(map[string]any{
		"status": "error",
		"error":  r.err.Error(),
		"stack":  r.err.stack,
	})
}

// Collect functions

// Collect converts a slice of Results into a Result of slice
// Returns the first error encountered, or Ok with all values
func Collect[T any](results []Result[T]) Result[[]T] {
	values := make([]T, len(results))
	for i, result := range results {
		if result.IsErr() {
			return Result[[]T]{err: result.err}
		}
		values[i] = result.value
	}
	return Ok(values)
}

// CollectAll converts a slice of Results into a Result of slice
// Collects all errors if any exist
func CollectAll[T any](results []Result[T]) Result[[]T] {
	var errors []*Error
	values := make([]T, 0, len(results))

	for _, result := range results {
		if result.IsErr() {
			errors = append(errors, result.err)
		} else {
			values = append(values, result.value)
		}
	}

	if len(errors) > 0 {
		return Result[[]T]{err: combineErrors(errors)}
	}
	return Ok(values)
}

// Error implementation

// Error implements the error interface
func (e *Error) Error() string {
	return e.message
}

// Unwrap returns the underlying cause
func (e *Error) Unwrap() error {
	return e.cause
}

// Stack returns the stack trace
func (e *Error) Stack() []StackFrame {
	return e.stack
}

// StackTrace returns a formatted stack trace string
func (e *Error) StackTrace() string {
	var sb strings.Builder
	sb.WriteString(e.message)
	sb.WriteString("\n\nStack trace:\n")

	for i, frame := range e.stack {
		sb.WriteString(fmt.Sprintf("  %d. %s\n", i+1, frame.Function))
		sb.WriteString(fmt.Sprintf("     %s:%d\n", frame.File, frame.Line))
	}

	if e.cause != nil {
		sb.WriteString(fmt.Sprintf("\nCaused by: %v", e.cause))
	}

	return sb.String()
}

// String implements the Stringer interface
func (e *Error) String() string {
	return e.StackTrace()
}

// MarshalJSON implements json.Marshaler for Error
func (e *Error) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"message": e.message,
		"stack":   e.stack,
		"cause":   e.cause,
	})
}

// Internal helper functions

func newError(err error) *Error {
	if err == nil {
		return nil
	}

	// If it's already our Error type, return it
	if e, ok := err.(*Error); ok {
		return e
	}

	return &Error{
		message: err.Error(),
		cause:   err,
		stack:   captureStack(2), // Skip newError and the calling function
	}
}

func wrapError(err error, message string) *Error {
	if err == nil {
		return nil
	}

	return &Error{
		message: message,
		cause:   err,
		stack:   captureStack(2), // Skip wrapError and the calling function
	}
}

func captureStack(skip int) []StackFrame {
	var frames []StackFrame

	// Capture up to 32 frames
	pcs := make([]uintptr, 32)
	n := runtime.Callers(skip+1, pcs)

	if n > 0 {
		callersFrames := runtime.CallersFrames(pcs[:n])
		for {
			frame, more := callersFrames.Next()

			// Skip runtime frames
			if !strings.Contains(frame.File, "runtime/") {
				frames = append(frames, StackFrame{
					Function: frame.Function,
					File:     frame.File,
					Line:     frame.Line,
				})
			}

			if !more {
				break
			}
		}
	}

	return frames
}

func combineErrors(errors []*Error) *Error {
	if len(errors) == 0 {
		return nil
	}
	if len(errors) == 1 {
		return errors[0]
	}

	var messages []string
	for _, err := range errors {
		messages = append(messages, err.Error())
	}

	return &Error{
		message: fmt.Sprintf("Multiple errors: %s", strings.Join(messages, "; ")),
		cause:   errors[0], // Use first error as primary cause
		stack:   captureStack(2),
	}
}

// Convenience functions for common patterns

// From converts a Go-style (value, error) return to Result in one line
func From[T any](value T, err error) Result[T] {
	if err != nil {
		return Err[T](err)
	}
	return Ok(value)
}

// Must converts a Go-style (value, error) to just the value, panicking on error
// Use this only when you're certain there won't be an error
func Must[T any](value T, err error) T {
	if err != nil {
		panic(fmt.Sprintf("Must failed: %v", err))
	}
	return value
}

// Safe converts a Go-style (value, error) to just the value, returning zero value on error
func Safe[T any](value T, err error) T {
	if err != nil {
		var zero T
		return zero
	}
	return value
}

// SafePtr converts a Go-style (value, error) to a pointer to value, returning nil on error
// This is perfect for pointer types to avoid double pointers
func SafePtr[T any](value T, err error) *T {
	if err != nil {
		return nil
	}
	return &value
}

// Try executes a function and converts panic to Result
func Try[T any](fn func() T) Result[T] {
	defer func() {
		if r := recover(); r != nil {
			// This will be handled by the return value
		}
	}()

	// Use a channel to safely capture the result or panic
	resultChan := make(chan Result[T], 1)

	go func() {
		defer func() {
			if r := recover(); r != nil {
				resultChan <- Err[T](fmt.Errorf("panic: %v", r))
			}
		}()

		result := fn()
		resultChan <- Ok(result)
	}()

	return <-resultChan
}

// FromGoError converts a Go-style (value, error) return to Result
func FromGoError[T any](value T, err error) Result[T] {
	if err != nil {
		return Err[T](err)
	}
	return Ok(value)
}

// ToGoError converts Result to Go-style (value, error) return
func (r Result[T]) ToGoError() (T, error) {
	if r.IsErr() {
		return *new(T), r.err
	}
	return r.value, nil
}
