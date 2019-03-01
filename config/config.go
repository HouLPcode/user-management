package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"time"
	//"github.com/stackcats/gosh"   consul使用的库
)

// MysqlHelper ...
type MysqlConfig struct {
	Host            string
	Net             string
	UserName        string
	Password        string
	Database        string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
	Parameters      string
}

type MgoConfig struct {
}

type RedisConfig struct {
	RedisURI      string `consul:"redis/cache/host"`
	RedisPassword string `consul:"redis/cache/password"`
}

// Config ...
type Config struct {
	MgoConfig
	RedisConfig
	MysqlConfig
}

type Database struct {
	Mysql *gorm.DB
	//Redis
	//Mongodb
}

var config *Config
var MyDatabase Database

func newConfig(){
	config = &Config{

	}
}

func init() {
	newConfig()
	config.initMysql()
	config.initRedis()
	config.initMgo()
	MyDatabase.NewMysql()
}

func (d *Database)NewMysql(){
	dburl := fmt.Sprintf("%s:%s@%s(%s)/%s?%s", config.MysqlConfig.UserName,config.MysqlConfig.Password,config.MysqlConfig.Net,config.MysqlConfig.Host,config.MysqlConfig.Database,config.MysqlConfig.Parameters)
	var err error
	d.Mysql ,err = gorm.Open("mysql", dburl)
	if err != nil {
		log.Fatalf("数据库连接异常：%v", err)
		panic(err)
	}
	d.Mysql.LogMode(true)
}

func (c *Config) initMysql() {
	c.MysqlConfig = MysqlConfig{
		Host:            "host",
		Net:             "tcp",
		UserName:        "plat",
		Password:        "ookai",
		Database:        "plat",
		MaxOpenConns:    100,
		MaxIdleConns:    10,
		ConnMaxLifetime: time.Second * 10,
		Parameters:      "charset=utf8&parseTime=True&loc=Local",
	}


}
func (c *Config) initMgo() {
	c.MgoConfig = MgoConfig{}
}

func (c *Config) initRedis() {
	c.RedisConfig = RedisConfig{}
}

func NewConfig() *Config {
	return config
}

