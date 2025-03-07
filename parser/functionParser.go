package parser

import (
	"errors"
	"strings"
)

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

func (f *functionParser) GetOnlyStructureNameFromFunc(
	l string,
) (
	string,
	error,
) {
	fn := "functionParser:GetOnlyStructureNameFromFunc"
	l = f.GetStructFromFunc(l)

	l = strings.ReplaceAll(l, "(", "")
	l = strings.ReplaceAll(l, ")", "")
	l = strings.ReplaceAll(l, "*", "")

	arr := strings.Split(l, " ")

	if len(arr) == 0 {
		return "", errors.New(fn + " error parser")
	}

	structName := ""

	if len(arr) == 1 {
		structName = arr[0]
	} else if len(arr) == 2 {
		structName = arr[1]
	} else {
		return "", errors.New(fn + " error parser")
	}

	structName = strings.Trim(structName, " ")

	return structName, nil
}

func (f *functionParser) GetFuncName(l string) string {
	funcName := ""

	l = f.RemoveStructFromFunc(l, false)
	l = strings.Replace(l, "func ", "", 1)
	l = strings.ReplaceAll(l, "{", "")
	arr := strings.Split(l, "(")
	funcName = arr[0]

	funcName = strings.Trim(funcName, " ")

	return funcName
}

func (f *functionParser) RemoveStructFromFunc(
	l string,
	isNeedTemplate bool,
) string {

	structFromFunc := f.GetStructFromFunc(l)
	if isNeedTemplate {
		l = strings.ReplaceAll(l, structFromFunc, "{{FUNC_STRUCT}}")
	} else {
		l = strings.ReplaceAll(l, structFromFunc, "")
	}

	return l
}
