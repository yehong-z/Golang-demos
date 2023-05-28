package config

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Mysql struct {
	Host      string `yaml:"host"`
	Port      int    `yaml:"port"`
	Database  string `yaml:"database"`
	Username  string `yaml:"username"`
	Password  string `yaml:"password"`
	Charset   string `yaml:"charset"`
	ParseTime bool   `yaml:"parse_time"`
	Loc       string `yaml:"loc"`
}

type Server struct {
	IP   string `yaml:"ip"`
	Port int    `yaml:"port"`
}

type Config struct {
	DB     Mysql `yaml:"mysql"`
	Server `yaml:"server"`
}

var Info Config

func init() {
	data, err := ioutil.ReadFile("./config/config.yaml") // 读取 YAML 文件
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(data, &Info) // 解析 YAML 文件到结构体中
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", Info) // 打印解析后的配置文件
}

// DBConnectString 填充得到数据库连接字符串
func DBConnectString() string {
	arg := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%v&loc=%s",
		Info.DB.Username, Info.DB.Password, Info.DB.Host, Info.DB.Port, Info.DB.Database,
		Info.DB.Charset, Info.DB.ParseTime, Info.DB.Loc)
	log.Println("connect to db")
	return arg
}
