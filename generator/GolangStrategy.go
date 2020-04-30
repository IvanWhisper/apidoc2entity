package generator

import (
	"apidoc2entity/model"
	"fmt"
)

type GolangStrategy struct {
}

// 构建头部
func (gogen *GolangStrategy) FormatHead(pkg *model.BuilderPackage) string {
	return fmt.Sprintf(TEMPLATE_GO_STRUCT_HEAD, pkg.PackageCode)
}

// 构建身体
func (gogen *GolangStrategy) FormatBodyBeg(st *model.BuilderStruct) string {
	fcode := TypeFormat(st.StructCode)
	return fmt.Sprintf(TEMPLATE_GO_STRUCT_BODY_BEG, st.StructName, fcode)
}

// 构建身体
func (gogen *GolangStrategy) FormatBodyEnd() string {
	return fmt.Sprintf(TEMPLATE_GO_STRUCT_BODY_END)
}

// 构建字段
func (gogen *GolangStrategy) FormatField(field *model.BuilderField) string {
	falias := field.FieldAlias
	if len(falias) == 0 {
		falias = field.FieldCode
	}
	fcode := FieldFormat(field.FieldCode)
	var ftype string
	ftype = gogen.TypeConvert(field.FieldType, field.Fields)
	return fmt.Sprintf(TEMPLATE_GO_STRUCT_FIELD, field.FieldName, field.FieldAnno, falias, ftype, fcode)
}

// 构建脚部
func (gogen *GolangStrategy) FormatFoot() string {
	return TEMPLATE_GO_STRUCT_FOOT
}

func (gogen *GolangStrategy) Suffix() string {
	return SUFFIX_GO
}

func (gogen *GolangStrategy) TypeConvert(ftype string, isMore bool) string {
	res := ConvertMapForGo[ftype]
	isCus := res == ""
	if isCus {
		res = TypeFormat(ftype)
	}
	if isMore {
		res = fmt.Sprintf("[]%s", res)
	} else {
		if isCus {
			res = fmt.Sprintf("*%s", res)
		} else {
			res = fmt.Sprintf("%s", res)
		}
	}
	return res
}
