package funcnameforlog

import (
	"log"

	"github.com/probeldev/go-multiline-formatter/parser"
)

type funcNameForLog struct{}

func GetFuncNameForLog() *funcNameForLog {
	f := funcNameForLog{}

	return &f
}

func (f *funcNameForLog) AddFuncName(
	body string,
) string {
	fn := "funcNameForLog:AddFuncName"

	response := body

	parserFunc := parser.GetFunctionParser()

	if !parserFunc.IsFuncDefinition(body) {
		return response
	}

	response += "\n\t"
	response += `fn := "`

	if parserFunc.IsStuctFunc(body) {
		structFromFunc, err := parserFunc.GetOnlyStructureNameFromFunc(body)
		if err != nil {
			log.Println(fn, err)
			return ""
		}
		response += structFromFunc
		response += ":"
	}

	response += parserFunc.GetFuncName(body)

	response += `"`
	response += "\n"

	return response
}
