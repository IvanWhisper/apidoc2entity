package generator

import (
	"apidoc2entity/model"
	"strings"
)

type Generator struct {
	Strategy IStrategy
}

func (gen *Generator) Generate(pkg *model.BuilderPackage) []model.Product {
	res := make([]model.Product, 0)
	// 构建头部
	head := gen.Strategy.FormatHead(pkg)
	foot := gen.Strategy.FormatFoot()
	for _, v := range pkg.BStructs {
		out := generateSingle(head, foot, v, gen.Strategy.FormatBodyBeg, gen.Strategy.FormatField, gen.Strategy.FormatBodyEnd)
		// 输出
		res = append(res, model.Product{PackageCode: pkg.PackageCode, FileCode: TypeFormat(v.StructCode), Data: out})
	}
	return res
}

func generateSingle(head, foot string, bstruct model.BuilderStruct,
	bodyfunc func(st *model.BuilderStruct) string,
	fieldfunc func(field *model.BuilderField) string,
	bodyendfunc func() string) string {
	var buf strings.Builder
	// 头部
	buf.WriteString(head)
	// 身体
	buf.WriteString(bodyfunc(&bstruct))
	// 字段
	for _, field := range bstruct.BFields {
		buf.WriteString(fieldfunc(&field))
	}
	// }
	buf.WriteString(bodyendfunc())
	// 脚部
	if len(foot) != 0 {
		buf.WriteString(foot)
	}
	return buf.String()
}
