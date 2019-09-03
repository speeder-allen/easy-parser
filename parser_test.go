package easy_parser_test

import (
	"context"
	easy_parser "github.com/speeder-allen/easy-parser"
	"gotest.tools/assert"
	"log"
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

func TestParserContext(t *testing.T) {
	lg := log.New(os.Stdout, "[debug]", log.LstdFlags)
	ctx := context.WithValue(context.Background(), "logger", lg)
	conf := struct {
		Log      *log.Logger `ctxkey:"logger"`
		ErrorLog *log.Logger `ctxkey:"error_logger"`
	}{}
	err := easy_parser.ParserContext(ctx, &conf)
	assert.NilError(t, err)

}
