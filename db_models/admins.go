package mysql

import (
	"code-gen/db_models/config"
	"code-gen/db_models/entity"
	"database/sql"
	"fmt"
	"math"
	"time"
)

type AdminsModel struct {
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

//预编译
func (m *AdminsModel) Prepare(sqlTxt string, value ...interface{}) (result sql.Result, err error) {
	stmt, err := m.DB.Prepare(sqlTxt)
	if err != nil {
		return
	}
	defer stmt.Close()
	result, err = stmt.Exec(value...)

	return
}

//预编译
func (m *AdminsModel) PrepareTx(sqlTxt string, value ...interface{}) (result sql.Result, err error) {
	stmt, err := m.Tx.Prepare(sqlTxt)
	if err != nil {
		return
	}
	defer stmt.Close()
	result, err = stmt.Exec(value...)

	return
}

// 新增信息
func (m *AdminsModel) Create(value *entity.Admins) (lastId int64, err error) {
	const sqlText = "INSERT INTO " + config.TABLE_ADMINS + " (`user_name`,`real_name`,`password`,`email`,`phone`,`status`) VALUES (?,?,?,?,?,?)"

	params := make([]interface{}, 0)
	params = append(params,
		value.UserName, // 管理员用户名
		value.RealName, // 管理员真实姓名
		value.Password, // 密码(sha1(md5(明文))加密)
		value.Email,    // 邮箱
		value.Phone,    // 手机
		value.Status,   // 状态(1:正常启用 2:暂停使用)

	)
	result, err := m.Prepare(sqlText, params...)
	if err != nil {
		return
	}
	lastId, err = result.LastInsertId()
	if err != nil {
		return
	}
	return
}

// 新增信息 tx
func (m *AdminsModel) CreateTx(value *entity.Admins) (lastId int64, err error) {
	const sqlText = "INSERT INTO " + config.TABLE_ADMINS + " (`user_name`,`real_name`,`password`,`email`,`phone`,`status`) VALUES (?,?,?,?,?,?)"
	params := make([]interface{}, 0)
	params = append(params,
		value.UserName, // 管理员用户名
		value.RealName, // 管理员真实姓名
		value.Password, // 密码(sha1(md5(明文))加密)
		value.Email,    // 邮箱
		value.Phone,    // 手机
		value.Status,   // 状态(1:正常启用 2:暂停使用)

	)

	result, err := m.PrepareTx(sqlText, params...)
	if err != nil {
		return
	}
	lastId, err = result.LastInsertId()
	if err != nil {
		return
	}
	return
}

// 更新数据
func (m *AdminsModel) Update(value *entity.Admins) (b bool, err error) {
	const sqlText = "UPDATE " + config.TABLE_ADMINS + " SET `user_name`=?,`real_name`=?,`password`=?,`email`=?,`phone`=?,`status`=? WHERE `id` = ?"
	params := make([]interface{}, 0)
	params = append(params,
		value.UserName, // 管理员用户名
		value.RealName, // 管理员真实姓名
		value.Password, // 密码(sha1(md5(明文))加密)
		value.Email,    // 邮箱
		value.Phone,    // 手机
		value.Status,   // 状态(1:正常启用 2:暂停使用)
		value.Id,       //

	)

	result, err := m.Prepare(sqlText, params...)
	if err != nil {
		return
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return
	}

	return affected > 0, nil
}

// 更新数据 tx
func (m *AdminsModel) UpdateTx(value *entity.Admins) (b bool, err error) {
	const sqlText = "UPDATE " + config.TABLE_ADMINS + " SET `user_name`=?,`real_name`=?,`password`=?,`email`=?,`phone`=?,`status`=? WHERE `id` = ?"
	params := make([]interface{}, 0)
	params = append(params,
		value.UserName, // 管理员用户名
		value.RealName, // 管理员真实姓名
		value.Password, // 密码(sha1(md5(明文))加密)
		value.Email,    // 邮箱
		value.Phone,    // 手机
		value.Status,   // 状态(1:正常启用 2:暂停使用)
		value.Id,       //

	)

	result, err := m.PrepareTx(sqlText, params...)
	if err != nil {
		return
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return
	}

	return affected > 0, nil
}

func (m *AdminsModel) Delete(id int64) (b bool, err error) {
	const sqlText = "UPDATE " + config.TABLE_ADMINS + " SET `deleted_at`=? WHERE `id` = ?"
	result, err := m.Prepare(sqlText, time.Now(), id)
	if err != nil {
		return
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return
	}
	return affected > 0, nil
}

// 查询多行数据 数据量大时Count数据需另外请求接口
func (m *AdminsModel) Find(page *Pagination, query string, args ...interface{}) (resList []*entity.Admins, err error) {
	var sqlText string
	if page != nil && page.PageSize > 0 && page.Page > 0 { // 分页查询
		if page.Total, err = m.Count(query, args...); err != nil {
			return
		}
		page.TotalPage = int64(math.Ceil(float64(page.Total) / float64(page.PageSize)))
		//sqlText = "SELECT" + m.getColumns() + "FROM " + config.TABLE_ADMINS + " WHERE " + query
		sqlText = fmt.Sprintf("select %v from %v where %v limit %v, %v",
			m.getColumns(),
			config.TABLE_ADMINS,
			query,
			page.PageSize*(page.Page-1),
			page.PageSize,
		)
	} else {
		sqlText = "SELECT" + m.getColumns() + "FROM " + config.TABLE_ADMINS + " WHERE " + query
	}

	resList, err = m.getRows(sqlText, args...)
	return
}

// 获取单行数据
func (m *AdminsModel) First(query string, args ...interface{}) (result *entity.Admins, err error) {
	sqlText := "SELECT" + m.getColumns() + "FROM " + config.TABLE_ADMINS + " WHERE " + query
	result, err = m.getRow(sqlText, args...)
	return
}

// 获取行数
// Get line count
func (m *AdminsModel) Count(query string, args ...interface{}) (count int64, err error) {
	var sqlText string
	if query != "" {
		sqlText = "SELECT COUNT(*) FROM " + config.TABLE_ADMINS + " WHERE " + query
	} else {
		sqlText = "SELECT COUNT(*) FROM " + config.TABLE_ADMINS
	}
	err = m.DB.QueryRow(sqlText, args...).Scan(&count)
	if err != nil && err != sql.ErrNoRows {
		return
	}
	return
}
