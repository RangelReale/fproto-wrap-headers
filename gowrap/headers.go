package fproto_wrap_headers

import (
	"github.com/RangelReale/fproto-wrap/gowrap"
	"github.com/RangelReale/fproto/fdep"
)

// Converts between fproto_gowrap_headers.Headers and map[string][]string
type TypeConverterPlugin_Headers struct {
}

func (t *TypeConverterPlugin_Headers) GetTypeConverter(tp *fdep.DepType) fproto_gowrap.TypeConverter {
	if tp.FileDep.FilePath == "github.com/RangelReale/fproto-wrap-headers/headers.proto" &&
		tp.FileDep.ProtoFile.PackageName == "fproto_wrap_headers" &&
		tp.Name == "Headers" {
		return &TypeConverter_Headers{}
	}
	return nil
}

// Converter
type TypeConverter_Headers struct {
}

func (t *TypeConverter_Headers) TypeName(g *fproto_gowrap.Generator, tntype fproto_gowrap.TypeConverterTypeNameType) string {
	return "map[string][]string"
}

func (t *TypeConverter_Headers) IsPointer() bool {
	return true
}

func (t *TypeConverter_Headers) GenerateImport(g *fproto_gowrap.Generator, varSrc string, varDest string, varError string) (checkError bool, err error) {
	alias := g.Dep("github.com/RangelReale/fproto-wrap-headers/gowrap/ptypes", "ptypes_headers")

	g.P("if ", varSrc, " != nil {")
	g.In()

	g.P(varDest, " = ", alias, ".Import(", varSrc, ")")

	g.Out()
	g.P("}")

	return false, nil
}

func (t *TypeConverter_Headers) GenerateExport(g *fproto_gowrap.Generator, varSrc string, varDest string, varError string) (checkError bool, err error) {
	alias := g.Dep("github.com/RangelReale/fproto-wrap-headers/gowrap/ptypes", "ptypes_headers")

	g.P("if ", varSrc, " != nil {")
	g.In()

	g.P(varDest, " = ", alias, ".Export(", varSrc, ")")

	g.Out()
	g.P("}")

	return false, nil
}
