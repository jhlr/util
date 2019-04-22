package util

import "reflect"

// TypeError is used when a function
// is called with an unsupported type for it
type TypeError struct{}

func (e TypeError) Error() string {
	return "unexpected argument type"
}

// BreakSignal will break any loop if panicked
type BreakSignal struct{}

// Break will seemlessly stop the For
func Break() {
	panic(BreakSignal{})
}

// NoCopy will unable making copies of the
// struct it is part of
type NoCopy struct{}

// Lock is to implement sync.Locker
func (l NoCopy) Lock() {}

// UnLock is to implement sync.Locker
func (l NoCopy) UnLock() {}

// valueOf will return the reflect.Value of the object
// or return it back if it was already a reflect.Value
func valueOf(i interface{}) reflect.Value {
	v, ok := i.(reflect.Value)
	if !ok {
		v = reflect.ValueOf(i)
	}
	return v
}
