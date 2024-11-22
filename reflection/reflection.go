package reflection

import "reflect"

type Profile struct {
	Age  int
	City string
}

type Person struct {
	Name    string
	Profile Profile
}

func walk(x interface{}, fn func(string)) {
	val := getValue(x)

	walkVal := func(value reflect.Value) {
		walk(value.Interface(), fn)
	}
	numOfVal := 0
	var getField func(int) reflect.Value

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		numOfVal = val.NumField()
		getField = val.Field
	case reflect.Slice, reflect.Array:
		numOfVal = val.Len()
		getField = val.Index
	case reflect.Map:
		for _, k := range val.MapKeys() {
			walk(val.MapIndex(k).Interface(), fn)
		}
	case reflect.Chan:
		for {
			if v, ok := val.Recv(); ok {
				walkVal(v)
			} else {
				break
			}
		}
	}

	for i := 0; i < numOfVal; i++ {
		walk(getField(i).Interface(), fn)
	}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}

	return val
}
