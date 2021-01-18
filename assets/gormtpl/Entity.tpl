package {{.PackageName}}

import (
	"math"
)


{{.TableComment}}
type {{.Table}} struct {
    Base
    {{range $j, $item := .Fields}}{{$item.Name}}       {{$item.Type}}    {{$item.FormatFields}}        {{$item.Remark}}
    {{end -}}
}


// 表名
func ({{.Table}}) TableName() string {
	return "{{.TableName}}"
}

// 添加记录
func (r *{{.Table}}) Add(model *{{.Table}}) (err error) {
	err = r.Db.Create(model).Error
	return
}

// 更新保存记录
func (r *{{.Table}}) Save(model *{{.Table}}) (err error) {
	err = r.Db.Save(model).Error
	return
}

// 软删除：结构体需要继承Base model 有delete_at字段
func (r *{{.Table}}) Delete(query interface{}, args ...interface{}) (err error) {
	//return r.Db.Unscoped().Where(query, args...).Delete(&{{.Table}}{}).Error //硬删除
	return r.Db.Where(query, args...).Delete(&{{.Table}}{}).Error
}

// 根据条件获取单挑记录
func (r *{{.Table}}) First(query interface{}, args ...interface{}) (model {{.Table}}, err error) {
	err = r.Db.Where(query, args...).First(&model).Error
	//err = SlaveDB.Where(query, args...).First(&model).Error //从
	return
}

// 获取列表 数据量大时Count数据需另外请求接口
func (r *{{.Table}}) Find(query interface{}, page *Pagination, args ...interface{}) (models []{{.Table}}, err error) {
	if page == nil {
		err = r.Db.Find(&models).Error
		//err = SlaveDB.Find(&models).Error //从
	} else {
		err = r.Db.Model({{.Table}}{}).Where(query, args...).
			//err = SlaveDB.Model({{.Table}}{}).Where(query, args...). //从
			Count(&page.Total).Offset((page.Page - 1) * page.PageSize).
			Limit(page.PageSize).Find(&models).Error
		// 总条数
		page.TotalPage = int64(math.Ceil(float64(page.Total / page.PageSize)))
	}

	return
}

// 获取总记录条数
func (r *{{.Table}}) Count(where interface{}, args ...interface{}) (count int64, err error) {
	err = r.Db.Model(&{{.Table}}{}).Where(where, args...).Count(&count).Error
	return
}