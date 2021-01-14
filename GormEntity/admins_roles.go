package entity

import (
	"math"
)

type AdminsRoles struct {
	Base
	AdminsId int64 `json:"admins_id" gorm:"admins_id" xml:"admins_id"` // 管理员ID
	RoleId   int64 `json:"role_id" gorm:"role_id" xml:"role_id"`       // 角色ID
}

// 表名
func (AdminsRoles) TableName() string {
	return "admins_roles"
}

// 添加记录
func (r *AdminsRoles) Add(model *AdminsRoles) (err error) {
	err = r.Db.Create(model).Error
	return
}

// 更新保存记录
func (r *AdminsRoles) Save(model *AdminsRoles) (err error) {
	err = r.Db.Save(model).Error
	return
}

// 软删除：结构体需要继承Base model 有delete_at字段
func (r *AdminsRoles) Delete(query interface{}, args ...interface{}) (err error) {
	//return r.Db.Unscoped().Where(query, args).Delete(&AdminsRoles{}).Error //硬删除
	return r.Db.Where(query, args).Delete(&AdminsRoles{}).Error
}

// 根据条件获取单挑记录
func (r *AdminsRoles) First(query interface{}, args ...interface{}) (model AdminsRoles, err error) {
	err = r.Db.Where(query, args...).First(&model).Error
	//err = SlaveDB.Where(query, args...).First(&model).Error //从
	return
}

// 获取列表
func (r *AdminsRoles) Find(query interface{}, page *Pagination, args ...interface{}) (models []AdminsRoles, err error) {
	if page == nil {
		err = r.Db.Find(&models).Error
		//err = SlaveDB.Find(&models).Error //从
	} else {
		err = r.Db.Model(AdminsRoles{}).Where(query, args...).
			//err = SlaveDB.Model(AdminsRoles{}).Where(query, args...). //从
			Count(&page.Total).Offset((page.Page - 1) * page.PageSize).
			Limit(page.PageSize).Find(&models).Error
		// 总条数
		page.TotalPage = int64(math.Ceil(float64(page.Total / page.PageSize)))
	}

	return
}
