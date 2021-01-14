package entity

import (
	"math"
)

type RoleMenus struct {
	Base
	RoleId int64 `json:"role_id" gorm:"role_id" xml:"role_id"` // 角色ID
	MenuId int64 `json:"menu_id" gorm:"menu_id" xml:"menu_id"` // 菜单ID
}

// 表名
func (RoleMenus) TableName() string {
	return "role_menus"
}

// 添加记录
func (r *RoleMenus) Add(model *RoleMenus) (err error) {
	err = r.Db.Create(model).Error
	return
}

// 更新保存记录
func (r *RoleMenus) Save(model *RoleMenus) (err error) {
	err = r.Db.Save(model).Error
	return
}

// 软删除：结构体需要继承Base model 有delete_at字段
func (r *RoleMenus) Delete(query interface{}, args ...interface{}) (err error) {
	//return r.Db.Unscoped().Where(query, args).Delete(&RoleMenus{}).Error //硬删除
	return r.Db.Where(query, args).Delete(&RoleMenus{}).Error
}

// 根据条件获取单挑记录
func (r *RoleMenus) First(query interface{}, args ...interface{}) (model RoleMenus, err error) {
	err = r.Db.Where(query, args...).First(&model).Error
	//err = SlaveDB.Where(query, args...).First(&model).Error //从
	return
}

// 获取列表
func (r *RoleMenus) Find(query interface{}, page *Pagination, args ...interface{}) (models []RoleMenus, err error) {
	if page == nil {
		err = r.Db.Find(&models).Error
		//err = SlaveDB.Find(&models).Error //从
	} else {
		err = r.Db.Model(RoleMenus{}).Where(query, args...).
			//err = SlaveDB.Model(RoleMenus{}).Where(query, args...). //从
			Count(&page.Total).Offset((page.Page - 1) * page.PageSize).
			Limit(page.PageSize).Find(&models).Error
		// 总条数
		page.TotalPage = int64(math.Ceil(float64(page.Total / page.PageSize)))
	}

	return
}
