package load

import (
	"apidoc2entity/model"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func Read(in string) []model.BuilderPackage {
	f, err := excelize.OpenFile(in)
	if err != nil {
		println(err.Error())
		return nil
	}
	result := make([]model.BuilderPackage, 0)
	for _, s := range f.GetSheetMap() {
		b := model.BuilderPackage{}
		b.PackageCode = s
		rows := f.GetRows(s)
		gp := make(map[string]*model.BuilderStruct)
		// 分组
		for i, row := range rows {
			if i == 0 {
			} else {
				var structCode string
				structCode = row[0]
				if gp[structCode] == nil {
					gp[structCode] = &model.BuilderStruct{StructCode: structCode}
				}
				if len(gp[structCode].StructName) == 0 {
					gp[structCode].StructName = structCode
				}
				fd := model.BuilderField{}
				fd.FieldCode = row[2]
				fd.FieldName = row[3]
				fd.FieldType = row[4]
				fd.FieldAnno = row[5]
				l, _ := strconv.Atoi(row[7])
				fd.FieldLength = int32(l)
				if len(row) > 8 && row[8] == "1" {
					fd.Fields = true
				}
				gp[structCode].BFields = append(gp[structCode].BFields, fd)
			}
		}
		// 装配
		for _, v := range gp {
			b.BStructs = append(b.BStructs, *v)
		}
		println(gp)
		result = append(result, b)
	}
	return result
}
