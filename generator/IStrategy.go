package generator

import "apidoc2entity/model"

type IStrategy interface {
	FormatHead(pkg *model.BuilderPackage) string
	FormatBodyBeg(st *model.BuilderStruct) string
	FormatBodyEnd() string
	FormatField(field *model.BuilderField) string
	FormatFoot() string
	Suffix() string
	TypeConvert(ftype string, isMore bool) string
}
