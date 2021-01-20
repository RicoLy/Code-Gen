package entity

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"time"
)

type Menus struct {
	Id            int64     `json:"id" db:"id"`
	CreatedAt     int64     `json:"created_at" db:"created_at"`         // 创建时间
	UpdatedAt     int64     `json:"updated_at" db:"updated_at"`         // 更新时间
	DeletedAt     time.Time `json:"deleted_at" db:"deleted_at"`         // 删除时间
	Status        int32     `json:"status" db:"status"`                 // 状态(1:启用 2:不启用)
	ParentId      int64     `json:"parent_id" db:"parent_id"`           // 父级ID
	FrontPath     string    `json:"front_path" db:"front_path"`         // 前端文件路径
	Url           string    `json:"url" db:"url"`                       // 菜单api路径URL
	Method        string    `json:"method" db:"method"`                 // 操作方法 /GET/POST/PUT/DELETE/
	Name          string    `json:"name" db:"name"`                     // 菜单名称
	InterfaceName string    `json:"interface_name" db:"interface_name"` // 接口名称
	MenuType      int32     `json:"menu_type" db:"menu_type"`           // 菜单类型(1:模块 2:菜单 3:操作)

}

type MenusNull struct {
	Id            sql.NullInt64
	CreatedAt     sql.NullInt64  // 创建时间
	UpdatedAt     sql.NullInt64  // 更新时间
	DeletedAt     mysql.NullTime // 删除时间
	Status        sql.NullInt32  // 状态(1:启用 2:不启用)
	ParentId      sql.NullInt64  // 父级ID
	FrontPath     sql.NullString // 前端文件路径
	Url           sql.NullString // 菜单api路径URL
	Method        sql.NullString // 操作方法 /GET/POST/PUT/DELETE/
	Name          sql.NullString // 菜单名称
	InterfaceName sql.NullString // 接口名称
	MenuType      sql.NullInt32  // 菜单类型(1:模块 2:菜单 3:操作)

}
