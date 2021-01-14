package generater

import (
	"bytes"
	"code-gen/assets"
	"code-gen/config"
	"code-gen/entity"
	"code-gen/tools"
	"strings"
	"text/template"
)

type Generate struct {
	T tools.Tools
}

// 创建结构实体
func (g *Generate) MyGenerateDBEntity(req *entity.EntityReq) (err error) {

	// 声明表结构变量
	TableData := new(entity.TableInfo)
	TableData.PackageName = req.EntityPkg
	TableData.Table = g.T.Capitalize(req.TableName)
	TableData.TableName = req.TableName
	TableData.NullTable = TableData.Table + config.DbNullPrefix
	TableData.TableComment = tools.AddToComment(req.TableComment, "")
	TableData.TableCommentNull = tools.AddToComment(req.TableComment, " Null Entity")
	// 判断表结构是否加载过
	//if l.T.CheckFileContainsChar(req.EntityPath, "type "+TableData.Table+" struct") == true {
	//	log.Println(req.EntityPath + "It already exists. Please delete it and regenerate it")
	//	return
	//}
	// 加载模板文件
	tplByte, err := assets.Asset(config.GormEntityTpl)
	if err != nil {
		return
	}
	tpl, err := template.New("entity").Parse(string(tplByte))
	if err != nil {
		return
	}
	// 装载表字段信息
	for _, val := range req.TableDesc {
		TableData.Fields = append(TableData.Fields, &entity.FieldsInfo{
			Name:         g.T.Capitalize(val.ColumnName),
			Type:         val.GolangType,
			NullType:     val.MysqlNullType,
			DbOriField:   val.ColumnName,
			FormatFields: tools.FormatField(val.ColumnName, req.FormatList),
			Remark:       tools.AddToComment(val.ColumnComment, ""),
		})
	}
	content := bytes.NewBuffer([]byte{})
	_ = tpl.Execute(content, TableData)
	// 表信息写入文件
	con := strings.Replace(content.String(), "&#34;", `"`, -1)
	err = tools.WriteFile(req.EntityPath, con)
	if err != nil {
		return
	}
	return
}
