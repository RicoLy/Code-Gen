package entity

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"time"
)

type RoleMenus struct {
	Id        int64     `json:"id" db:"id"`
	CreatedAt int64     `json:"created_at" db:"created_at"` // 创建时间
	UpdatedAt int64     `json:"updated_at" db:"updated_at"` // 更新时间
	DeletedAt time.Time `json:"deleted_at" db:"deleted_at"` // 删除时间
	RoleId    int64     `json:"role_id" db:"role_id"`       // 角色ID
	MenuId    int64     `json:"menu_id" db:"menu_id"`       // 菜单ID

}

type RoleMenusNull struct {
	Id        sql.NullInt64
	CreatedAt sql.NullInt64  // 创建时间
	UpdatedAt sql.NullInt64  // 更新时间
	DeletedAt mysql.NullTime // 删除时间
	RoleId    sql.NullInt64  // 角色ID
	MenuId    sql.NullInt64  // 菜单ID

}
