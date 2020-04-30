package core

import (
	"apidoc2entity/export"
	"apidoc2entity/generator"
	"apidoc2entity/load"
	"apidoc2entity/util"
	"fmt"
	"os"
	"path"
)

func Excel2Entities(in, out, lang string) error {
	fmt.Printf("源文件：%s\n输出目录：%s\n语言：%s", in, out, lang)
	// 抽取数据
	bs := load.Read(in)
	// 组合
	for _, b := range bs {
		gen := &generator.Generator{Strategy: ChooseStrategy(lang)}
		pd := gen.Generate(&b)
		for _, p := range pd {
			// 创建包文件夹
			dir := path.Join(out, p.PackageCode)
			e, er := util.Exists(dir)
			if er != nil {
				fmt.Println(er)
			}
			if !e {
				err := os.Mkdir(dir, os.ModePerm)
				if err != nil {
					fmt.Println(err)
				}
			}
			outfile := path.Join(dir, fmt.Sprintf("%s.%s", p.FileCode, gen.Strategy.Suffix()))
			// 写入
			export.Write2File(outfile, p.Data)
		}
	}
	return nil
}

// 选择语言策略
func ChooseStrategy(lang string) generator.IStrategy {
	switch lang {
	case "GO":
		return &generator.GolangStrategy{}
	default:
		return &generator.CSharpStrategy{}
	}
}
