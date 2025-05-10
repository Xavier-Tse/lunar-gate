package maps

import "reflect"

func Struct2Map(data any, tag string) map[string]any {
	mps := make(map[string]any)
	v := reflect.ValueOf(data)
	t := reflect.TypeOf(data)
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		col := t.Field(i).Tag.Get(tag)
		if col == "" || col == "-" {
			continue
		}

		if field.Kind() == reflect.Pointer {
			if field.IsNil() {
				continue
			}
			mps[col] = field.Elem().Interface()
			continue
		}
		mps[col] = field.Interface()
	}
	return mps
}
