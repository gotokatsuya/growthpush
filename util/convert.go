package util

import (
	"fmt"
	"reflect"
)

func ToMap(in interface{}, tag string) (map[string]interface{}, error) {
	out := make(map[string]interface{})

	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("ToMap only accepts struct; got %T", v)
	}

	typ := v.Type()
	for i := 0; i < v.NumField(); i++ {
		key := typ.Field(i).Tag.Get(tag)
		if key == "" {
			continue
		}
		out[key] = v.Field(i).Interface()
	}
	return out, nil
}

func JSONToMap(in interface{}) (map[string]interface{}, error) {
	return ToMap(in, "json")
}

func ToMapString(in interface{}, tag string) (map[string]string, error) {
	out := make(map[string]string)

	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("ToMapString only accepts struct; got %T", v)
	}

	typ := v.Type()
	for i := 0; i < v.NumField(); i++ {
		key := typ.Field(i).Tag.Get(tag)
		if key == "" {
			continue
		}
		strValue, ok := v.Field(i).Interface().(string)
		if !ok {
			return nil, fmt.Errorf("ToMapString only accepts string interface; got %T", v.Field(i).Interface())
		}
		if strValue == "" {
			continue
		}
		out[key] = strValue
	}
	return out, nil
}

func JSONToMapString(in interface{}) (map[string]string, error) {
	return ToMapString(in, "json")
}
