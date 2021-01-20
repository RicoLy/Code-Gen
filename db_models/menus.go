package mysql

import (
	"database/sql"
	"fmt"
	"math"
	"time"
)

type MenusModel struct {
	DB *sql.DB
	Tx *sql.Tx
}

// not transaction
func NewMenus(db ...*sql.DB) *MenusModel {
	if len(db) > 0 {
		return &MenusModel{
			DB: db[0],
		}
	}
	return &MenusModel{
		DB: masterDB,
	}
}

// transaction object
func NewMenusTx(tx *sql.Tx) *MenusModel {
	return &MenusModel{
		Tx: tx,
	}
}

// 获取所有的表字段
func (m *MenusModel) getColumns() string {
	return " `id`,`created_at`,`updated_at`,`deleted_at`,`status`,`parent_id`,`front_path`,`url`,`method`,`name`,`interface_name`,`menu_type` "
}

// 获取多行数据.
func (m *MenusModel) getRows(sqlTxt string, params ...interface{}) (rowsResult []*entity.Menus, err error) {
	query, err := m.DB.Query(sqlTxt, params...)
	if err != nil {
		return
	}
	defer query.Close()
	for query.Next() {
		row := entity.MenusNull{}
		err = query.Scan(
			&row.Id,            //
			&row.CreatedAt,     // 创建时间
			&row.UpdatedAt,     // 更新时间
			&row.DeletedAt,     // 删除时间
			&row.Status,        // 状态(1:启用 2:不启用)
			&row.ParentId,      // 父级ID
			&row.FrontPath,     // 前端文件路径
			&row.Url,           // 菜单api路径URL
			&row.Method,        // 操作方法 /GET/POST/PUT/DELETE/
			&row.Name,          // 菜单名称
			&row.InterfaceName, // 接口名称
			&row.MenuType,      // 菜单类型(1:模块 2:菜单 3:操作)

		)
		if err != nil {
			return
		}
		rowsResult = append(rowsResult, &entity.Menus{
			Id:            row.Id.Int64,             //
			CreatedAt:     row.CreatedAt.Int64,      // 创建时间
			UpdatedAt:     row.UpdatedAt.Int64,      // 更新时间
			DeletedAt:     row.DeletedAt.Time,       // 删除时间
			Status:        row.Status.Int32,         // 状态(1:启用 2:不启用)
			ParentId:      row.ParentId.Int64,       // 父级ID
			FrontPath:     row.FrontPath.String,     // 前端文件路径
			Url:           row.Url.String,           // 菜单api路径URL
			Method:        row.Method.String,        // 操作方法 /GET/POST/PUT/DELETE/
			Name:          row.Name.String,          // 菜单名称
			InterfaceName: row.InterfaceName.String, // 接口名称
			MenuType:      row.MenuType.Int32,       // 菜单类型(1:模块 2:菜单 3:操作)

		})
	}
	return
}

// 获取单行数据
func (m *MenusModel) getRow(sqlText string, params ...interface{}) (rowResult *entity.Menus, err error) {
	query := m.DB.QueryRow(sqlText, params...)
	row := entity.MenusNull{}
	err = query.Scan(
		&row.Id,            //
		&row.CreatedAt,     // 创建时间
		&row.UpdatedAt,     // 更新时间
		&row.DeletedAt,     // 删除时间
		&row.Status,        // 状态(1:启用 2:不启用)
		&row.ParentId,      // 父级ID
		&row.FrontPath,     // 前端文件路径
		&row.Url,           // 菜单api路径URL
		&row.Method,        // 操作方法 /GET/POST/PUT/DELETE/
		&row.Name,          // 菜单名称
		&row.InterfaceName, // 接口名称
		&row.MenuType,      // 菜单类型(1:模块 2:菜单 3:操作)

	)
	if err != nil {
		return
	}
	rowResult = &entity.Menus{
		Id:            row.Id.Int64,             //
		CreatedAt:     row.CreatedAt.Int64,      // 创建时间
		UpdatedAt:     row.UpdatedAt.Int64,      // 更新时间
		DeletedAt:     row.DeletedAt.Time,       // 删除时间
		Status:        row.Status.Int32,         // 状态(1:启用 2:不启用)
		ParentId:      row.ParentId.Int64,       // 父级ID
		FrontPath:     row.FrontPath.String,     // 前端文件路径
		Url:           row.Url.String,           // 菜单api路径URL
		Method:        row.Method.String,        // 操作方法 /GET/POST/PUT/DELETE/
		Name:          row.Name.String,          // 菜单名称
		InterfaceName: row.InterfaceName.String, // 接口名称
		MenuType:      row.MenuType.Int32,       // 菜单类型(1:模块 2:菜单 3:操作)

	}
	return
}

//预编译
func (m *MenusModel) Prepare(sqlTxt string, value ...interface{}) (result sql.Result, err error) {
	stmt, err := m.DB.Prepare(sqlTxt)
	if err != nil {
		return
	}
	defer stmt.Close()
	result, err = stmt.Exec(value...)

	return
}

//预编译
func (m *MenusModel) PrepareTx(sqlTxt string, value ...interface{}) (result sql.Result, err error) {
	stmt, err := m.Tx.Prepare(sqlTxt)
	if err != nil {
		return
	}
	defer stmt.Close()
	result, err = stmt.Exec(value...)

	return
}

