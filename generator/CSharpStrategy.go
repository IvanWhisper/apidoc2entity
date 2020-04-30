package generator

import (
	"apidoc2entity/model"
	"fmt"
)

type CSharpStrategy struct {
}

// 构建头部
func (cs *CSharpStrategy) FormatHead(pkg *model.BuilderPackage) string {
	return fmt.Sprintf(TEMPLATE_CS_CLASS_HEAD, pkg.PackageCode)
}

// 构建身体
func (cs *CSharpStrategy) FormatBodyBeg(st *model.BuilderStruct) string {
	fcode := TypeFormat(st.StructCode)
	return fmt.Sprintf(TEMPLATE_CS_CLASS_BODY_BEG, st.StructName, fcode)
}

// 构建身体
func (cs *CSharpStrategy) FormatBodyEnd() string {
	return fmt.Sprintf(TEMPLATE_CS_CLASS_BODY_END)
}

// 构建字段
func (cs *CSharpStrategy) FormatField(field *model.BuilderField) string {
	falias := field.FieldAlias
	if len(falias) == 0 {
		falias = field.FieldCode
	}
	fcode := FieldFormat(field.FieldCode)
	var ftype string
	ftype = cs.TypeConvert(field.FieldType, field.Fields)
	return fmt.Sprintf(TEMPLATE_CS_CLASS_FIELD, field.FieldName, field.FieldAnno, falias, ftype, fcode)
}

// 构建脚部
func (cs *CSharpStrategy) FormatFoot() string {
	return TEMPLATE_CS_CLASS_FOOT
}

func (cs *CSharpStrategy) Suffix() string {
	return SUFFIX_CS
}

func (cs *CSharpStrategy) TypeConvert(ftype string, isMore bool) string {
	res := ConvertMapForCs[ftype]
	if res == "" {
		res = TypeFormat(ftype)
	}
	if isMore {
		res = fmt.Sprintf("IEnumerable<%s>", res)
	}
	return res
}
