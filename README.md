# easy-parser
a easy parser for golang 

## test

[![Build Status](https://travis-ci.org/speeder-allen/easy-parser.svg?branch=master)](https://travis-ci.org/speeder-allen/easy-parser)

## install

```shell
go get -u github.com/speeder-allen/easy-parser
```

## how to use

```go
package main

import (
    "context"
    "encoding/xml"
    "fmt"
    parser "github.com/speeder-allen/easy-parser"
    "log"
    "os"
)

// MysqlConf is a struct for parser environment
// envkey is define the key of environment variable name
// envtype is advanced for json、xml struct
type MysqlConf struct {
	Host     string `envkey:"MYSQL_HOST"`
	Port     uint32 `envkey:"MYSQL_PORT"`
	Username string `envkey:"MYSQL_USERNAME"`
	Password string `envkey:"MYSQL_PASSWORD"`
	MetaData struct {
		CharSet  string `json:"charset"`
		TimeZone string `json:"timezone"`
	} `envkey:"MYSQL_META" envtype:"json"`
	XmlTest struct {
		XMLName xml.Name `xml:"persons"`
		Persons []struct {
			Name      string   `xml:"name,attr"`
			Age       string   `xml:"age,attr"`
			Career    string   `xml:"career"`
			Interests []string `xml:"interests>interest"`
		} `xml:"person"`
	} `envkey:"TEST_XML" envtype:"xml"`
}

var xmlstring = `<?xml version="1.0" encoding="UTF-8"?>
<persons>
    <person name="Jack" age="35">
        <career>Leader</career>
        <interests>
            <interest>travel</interest>
            <interest>fitness</interest>
        </interests>
    </person>
    <person name="John" age="27">
        <career>programmer</career>
        <interests>
            <interest>reading</interest>
            <interest>game</interest>
        </interests>
    </person>
</persons>`

func main() {
	// define some environment for test
	os.Setenv("MYSQL_HOST", "127.0.0.1")
	os.Setenv("MYSQL_PORT", "3306")
	os.Setenv("MYSQL_USERNAME", "root")
	os.Setenv("MYSQL_PASSWORD", "123456")
	os.Setenv("MYSQL_META", `{"charset":"utf8","timezone":"UTC"}`)
	os.Setenv("TEST_XML", xmlstring)

	// create a struct
	conf := MysqlConf{}

	// parser environment to struct
	parser.ParserEnvironment(&conf)

	//print result
	fmt.Println(conf.Host)                              // 127.0.0.1
	fmt.Println(conf.Port)                              // 3306
	fmt.Println(conf.Username)                          // root
	fmt.Println(conf.Password)                          // 123456
	fmt.Println(conf.MetaData.CharSet)                  // utf8
	fmt.Println(conf.MetaData.TimeZone)                 // UTC
	fmt.Println(conf.XmlTest.XMLName.Local)             // persons
	fmt.Println(len(conf.XmlTest.Persons))              // 2
	fmt.Println(conf.XmlTest.Persons[0].Name)           // Jack
	fmt.Println(conf.XmlTest.Persons[1].Career)         // programmer
	fmt.Println(len(conf.XmlTest.Persons[0].Interests)) //2
	fmt.Println(conf.XmlTest.Persons[0].Interests[0])   //travel
	fmt.Println(conf.XmlTest.Persons[1].Age)            // 27

	// test context parser
	log1 := log.New(os.Stdout, "[normal]", log.LstdFlags)
	log2 := log.New(os.Stderr, "[error]", log.LstdFlags)
	ctx := context.WithValue(context.Background(), "normal_logger", log1)
	ctx = context.WithValue(ctx, "error_logger", log2)
	some(ctx)
}

func some(ctx context.Context) {
	instance := struct {
		NormalLog *log.Logger `ctxkey:"normal_logger"`
		ErrorLog  *log.Logger `ctxkey:"error_logger"`
	}{}

	// parser context value to struct
	parser.ParserContext(ctx, &instance)

	//print result
	fmt.Println(instance.NormalLog)
	fmt.Println(instance.ErrorLog)

	// print log
	instance.NormalLog.Println("12345")  // [normal]12345
	instance.ErrorLog.Println("error!!") // [error]error!!
}

```