package easy_parser_test

import (
	easy_parser "github.com/speeder-allen/easy-parser"
	"gotest.tools/assert"
	"reflect"
	"testing"
)

func TestStr2Bool(t *testing.T) {
	s := "true"
	r, err := easy_parser.Str2Bool(s)
	assert.NilError(t, err)
	assert.Equal(t, reflect.TypeOf(r).Kind(), reflect.Bool)
	assert.Equal(t, r, true)
}

func TestStr2Bytes(t *testing.T) {
	s := "test"
	r, err := easy_parser.Str2Bytes(s)
	assert.NilError(t, err)
	assert.Equal(t, reflect.TypeOf(r), reflect.TypeOf([]byte{}))
	assert.DeepEqual(t, r, []byte("test"))
}

func TestStr2Float32(t *testing.T) {
	s := "3.1415926"
	r, err := easy_parser.Str2Float32(s)
	assert.NilError(t, err)
	assert.Equal(t, reflect.TypeOf(r).Kind(), reflect.Float32)
	assert.Equal(t, r, float32(3.1415926))
}

func TestStr2Float64(t *testing.T) {
	s := "99993333.1111111"
	r, err := easy_parser.Str2Float64(s)
	assert.NilError(t, err)
	assert.Equal(t, reflect.TypeOf(r).Kind(), reflect.Float64)
	assert.Equal(t, r, 99993333.1111111)
}

func TestStr2Int(t *testing.T) {
	s := "-134561234"
	r, err := easy_parser.Str2Int(s)
	assert.NilError(t, err)
	assert.Equal(t, reflect.TypeOf(r).Kind(), reflect.Int)
	assert.Equal(t, r, int(-134561234))
}

func TestStr2Uint(t *testing.T) {
	s := "134561234"
	r, err := easy_parser.Str2Int(s)
	assert.NilError(t, err)
	assert.Equal(t, reflect.TypeOf(r).Kind(), reflect.Int)
	assert.Equal(t, r, int(134561234))
}

func TestStr2Int8(t *testing.T) {
	s := "128"
	_, err := easy_parser.Str2Int8(s)
	if err == nil {
		t.Fatalf("128 should not convert int8")
	}
	s = "-127"
	r, err := easy_parser.Str2Int8(s)
	assert.NilError(t, err)
	assert.Equal(t, reflect.TypeOf(r).Kind(), reflect.Int8)
	assert.Equal(t, r, int8(-127))
}

func TestStr2Uint8(t *testing.T) {
	s := "256"
	_, err := easy_parser.Str2Uint8(s)
	if err == nil {
		t.Fatalf("256 should not convert uint8")
	}
	s = "-1"
	_, err = easy_parser.Str2Uint8(s)
	if err == nil {
		t.Fatalf("-1 should not convert uint8")
	}
	s = "255"
	r, err := easy_parser.Str2Uint8(s)
	assert.NilError(t, err)
	assert.Equal(t, reflect.TypeOf(r).Kind(), reflect.Uint8)
	assert.Equal(t, r, uint8(255))
}

func TestStr2Int16(t *testing.T) {
	s := "32768"
	_, err := easy_parser.Str2Int16(s)
	if err == nil {
		t.Fatalf("32768 should not convert int16")
	}
	s = "-32767"
	r, err := easy_parser.Str2Int16(s)
	assert.NilError(t, err)
	assert.Equal(t, reflect.TypeOf(r).Kind(), reflect.Int16)
	assert.Equal(t, r, int16(-32767))
}

func TestStr2Uint16(t *testing.T) {
	s := "65536"
	_, err := easy_parser.Str2Uint16(s)
	if err == nil {
		t.Fatalf("65536 should not convert uint16")
	}
	s = "-1"
	_, err = easy_parser.Str2Uint16(s)
	if err == nil {
		t.Fatalf("-1 should not convert uint16")
	}
	s = "65535"
	r, err := easy_parser.Str2Uint16(s)
	assert.NilError(t, err)
	assert.Equal(t, reflect.TypeOf(r).Kind(), reflect.Uint16)
	assert.Equal(t, r, uint16(65535))
}

func TestStr2Int32(t *testing.T) {
	s := "131313"
	r, err := easy_parser.Str2Int32(s)
	assert.NilError(t, err)
	assert.Equal(t, reflect.TypeOf(r).Kind(), reflect.Int32)
	assert.Equal(t, r, int32(131313))
}

func TestStr2Uint32(t *testing.T) {
	s := "131313"
	r, err := easy_parser.Str2Uint32(s)
	assert.NilError(t, err)
	assert.Equal(t, reflect.TypeOf(r).Kind(), reflect.Uint32)
	assert.Equal(t, r, uint32(131313))
}

func TestStr2Int64(t *testing.T) {
	s := "131313"
	r, err := easy_parser.Str2Int64(s)
	assert.NilError(t, err)
	assert.Equal(t, reflect.TypeOf(r).Kind(), reflect.Int64)
	assert.Equal(t, r, int64(131313))
}

func TestStr2Uint64(t *testing.T) {
	s := "131313"
	r, err := easy_parser.Str2Uint64(s)
	assert.NilError(t, err)
	assert.Equal(t, reflect.TypeOf(r).Kind(), reflect.Uint64)
	assert.Equal(t, r, uint64(131313))
}

func TestString2Type(t *testing.T) {
	s := "12345"
	r, err := easy_parser.String2Type(s, "int32")
	assert.NilError(t, err)
	assert.Equal(t, reflect.TypeOf(r).Kind(), reflect.Int32)
	assert.Equal(t, r, int32(12345))
	r, err = easy_parser.String2Type(s, "invalid")
	assert.Error(t, err, easy_parser.ErrorInvalidType.Error())
}

func TestStr2Json(t *testing.T) {
	s := `{"name":"test","age":32}`
	r, err := easy_parser.Str2Json(s)
	assert.NilError(t, err)
	assert.Equal(t, r.(map[string]interface{})["name"], "test")
}
