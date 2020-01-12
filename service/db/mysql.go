package db


import (
	"cifs/service/config"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

var (
	DbConn *sql.DB
)

type MySQLDriver struct {
	config.MySQL
}

func NewMysql(conf *config.Config) *MySQLDriver {
	return &MySQLDriver{
		MySQL: config.MySQL{
			Username:     conf.Data.Mysql.Username,
			Password:     conf.Data.Mysql.Password,
			Host:         conf.Data.Mysql.Host,
			Port:         conf.Data.Mysql.Port,
			Charset:      conf.Data.Mysql.Charset,
			DbName:       conf.Data.Mysql.DbName,
			MaxOpenConns: conf.Data.Mysql.MaxOpenConns,
		},
	}
}

func (db *MySQLDriver) Init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&timeout=10s",
		db.Username,
		db.Password,
		db.Host,
		db.Port,
		db.DbName,
		db.Charset)

	var err error
	DbConn, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatalf("open mysql connection error: %v", err)
	}
	DbConn.SetMaxIdleConns(2)
	DbConn.SetMaxOpenConns(db.MaxOpenConns)
	DbConn.SetConnMaxLifetime(time.Second * time.Duration(60))

	log.Println("init db...")
}

