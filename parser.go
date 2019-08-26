package easy_parser

import (
	"errors"
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
	return nil
}
