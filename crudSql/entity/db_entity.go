// 判断package是否加载过
package entity

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"time"
)

type Admins struct {
	Id        int64
	CreatedAt int64     // 创建时间
	UpdatedAt int64     // 更新时间
	DeletedAt time.Time // 删除时间
	UserName  string    // 管理员用户名
	RealName  string    // 管理员真实姓名
	Password  string    // 密码(sha1(md5(明文))加密)
	Email     string    // 邮箱
	Phone     string    // 手机
	Status    int32     // 状态(1:正常启用 2:暂停使用)
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

type AdminsRoles struct {
	Id        int64
	CreatedAt int64     // 创建时间
	UpdatedAt int64     // 更新时间
	DeletedAt time.Time // 删除时间
	AdminsId  int64     // 管理员ID
	RoleId    int64     // 角色ID
}

type AdminsRolesNull struct {
	Id        sql.NullInt64
	CreatedAt sql.NullInt64  // 创建时间
	UpdatedAt sql.NullInt64  // 更新时间
	DeletedAt mysql.NullTime // 删除时间
	AdminsId  sql.NullInt64  // 管理员ID
	RoleId    sql.NullInt64  // 角色ID
}

type Menus struct {
	Id            int64
	CreatedAt     int64     // 创建时间
	UpdatedAt     int64     // 更新时间
	DeletedAt     time.Time // 删除时间
	Status        int32     // 状态(1:启用 2:不启用)
	ParentId      int64     // 父级ID
	FrontPath     string    // 前端文件路径
	Url           string    // 菜单api路径URL
	Method        string    // 操作方法 /GET/POST/PUT/DELETE/
	Name          string    // 菜单名称
	InterfaceName string    // 接口名称
	MenuType      int32     // 菜单类型(1:模块 2:菜单 3:操作)
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

type RoleMenus struct {
	Id        int64
	CreatedAt int64     // 创建时间
	UpdatedAt int64     // 更新时间
	DeletedAt time.Time // 删除时间
	RoleId    int64     // 角色ID
	MenuId    int64     // 菜单ID
}

type RoleMenusNull struct {
	Id        sql.NullInt64
	CreatedAt sql.NullInt64  // 创建时间
	UpdatedAt sql.NullInt64  // 更新时间
	DeletedAt mysql.NullTime // 删除时间
	RoleId    sql.NullInt64  // 角色ID
	MenuId    sql.NullInt64  // 菜单ID
}

type Roles struct {
	Id        int64
	CreatedAt int64     // 创建时间
	UpdatedAt int64     // 更新时间
	DeletedAt time.Time // 删除时间
	Name      string    // 角色名称
	ParentId  int64     // 父级ID
	Status    int32     // 状态(1:正常 2:暂停使用)
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
