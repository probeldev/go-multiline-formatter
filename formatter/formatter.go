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

func (f *formatter) removeEmptyLines(s string) string {
	lines := strings.Split(s, "\n")
	var nonEmptyLines []string
	for _, line := range lines {
		if strings.TrimSpace(line) != "" {
			nonEmptyLines = append(nonEmptyLines, line)
		}
	}
	return strings.Join(nonEmptyLines, "\n")
}

func (f *formatter) formateOneLineToMultiLine(l string) string {
	// 1. Обрабатываем открывающую скобку параметров
	l = strings.ReplaceAll(l, "(", "(\n\t")

	// 2. Обрабатываем запятые между параметрами
	l = strings.ReplaceAll(l, ", ", ",\n\t")

	// 3. Специальная обработка для возвращаемых значений
	if strings.Contains(l, ") (") {
		// Случай с возвращаемыми значениями в скобках
		parts := strings.SplitN(l, ") (", 2)
		parts[0] = parts[0] + ",\n)"
		parts[1] = "(\n\t" + parts[1]
		l = parts[0] + " " + parts[1]
	} else if strings.Contains(l, ")") {
		// Общий случай для закрывающей скобки
		l = strings.ReplaceAll(l, ")", ",\n)")
	}

	// 4. Обрабатываем фигурную скобку
	l = strings.ReplaceAll(l, ") {", ",\n) {")

	l = f.removeEmptyLines(l)

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
