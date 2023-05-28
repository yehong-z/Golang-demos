package config_demo

import (
	"fmt"
	"io/ioutil"
	"path"

	"github.com/BurntSushi/toml"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

type Config struct {
	Title  string
	Server struct {
		Host    string
		Port    int
		Timeout int
	}
	Database struct {
		Name     string
		Username string
		Password string
		Port     int
	}
}

func fromTomlFile(c Config, p string) error {
	content, err := ioutil.ReadFile(path.Clean(p))
	if err != nil {
		return err
	}

	dec := unicode.BOMOverride(transform.Nop)
	content, _, err = transform.Bytes(dec, content)
	if err != nil {
		return err
	}

	_, err = toml.Decode(string(content), c)
	return err
}

func PrintTomlFromFile() {
	var conf Config
	err := fromTomlFile(conf, "./test.conf")
	if err != nil {
		return
	}
	fmt.Printf("%+v\n", conf)
}

func PrintConfig() {
	var conf Config
	if _, err := toml.DecodeFile("test.conf", &conf); err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", conf)
}
