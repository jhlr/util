package util

import "reflect"

// CheckFunc returns true if Foo meets conditions
//
// in
//   can have more elem than f.NumIn() if it is variadic
//   each elem should be AssignableTo its respective arg
// out
//   can have less elem than f.NumOut()
//   f.Out(i) should be AssignableTo out[i]
func CheckFunc(foo interface{}, in, out []reflect.Type) (ok bool) {
	defer func() {
		r := recover()
		if r != nil {
			ok = false
		}
	}()
	fv := valueOf(foo)
	ft := fv.Type()
	nin := ft.NumIn()

	if ft.IsVariadic() && len(in) > nin {
		varg := ft.In(nin - 1)
		ve := varg.Elem()
		for i := nin; i < len(in); i++ {
			if !in[i].AssignableTo(ve) {
				return false
			}
		}
	}
	for i := 0; i < nin; i++ {
		if !in[i].AssignableTo(ft.In(i)) {
			return false
		}
	}
	for i := range out {
		if !ft.Out(i).AssignableTo(out[i]) {
			return false
		}
	}
	return true
}
