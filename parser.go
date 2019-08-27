package easy_parser

import (
	"errors"
	"os"
	"reflect"
)

var (
	ErrorArgumentMustPointer = errors.New("argument must a pointer")
)

func ParserEnvironment(inf interface{}) error {
	t := reflect.TypeOf(inf)
	if t.String()[0:1] != "*" {
		return ErrorArgumentMustPointer
	}
	t = t.Elem()
	v := reflect.ValueOf(inf).Elem()
	for i := 0; i < t.NumField(); i++ {
		if envKey, ok := t.Field(i).Tag.Lookup("envkey"); ok {
			envType, ok := t.Field(i).Tag.Lookup("envtype")
			if !ok {
				envType = t.Field(i).Type.String()
			}
			if envVal := os.Getenv(envKey); envVal != "" {
				val, err := String2Type(envVal, envType)
				if err == nil && val != nil {
					v.Field(i).Set(reflect.ValueOf(val))
				}
			}
		}
	}
	return nil
}
