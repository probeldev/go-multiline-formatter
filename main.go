package main

import (
	"bufio"
	"fmt"
	"go_multiline_formatter/formatter"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		input := scanner.Text()
		f := formatter.GetFormatter()
		output := f.OneLine2MultiLine(input)
		fmt.Println(output)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error: ", err)
	}
}
