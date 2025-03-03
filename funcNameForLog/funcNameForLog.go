package funcnameforlog

import "github.com/probeldev/go-multiline-formatter/parser"

type funcNameForLog struct{}

func GetFuncNameForLog() *funcNameForLog {
	f := funcNameForLog{}

	return &f
}

func (f *funcNameForLog) AddFuncName(
	body string,
) string {
	response := body

	parserFunc := parser.GetFunctionParser()

	if !parserFunc.IsFuncDefinition(body) {
		return response
	}

	response += "\n"
	response += `fn :="`

	if parserFunc.IsStuctFunc(body) {
		structFromFunc := parserFunc.GetStructFromFunc(body)
		response += structFromFunc
		response += ":"
	}

	// TODO

	response += `"`

	return response
}
