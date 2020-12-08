package functools

import (
	"reflect"
)

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
			rfn := reflect.ValueOf(fn)
			rargs = rfn.Call(rargs)

			// functools.Unwrap() returns a slice of interfaces that captures
			// the arguments for the next function. Do some reflection magic to
			// convert it to a slice of reflect.Value.
			if rfn.Pointer() == reflect.ValueOf(Unwrap).Pointer() {
				vals := rargs[0].Interface().([]interface{})
				rargs = make([]reflect.Value, len(vals))
				for i, val := range vals {
					rargs[i] = reflect.ValueOf(val)
				}
			}
		}
	}
}

// Unwrap takes variadic arguments where the last argument is an error. If the
// error is nil, it returns the rest of the arguments back. If the error is
// non-nil, it panics.
func Unwrap(args ...interface{}) []interface{} {
	l := len(args)
	if args[l-1] != nil {
		err := args[l-1].(error)
		if err != nil {
			panic(err)
		}
	}

	return args[:l-1]
}
