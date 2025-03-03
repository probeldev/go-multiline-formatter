package formatter

import (
	"strings"

	"github.com/probeldev/go-multiline-formatter/parser"
)

type formatter struct{}

func GetFormatter() *formatter {
	f := formatter{}

	return &f
}

func (f *formatter) OneLine2MultiLine(
	l string,
) string {

	funcParser := parser.GetFunctionParser()

	isStructFunc := funcParser.IsStuctFunc(l)
	structFromFunc := ""
	declaration := ""

	if isStructFunc {
		structFromFunc = funcParser.GetStructFromFunc(l)
		l = f.removeStructFromFunc(l)
	}

	declaration = f.getDeclarationVariable(l)

	l = f.removeDeclarationVariable(l)

	l = f.formateOneLineToMultiLine(l)

	if isStructFunc {
		l = strings.ReplaceAll(l, "{{FUNC_STRUCT}}", structFromFunc)
	}
	l = strings.ReplaceAll(l, "{{DECLARATION}}", declaration)

	l = f.removeDoubleSpaces(l)

	return l
}

func (f *formatter) formateOneLineToMultiLine(
	l string,
) string {
	l = strings.ReplaceAll(l, "(", "(\n\t")
	l = strings.ReplaceAll(l, ",", ",\n\t")
	l = strings.ReplaceAll(l, ")", ",\n)")
	l = strings.ReplaceAll(l, "\t ", "\t")

	return l
}

func (f *formatter) removeStructFromFunc(
	l string,
) string {
	funcParser := parser.GetFunctionParser()

	structFromFunc := funcParser.GetStructFromFunc(l)
	l = strings.ReplaceAll(l, structFromFunc, "{{FUNC_STRUCT}}")

	return l
}

func (f *formatter) removeDoubleSpaces(
	l string,
) string {
	l = strings.ReplaceAll(l, "  ", " ")
	return l
}

func (f *formatter) getDeclarationVariable(
	l string,
) string {
	arr := strings.Split(l, "=")
	return arr[0]
}

func (f *formatter) removeDeclarationVariable(
	l string,
) string {
	arr := strings.Split(l, "=")

	responseArr := []string{}
	for i, r := range arr {
		if i == 0 {
			responseArr = append(responseArr, "{{DECLARATION}}")
			continue
		} else {
			responseArr = append(responseArr, r)
		}
	}

	return strings.Join(responseArr, "=")

}
