package easy_parser_test

import (
	"context"
	easy_parser "github.com/speeder-allen/easy-parser"
	"gotest.tools/assert"
	"log"
	"os"
	"testing"
)

type testStruct struct {
	Name  string `envKey:"TEST_NAME"`
	Age   uint8  `envKey:"TEST_AGE"`
	Attrs struct {
		Level int `json:"level"`
	} `envKey:"TEST_ATTRS" envType:"json"`
}

func TestParserEnvironment(t *testing.T) {
	s := testStruct{}
	err := easy_parser.ParserEnvironment(s)
	assert.Equal(t, err, easy_parser.ErrorArgumentMustPointer)
	m := struct {
		Name         string `envKey:"TEST_NAME"`
		notExportKey string `envKey:"TEST_EXPORT"`
	}{}
	assert.ErrorContains(t, easy_parser.ParserEnvironment(&m), "unexported")
	os.Setenv("TEST_NAME", "speeder-allen")
	os.Setenv("TEST_AGE", "15")
	os.Setenv("TEST_ATTRS", "{\"level\": 11}")
	err = easy_parser.ParserEnvironment(&s)
	assert.NilError(t, err)
	t.Log(s)
	assert.Equal(t, s.Name, "speeder-allen")
	assert.Equal(t, s.Age, uint8(15))
	assert.Equal(t, s.Attrs.Level, 11)
}

func TestParserContext(t *testing.T) {
	lg := log.New(os.Stdout, "[debug]", log.LstdFlags)
	ctx := context.WithValue(context.Background(), "logger", lg)
	conf := struct {
		Log      *log.Logger `ctxKey:"logger"`
		ErrorLog *log.Logger `ctxKey:"error_logger"`
	}{}
	conf1 := struct {
		Log          *log.Logger `ctxKey:"logger"`
		notExportKey string      `ctxKey:"not_export"`
	}{}
	assert.Equal(t, easy_parser.ParserContext(ctx, conf), easy_parser.ErrorArgumentMustPointer)
	assert.ErrorContains(t, easy_parser.ParserContext(ctx, &conf1), "unexported")
	err := easy_parser.ParserContext(ctx, &conf)
	assert.NilError(t, err)

}

func BenchmarkParserEnvironment(b *testing.B) {
	s := testStruct{}
	os.Setenv("TEST_NAME", "speeder-allen")
	os.Setenv("TEST_AGE", "15")
	os.Setenv("TEST_ATTRS", "{\"level\": 11}")
	for i := 0; i < b.N; i++ {
		_ = easy_parser.ParserEnvironment(&s)
	}
}

func BenchmarkParserContext(b *testing.B) {
	conf := struct {
		Log      *log.Logger `ctxKey:"logger"`
		ErrorLog *log.Logger `ctxKey:"error_logger"`
	}{}
	ctx := context.WithValue(context.Background(), "logger", log.New(os.Stdout, "[debug]", log.LstdFlags))
	for i := 0; i < b.N; i++ {
		_ = easy_parser.ParserContext(ctx, &conf)
	}
}
