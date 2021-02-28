package entity

import (
	"math"
)

type Roles struct {
	Base
	Name     string `json:"name" gorm:"name"`           // 角色名称
	ParentId int64  `json:"parent_id" gorm:"parent_id"` // 父级ID
	Status   int32  `json:"status" gorm:"status"`       // 状态(1:正常 2:暂停使用)
}

// 表名
func (Roles) TableName() string {
	return "roles"
}

// 添加记录
func (r *Roles) Add(model *Roles) (err error) {
	err = r.Db.Create(model).Error
	return
}

// 更新保存记录
func (r *Roles) Save(model *Roles) (err error) {
	err = r.Db.Save(model).Error
	return
}

// 软删除：结构体需要继承Base model 有delete_at字段
func (r *Roles) Delete(query interface{}, args ...interface{}) (err error) {
	//return r.Db.Unscoped().Where(query, args...).Delete(&Roles{}).Error //硬删除
	return r.Db.Where(query, args...).Delete(&Roles{}).Error
}

// 根据条件获取单挑记录
func (r *Roles) First(query interface{}, args ...interface{}) (model Roles, err error) {
	err = r.Db.Where(query, args...).First(&model).Error
	//err = SlaveDB.Where(query, args...).First(&model).Error //从
	return
}

// 获取列表 数据量大时Count数据需另外请求接口
func (r *Roles) Find(query interface{}, page *Pagination, args ...interface{}) (models []Roles, err error) {
	if page == nil {
		err = r.Db.Find(&models).Error
		//err = SlaveDB.Find(&models).Error //从
	} else {
		err = r.Db.Model(Roles{}).Where(query, args...).
			//err = SlaveDB.Model(Roles{}).Where(query, args...). //从
			Count(&page.Total).Offset((page.Page - 1) * page.PageSize).
			Limit(page.PageSize).Find(&models).Error
		// 总条数
		page.TotalPage = int64(math.Ceil(float64(page.Total / page.PageSize)))
	}

	return
}

// 获取总记录条数
func (r *Roles) Count(where interface{}, args ...interface{}) (count int64, err error) {
	err = r.Db.Model(&Roles{}).Where(where, args...).Count(&count).Error
	return
}
