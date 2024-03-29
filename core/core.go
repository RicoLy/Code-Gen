package core

import (
	"bufio"
	"code-gen/cmds"
	"code-gen/config"
	"code-gen/db"
	"code-gen/entity"
	"code-gen/generater"
	"code-gen/logic"
	"code-gen/tools"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/howeyc/gopass"

	"github.com/urfave/cli/v2"
)

var (
	Conn   *sql.DB         //连接对象
	DbConn entity.DBConfig //db config

)

//命令行实现
func Start(app *cli.App, stop chan bool) {
	usage(app)
	go release(stop)
	DbConn.MaxIdleConn = 5
	DbConn.MaxOpenConn = 10

	app.Action = func(c *cli.Context) error {
		if c.Bool("debug") {
			log.SetFlags(log.Lshortfile | log.LstdFlags)
		}
		DbConn.Host = c.String("h")    // 数据库地址
		DbConn.Name = c.String("u")    // 数据库用户名称
		DbConn.Port = c.Int("P")       // 端口号
		DbConn.Pass = c.String("p")    // 密码
		DbConn.Charset = c.String("c") // 编码格式
		if c.NumFlags() > 0 {
			dbName := c.String("d") // 数据库名称
			if dbName == "" {
				return cli.Exit("database is null, please use -d params", 9)
			}
			DbConn.DBName = dbName
			if DbConn.Pass == "" { // Use an asterisk instead of the password you entered
				pass, err := gopass.GetPasswdPrompt("Enter password:", false, os.Stdin, os.Stdout)
				if err != nil {
					return err
				}
				DbConn.Pass = string(pass)
				tools.Clean()
			}
			if err := Commands(stop); err != nil {
				return cli.NewExitError(err, 1)
			}
		}
		return nil
	}
	// cli start run
	if err := app.Run(os.Args); err != nil {
		log.Fatal("RUN>>", err)
	}
}

//释放资源
func release(stop chan bool) {
	<-stop
	_ = Conn.Close()
}

//构建命令使用说明
func usage(app *cli.App) {
	app.Name = config.ProjectName //项目名称
	app.Authors = []*cli.Author{{
		Name:  config.Author,
		Email: config.AuthorEmail,
	}}
	app.Version = config.Version                                                         //版本号
	app.Copyright = "@Copyright " + config.Copyright                                     //版权保护
	app.Usage = "Quickly generate CURD and documentation for operating MYSQL.etc" //说明
	app.Commands = []*cli.Command{
		{
			Name:    "help",
			Aliases: []string{"h", "?"},
			Usage:   "show help",
			Action: func(c *cli.Context) error {
				_ = cli.ShowAppHelp(c)
				return nil
			},
		},
		{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "print the version",
			Action: func(c *cli.Context) error {
				cli.ShowVersion(c)
				return nil
			},
		},
	}
	app.HideVersion = true
	app.HideHelp = true
	app.Flags = []cli.Flag{
		&cli.StringFlag{Name: "h", Value: "127.0.0.1", Usage: "Database address"},
		&cli.IntFlag{Name: "P", Value: 3306, Usage: "port number"},
		&cli.StringFlag{Name: "u", Value: "root", Usage: "database username", Required: true},
		&cli.StringFlag{Name: "p", Value: "", Usage: "database password", Required: true, DefaultText: ""},
		&cli.StringFlag{Name: "c", Value: "utf8mb4", Usage: "database format"},
		&cli.StringFlag{Name: "d", Usage: "database name"},
		&cli.BoolFlag{Name: "debug", Usage: "debug", Value: false},
	}
}

// 实现命令功能
func Commands(stop chan bool) error {
	var err error
	Conn, err = db.InitDB(DbConn)
	if Conn == nil || err != nil {
		return errors.New("database connect failed>>" + err.Error())
	}
	//初使工作
	DbModel := db.NewDB()
	DbModel.Using(Conn)
	DbModel.DBName = DbConn.DBName

	lgc := &logic.Logic{
		DB:   DbModel,
		Path: tools.GetExeRootDir() + config.DefaultSavePath + config.DS, //默认当前命令所在目录
	}
	err = lgc.DB.GetTableNameAndComment()
	if err != nil {
		return err
	}

	commands := cmds.NewCommands(lgc, &generater.Generate{})
	commands.Help(nil)
	handlers := commands.Handlers()

	br := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("-> # ")
		line, _, _ := br.ReadLine()
		if len(line) == 0 {
			continue
		}
		tokens := strings.Split(string(line), " ")
		if handler, ok := handlers[strings.ToLower(tokens[0])]; ok {
			ret := handler(tokens)
			if ret != 0 {
				break
			}
		} else {
			fmt.Println("Unknown command>>", tokens[0])
		}
	}
	stop <- true
	return nil
}
