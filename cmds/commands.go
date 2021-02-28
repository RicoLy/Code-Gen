package cmds

import (
	"bufio"
	"code-gen/config"
	"code-gen/entity"
	"code-gen/generater"
	"code-gen/logic"
	"code-gen/tools"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type commands struct {
	l *logic.Logic
	g *generater.Generate
}

func NewCommands(logic *logic.Logic, gen *generater.Generate) *commands {
	return &commands{
		l: logic,
		g: gen,
	}
}

//映射相应的命令
func (c *commands) Handlers() map[string]func(args []string) int {
	return map[string]func(args []string) int{
		"0":     c.CustomDir,
		"1":     c.MarkDown,
		"2":     c.GenerateEntry,
		"21":    c.GormGenerateEntry,
		"3":     c.GenerateCURD,
		"31":    c.SQLGenerateCURD,
		"32":    c.SqlGenerateEntityAndCURD,
		"4":     c.CustomFormat,
		"5":     c.ShowTableList,
		"7":     c.Clean,
		"clear": c.Clean,
		"c":     c.Clean,
		"8":     c.Help,
		"h":     c.Help,
		"help":  c.Help,
		"ll":    c.Help,
		"ls":    c.Help,
		"quit":  c.Quit,
		"q":     c.Quit,
		"exit":  c.Quit,
	}
}

//生成数据库表的markdown文档
func (c *commands) MarkDown(args []string) int {
	fmt.Println("Preparing to generate the markdown document...")
	//检查目录是否存在
	tools.CreateDir(c.l.Path)
	err := c.l.CreateMarkdown()
	if err != nil {
		log.Println("MarkDown>>", err)
	}
	return 0
}

//help list
func (c *commands) Help(args []string) int {
	for _, row := range config.CmdHelp {
		s := fmt.Sprintf("%s %s\n", "NO:"+row.No, row.Msg)
		fmt.Print(s)
	}
	return 0
}

//生成golang表对应的结构实体
func (c *commands) GenerateEntry(args []string) int {
	fmt.Print("Do you need to set the format of the structure?(Yes|No)>")
	line, _, _ := bufio.NewReader(os.Stdin).ReadLine()
	switch strings.ToLower(string(line)) {
	case "yes", "y":
		config.Formats = c._setFormat()
	}
	err := c.l.CreateEntity(config.Formats)
	if err != nil {
		log.Println("GenerateEntry>>", err.Error())
	}
	go tools.Gofmt(tools.GetExeRootDir()) //格式化代码
	return 0
}

//生成golang表对应的结构实体
func (c *commands) SQLGenerateEntry(args []string) int {
	fmt.Print("Do you need to set the format of the structure?(Yes|No)>")
	line, _, _ := bufio.NewReader(os.Stdin).ReadLine()
	switch strings.ToLower(string(line)) {
	case "yes", "y":
		config.Formats = c._setFormat()
	}
	err := c.l.CreateEntity(config.Formats)
	if err != nil {
		log.Println("GenerateEntry>>", err.Error())
	}
	go tools.Gofmt(tools.GetExeRootDir()) //格式化代码
	return 0
}

//生成golang Gorm表对应的结构实体
func (c *commands) GormGenerateEntry(args []string) int {
	fmt.Print("Do you need to set the format of the structure?(Yes|No)>")
	line, _, _ := bufio.NewReader(os.Stdin).ReadLine()
	switch strings.ToLower(string(line)) {
	case "yes", "y":
		config.Formats = c._setFormat()
	}
	err := c.l.GormCreateEntity(config.Formats)
	if err != nil {
		log.Println("GenerateEntry>>", err.Error())
	}
	//go tools.Gofmt(tools.GetExeRootDir()) //格式化代码

	go tools.Gofmt(tools.CreateDir(c.l.Path+config.GormDirEntity) + config.DS) //格式化代码
	return 0
}

//还可以自定义结构体解析实体,如json,gorm,xml
func (c *commands) CustomFormat(args []string) int {
	config.Formats = c._setFormat()
	return 0
}

//生成golang操作mysql的CRUD增删改查语句
func (c *commands) GenerateCURD(args []string) int {
	err := c.l.CreateCURD(config.Formats)
	if err != nil {
		log.Println("GenerateCURD>>", err.Error())
	}
	go tools.Gofmt(tools.GetExeRootDir())
	return 0
}

//生成golang操作mysql的CRUD增删改查语句
func (c *commands) SQLGenerateCURD(args []string) int {
	fmt.Print("Do you need to set the format of the structure?(Yes|No)>")
	line, _, _ := bufio.NewReader(os.Stdin).ReadLine()
	switch strings.ToLower(string(line)) {
	case "yes", "y":
		config.Formats = c._setFormat()
	}

	err := c.l.SQLCreateCURD(config.Formats)
	if err != nil {
		log.Println("GenerateCURD>>", err.Error())
	}
	go tools.Gofmt(tools.GetExeRootDir())
	return 0
}

//自定义生成目录
func (c *commands) CustomDir(args []string) int {
	fmt.Print("Please set the build directory>")
	line, _, _ := bufio.NewReader(os.Stdin).ReadLine()
	if string(line) != "" {
		path, err := c.l.T.GenerateDir(string(line))
		if err == nil {
			c.l.Path = path
			fmt.Println("Directory success:", path)
		} else {
			log.Println("Set directory failed>>", err)
		}
	}
	return 0
}

//显示所有的表名
func (c *commands) ShowTableList(args []string) int {
	if len(c.l.DB.Tables) == 0 {
		fmt.Println("Whoops, Nothing at all!!!")
		return 0
	}
	c._showTableList(c.l.DB.Tables)
	fmt.Print("Select the table sequence number you need?(By default all, comma separated,all represents all)>")
	line, _, _ := bufio.NewReader(os.Stdin).ReadLine()
	if !strings.EqualFold(string(line), "") {
		c.l.DB.DoTables = c._filterTables(string(line), c.l.DB.Tables)
	}
	return 0
}

//清屏
func (c *commands) Clean(args []string) int {
	tools.Clean()
	return 0
}

//退出
func (c *commands) Quit(args []string) int {
	return 1
}

//生成数据库表的Sql struct Crud
func (c *commands) SqlGenerateEntityAndCURD(args []string) int {
	fmt.Println("Testing .............")
	fmt.Print("Do you need to set the format of the structure?(Yes|No)>")
	line, _, _ := bufio.NewReader(os.Stdin).ReadLine()
	switch strings.ToLower(string(line)) {
	case "yes", "y":
		config.Formats = c._setFormat()
	}

	//获取实列渲染数据
	tableInfoList, err := c.l.GetTableInfoList()
	if err != nil {
		fmt.Println(err)
		return 0
	}

	//生成entity
	path := tools.CreateDir(c.l.Path + config.GODIR_MODELS + config.DS + config.GODIR_Entity)
	for _, en := range tableInfoList.TableInfos { // 生成结构体
		if err := c.g.GenerateFiles(config.SQLTPL_ENTITY, en, path+config.DS+en.TableName+".go"); err != nil {
			fmt.Println(err)
			return 0
		}
	}
	go tools.Gofmt(path) // 代码格式化
	path = tools.CreateDir(c.l.Path + config.GODIR_MODELS)
	for _, sqlInfo := range tableInfoList.SQLInfo { // 生成CRUD
		if err := c.g.GenerateFiles(config.SQLTPL_CURD, sqlInfo, path+config.DS+sqlInfo.TableName+".go"); err != nil {
			fmt.Println(err)
			return 0
		}
	}
	// 生成Init
	if err := c.g.GenerateFiles(config.SQLTPL_INIT, config.PkgDbModels, path+config.DS+config.GoFile_Init); err != nil {
		fmt.Println(err)
		return 0
	}
	go tools.Gofmt(path) // 代码格式化

	// 生成config/tables文件
	path = tools.CreateDir(c.l.Path + config.GODIR_MODELS + config.DS + config.GODIR_Config)

	if err := c.g.GenerateFiles(config.SQLTPL_CONF, tableInfoList.TableNames, path+config.DS+config.GoFile_TableList); err != nil {
		fmt.Println(err)
		return 0
	}
	go tools.Gofmt(path) // 代码格式化

	return 0
}

//过滤表名
func (c *commands) _filterTables(ids string, tables []entity.TableNameAndComment) []entity.TableNameAndComment {
	lst := strings.Split(ids, ",")
	result := make([]entity.TableNameAndComment, 0)
	if strings.ToLower(ids) == "all" {
		return tables
	}
	for _, id := range lst {
		id = strings.TrimSpace(id)
		for _, t := range tables {
			if strconv.Itoa(t.Index) == id || id == t.Name {
				result = append(result, t)
			}
		}
	}
	return result
}

//显示所有名视图
func (c *commands) _showTableList(NameAndComment []entity.TableNameAndComment) {
	for idx, table := range NameAndComment {
		idx++
		info := fmt.Sprintf("%s:%s", strconv.Itoa(idx), table.Name)
		if table.Comment != "" {
			info += fmt.Sprintf("(%s)", table.Comment)
		}
		fmt.Println(info)
	}
	fmt.Println("Total " + strconv.Itoa(len(NameAndComment)) + " tables\n")
}

//set struct format
func (c *commands) _setFormat() []string {
	fmt.Print("Set the mapping name of the structure, separated by a comma (example :json,gorm)>")
	input, _, _ := bufio.NewReader(os.Stdin).ReadLine()
	if string(input) != "" {
		formatList := tools.CheckCharDoSpecialArr(string(input), ',', `[\w\,\-]+`)
		if len(formatList) > 0 {
			fmt.Printf("Set format success: %v\n", formatList)
			return formatList
		}
	}
	fmt.Println("Set failed")
	return nil
}
