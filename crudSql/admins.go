//
package mysql

import (
	"database/sql"
	"fmt"
	"math"
	"tpldemo/crudSql/config"
	"tpldemo/crudSql/entity"

	_ "github.com/go-sql-driver/mysql"
)

type AdminsModel struct {
	E
	DB *sql.DB
	Tx *sql.Tx
}

// not transaction
func NewAdmins(db ...*sql.DB) *AdminsModel {
	if len(db) > 0 {
		return &AdminsModel{
			DB: db[0],
		}
	}
	return &AdminsModel{
		DB: masterDB,
	}
}

// transaction object
func NewAdminsTx(tx *sql.Tx) *AdminsModel {
	return &AdminsModel{
		Tx: tx,
	}
}

// 获取所有的表字段
func (m *AdminsModel) getColumns() string {
	return " `id`,`created_at`,`updated_at`,`deleted_at`,`user_name`,`real_name`,`password`,`email`,`phone`,`status` "
}

// 获取多行数据.
func (m *AdminsModel) getRows(sqlTxt string, params ...interface{}) (rowsResult []*entity.Admins, err error) {
	query, err := m.DB.Query(sqlTxt, params...)
	if err != nil {
		err = m.E.Stack(err)
		return
	}
	defer query.Close()
	for query.Next() {
		row := entity.AdminsNull{}
		err = query.Scan(
			&row.Id,        //
			&row.CreatedAt, // 创建时间
			&row.UpdatedAt, // 更新时间
			&row.DeletedAt, // 删除时间
			&row.UserName,  // 管理员用户名
			&row.RealName,  // 管理员真实姓名
			&row.Password,  // 密码(sha1(md5(明文))加密)
			&row.Email,     // 邮箱
			&row.Phone,     // 手机
			&row.Status,    // 状态(1:正常启用 2:暂停使用)
		)
		if err != nil {
			err = m.E.Stack(err)
			return
		}
		rowsResult = append(rowsResult, &entity.Admins{
			Id:        row.Id.Int64,        //
			CreatedAt: row.CreatedAt.Int64, // 创建时间
			UpdatedAt: row.UpdatedAt.Int64, // 更新时间
			DeletedAt: row.DeletedAt.Time,  // 删除时间
			UserName:  row.UserName.String, // 管理员用户名
			RealName:  row.RealName.String, // 管理员真实姓名
			Password:  row.Password.String, // 密码(sha1(md5(明文))加密)
			Email:     row.Email.String,    // 邮箱
			Phone:     row.Phone.String,    // 手机
			Status:    row.Status.Int32,    // 状态(1:正常启用 2:暂停使用)
		})
	}
	return
}

// 获取单行数据
func (m *AdminsModel) getRow(sqlText string, params ...interface{}) (rowResult *entity.Admins, err error) {
	query := m.DB.QueryRow(sqlText, params...)
	row := entity.AdminsNull{}
	err = query.Scan(
		&row.Id,        //
		&row.CreatedAt, // 创建时间
		&row.UpdatedAt, // 更新时间
		&row.DeletedAt, // 删除时间
		&row.UserName,  // 管理员用户名
		&row.RealName,  // 管理员真实姓名
		&row.Password,  // 密码(sha1(md5(明文))加密)
		&row.Email,     // 邮箱
		&row.Phone,     // 手机
		&row.Status,    // 状态(1:正常启用 2:暂停使用)
	)
	if err != nil {
		err = m.E.Stack(err)
		return
	}
	rowResult = &entity.Admins{
		Id:        row.Id.Int64,        //
		CreatedAt: row.CreatedAt.Int64, // 创建时间
		UpdatedAt: row.UpdatedAt.Int64, // 更新时间
		DeletedAt: row.DeletedAt.Time,  // 删除时间
		UserName:  row.UserName.String, // 管理员用户名
		RealName:  row.RealName.String, // 管理员真实姓名
		Password:  row.Password.String, // 密码(sha1(md5(明文))加密)
		Email:     row.Email.String,    // 邮箱
		Phone:     row.Phone.String,    // 手机
		Status:    row.Status.Int32,    // 状态(1:正常启用 2:暂停使用)
	}
	return
}

// _更新数据
func (m *AdminsModel) Save(sqlTxt string, value ...interface{}) (b bool, err error) {
	stmt, err := m.DB.Prepare(sqlTxt)
	if err != nil {
		err = m.E.Stack(err)
		return
	}
	defer stmt.Close()
	result, err := stmt.Exec(value...)
	if err != nil {
		err = m.E.Stack(err)
		return
	}
	var affectCount int64
	affectCount, err = result.RowsAffected()
	if err != nil {
		err = m.E.Stack(err)
		return
	}
	b = affectCount > 0
	return
}

