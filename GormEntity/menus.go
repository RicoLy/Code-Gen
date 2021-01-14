package entity

import (
	"math"
)

type Menus struct {
	Base
	Status        int32  `json:"status" gorm:"status" xml:"status"`                         // 状态(1:启用 2:不启用)
	ParentId      int64  `json:"parent_id" gorm:"parent_id" xml:"parent_id"`                // 父级ID
	FrontPath     string `json:"front_path" gorm:"front_path" xml:"front_path"`             // 前端文件路径
	Url           string `json:"url" gorm:"url" xml:"url"`                                  // 菜单api路径URL
	Method        string `json:"method" gorm:"method" xml:"method"`                         // 操作方法 /GET/POST/PUT/DELETE/
	Name          string `json:"name" gorm:"name" xml:"name"`                               // 菜单名称
	InterfaceName string `json:"interface_name" gorm:"interface_name" xml:"interface_name"` // 接口名称
	MenuType      int32  `json:"menu_type" gorm:"menu_type" xml:"menu_type"`                // 菜单类型(1:模块 2:菜单 3:操作)
}

// 表名
func (Menus) TableName() string {
	return "menus"
}

// 添加记录
func (r *Menus) Add(model *Menus) (err error) {
	err = r.Db.Create(model).Error
	return
}

// 更新保存记录
func (r *Menus) Save(model *Menus) (err error) {
	err = r.Db.Save(model).Error
	return
}

// 软删除：结构体需要继承Base model 有delete_at字段
func (r *Menus) Delete(query interface{}, args ...interface{}) (err error) {
	//return r.Db.Unscoped().Where(query, args).Delete(&Menus{}).Error //硬删除
	return r.Db.Where(query, args).Delete(&Menus{}).Error
}

// 根据条件获取单挑记录
func (r *Menus) First(query interface{}, args ...interface{}) (model Menus, err error) {
	err = r.Db.Where(query, args...).First(&model).Error
	//err = SlaveDB.Where(query, args...).First(&model).Error //从
	return
}

// 获取列表
func (r *Menus) Find(query interface{}, page *Pagination, args ...interface{}) (models []Menus, err error) {
	if page == nil {
		err = r.Db.Find(&models).Error
		//err = SlaveDB.Find(&models).Error //从
	} else {
		err = r.Db.Model(Menus{}).Where(query, args...).
			//err = SlaveDB.Model(Menus{}).Where(query, args...). //从
			Count(&page.Total).Offset((page.Page - 1) * page.PageSize).
			Limit(page.PageSize).Find(&models).Error
		// 总条数
		page.TotalPage = int64(math.Ceil(float64(page.Total / page.PageSize)))
	}

	return
}
