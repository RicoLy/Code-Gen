import (
	"database/sql"
	"fmt"
	"net/url"
	"strings"
)

var (
	masterDB *sql.DB
	slaveDB *sql.DB
)

func init() {
	cfg := dbConfig{
		Host:    "localhost",
		Port:    3306,
		Name:    "root",
		Pass:    "123456",
		DBName:  "kindled",
		Charset: "utf8mb4",
		MaxOpen: 100,
		MaxIdle: 50,
	}
	masterDB = initDB(cfg)
	//slaveDB = initDB(cfg)
}

type dbConfig struct {
	Host     string //地址
	Port     int    //端口
	Name     string //用户
	Pass     string //密码
	DBName   string //库名
	Charset  string //编码
	Timezone string //时区
	MaxIdle  int    //最大空间连接
	MaxOpen  int    //最大连接数
}

// 连接数据库
func initDB(cfg dbConfig) *sql.DB {
	if strings.EqualFold(cfg.Timezone, "") {
		cfg.Timezone = "'Asia/Shanghai'"
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=true&loc=Local&time_zone=%s",
		cfg.Name,
		cfg.Pass,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
		cfg.Charset,
		url.QueryEscape(cfg.Timezone),
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