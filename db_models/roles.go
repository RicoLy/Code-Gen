package mysql

import (
	"database/sql"
	"fmt"
	"math"
	"time"
)

type RolesModel struct {
	DB *sql.DB
	Tx *sql.Tx
}

// not transaction
func NewRoles(db ...*sql.DB) *RolesModel {
	if len(db) > 0 {
		return &RolesModel{
			DB: db[0],
		}
	}
	return &RolesModel{
		DB: masterDB,
	}
}

// transaction object
func NewRolesTx(tx *sql.Tx) *RolesModel {
	return &RolesModel{
		Tx: tx,
	}
}

// 获取所有的表字段
func (m *RolesModel) getColumns() string {
	return " `id`,`created_at`,`updated_at`,`deleted_at`,`name`,`parent_id`,`status` "
}

// 获取多行数据.
func (m *RolesModel) getRows(sqlTxt string, params ...interface{}) (rowsResult []*entity.Roles, err error) {
	query, err := m.DB.Query(sqlTxt, params...)
	if err != nil {
		return
	}
	defer query.Close()
	for query.Next() {
		row := entity.RolesNull{}
		err = query.Scan(
			&row.Id,        //
			&row.CreatedAt, // 创建时间
			&row.UpdatedAt, // 更新时间
			&row.DeletedAt, // 删除时间
			&row.Name,      // 角色名称
			&row.ParentId,  // 父级ID
			&row.Status,    // 状态(1:正常 2:暂停使用)

		)
		if err != nil {
			return
		}
		rowsResult = append(rowsResult, &entity.Roles{
			Id:        row.Id.Int64,        //
			CreatedAt: row.CreatedAt.Int64, // 创建时间
			UpdatedAt: row.UpdatedAt.Int64, // 更新时间
			DeletedAt: row.DeletedAt.Time,  // 删除时间
			Name:      row.Name.String,     // 角色名称
			ParentId:  row.ParentId.Int64,  // 父级ID
			Status:    row.Status.Int32,    // 状态(1:正常 2:暂停使用)

		})
	}
	return
}

// 获取单行数据
func (m *RolesModel) getRow(sqlText string, params ...interface{}) (rowResult *entity.Roles, err error) {
	query := m.DB.QueryRow(sqlText, params...)
	row := entity.RolesNull{}
	err = query.Scan(
		&row.Id,        //
		&row.CreatedAt, // 创建时间
		&row.UpdatedAt, // 更新时间
		&row.DeletedAt, // 删除时间
		&row.Name,      // 角色名称
		&row.ParentId,  // 父级ID
		&row.Status,    // 状态(1:正常 2:暂停使用)

	)
	if err != nil {
		return
	}
	rowResult = &entity.Roles{
		Id:        row.Id.Int64,        //
		CreatedAt: row.CreatedAt.Int64, // 创建时间
		UpdatedAt: row.UpdatedAt.Int64, // 更新时间
		DeletedAt: row.DeletedAt.Time,  // 删除时间
		Name:      row.Name.String,     // 角色名称
		ParentId:  row.ParentId.Int64,  // 父级ID
		Status:    row.Status.Int32,    // 状态(1:正常 2:暂停使用)

	}
	return
}

//预编译
func (m *RolesModel) Prepare(sqlTxt string, value ...interface{}) (result sql.Result, err error) {
	stmt, err := m.DB.Prepare(sqlTxt)
	if err != nil {
		return
	}
	defer stmt.Close()
	result, err = stmt.Exec(value...)

	return
}

//预编译
func (m *RolesModel) PrepareTx(sqlTxt string, value ...interface{}) (result sql.Result, err error) {
	stmt, err := m.Tx.Prepare(sqlTxt)
	if err != nil {
		return
	}
	defer stmt.Close()
	result, err = stmt.Exec(value...)

	return
}

// 新增信息
func (m *RolesModel) Create(value *entity.Roles) (lastId int64, err error) {
	const sqlText = "INSERT INTO " + config.TABLE_ROLES + " (`name`,`parent_id`,`status`) VALUES (?,?,?)"

	params := make([]interface{}, 0)
	params = append(params,
		value.Name,     // 角色名称
		value.ParentId, // 父级ID
		value.Status,   // 状态(1:正常 2:暂停使用)

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
func (m *RolesModel) CreateTx(value *entity.Roles) (lastId int64, err error) {
	const sqlText = "INSERT INTO " + config.TABLE_ROLES + " (`name`,`parent_id`,`status`) VALUES (?,?,?)"
	params := make([]interface{}, 0)
	params = append(params,
		value.Name,     // 角色名称
		value.ParentId, // 父级ID
		value.Status,   // 状态(1:正常 2:暂停使用)

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
func (m *RolesModel) Update(value *entity.Roles) (b bool, err error) {
	const sqlText = "UPDATE " + config.TABLE_ROLES + " SET `name`=?,`parent_id`=?,`status`=? WHERE `id` = ?"
	params := make([]interface{}, 0)
	params = append(params,
		value.Name,     // 角色名称
		value.ParentId, // 父级ID
		value.Status,   // 状态(1:正常 2:暂停使用)
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
func (m *RolesModel) UpdateTx(value *entity.Roles) (b bool, err error) {
	const sqlText = "UPDATE " + config.TABLE_ROLES + " SET `name`=?,`parent_id`=?,`status`=? WHERE `id` = ?"
	params := make([]interface{}, 0)
	params = append(params,
		value.Name,     // 角色名称
		value.ParentId, // 父级ID
		value.Status,   // 状态(1:正常 2:暂停使用)
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

func (m *RolesModel) Delete(id int64) (b bool, err error) {
	const sqlText = "UPDATE " + config.TABLE_ROLES + " SET `deleted_at`=? WHERE `id` = ?"
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
func (m *RolesModel) Find(page *Pagination, query string, args ...interface{}) (resList []*entity.Roles, err error) {
	var sqlText string
	if page != nil && page.PageSize > 0 && page.Page > 0 { // 分页查询
		if page.Total, err = m.Count(query, args...); err != nil {
			return
		}
		page.TotalPage = int64(math.Ceil(float64(page.Total) / float64(page.PageSize)))
		//sqlText = "SELECT" + m.getColumns() + "FROM " + config.TABLE_ROLES + " WHERE " + query
		sqlText = fmt.Sprintf("select %v from %v where %v limit %v, %v",
			m.getColumns(),
			config.TABLE_ROLES,
			query,
			page.PageSize*(page.Page-1),
			page.PageSize,
		)
	} else {
		sqlText = "SELECT" + m.getColumns() + "FROM " + config.TABLE_ROLES + " WHERE " + query
	}

	resList, err = m.getRows(sqlText, args...)
	return
}

// 获取单行数据
func (m *RolesModel) First(query string, args ...interface{}) (result *entity.Roles, err error) {
	sqlText := "SELECT" + m.getColumns() + "FROM " + config.TABLE_ROLES + " WHERE " + query
	result, err = m.getRow(sqlText, args...)
	return
}

// 获取行数
// Get line count
func (m *RolesModel) Count(query string, args ...interface{}) (count int64, err error) {
	var sqlText string
	if query != "" {
		sqlText = "SELECT COUNT(*) FROM " + config.TABLE_ROLES + " WHERE " + query
	} else {
		sqlText = "SELECT COUNT(*) FROM " + config.TABLE_ROLES
	}
	err = m.DB.QueryRow(sqlText, args...).Scan(&count)
	if err != nil && err != sql.ErrNoRows {
		return
	}
	return
}
