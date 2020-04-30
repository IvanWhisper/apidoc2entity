package cmd

import (
	"apidoc2entity/core"
	"apidoc2entity/util"
	"fmt"
	"github.com/urfave/cli"
	"os"
)

func Run() error {
	app := cli.NewApp()
	app.Name = "Generator"
	app.Usage = "简单对象生成器"
	app.Version = "1.0.0"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "source, s",
			Value: "./source.xlsx",
			Usage: "target file",
		},
		cli.StringFlag{
			Name:  "target, t",
			Value: "./output",
			Usage: "target dir",
		},
		cli.StringFlag{
			Name:  "lang, l",
			Value: "CS",
			Usage: "language",
		},
	}
	app.Action = func(c *cli.Context) error {
		source := c.String("source")
		target := c.String("target")
		lang := c.String("lang")
		sIsEx, err := util.Exists(source)
		// 检查源文件
		if err != nil {
			return err
		}
		if !sIsEx {
			panic("源文件不存在")
		}

		// 创建包文件夹
		e, er := util.Exists(target)
		if er != nil {
			fmt.Println(er)
		}
		if !e {
			err := os.MkdirAll(target, os.ModePerm)
			if err != nil {
				fmt.Println(err)
			}
		}
		core.Excel2Entities(source, target, lang)
		return nil
	}
	return app.Run(os.Args)
}
