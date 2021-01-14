package assets

import (
	"bytes"
	"code-gen/config"
	"fmt"
	"text/template"

	"testing"
)

func TestAssetNames(t *testing.T) {
	tplByte, err := Asset(config.GormToolsTpl)
	if err != nil {
		fmt.Println(err)
		return
	}
	//tpl, err := template.New("tools").Parse(template.HTMLEscapeString(string(tplByte)))
	tpl, err := template.New("tools").Parse(string(tplByte))
	if err != nil {
		fmt.Println(err)
		return
	}
	// 解析
	content := bytes.NewBuffer([]byte{})
	err = tpl.Execute(content, config.PkgEntity)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(content.String())
}
