package entity

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"time"
)

type AdminsRoles struct {
	Id        int64     `json:"id" db:"id"`
	CreatedAt int64     `json:"created_at" db:"created_at"` // 创建时间
	UpdatedAt int64     `json:"updated_at" db:"updated_at"` // 更新时间
	DeletedAt time.Time `json:"deleted_at" db:"deleted_at"` // 删除时间
	AdminsId  int64     `json:"admins_id" db:"admins_id"`   // 管理员ID
	RoleId    int64     `json:"role_id" db:"role_id"`       // 角色ID

}

type AdminsRolesNull struct {
	Id        sql.NullInt64
	CreatedAt sql.NullInt64  // 创建时间
	UpdatedAt sql.NullInt64  // 更新时间
	DeletedAt mysql.NullTime // 删除时间
	AdminsId  sql.NullInt64  // 管理员ID
	RoleId    sql.NullInt64  // 角色ID

}
