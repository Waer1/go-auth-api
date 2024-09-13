package utils

import "reflect"

// isZeroValue checks if the given value is the zero value for its type.
func isZeroValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Ptr, reflect.Interface:
		return v.IsNil()
	case reflect.Array, reflect.Slice, reflect.Map, reflect.String:
		return v.Len() == 0
	case reflect.Struct:
		return reflect.DeepEqual(v.Interface(), reflect.Zero(v.Type()).Interface())
	default:
		// Check for numeric and other primitive types
		return v.Interface() == reflect.Zero(v.Type()).Interface()
	}
}

// Copy assigns fields from src to dest but skips fields with default values in src.
func Copy(dest, src interface{}) {
	destVal := reflect.ValueOf(dest).Elem()
	srcVal := reflect.ValueOf(src).Elem()

	for i := 0; i < srcVal.NumField(); i++ {
		field := srcVal.Type().Field(i)
		destField := destVal.FieldByName(field.Name)
		srcField := srcVal.Field(i)

		// Ensure destination field is valid, settable, and source field is not default value
		if destField.IsValid() && destField.CanSet() && !isZeroValue(srcField) {
			destField.Set(srcField)
		}
	}
}
