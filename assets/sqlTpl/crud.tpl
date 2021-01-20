package {{.PackageName}}

import (
	"database/sql"
	"fmt"
	"math"
	"time"
)

type {{.StructTableName}}Model struct {
	DB *sql.DB
	Tx *sql.Tx
}

// not transaction
func New{{.StructTableName}}(db ...*sql.DB) *{{.StructTableName}}Model {
	if len(db) > 0 {
		return &{{.StructTableName}}Model{
			DB: db[0],
		}
	}
	return &{{.StructTableName}}Model{
		DB: masterDB,
	}
}

// transaction object
func New{{.StructTableName}}Tx(tx *sql.Tx) *{{.StructTableName}}Model {
	return &{{.StructTableName}}Model{
		Tx: tx,
	}
}

// 获取所有的表字段
func (m *{{.StructTableName}}Model) getColumns() string {
	return " {{.AllFieldList}} "
}

// 获取多行数据.
func (m *{{.StructTableName}}Model) getRows(sqlTxt string, params ...interface{}) (rowsResult []*{{.PkgEntity}}{{.StructTableName}}, err error) {
	query, err := m.DB.Query(sqlTxt, params...)
	if err != nil {
		return
	}
	defer query.Close()
	for query.Next() {
		row := {{.PkgEntity}}{{.NullStructTableName}}{}
		err = query.Scan(
			{{range .NullFieldsInfo}}&row.{{.HumpName}},// {{.Comment}}
            {{end}}
		)
		if err != nil {
			return
		}
		rowsResult = append(rowsResult, &{{.PkgEntity}}{{.StructTableName}}{
			{{range .NullFieldsInfo}}{{if eq .GoType "float64"}}{{.HumpName}}:row.{{.HumpName}}.Float64,// {{.Comment}}
            {{else if eq .GoType "int64"}}{{.HumpName}}:row.{{.HumpName}}.Int64,// {{.Comment}}
            {{else if eq .GoType "time.Time"}}{{.HumpName}}:row.{{.HumpName}}.Time,// {{.Comment}}
            {{else if eq .GoType "int32"}}{{.HumpName}}:row.{{.HumpName}}.Int32,// {{.Comment}}
            {{else}}{{.HumpName}}:row.{{.HumpName}}.String,// {{.Comment}}
            {{end}}{{end}}
		})
	}
	return
}

// 获取单行数据
func (m *{{.StructTableName}}Model) getRow(sqlText string, params ...interface{}) (rowResult *{{.PkgEntity}}{{.StructTableName}}, err error) {
	query := m.DB.QueryRow(sqlText, params...)
	row := {{.PkgEntity}}{{.NullStructTableName}}{}
	err = query.Scan(
		{{range .NullFieldsInfo}}&row.{{.HumpName}},// {{.Comment}}
        {{end}}
	)
	if err != nil {
		return
	}
	rowResult = &{{.PkgEntity}}{{.StructTableName}}{
		{{range .NullFieldsInfo}}{{if eq .GoType "float64"}}{{.HumpName}}:row.{{.HumpName}}.Float64, // {{.Comment}}
        {{else if eq .GoType "int64"}}{{.HumpName}}:row.{{.HumpName}}.Int64,// {{.Comment}}
        {{else if eq .GoType "time.Time"}}{{.HumpName}}:row.{{.HumpName}}.Time,// {{.Comment}}
        {{else if eq .GoType "int32"}}{{.HumpName}}:row.{{.HumpName}}.Int32,// {{.Comment}}
        {{else}}{{.HumpName}}:row.{{.HumpName}}.String,// {{.Comment}}
        {{end}}{{end}}
	}
	return
}

//预编译
func (m *{{.StructTableName}}Model) Prepare(sqlTxt string, value ...interface{}) (result sql.Result, err error) {
	stmt, err := m.DB.Prepare(sqlTxt)
	if err != nil {
		return
	}
	defer stmt.Close()
	result, err = stmt.Exec(value...)

	return
}

//预编译
func (m *{{.StructTableName}}Model) PrepareTx(sqlTxt string, value ...interface{}) (result sql.Result, err error) {
	stmt, err := m.Tx.Prepare(sqlTxt)
	if err != nil {
		return
	}
	defer stmt.Close()
	result, err = stmt.Exec(value...)

	return
}

