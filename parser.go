package easy_parser

import (
	"context"
	"errors"
	"fmt"
	"os"
	"reflect"
)

var (
	ErrorArgumentMustPointer = errors.New("argument must a pointer")
)

const EnvKeyName = "envKey"
const EnvKeyType = "envType"
const ContextKeyName = "ctxKey"

// ParserEnvironment is parser to struct form environment
// inf is pointer to struct
// struct field tag envKey is environment variable name
// struct field tag envType can custom field decode type, example: json
func ParserEnvironment(inf interface{}) error {
	t := reflect.TypeOf(inf)
	if !isPointer(t) {
		return ErrorArgumentMustPointer
	}
	t = t.Elem()
	v := reflect.ValueOf(inf).Elem()
	for i := 0; i < t.NumField(); i++ {
		if envKey, ok := t.Field(i).Tag.Lookup(EnvKeyName); ok {
			if !isExportColumn(t.Field(i)) {
				return fmt.Errorf("%s is unexported field", t.Field(i).Name)
			}
			if envVal := os.Getenv(envKey); envVal != "" {
				envType := t.Field(i).Type
				envDecode, ok := t.Field(i).Tag.Lookup(EnvKeyType)
				if !ok {
					envDecode = envType.String()
				}
				value := reflect.New(envType)
				err := StringParser(envVal, envDecode, &value)
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
// struct field tag ctxKey is context value key name
func ParserContext(ctx context.Context, inf interface{}) error {
	t := reflect.TypeOf(inf)
	if !isPointer(t) {
		return ErrorArgumentMustPointer
	}
	t = t.Elem()
	v := reflect.ValueOf(inf).Elem()
	for i := 0; i < t.NumField(); i++ {
		if ctxKey, ok := t.Field(i).Tag.Lookup(ContextKeyName); ok {
			if !isExportColumn(t.Field(i)) {
				return fmt.Errorf("%s is unexported field", t.Field(i).Name)
			}
			if ctx.Value(ctxKey) != nil {
				v.Field(i).Set(reflect.ValueOf(ctx.Value(ctxKey)))
			}
		}
	}
	return nil
}
