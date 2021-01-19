package generater

import (
	"bytes"
	"code-gen/assets"
	"code-gen/config"
	"code-gen/tools"
	"text/template"
)

type Generate struct {
	T tools.Tools
}

// 生成init
func (g *Generate) GormGenerateInit(path string, data string) (err error) {
	file := path + config.GormInitFileName

	tplByte, err := assets.Asset(config.GormInitTpl)
	if err != nil {
		return
	}
	tpl, err := template.New("init").Parse(string(tplByte))
	if err != nil {
		return
	}
	// 解析
	content := bytes.NewBuffer([]byte{})
	err = tpl.Execute(content, data)
	if err != nil {
		return
	}
	// 表信息写入文件
	err = tools.WriteFile(file, content.String())
	if err != nil {
		return
	}
	return
}

func (g *Generate) GormGenerateTools(path string, data string) (err error) {
	file := path + config.GormToolsFileName

	tplByte, err := assets.Asset(config.GormToolsTpl)
	if err != nil {
		return
	}
	tpl, err := template.New("tools").Parse(string(tplByte))
	if err != nil {
		return
	}
	// 解析
	content := bytes.NewBuffer([]byte{})
	err = tpl.Execute(content, data)
	if err != nil {
		return
	}
	// 表信息写入文件
	err = tools.WriteFile(file, content.String())
	if err != nil {
		return
	}
	return
}

// 根据模板生成文件
func (g *Generate) GenerateFiles(temp string, data interface{}, path string) (err error) {

	tplByte, err := assets.Asset(temp)
	if err != nil {
		return
	}
	tpl, err := template.New("tpl").Parse(string(tplByte))
	if err != nil {
		return
	}
	// 解析
	content := bytes.NewBuffer([]byte{})
	err = tpl.Execute(content, data)
	if err != nil {
		return
	}
	// 表信息写入文件
	err = tools.WriteFile(path, content.String())
	if err != nil {
		return
	}
	return
}