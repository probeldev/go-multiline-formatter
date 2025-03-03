package parser

import "strings"

type functionParser struct{}

func GetFunctionParser() *functionParser {
	f := functionParser{}

	return &f
}

func (f *functionParser) IsFuncDefinition(l string) bool {
	return strings.Contains(l, "func")
}

func (f *functionParser) IsStuctFunc(
	l string,
) bool {
	l = strings.ReplaceAll(l, " ", "")

	return strings.HasPrefix(l, "func(")
}

func (f *functionParser) GetStructFromFunc(
	l string,
) string {
	arr := strings.Split(l, "(")
	l = arr[1]
	arr = strings.Split(l, ")")
	l = arr[0]

	return "(" + l + ")"
}

func (f *functionParser) GetFuncName(l string) string {
	funcName := ""

	return funcName
}
