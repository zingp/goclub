package reconf

// 实现一个解析配置文件的包
import (
	"time"
	"io"
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
	"sync"
)

type Config struct{
	filename string
	data map[string]string
	lastModifyTime int64
	rwLock sync.RWMutex
	notifyList []Notifyer
}

func NewConfig(file string)(conf *Config, err error){
	conf = &Config{
		filename: file,
		data: make(map[string]string, 1024),
	}

	m, err := conf.parse()
	if err != nil {
		fmt.Printf("parse conf error:%v\n", err)
		return
	}

	// 将解析配置文件后的数据更新到结构体的map中，写锁
	conf.rwLock.Lock()
	conf.data = m
	conf.rwLock.Unlock()

	// 启一个后台线程去检测配置文件是否更改
	go conf.reload()
	return
}

// 添加观察者
func (c *Config) AddObserver(n Notifyer) {
	c.notifyList = append(c.notifyList, n)
}

func (c *Config) reload(){
	// 定时器
	ticker := time.NewTicker(time.Second * 5) 
	for _ = range ticker.C {
		// 打开文件
		// 为什么使用匿名函数？ 当匿名函数退出时可用defer去关闭文件
		// 如果不用匿名函数，在循环中不好关闭文件，一不小心就内存泄露
		func (){
			f, err := os.Open(c.filename)
			if err != nil {
				fmt.Printf("open file error:%s\n", err)
				return
			}
			defer f.Close()

			fileInfo, err := f.Stat()
			if err != nil {
				fmt.Printf("stat file error:%s\n", err)
				return
			}
			// 或取当前文件修改时间
			curModifyTime := fileInfo.ModTime().Unix() 
			if curModifyTime > c.lastModifyTime {
				// 重新解析时，要考虑应用程序正在读取这个配置因此应该加锁
				m, err := c.parse()
				if err != nil {
					fmt.Printf("parse config error:%v\n", err)
					return
				}

				c.rwLock.Lock()
				c.data = m
				c.rwLock.Unlock()

				c.lastModifyTime = curModifyTime

				// 配置更新通知所有观察者
				for _, n := range c.notifyList {
					n.Callback(c)
				}
			}
		}()
	}
}

func (c *Config) parse() (m map[string]string, err error) {
	// 如果在parse()中定义一个map，这样就是一个新的map不用加锁
	m = make(map[string]string, 1024)

	f, err := os.Open(c.filename)
	if err != nil {
		return
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	// 声明一个变量存放读取行数
	var lineNo int
	for {
		line, errRet := reader.ReadString('\n')
		if errRet == io.EOF {
			// 这里有一个坑，最后一行如果不是\n结尾会漏读
			lineParse(&lineNo, &line, &m)
			break
		}
		if errRet != nil {
			err = errRet
			return
		}
		
		lineParse(&lineNo, &line, &m)
	}

	return 
}

func lineParse(lineNo *int, line *string, m *map[string]string) {
		*lineNo++

		l := strings.TrimSpace(*line)
		// 如果空行 或者 是注释 跳过
		if len(l) == 0 || l[0] =='\n' || l[0]=='#' || l[0]==';' {
			return
		}

		itemSlice := strings.Split(l, "=")
		// =
		if len(itemSlice) == 0 {
			fmt.Printf("invalid config, line:%d", lineNo)
			return
		}

		key := strings.TrimSpace(itemSlice[0])
		if len(key) == 0 {
			fmt.Printf("invalid config, line:%d", lineNo)
			return
		}
		if len(key) == 1 {
			(*m)[key] = ""
			return
		}

		value := strings.TrimSpace(itemSlice[1])
		(*m)[key] = value	

		return
}


func (c *Config) GetInt(key string)(value int, err error){
	c.rwLock.RLock()
	defer c.rwLock.RUnlock()

	str, ok := c.data[key]
	if !ok {
		err = fmt.Errorf("key [%s] not found", key)
	}
	value, err = strconv.Atoi(str)
	return
}

func (c *Config) GetIntDefault(key string, defaultInt int)(value int){
	c.rwLock.RLock()
	defer c.rwLock.RUnlock()

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
	c.rwLock.RLock()
	defer c.rwLock.RUnlock()

	value, ok := c.data[key]
	if !ok {
		err = fmt.Errorf("key [%s] not found", key)
	}
	return
}

func (c *Config) GetIStringDefault(key string, defaultStr string)(value string){
	c.rwLock.RLock()
	defer c.rwLock.RUnlock()

	value, ok := c.data[key]
	if !ok {
		value = defaultStr
		return
	}
	return
}