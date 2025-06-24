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

	if isStructFunc {
		structFromFunc = funcParser.GetStructFromFunc(l)
		l = funcParser.RemoveStructFromFunc(l, true)
	}

	l = f.removeDoubleSpaces(l)
	l = strings.TrimSpace(l)

	// Проверяем, есть ли объявление с =
	if strings.Contains(l, "=") {
		declaration := f.getDeclarationVariable(l)
		l = f.removeDeclarationVariable(l)
		l = f.formateOneLineToMultiLine(l)
		l = strings.ReplaceAll(l, "{{DECLARATION}}", declaration)
	} else {
		// Если нет =, обрабатываем как чистую функцию
		l = f.formateOneLineToMultiLine(l)
	}

	if isStructFunc {
		l = strings.ReplaceAll(l, "{{FUNC_STRUCT}}", structFromFunc)
	}

	l = f.removeDoubleSpaces(l)

	return l
}

func (f *formatter) formateOneLineToMultiLine(l string) string {
	// Обрабатываем параметры функции
	l = strings.Replace(l, "(", "(\n\t", 1)
	l = strings.ReplaceAll(l, ", ", ",\n\t")

	// Обрабатываем возвращаемые значения
	if strings.Contains(l, ") (") {
		l = strings.Replace(l, ") (", ",\n) (\n\t", 1)
	} else if strings.Contains(l, ")(") {
		l = strings.Replace(l, ")(", ",\n) (\n\t", 1)
	}

	// Закрывающая скобка
	l = strings.Replace(l, ") {", ",\n) {", 1)

	return l
}

func (f *formatter) removeDoubleSpaces(
	l string,
) string {
	for strings.Contains(l, "  ") {
		l = strings.ReplaceAll(l, "  ", " ")
	}
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
