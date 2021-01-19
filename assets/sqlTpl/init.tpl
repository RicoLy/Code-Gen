package {{.}}

import (
	"database/sql"
	"fmt"
	"strings"
)

var (
	masterDB *sql.DB
	slaveDB  *sql.DB //从数据库
)

func init() {
	cfg := DbConfig{
		Host:    "localhost",
		Port:    3306,
		Name:    "root",
		Pass:    "ly123456",
		DBName:  "temp",
		Charset: "utf8mb4",
		MaxOpen: 100,
		MaxIdle: 50,
	}
	masterDB = initDB(cfg)
	//slaveDB = initDB(cfg)
}

type DbConfig struct {
	Host    string //地址
	Port    int    //端口
	Name    string //用户
	Pass    string //密码
	DBName  string //库名
	Charset string //编码
	MaxIdle int    //最大空闲连接
	MaxOpen int    //最大连接数
}

// 页码结构体
type Pagination struct {
	Page      int64 `json:"page" example:"0"`      // 当前页
	PageSize  int64 `json:"pageSize" example:"20"` // 每页条数
	Total     int64 `json:"total"`                 // 总条数
	TotalPage int64 `json:"totalPage"`             // 总页数
}

// 连接数据库
func initDB(cfg DbConfig) *sql.DB {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&",
		cfg.Name,
		cfg.Pass,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
		cfg.Charset,
	)

	defer func() {
		if err := recover(); err != nil {
			if err1 := MasterDBClose(); err1 != nil {
				panic(err1)
			}
			if err2 := SlaveDBClose(); err2 != nil {
				panic(err2)
			}
			panic(err)
		}
	}()
	connection, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	connection.SetMaxOpenConns(cfg.MaxOpen)
	connection.SetMaxIdleConns(cfg.MaxIdle)
	return connection
}

// close db
func MasterDBClose() error {
	if masterDB != nil {
		return masterDB.Close()
	}
	return nil
}

// close db
func SlaveDBClose() error {
	if slaveDB != nil {
		return slaveDB.Close()
	}
	return nil
}

// transaction start
func TxBegin() (*sql.Tx, error) {
	return masterDB.Begin()
}

// repeat response to ?,?,?
func RepeatQuestionMark(count int) string {
	return strings.TrimRight(strings.Repeat("?,", count), ",")
}