// _更新数据
func (m *AdminsModel) SaveTx(sqlTxt string, value ...interface{}) (b bool, err error) {
	stmt, err := m.Tx.Prepare(sqlTxt)
	if err != nil {
		err = m.E.Stack(err)
		return
	}
	defer stmt.Close()
	result, err := stmt.Exec(value...)
	if err != nil {
		err = m.E.Stack(err)
		return
	}
	var affectCount int64
	affectCount, err = result.RowsAffected()
	if err != nil {
		err = m.E.Stack(err)
		return
	}
	b = affectCount > 0
	return
}

// 新增信息
func (m *AdminsModel) Create(value *entity.Admins) (lastId int64, err error) {
	const sqlText = "INSERT INTO " + config.TABLE_ADMINS + " (`deleted_at`,`user_name`,`real_name`,`password`,`email`,`phone`,`status`) VALUES (?,?,?,?,?,?,?)"
	stmt, err := m.DB.Prepare(sqlText)
	if err != nil {
		err = m.E.Stack(err)
		return
	}
	defer stmt.Close()
	result, err := stmt.Exec(
		value.DeletedAt, // 删除时间
		value.UserName,  // 管理员用户名
		value.RealName,  // 管理员真实姓名
		value.Password,  // 密码(sha1(md5(明文))加密)
		value.Email,     // 邮箱
		value.Phone,     // 手机
		value.Status,    // 状态(1:正常启用 2:暂停使用)
	)
	if err != nil {
		err = m.E.Stack(err)
		return
	}
	lastId, err = result.LastInsertId()
	if err != nil {
		err = m.E.Stack(err)
		return
	}
	return
}

// 新增信息 tx
func (m *AdminsModel) CreateTx(value *entity.Admins) (lastId int64, err error) {
	const sqlText = "INSERT INTO " + config.TABLE_ADMINS + " (`deleted_at`,`user_name`,`real_name`,`password`,`email`,`phone`,`status`) VALUES (?,?,?,?,?,?,?)"
	stmt, err := m.Tx.Prepare(sqlText)
	if err != nil {
		err = m.E.Stack(err)
		return
	}
	defer stmt.Close()
	result, err := stmt.Exec(
		value.DeletedAt, // 删除时间
		value.UserName,  // 管理员用户名
		value.RealName,  // 管理员真实姓名
		value.Password,  // 密码(sha1(md5(明文))加密)
		value.Email,     // 邮箱
		value.Phone,     // 手机
		value.Status,    // 状态(1:正常启用 2:暂停使用)
	)
	if err != nil {
		err = m.E.Stack(err)
		return
	}
	lastId, err = result.LastInsertId()
	if err != nil {
		err = m.E.Stack(err)
		return
	}
	return
}

// 更新数据
func (m *AdminsModel) Update(value *entity.Admins) (b bool, err error) {
	const sqlText = "UPDATE " + config.TABLE_ADMINS + " SET `deleted_at`=?,`user_name`=?,`real_name`=?,`password`=?,`email`=?,`phone`=?,`status`=? WHERE `id` = ?"
	params := make([]interface{}, 0)
	params = append(params, value.DeletedAt)
	params = append(params, value.UserName)
	params = append(params, value.RealName)
	params = append(params, value.Password)
	params = append(params, value.Email)
	params = append(params, value.Phone)
	params = append(params, value.Status)
	params = append(params, value.Id)

	return m.Save(sqlText, params...)
}

// 更新数据 tx
func (m *AdminsModel) UpdateTx(value *entity.Admins) (b bool, err error) {
	const sqlText = "UPDATE " + config.TABLE_ADMINS + " SET `deleted_at`=?,`user_name`=?,`real_name`=?,`password`=?,`email`=?,`phone`=?,`status`=? WHERE `id` = ?"
	params := make([]interface{}, 0)
	params = append(params, value.DeletedAt)
	params = append(params, value.UserName)
	params = append(params, value.RealName)
	params = append(params, value.Password)
	params = append(params, value.Email)
	params = append(params, value.Phone)
	params = append(params, value.Status)
	params = append(params, value.Id)

	return m.SaveTx(sqlText, params...)
}

