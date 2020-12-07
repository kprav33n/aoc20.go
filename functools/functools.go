package functools

import "reflect"

// Compose composes the given functions by passing the returned values of the
// preceeding function as arguments to the subsequent function. This is a
// minimal implementation and doesn't perform any safety checks.
func Compose(fns ...interface{}) func(args ...interface{}) {
	return func(args ...interface{}) {
		rargs := make([]reflect.Value, len(args))
		for i, arg := range args {
			rargs[i] = reflect.ValueOf(arg)
		}

		for _, fn := range fns {
			rargs = reflect.ValueOf(fn).Call(rargs)
		}
	}
}
