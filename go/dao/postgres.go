package dao

import (
	"fmt"
	"io/ioutil"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gopkg.in/yaml.v2"
)

const DRIVER = "postgres"

var SqlSession *gorm.DB

type conf struct {
	Url      string `yaml:"Url"`
	Port     int    `yaml:"Port"`
	DbName   string `yaml:"DbName"`
	UserName string `yaml:"UserName"`
	Password string `yaml:"Password"`
}

func (c *conf) getConf() *conf {
	//讀取 resources/application.yaml 檔案
	yamlFile, err := ioutil.ReadFile("resources/application.yaml")
	//若出現錯誤，列印錯誤提示
	if err != nil {
		fmt.Println(err.Error())
	}
	//將讀取的字串轉換成結構體 conf
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Println(err.Error())
	}
	return c
}

func InitPostgres() (err error) {
	var c conf
	// 獲取 yaml 配置引數
	conf := c.getConf()
	// 將 yaml 配置引數拼接成連線資料庫 url 並連線資料庫
	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		conf.Url,
		conf.Port,
		conf.DbName,
		conf.UserName,
		conf.Password,
	)
	// 連線資料庫
	SqlSession, err = gorm.Open(DRIVER, dsn)
	if err != nil {
		panic(err)
	}
	// 驗證資料庫連線是否成功，若成功，則無異常
	return SqlSession.DB().Ping()
}
func Close() {
	SqlSession.Close()
}