// 查询多行数据
func (m *AdminsModel) Find(where string, page *Pagination, args ...interface{}) (resList []*entity.Admins, err error) {
	sqlText := m.getQuerySql(where, page)
	if resList, err = m.getRows(sqlText, args...); err != nil {
		return
	}
	if page != nil {
		if page.Total, err = m.Count(); err != nil {
			return
		}
		page.TotalPage = int64(math.Ceil(float64(page.Total)/float64(page.PageSize)))
	}
	return
}

// select * from '表名' where '限制条件'  group by '分组依据' having '过滤条件' order by  limit '展示条数'
func (m *AdminsModel) getQuerySql(where string, page *Pagination) string {
	if page != nil && page.Page > 0 {
		return fmt.Sprintf("select distinct %v from %v where %v limit %v, %v",
			m.getColumns(),
			config.TABLE_ADMINS,
			where,
			page.PageSize*(page.Page-1),
			page.PageSize,
		)
	}
	return fmt.Sprintf("select distinct %v from %v where %v",
		m.getColumns(),
		config.TABLE_ADMINS,
		where,
	)
}

// 获取单行数据
func (m *AdminsModel) First(id int64) (result *entity.Admins, err error) {
	sqlText := "SELECT" + m.getColumns() + "FROM " + config.TABLE_ADMINS + " WHERE `id` = ? LIMIT 1"
	result, err = m.getRow(sqlText, id)
	return
}

// 获取最后一行数据
func (m *AdminsModel) Last() (result *entity.Admins, err error) {
	sqlText := "SELECT" + m.getColumns() + "FROM " + config.TABLE_ADMINS + " ORDER BY ID DESC LIMIT 1"
	result, err = m.getRow(sqlText)
	return
}

// 单列数据
func (m *AdminsModel) Pluck(id int64) (result map[int64]interface{}, err error) {
	const sqlText = "SELECT `id`, `created_at` FROM " + config.TABLE_ADMINS + " where `id` = ?"
	rows, err := m.DB.Query(sqlText, id)
	if err != nil {
		err = m.E.Stack(err)
		return
	}
	defer rows.Close()
	result = make(map[int64]interface{})
	var (
		_id  int64
		_val interface{}
	)
	for rows.Next() {
		err = rows.Scan(&_id, &_val)
		if err != nil {
			err = m.E.Stack(err)
			return
		}
		result[_id] = _val
	}
	return
}

// 单列数据 by 支持切片传入
// Get column data
func (m *AdminsModel) Plucks(ids []int64) (result map[int64]interface{}, err error) {
	result = make(map[int64]interface{})
	if len(ids) == 0 {
		return
	}
	sqlText := "SELECT `id`, `created_at` FROM " + config.TABLE_ADMINS + " where " +
		"`id` in (" + RepeatQuestionMark(len(ids)) + ")"
	params := make([]interface{}, len(ids))
	for idx, id := range ids {
		params[idx] = id
	}
	rows, err := m.DB.Query(sqlText, params...)
	if err != nil {
		err = m.E.Stack(err)
		return
	}
	defer rows.Close()
	var (
		_id  int64
		_val interface{}
	)
	for rows.Next() {
		err = rows.Scan(&_id, &_val)
		if err != nil {
			err = m.E.Stack(err)
			return
		}
		result[_id] = _val
	}
	return
}

// 获取单个数据
// Get one data
func (m *AdminsModel) One(id int64) (result int64, err error) {
	sqlText := "SELECT `id` FROM " + config.TABLE_ADMINS + " where `id`=?"
	err = m.DB.QueryRow(sqlText, id).Scan(&result)
	if err != nil && err != sql.ErrNoRows {
		err = m.E.Stack(err)
		return
	}
	return
}

// 获取行数
// Get line count
func (m *AdminsModel) Count() (count int64, err error) {
	sqlText := "SELECT COUNT(*) FROM " + config.TABLE_ADMINS
	err = m.DB.QueryRow(sqlText).Scan(&count)
	if err != nil && err != sql.ErrNoRows {
		err = m.E.Stack(err)
		return
	}
	return
}

// 判断数据是否存在
// Check the data is have?
func (m *AdminsModel) Has(id int64) (b bool, err error) {
	sqlText := "SELECT `id` FROM " + config.TABLE_ADMINS + " where `id` = ?"
	var count int64
	err = m.DB.QueryRow(sqlText, id).Scan(&count)
	if err != nil && err != sql.ErrNoRows {
		err = m.E.Stack(err)
		return
	}
	return count > 0, nil
}
