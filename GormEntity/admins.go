package entity

import (
	"math"
)

type Admins struct {
	Base
	UserName string `json:"user_name" gorm:"user_name"` // 管理员用户名
	RealName string `json:"real_name" gorm:"real_name"` // 管理员真实姓名
	Password string `json:"password" gorm:"password"`   // 密码(sha1(md5(明文))加密)
	Email    string `json:"email" gorm:"email"`         // 邮箱
	Phone    string `json:"phone" gorm:"phone"`         // 手机
	Status   int32  `json:"status" gorm:"status"`       // 状态(1:正常启用 2:暂停使用)
}

// 表名
func (Admins) TableName() string {
	return "admins"
}

// 添加记录
func (r *Admins) Add(model *Admins) (err error) {
	err = r.Db.Create(model).Error
	return
}

// 更新保存记录
func (r *Admins) Save(model *Admins) (err error) {
	err = r.Db.Save(model).Error
	return
}

// 软删除：结构体需要继承Base model 有delete_at字段
func (r *Admins) Delete(query interface{}, args ...interface{}) (err error) {
	//return r.Db.Unscoped().Where(query, args...).Delete(&Admins{}).Error //硬删除
	return r.Db.Where(query, args...).Delete(&Admins{}).Error
}

// 根据条件获取单挑记录
func (r *Admins) First(query interface{}, args ...interface{}) (model Admins, err error) {
	err = r.Db.Where(query, args...).First(&model).Error
	//err = SlaveDB.Where(query, args...).First(&model).Error //从
	return
}

// 获取列表 数据量大时Count数据需另外请求接口
func (r *Admins) Find(query interface{}, page *Pagination, args ...interface{}) (models []Admins, err error) {
	if page == nil {
		err = r.Db.Find(&models).Error
		//err = SlaveDB.Find(&models).Error //从
	} else {
		err = r.Db.Model(Admins{}).Where(query, args...).
			//err = SlaveDB.Model(Admins{}).Where(query, args...). //从
			Count(&page.Total).Offset((page.Page - 1) * page.PageSize).
			Limit(page.PageSize).Find(&models).Error
		// 总条数
		page.TotalPage = int64(math.Ceil(float64(page.Total / page.PageSize)))
	}

	return
}

// 获取总记录条数
func (r *Admins) Count(where interface{}, args ...interface{}) (count int64, err error) {
	err = r.Db.Model(&Admins{}).Where(where, args...).Count(&count).Error
	return
}
