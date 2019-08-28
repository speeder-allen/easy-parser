package easy_parser

import (
	"context"
	"errors"
	"os"
	"reflect"
)

var (
	ErrorArgumentMustPointer = errors.New("argument must a pointer")
)

// ParserEnvironment is parser to struct form environment
// inf is pointer to struct
// struct field tag envkey is environment variable name
// struct field tag envtype can custom field decode type, example: json
func ParserEnvironment(inf interface{}) error {
	t := reflect.TypeOf(inf)
	if t.String()[0:1] != "*" {
		return ErrorArgumentMustPointer
	}
	t = t.Elem()
	v := reflect.ValueOf(inf).Elem()
	for i := 0; i < t.NumField(); i++ {
		if envKey, ok := t.Field(i).Tag.Lookup("envkey"); ok {
			if envVal := os.Getenv(envKey); envVal != "" {
				envType := t.Field(i).Type
				envDecode, ok := t.Field(i).Tag.Lookup("envtype")
				if !ok {
					envDecode = envType.String()
				}
				value := reflect.New(envType)
				err := StringParser(envVal, envDecode, &value)
				//log.Println(err, envType, value.Elem())
				if err == nil {
					v.Field(i).Set(value.Elem())
				}
			}
		}
	}
	return nil
}

// ParserContext is parser to struct from context
// inf is pointer to struct
// struct field tag ctxkey is context value key name
func ParserContext(ctx context.Context, inf interface{}) error {
	t := reflect.TypeOf(inf)
	if t.String()[0:1] != "*" {
		return ErrorArgumentMustPointer
	}
	t = t.Elem()
	v := reflect.ValueOf(inf).Elem()
	for i := 0; i < t.NumField(); i++ {
		if ctxKey, ok := t.Field(i).Tag.Lookup("ctxkey"); ok {
			v.Field(i).Set(reflect.ValueOf(ctx.Value(ctxKey)))
		}
	}
	return nil
}
