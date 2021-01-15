package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
	"testing"
	"tpldemo/crudSql/entity"
)

type DBConfigEntity struct {
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
func InitDB(cfg DBConfigEntity) *sql.DB {
	if strings.EqualFold(cfg.Timezone, "") {
		cfg.Timezone = "'Asia/Shanghai'"
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s",
		cfg.Name,
		cfg.Pass,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
		cfg.Charset,
	)
	connection, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	connection.SetMaxOpenConns(cfg.MaxOpen)
	connection.SetMaxIdleConns(cfg.MaxIdle)
	return connection
}

var DB *sql.DB
var AdminsDb *AdminsModel

func initConnection() {
	cfg := DBConfigEntity{
		Host:    "localhost",
		Port:    3306,
		Name:    "root",
		Pass:    "ly123456",
		DBName:  "temp",
		Charset: "utf8mb4",
		MaxOpen: 10,
		MaxIdle: 5,
	}
	DB = InitDB(cfg)
	AdminsDb = NewAdmins(DB)
}

// 查询所有的数据
func TestAdminsFindWhere(t *testing.T) {
	initConnection()

	result, err := AdminsDb.Find("id > ? and user_name like ?", &Pagination{PageSize:2,Page:1}, 321231321, "王%") //函数里需要拼接一下sql
	if err != nil {
		t.Error(err)
	}
	for key, value := range result {
		fmt.Printf("key: %v, value: %+v \n", key, value)
	}
}

// 获取最后一条数据
func TestAdminsLast(t *testing.T) {
	initConnection()
	result, err := AdminsDb.First(1)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(result)
}

// 获取总数量
func TestAdminsCount(t *testing.T) {
	initConnection()
	result, err := AdminsDb.First(1)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(result)
}

// 创建数据
func TestAdminsCreate(t *testing.T) {
	initConnection()

	result, err := AdminsDb.Create(&entity.Admins{})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(result)
}

// 创建数据
func TestExampleUpdate(t *testing.T) {
	initConnection()

	result, err := AdminsDb.Update(&entity.Admins{})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(result)
}
