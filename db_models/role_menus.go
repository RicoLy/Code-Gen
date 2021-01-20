package mysql

import (
	"database/sql"
	"fmt"
	"math"
	"time"
)

type RoleMenusModel struct {
	DB *sql.DB
	Tx *sql.Tx
}

// not transaction
func NewRoleMenus(db ...*sql.DB) *RoleMenusModel {
	if len(db) > 0 {
		return &RoleMenusModel{
			DB: db[0],
		}
	}
	return &RoleMenusModel{
		DB: masterDB,
	}
}

// transaction object
func NewRoleMenusTx(tx *sql.Tx) *RoleMenusModel {
	return &RoleMenusModel{
		Tx: tx,
	}
}

// 获取所有的表字段
func (m *RoleMenusModel) getColumns() string {
	return " `id`,`created_at`,`updated_at`,`deleted_at`,`role_id`,`menu_id` "
}

// 获取多行数据.
func (m *RoleMenusModel) getRows(sqlTxt string, params ...interface{}) (rowsResult []*entity.RoleMenus, err error) {
	query, err := m.DB.Query(sqlTxt, params...)
	if err != nil {
		return
	}
	defer query.Close()
	for query.Next() {
		row := entity.RoleMenusNull{}
		err = query.Scan(
			&row.Id,        //
			&row.CreatedAt, // 创建时间
			&row.UpdatedAt, // 更新时间
			&row.DeletedAt, // 删除时间
			&row.RoleId,    // 角色ID
			&row.MenuId,    // 菜单ID

		)
		if err != nil {
			return
		}
		rowsResult = append(rowsResult, &entity.RoleMenus{
			Id:        row.Id.Int64,        //
			CreatedAt: row.CreatedAt.Int64, // 创建时间
			UpdatedAt: row.UpdatedAt.Int64, // 更新时间
			DeletedAt: row.DeletedAt.Time,  // 删除时间
			RoleId:    row.RoleId.Int64,    // 角色ID
			MenuId:    row.MenuId.Int64,    // 菜单ID

		})
	}
	return
}

// 获取单行数据
func (m *RoleMenusModel) getRow(sqlText string, params ...interface{}) (rowResult *entity.RoleMenus, err error) {
	query := m.DB.QueryRow(sqlText, params...)
	row := entity.RoleMenusNull{}
	err = query.Scan(
		&row.Id,        //
		&row.CreatedAt, // 创建时间
		&row.UpdatedAt, // 更新时间
		&row.DeletedAt, // 删除时间
		&row.RoleId,    // 角色ID
		&row.MenuId,    // 菜单ID

	)
	if err != nil {
		return
	}
	rowResult = &entity.RoleMenus{
		Id:        row.Id.Int64,        //
		CreatedAt: row.CreatedAt.Int64, // 创建时间
		UpdatedAt: row.UpdatedAt.Int64, // 更新时间
		DeletedAt: row.DeletedAt.Time,  // 删除时间
		RoleId:    row.RoleId.Int64,    // 角色ID
		MenuId:    row.MenuId.Int64,    // 菜单ID

	}
	return
}

//预编译
func (m *RoleMenusModel) Prepare(sqlTxt string, value ...interface{}) (result sql.Result, err error) {
	stmt, err := m.DB.Prepare(sqlTxt)
	if err != nil {
		return
	}
	defer stmt.Close()
	result, err = stmt.Exec(value...)

	return
}

//预编译
func (m *RoleMenusModel) PrepareTx(sqlTxt string, value ...interface{}) (result sql.Result, err error) {
	stmt, err := m.Tx.Prepare(sqlTxt)
	if err != nil {
		return
	}
	defer stmt.Close()
	result, err = stmt.Exec(value...)

	return
}

// 新增信息
func (m *RoleMenusModel) Create(value *entity.RoleMenus) (lastId int64, err error) {
	const sqlText = "INSERT INTO " + config.TABLE_ROLE_MENUS + " (`role_id`,`menu_id`) VALUES (?,?)"

	params := make([]interface{}, 0)
	params = append(params,
		value.RoleId, // 角色ID
		value.MenuId, // 菜单ID

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
func (m *RoleMenusModel) CreateTx(value *entity.RoleMenus) (lastId int64, err error) {
	const sqlText = "INSERT INTO " + config.TABLE_ROLE_MENUS + " (`role_id`,`menu_id`) VALUES (?,?)"
	params := make([]interface{}, 0)
	params = append(params,
		value.RoleId, // 角色ID
		value.MenuId, // 菜单ID

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
func (m *RoleMenusModel) Update(value *entity.RoleMenus) (b bool, err error) {
	const sqlText = "UPDATE " + config.TABLE_ROLE_MENUS + " SET `role_id`=?,`menu_id`=? WHERE `id` = ?"
	params := make([]interface{}, 0)
	params = append(params,
		value.RoleId, // 角色ID
		value.MenuId, // 菜单ID
		value.Id,     //

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
func (m *RoleMenusModel) UpdateTx(value *entity.RoleMenus) (b bool, err error) {
	const sqlText = "UPDATE " + config.TABLE_ROLE_MENUS + " SET `role_id`=?,`menu_id`=? WHERE `id` = ?"
	params := make([]interface{}, 0)
	params = append(params,
		value.RoleId, // 角色ID
		value.MenuId, // 菜单ID
		value.Id,     //

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

func (m *RoleMenusModel) Delete(id int64) (b bool, err error) {
	const sqlText = "UPDATE " + config.TABLE_ROLE_MENUS + " SET `deleted_at`=? WHERE `id` = ?"
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
func (m *RoleMenusModel) Find(page *Pagination, query string, args ...interface{}) (resList []*entity.RoleMenus, err error) {
	var sqlText string
	if page != nil && page.PageSize > 0 && page.Page > 0 { // 分页查询
		if page.Total, err = m.Count(query, args...); err != nil {
			return
		}
		page.TotalPage = int64(math.Ceil(float64(page.Total) / float64(page.PageSize)))
		//sqlText = "SELECT" + m.getColumns() + "FROM " + config.TABLE_ROLE_MENUS + " WHERE " + query
		sqlText = fmt.Sprintf("select %v from %v where %v limit %v, %v",
			m.getColumns(),
			config.TABLE_ROLE_MENUS,
			query,
			page.PageSize*(page.Page-1),
			page.PageSize,
		)
	} else {
		sqlText = "SELECT" + m.getColumns() + "FROM " + config.TABLE_ROLE_MENUS + " WHERE " + query
	}

	resList, err = m.getRows(sqlText, args...)
	return
}

// 获取单行数据
func (m *RoleMenusModel) First(query string, args ...interface{}) (result *entity.RoleMenus, err error) {
	sqlText := "SELECT" + m.getColumns() + "FROM " + config.TABLE_ROLE_MENUS + " WHERE " + query
	result, err = m.getRow(sqlText, args...)
	return
}

// 获取行数
// Get line count
func (m *RoleMenusModel) Count(query string, args ...interface{}) (count int64, err error) {
	var sqlText string
	if query != "" {
		sqlText = "SELECT COUNT(*) FROM " + config.TABLE_ROLE_MENUS + " WHERE " + query
	} else {
		sqlText = "SELECT COUNT(*) FROM " + config.TABLE_ROLE_MENUS
	}
	err = m.DB.QueryRow(sqlText, args...).Scan(&count)
	if err != nil && err != sql.ErrNoRows {
		return
	}
	return
}
