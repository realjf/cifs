package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"sync"
)


type Redis struct {
	Host string
	Port int
}

type MySQL struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	DbName       string `json:"dbName"`
	Host         string `json:"host"`
	Port         int `json:"port"`
	Charset      string `json:"charset"`
	MaxOpenConns int `json:"maxOpenConns"`
}

type Config struct {
	Data struct{
		Mysql MySQL `json:"mysql"`
		Redis Redis `json:"redis"`
	}

	once sync.Once // 实现单例模式
	lock sync.RWMutex
	path string
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) LoadConfig(path string) *Config {
	c.once.Do(func() {
		if path == "" {
			log.Fatal("config path is empty")
		}
		c.path = path
		// 检查文件是否存在
		fileData, err := ioutil.ReadFile(path)
		if err != nil || len(fileData) <= 0 {
			log.Fatalf("read config file error: %v", err)
		}
		if err := json.Unmarshal(fileData, &c.Data); err != nil {
			log.Fatalf("json unmarshal error: %v", err)
		}
	})

	return c
}

