package easy_parser_test

import (
	easy_parser "github.com/speeder-allen/easy-parser"
	"gotest.tools/assert"
	"testing"
)

func TestParserEnvironment(t *testing.T) {
	s := struct {
		Name string
		Age  uint8
	}{}
	err := easy_parser.ParserEnvironment(s)
	assert.Equal(t, err, easy_parser.ErrorArgumentMustPointer)
}
