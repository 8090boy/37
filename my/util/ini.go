package util

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Config struct {
	filepath string
	conflist []map[string]map[string]string
}

var _conf *Config = nil

func SetConfig(filepath string) *Config {
	if _conf != nil {
		return _conf
	}
	_conf = new(Config)
	_conf.filepath = filepath
	return _conf
}

func NewConfig(filepath string) *Config {
	__conf := new(Config)
	__conf.filepath = filepath
	return __conf
}

func GetConfig() *Config {
	return _conf
}

func (c *Config) Get(section, name string) string {
	c.ReadList()
	conf := c.ReadList()
	for _, v := range conf {
		for key, value := range v {
			if key == section {
				return value[name]
			}
		}
	}
	return "no value"
}

func (c *Config) ReadList() []map[string]map[string]string {

	file, err := os.Open(c.filepath)
	if err != nil {
		CheckErr(err)
	}
	defer file.Close()
	var data map[string]map[string]string
	var section string
	buf := bufio.NewReader(file)
	for {
		l, err := buf.ReadString('\n')
		line := strings.TrimSpace(l)
		if err != nil {
			if err != io.EOF {
				CheckErr(err)
			}
			if len(line) == 0 {
				break
			}
		}
		switch {
		case len(line) == 0:
		case line[0] == '[' && line[len(line)-1] == ']':
			section = strings.TrimSpace(line[1 : len(line)-1])
			data = make(map[string]map[string]string)
			data[section] = make(map[string]string)
		default:
			i := strings.IndexAny(line, "=")
			value := strings.TrimSpace(line[i+1 : len(line)])
			data[section][strings.TrimSpace(line[0:i])] = value
			if c.uniquappend(section) == true {
				c.conflist = append(c.conflist, data)
			}
		}

	}

	return c.conflist
}

func CheckErr(err error) string {
	if err != nil {
		return fmt.Sprintf("Error is :'%s'", err.Error())
	}
	return "Notfound this error"
}

func (c *Config) uniquappend(conf string) bool {
	for _, v := range c.conflist {
		for k, _ := range v {
			if k == conf {
				return false
			}
		}
	}
	return true
}
