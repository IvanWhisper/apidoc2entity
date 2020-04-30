package generator

import (
	"apidoc2entity/util"
	"strings"
)

const (
	SUFFIX_GO = "go"
	SUFFIX_CS = "cs"

	B_TAB      = "    "
	B_SPACE    = " "
	B_LINE     = "\n"
	B_BODY_END = "}"

	TEMPLATE_GO_STRUCT_HEAD     = "package %s\n"
	TEMPLATE_GO_STRUCT_BODY_BEG = "// %s\ntype %s struct {\n"
	TEMPLATE_GO_STRUCT_FIELD    = "    // %s：%s\n    %s %s `json:\"%s\"`\n"
	TEMPLATE_GO_STRUCT_BODY_END = "}"
	TEMPLATE_GO_STRUCT_FOOT     = ""

	TEMPLATE_CS_CLASS_HEAD     = "using System;\nusing System.Collections.Generic;\nusing Newtonsoft.Json;\n\nnamespace %s\n{\n"
	TEMPLATE_CS_CLASS_BODY_BEG = "    /// <summary>\n    /// %s\n    /// </summary>\n    public class %s\n    {\n"
	TEMPLATE_CS_CLASS_FIELD    = "        /// <summary>\n        /// %s:%s\n        /// </summary>\n        [JsonProperty(\"%s\")]\n        public %s %s { get; set; }\n"
	TEMPLATE_CS_CLASS_BODY_END = "    }\n"
	TEMPLATE_CS_CLASS_FOOT     = "}"
)

var ConvertMapForGo map[string]string
var ConvertMapForCs map[string]string

func init() {
	ConvertMapForGo = make(map[string]string)
	ConvertMapForGo["string"] = "string"
	ConvertMapForGo["int"] = "int32"
	ConvertMapForGo["int8"] = "int8"
	ConvertMapForGo["int16"] = "int16"
	ConvertMapForGo["int32"] = "int32"
	ConvertMapForGo["int64"] = "int64"
	ConvertMapForGo["bool"] = "bool"
	ConvertMapForGo["date"] = "string"
	ConvertMapForGo["float"] = "float32"
	ConvertMapForGo["double"] = "float64"

	ConvertMapForCs = make(map[string]string)
	ConvertMapForCs["string"] = "string"
	ConvertMapForCs["int"] = "int32"
	ConvertMapForCs["int8"] = "int8"
	ConvertMapForCs["int16"] = "int16"
	ConvertMapForCs["int32"] = "int32"
	ConvertMapForCs["int64"] = "int64"
	ConvertMapForCs["bool"] = "bool"
	ConvertMapForCs["date"] = "string"
	ConvertMapForCs["float"] = "float32"
	ConvertMapForCs["double"] = "double"
}

func FieldFormat(s string) string {
	return util.Capitalize(s)
}

func TypeFormat(s string) string {
	words := strings.Split(s, "_")
	var buf strings.Builder
	for _, w := range words {
		buf.WriteString(util.Capitalize(w))
	}
	return buf.String()
}
