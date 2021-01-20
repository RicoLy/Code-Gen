package mysql

import (
	"database/sql"
	"fmt"
	"math"
	"time"
)

type AdminsRolesModel struct {
	DB *sql.DB
	Tx *sql.Tx
}

// not transaction
func NewAdminsRoles(db ...*sql.DB) *AdminsRolesModel {
	if len(db) > 0 {
		return &AdminsRolesModel{
			DB: db[0],
		}
	}
	return &AdminsRolesModel{
		DB: masterDB,
	}
}

// transaction object
func NewAdminsRolesTx(tx *sql.Tx) *AdminsRolesModel {
	return &AdminsRolesModel{
		Tx: tx,
	}
}

// 获取所有的表字段
func (m *AdminsRolesModel) getColumns() string {
	return " `id`,`created_at`,`updated_at`,`deleted_at`,`admins_id`,`role_id` "
}

// 获取多行数据.
func (m *AdminsRolesModel) getRows(sqlTxt string, params ...interface{}) (rowsResult []*entity.AdminsRoles, err error) {
	query, err := m.DB.Query(sqlTxt, params...)
	if err != nil {
		return
	}
	defer query.Close()
	for query.Next() {
		row := entity.AdminsRolesNull{}
		err = query.Scan(
			&row.Id,        //
			&row.CreatedAt, // 创建时间
			&row.UpdatedAt, // 更新时间
			&row.DeletedAt, // 删除时间
			&row.AdminsId,  // 管理员ID
			&row.RoleId,    // 角色ID

		)
		if err != nil {
			return
		}
		rowsResult = append(rowsResult, &entity.AdminsRoles{
			Id:        row.Id.Int64,        //
			CreatedAt: row.CreatedAt.Int64, // 创建时间
			UpdatedAt: row.UpdatedAt.Int64, // 更新时间
			DeletedAt: row.DeletedAt.Time,  // 删除时间
			AdminsId:  row.AdminsId.Int64,  // 管理员ID
			RoleId:    row.RoleId.Int64,    // 角色ID

		})
	}
	return
}

// 获取单行数据
func (m *AdminsRolesModel) getRow(sqlText string, params ...interface{}) (rowResult *entity.AdminsRoles, err error) {
	query := m.DB.QueryRow(sqlText, params...)
	row := entity.AdminsRolesNull{}
	err = query.Scan(
		&row.Id,        //
		&row.CreatedAt, // 创建时间
		&row.UpdatedAt, // 更新时间
		&row.DeletedAt, // 删除时间
		&row.AdminsId,  // 管理员ID
		&row.RoleId,    // 角色ID

	)
	if err != nil {
		return
	}
	rowResult = &entity.AdminsRoles{
		Id:        row.Id.Int64,        //
		CreatedAt: row.CreatedAt.Int64, // 创建时间
		UpdatedAt: row.UpdatedAt.Int64, // 更新时间
		DeletedAt: row.DeletedAt.Time,  // 删除时间
		AdminsId:  row.AdminsId.Int64,  // 管理员ID
		RoleId:    row.RoleId.Int64,    // 角色ID

	}
	return
}

//预编译
func (m *AdminsRolesModel) Prepare(sqlTxt string, value ...interface{}) (result sql.Result, err error) {
	stmt, err := m.DB.Prepare(sqlTxt)
	if err != nil {
		return
	}
	defer stmt.Close()
	result, err = stmt.Exec(value...)

	return
}

//预编译
func (m *AdminsRolesModel) PrepareTx(sqlTxt string, value ...interface{}) (result sql.Result, err error) {
	stmt, err := m.Tx.Prepare(sqlTxt)
	if err != nil {
		return
	}
	defer stmt.Close()
	result, err = stmt.Exec(value...)

	return
}

// 新增信息
func (m *AdminsRolesModel) Create(value *entity.AdminsRoles) (lastId int64, err error) {
	const sqlText = "INSERT INTO " + config.TABLE_ADMINS_ROLES + " (`admins_id`,`role_id`) VALUES (?,?)"

	params := make([]interface{}, 0)
	params = append(params,
		value.AdminsId, // 管理员ID
		value.RoleId,   // 角色ID

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
func (m *AdminsRolesModel) CreateTx(value *entity.AdminsRoles) (lastId int64, err error) {
	const sqlText = "INSERT INTO " + config.TABLE_ADMINS_ROLES + " (`admins_id`,`role_id`) VALUES (?,?)"
	params := make([]interface{}, 0)
	params = append(params,
		value.AdminsId, // 管理员ID
		value.RoleId,   // 角色ID

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
func (m *AdminsRolesModel) Update(value *entity.AdminsRoles) (b bool, err error) {
	const sqlText = "UPDATE " + config.TABLE_ADMINS_ROLES + " SET `admins_id`=?,`role_id`=? WHERE `id` = ?"
	params := make([]interface{}, 0)
	params = append(params,
		value.AdminsId, // 管理员ID
		value.RoleId,   // 角色ID
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
func (m *AdminsRolesModel) UpdateTx(value *entity.AdminsRoles) (b bool, err error) {
	const sqlText = "UPDATE " + config.TABLE_ADMINS_ROLES + " SET `admins_id`=?,`role_id`=? WHERE `id` = ?"
	params := make([]interface{}, 0)
	params = append(params,
		value.AdminsId, // 管理员ID
		value.RoleId,   // 角色ID
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

func (m *AdminsRolesModel) Delete(id int64) (b bool, err error) {
	const sqlText = "UPDATE " + config.TABLE_ADMINS_ROLES + " SET `deleted_at`=? WHERE `id` = ?"
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
func (m *AdminsRolesModel) Find(page *Pagination, query string, args ...interface{}) (resList []*entity.AdminsRoles, err error) {
	var sqlText string
	if page != nil && page.PageSize > 0 && page.Page > 0 { // 分页查询
		if page.Total, err = m.Count(query, args...); err != nil {
			return
		}
		page.TotalPage = int64(math.Ceil(float64(page.Total) / float64(page.PageSize)))
		//sqlText = "SELECT" + m.getColumns() + "FROM " + config.TABLE_ADMINS_ROLES + " WHERE " + query
		sqlText = fmt.Sprintf("select %v from %v where %v limit %v, %v",
			m.getColumns(),
			config.TABLE_ADMINS_ROLES,
			query,
			page.PageSize*(page.Page-1),
			page.PageSize,
		)
	} else {
		sqlText = "SELECT" + m.getColumns() + "FROM " + config.TABLE_ADMINS_ROLES + " WHERE " + query
	}

	resList, err = m.getRows(sqlText, args...)
	return
}

// 获取单行数据
func (m *AdminsRolesModel) First(query string, args ...interface{}) (result *entity.AdminsRoles, err error) {
	sqlText := "SELECT" + m.getColumns() + "FROM " + config.TABLE_ADMINS_ROLES + " WHERE " + query
	result, err = m.getRow(sqlText, args...)
	return
}

// 获取行数
// Get line count
func (m *AdminsRolesModel) Count(query string, args ...interface{}) (count int64, err error) {
	var sqlText string
	if query != "" {
		sqlText = "SELECT COUNT(*) FROM " + config.TABLE_ADMINS_ROLES + " WHERE " + query
	} else {
		sqlText = "SELECT COUNT(*) FROM " + config.TABLE_ADMINS_ROLES
	}
	err = m.DB.QueryRow(sqlText, args...).Scan(&count)
	if err != nil && err != sql.ErrNoRows {
		return
	}
	return
}
