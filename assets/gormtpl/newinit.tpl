package {{.}}

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var (
	MasterDB *gorm.DB //主数据库连接
	//SlaveDB  *gorm.DB //从数据库连接
)

// 页码结构体
type Pagination struct {
	Page      int64 `json:"page" example:"0"`      // 当前页
	PageSize  int64 `json:"pageSize" example:"20"` // 每页条数
	Total     int64 `json:"total"`                 // 总条数
	TotalPage int64 `json:"totalPage"`             // 总页数
}

type Base struct {
	ID        uint64     `gorm:"column:id;primary_key;AUTO_INCREMENT:false;" json:"id,string" form:"id"`                  // 主键
	CreatedAt int64      `gorm:"column:created_at;not null;default:'0';comment:'创建时间'" json:"createdAt" form:"createdAt"` // 创建时间
	UpdatedAt int64      `gorm:"column:updated_at;not null;default:'0';comment:'更新时间'" json:"updatedAt" form:"updatedAt"` // 更新时间
	DeletedAt *time.Time `gorm:"index;null;default:null;comment:'删除时间'" json:"deletedAt" form:"deletedAt"`                // 删除时间
}

// 添加前
func (m *Base) BeforeCreate(scope *gorm.Scope) error {
	m.CreatedAt = time.Now().Unix()
	m.UpdatedAt = time.Now().Unix()
	return nil
}

// 更新前
func (m *Base) BeforeUpdate(scope *gorm.Scope) error {
	m.UpdatedAt = time.Now().Unix()
	return nil
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

//初始化
func init() {
	masterCfg := DbConfig{ //主数据库配置
		Host:    "localhost",
		Port:    3306,
		Name:    "root",
		Pass:    "ly123456",
		DBName:  "temp",
		Charset: "utf8mb4",
		MaxOpen: 100,
		MaxIdle: 50,
	}
	//slaveCfg := DbConfig{   //从数据库配置
	//	Host:    "localhost",
	//	Port:    3307,
	//	Name:    "root",
	//	Pass:    "123456",
	//	DBName:  "kindled",
	//	Charset: "utf8mb4",
	//	MaxOpen: 100,
	//	MaxIdle: 50,
	//}
	MasterDB, _ = initDB(masterCfg) //主
	//SlaveDB, _ = initDB(slaveCfg)  //从
}

//初始化数据库连接
func initDB(cfg DbConfig) (db *gorm.DB, err error) {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&",
		cfg.Name,
		cfg.Pass,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
		cfg.Charset,
	)

	defer func() { //报错释放资源
		if err := recover(); err != nil {
			if err1 := MasterDBClose(); err1 != nil {
				panic(err1)
			}
			//if err2 := SlaveDBClose(); err2 != nil {
			//	panic(err2)
			//}
			panic(err)
		}
	}()

	if db, err = gorm.Open("mysql", dsn); err != nil {
		panic(err)
	}
	sqlDb := db.DB()
	sqlDb.SetMaxIdleConns(cfg.MaxIdle) //空闲连接数
	sqlDb.SetMaxOpenConns(cfg.MaxOpen) //最大连接数

	db.LogMode(true) //打开sql执行日志
	db = db.Debug()  //debug模式
	//添加钩子函数
	addCallBackFunc(db)
	//数据迁移生成表
	//if err = migration(db); err != nil {
	//	panic(err)
	//}
	//db.SingularTable(true)	//数据迁移生成表结尾不带s

	return
}

//添加钩子函数
func addCallBackFunc(db *gorm.DB) {

	// 创建时雪花算法生成ID
	db.Callback().Create().Before("gorm:create").Register("auto_insert_id", func(scope *gorm.Scope) {
		// 若主键为空自动填充 ID
		isTrue := scope.PrimaryKeyZero()
		if isTrue {
			_ = scope.SetColumn("ID", NextId())
		}
	})

	// 修改时间 使用时间戳
	db.Callback().Update().Before("gorm:update").Register("update_at_to_stamp", func(scope *gorm.Scope) {
		if _, ok := scope.Get("gorm:update_column"); !ok {
			_ = scope.SetColumn("UpdatedAt", time.Now().Unix())
		}
	})
}

//数据迁移
func migration(db *gorm.DB) (err error) {
	//传入要迁移的模型指针
	if err = db.AutoMigrate(
	//new(Admin),
	//new(AdminsRole),
	//new(Menu),
	//new(Role),
	//new(RoleMenu),
	).Error; err != nil {
		// 错误
	}
	return
}

// close MasterDb
func MasterDBClose() error {
	if MasterDB != nil {
		return MasterDB.Close()
	}
	return nil
}

// close slaveDb
//func SlaveDBClose() error {
//	if SlaveDB != nil {
//		return SlaveDB.Close()
//	}
//	return nil
//}

//开启事务
func Begin(i int) *gorm.DB {
	fmt.Println(i, "begin")
	return MasterDB.Begin()
}

//提交回滚事务
func Commit(txDb *gorm.DB, err error) {

	if err != nil {
		//错误回滚
		txDb.Rollback()
		return
	}
	//提交
	txDb.Commit()
}