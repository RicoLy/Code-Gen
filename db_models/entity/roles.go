package entity

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"time"
)

type Roles struct {
	Id        int64     `json:"id" db:"id"`
	CreatedAt int64     `json:"created_at" db:"created_at"` // 创建时间
	UpdatedAt int64     `json:"updated_at" db:"updated_at"` // 更新时间
	DeletedAt time.Time `json:"deleted_at" db:"deleted_at"` // 删除时间
	Name      string    `json:"name" db:"name"`             // 角色名称
	ParentId  int64     `json:"parent_id" db:"parent_id"`   // 父级ID
	Status    int32     `json:"status" db:"status"`         // 状态(1:正常 2:暂停使用)

}

type RolesNull struct {
	Id        sql.NullInt64
	CreatedAt sql.NullInt64  // 创建时间
	UpdatedAt sql.NullInt64  // 更新时间
	DeletedAt mysql.NullTime // 删除时间
	Name      sql.NullString // 角色名称
	ParentId  sql.NullInt64  // 父级ID
	Status    sql.NullInt32  // 状态(1:正常 2:暂停使用)

}
