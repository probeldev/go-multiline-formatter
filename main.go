package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/probeldev/go-multiline-formatter/formatter"
	funcnameforlog "github.com/probeldev/go-multiline-formatter/funcNameForLog"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	multiline := flag.Bool("multiline", false, "Преоброазование одной строки в несколько")
	logfunc := flag.Bool("logfunc", false, "Добавление строковой переменной с названием структуры и функции")

	flag.Parse()

	if scanner.Scan() {
		input := scanner.Text()

		if *multiline {
			f := formatter.GetFormatter()
			output := f.OneLine2MultiLine(input)
			fmt.Println(output)
		} else if *logfunc {
			f := funcnameforlog.GetFuncNameForLog()
			output := f.AddFuncName(input)
			fmt.Println(output)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error: ", err)
	}
}
