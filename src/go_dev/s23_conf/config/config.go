package config

// 实现一个解析配置文件的包
import (
	"io"
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

type Config struct{
	filename string
	data map[string]string
}

func NewConfig(file string)(conf *Config, err error){
	conf = &Config{
		filename: file,
		data: make(map[string]string, 1024),
	}

	err = conf.parse()
	return
}

func (c *Config) parse() (err error) {
	f, err := os.Open(c.filename)
	if err != nil {
		return
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	var lineNo int
	for {
		line, errRet := reader.ReadString('\n')
		if errRet == io.EOF {
			break
		}
		if errRet != nil {
			err = errRet
			return
		}
		lineNo++

		line = strings.TrimSpace(line)
		if len(line) == 0 || line[0] =='\n' || line[0]=='#' || line[0]==';' {
			continue
		}

		itemSlice := strings.Split(line, "=")
		if len(itemSlice) == 0 {
			fmt.Printf("invalid config, line:%d", lineNo)
			continue
		}

		key := strings.TrimSpace(itemSlice[0])
		if len(key) == 0 {
			fmt.Printf("invalid config, line:%d", lineNo)
			continue
		}
		if len(key) == 1 {
			c.data[key] = ""
			continue
		}

		value := strings.TrimSpace(itemSlice[1])
		c.data[key] = value	
	}

	return 
}


func (c *Config) GetInt(key string)(value int, err error){
	str, ok := c.data[key]
	if !ok {
		err = fmt.Errorf("key [%s] not found", key)
	}
	value, err = strconv.Atoi(str)
	return
}

func (c *Config) GetIntDefault(key string, defaultInt int)(value int){
	str, ok := c.data[key]
	if !ok {
		value = defaultInt
		return
	}
	value, err := strconv.Atoi(str)
	if err != nil {
		value = defaultInt
	}
	return
}

func (c *Config) GetString(key string)(value string, err error){
	value, ok := c.data[key]
	if !ok {
		err = fmt.Errorf("key [%s] not found", key)
	}
	return
}

func (c *Config) GetIStringDefault(key string, defStr string)(value string){
	value, ok := c.data[key]
	if !ok {
		value = defStr
		return
	}
	return
}