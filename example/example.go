package example

import (
	"fmt"
	parser "github.com/speeder-allen/easy-parser"
	"os"
)

type MysqlConf struct {
	Host     string `envkey:"MYSQL_HOST"`
	Port     uint32 `envkey:"MYSQL_PORT"`
	Username string `envkey:"MYSQL_USERNAME"`
	Password string `envkey:"MYSQL_PASSWORD"`
}

func Example() {
	os.Setenv("MYSQL_HOST", "127.0.0.1")
	os.Setenv("MYSQL_PORT", "3306")
	os.Setenv("MYSQL_USERNAME", "root")
	os.Setenv("MYSQL_PASSWORD", "123456")

	conf := MysqlConf{}
	parser.ParserEnvironment(&conf)
	fmt.Println(conf.Host)     // 127.0.0.1
	fmt.Println(conf.Port)     // 3306
	fmt.Println(conf.Username) // root
	fmt.Println(conf.Password) // 123456
}