// 新增信息
func (m *MenusModel) Create(value *entity.Menus) (lastId int64, err error) {
	const sqlText = "INSERT INTO " + config.TABLE_MENUS + " (`status`,`parent_id`,`front_path`,`url`,`method`,`name`,`interface_name`,`menu_type`) VALUES (?,?,?,?,?,?,?,?)"

	params := make([]interface{}, 0)
	params = append(params,
		value.Status,        // 状态(1:启用 2:不启用)
		value.ParentId,      // 父级ID
		value.FrontPath,     // 前端文件路径
		value.Url,           // 菜单api路径URL
		value.Method,        // 操作方法 /GET/POST/PUT/DELETE/
		value.Name,          // 菜单名称
		value.InterfaceName, // 接口名称
		value.MenuType,      // 菜单类型(1:模块 2:菜单 3:操作)

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
func (m *MenusModel) CreateTx(value *entity.Menus) (lastId int64, err error) {
	const sqlText = "INSERT INTO " + config.TABLE_MENUS + " (`status`,`parent_id`,`front_path`,`url`,`method`,`name`,`interface_name`,`menu_type`) VALUES (?,?,?,?,?,?,?,?)"
	params := make([]interface{}, 0)
	params = append(params,
		value.Status,        // 状态(1:启用 2:不启用)
		value.ParentId,      // 父级ID
		value.FrontPath,     // 前端文件路径
		value.Url,           // 菜单api路径URL
		value.Method,        // 操作方法 /GET/POST/PUT/DELETE/
		value.Name,          // 菜单名称
		value.InterfaceName, // 接口名称
		value.MenuType,      // 菜单类型(1:模块 2:菜单 3:操作)

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
func (m *MenusModel) Update(value *entity.Menus) (b bool, err error) {
	const sqlText = "UPDATE " + config.TABLE_MENUS + " SET `status`=?,`parent_id`=?,`front_path`=?,`url`=?,`method`=?,`name`=?,`interface_name`=?,`menu_type`=? WHERE `id` = ?"
	params := make([]interface{}, 0)
	params = append(params,
		value.Status,        // 状态(1:启用 2:不启用)
		value.ParentId,      // 父级ID
		value.FrontPath,     // 前端文件路径
		value.Url,           // 菜单api路径URL
		value.Method,        // 操作方法 /GET/POST/PUT/DELETE/
		value.Name,          // 菜单名称
		value.InterfaceName, // 接口名称
		value.MenuType,      // 菜单类型(1:模块 2:菜单 3:操作)
		value.Id,            //

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
func (m *MenusModel) UpdateTx(value *entity.Menus) (b bool, err error) {
	const sqlText = "UPDATE " + config.TABLE_MENUS + " SET `status`=?,`parent_id`=?,`front_path`=?,`url`=?,`method`=?,`name`=?,`interface_name`=?,`menu_type`=? WHERE `id` = ?"
	params := make([]interface{}, 0)
	params = append(params,
		value.Status,        // 状态(1:启用 2:不启用)
		value.ParentId,      // 父级ID
		value.FrontPath,     // 前端文件路径
		value.Url,           // 菜单api路径URL
		value.Method,        // 操作方法 /GET/POST/PUT/DELETE/
		value.Name,          // 菜单名称
		value.InterfaceName, // 接口名称
		value.MenuType,      // 菜单类型(1:模块 2:菜单 3:操作)
		value.Id,            //

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

func (m *MenusModel) Delete(id int64) (b bool, err error) {
	const sqlText = "UPDATE " + config.TABLE_MENUS + " SET `deleted_at`=? WHERE `id` = ?"
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
func (m *MenusModel) Find(page *Pagination, query string, args ...interface{}) (resList []*entity.Menus, err error) {
	var sqlText string
	if page != nil && page.PageSize > 0 && page.Page > 0 { // 分页查询
		if page.Total, err = m.Count(query, args...); err != nil {
			return
		}
		page.TotalPage = int64(math.Ceil(float64(page.Total) / float64(page.PageSize)))
		//sqlText = "SELECT" + m.getColumns() + "FROM " + config.TABLE_MENUS + " WHERE " + query
		sqlText = fmt.Sprintf("select %v from %v where %v limit %v, %v",
			m.getColumns(),
			config.TABLE_MENUS,
			query,
			page.PageSize*(page.Page-1),
			page.PageSize,
		)
	} else {
		sqlText = "SELECT" + m.getColumns() + "FROM " + config.TABLE_MENUS + " WHERE " + query
	}

	resList, err = m.getRows(sqlText, args...)
	return
}

// 获取单行数据
func (m *MenusModel) First(query string, args ...interface{}) (result *entity.Menus, err error) {
	sqlText := "SELECT" + m.getColumns() + "FROM " + config.TABLE_MENUS + " WHERE " + query
	result, err = m.getRow(sqlText, args...)
	return
}

// 获取行数
// Get line count
func (m *MenusModel) Count(query string, args ...interface{}) (count int64, err error) {
	var sqlText string
	if query != "" {
		sqlText = "SELECT COUNT(*) FROM " + config.TABLE_MENUS + " WHERE " + query
	} else {
		sqlText = "SELECT COUNT(*) FROM " + config.TABLE_MENUS
	}
	err = m.DB.QueryRow(sqlText, args...).Scan(&count)
	if err != nil && err != sql.ErrNoRows {
		return
	}
	return
}
