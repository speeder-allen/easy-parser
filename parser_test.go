package easy_parser_test

import (
	easy_parser "github.com/speeder-allen/easy-parser"
	"gotest.tools/assert"
	"os"
	"testing"
)

func TestParserEnvironment(t *testing.T) {
	s := struct {
		Name string `envkey:"TEST_NAME"`
		Age  uint8  `envkey:"TEST_AGE"`
	}{}
	err := easy_parser.ParserEnvironment(s)
	assert.Equal(t, err, easy_parser.ErrorArgumentMustPointer)
	os.Setenv("TEST_NAME", "speeder-allen")
	os.Setenv("TEST_AGE", "15")
	err = easy_parser.ParserEnvironment(&s)
	assert.NilError(t, err)
	t.Log(s)

}
