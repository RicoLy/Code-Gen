package entity

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"time"
)

type Admins struct {
	Id        int64     `json:"id" db:"id"`
	CreatedAt int64     `json:"created_at" db:"created_at"` // 创建时间
	UpdatedAt int64     `json:"updated_at" db:"updated_at"` // 更新时间
	DeletedAt time.Time `json:"deleted_at" db:"deleted_at"` // 删除时间
	UserName  string    `json:"user_name" db:"user_name"`   // 管理员用户名
	RealName  string    `json:"real_name" db:"real_name"`   // 管理员真实姓名
	Password  string    `json:"password" db:"password"`     // 密码(sha1(md5(明文))加密)
	Email     string    `json:"email" db:"email"`           // 邮箱
	Phone     string    `json:"phone" db:"phone"`           // 手机
	Status    int32     `json:"status" db:"status"`         // 状态(1:正常启用 2:暂停使用)

}

type AdminsNull struct {
	Id        sql.NullInt64
	CreatedAt sql.NullInt64  // 创建时间
	UpdatedAt sql.NullInt64  // 更新时间
	DeletedAt mysql.NullTime // 删除时间
	UserName  sql.NullString // 管理员用户名
	RealName  sql.NullString // 管理员真实姓名
	Password  sql.NullString // 密码(sha1(md5(明文))加密)
	Email     sql.NullString // 邮箱
	Phone     sql.NullString // 手机
	Status    sql.NullInt32  // 状态(1:正常启用 2:暂停使用)

}