// 新增信息
func (m *{{.StructTableName}}Model) Create(value *{{.PkgEntity}}{{.StructTableName}}) (lastId int64, err error) {
	const sqlText = "INSERT INTO " + {{.PkgTable}}{{.UpperTableName}} + " ({{.InsertFieldList}}) VALUES ({{.InsertMark}})"

	params := make([]interface{}, 0)
	params = append(params,
		{{range .InsertInfo}}value.{{.HumpName}},// {{.Comment}}
        {{end}}
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
func (m *{{.StructTableName}}Model) CreateTx(value *{{.PkgEntity}}{{.StructTableName}}) (lastId int64, err error) {
	const sqlText = "INSERT INTO " + {{.PkgTable}}{{.UpperTableName}} + " ({{.InsertFieldList}}) VALUES ({{.InsertMark}})"
	params := make([]interface{}, 0)
	params = append(params,
		{{range .InsertInfo}}value.{{.HumpName}},// {{.Comment}}
        {{end}}
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
func (m *{{.StructTableName}}Model) Update(value *{{.PkgEntity}}{{.StructTableName}}) (b bool, err error) {
	const sqlText = "UPDATE " + {{.PkgTable}}{{.UpperTableName}} + " SET {{.UpdateFieldList}} WHERE {{.PrimaryKey}} = ?"
	params := make([]interface{}, 0)
    params = append(params,
    		{{range .UpdateInfo}}value.{{.HumpName}},// {{.Comment}}
            {{end}}
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
func (m *{{.StructTableName}}Model) UpdateTx(value *{{.PkgEntity}}{{.StructTableName}}) (b bool, err error) {
	const sqlText = "UPDATE " + {{.PkgTable}}{{.UpperTableName}} + " SET {{.UpdateFieldList}} WHERE {{.PrimaryKey}} = ?"
    params := make([]interface{}, 0)
    params = append(params,
        		{{range .UpdateInfo}}value.{{.HumpName}},// {{.Comment}}
                {{end}}
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

func (m *{{.StructTableName}}Model) Delete(id int64) (b bool, err error) {
	const sqlText = "UPDATE " + {{.PkgTable}}{{.UpperTableName}} + " SET `deleted_at`=? WHERE {{.PrimaryKey}} = ?"
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
func (m *{{.StructTableName}}Model) Find(page *Pagination, query string, args ...interface{}) (resList []*{{.PkgEntity}}{{.StructTableName}}, err error) {
	var sqlText string
	if page != nil && page.PageSize > 0 && page.Page > 0 { // 分页查询
		if page.Total, err = m.Count(query, args...); err != nil {
			return
		}
		page.TotalPage = int64(math.Ceil(float64(page.Total) / float64(page.PageSize)))
		//sqlText = "SELECT" + m.getColumns() + "FROM " + {{.PkgTable}}{{.UpperTableName}} + " WHERE " + query
		sqlText = fmt.Sprintf("select %v from %v where %v limit %v, %v",
			m.getColumns(),
			{{.PkgTable}}{{.UpperTableName}},
			query,
			page.PageSize*(page.Page-1),
			page.PageSize,
		)
	} else {
		sqlText = "SELECT" + m.getColumns() + "FROM " + {{.PkgTable}}{{.UpperTableName}} + " WHERE " + query
	}

	resList, err = m.getRows(sqlText, args...)
	return
}

// 获取单行数据
func (m *{{.StructTableName}}Model) First(query string, args ...interface{}) (result *{{.PkgEntity}}{{.StructTableName}}, err error) {
	sqlText := "SELECT" + m.getColumns() + "FROM " + {{.PkgTable}}{{.UpperTableName}} + " WHERE " + query
	result, err = m.getRow(sqlText, args...)
	return
}


// 获取行数
// Get line count
func (m *{{.StructTableName}}Model) Count(query string, args ...interface{}) (count int64, err error) {
	var sqlText string
	if query != "" {
		sqlText = "SELECT COUNT(*) FROM " + {{.PkgTable}}{{.UpperTableName}} + " WHERE " + query
	} else {
		sqlText = "SELECT COUNT(*) FROM " + {{.PkgTable}}{{.UpperTableName}}
	}
	err = m.DB.QueryRow(sqlText, args...).Scan(&count)
	if err != nil && err != sql.ErrNoRows {
		return
	}
	return
}
