package formatter

import (
	"strings"
)

type formatter struct{}

func GetFormatter() *formatter {
	f := formatter{}

	return &f
}

func (f *formatter) OneLine2MultiLine(
	l string,
) string {

	isStructFunc := f.isStuctFunc(l)
	structFromFunc := ""
	declaration := ""

	if isStructFunc {
		structFromFunc = f.getStructFromFunc(l)
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

func (f *formatter) isStuctFunc(
	l string,
) bool {
	l = strings.ReplaceAll(l, " ", "")

	return strings.HasPrefix(l, "func(")
}

func (f *formatter) getStructFromFunc(
	l string,
) string {
	arr := strings.Split(l, "(")
	l = arr[1]
	arr = strings.Split(l, ")")
	l = arr[0]

	return "(" + l + ")"
}

func (f *formatter) removeStructFromFunc(
	l string,
) string {
	structFromFunc := f.getStructFromFunc(l)
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
