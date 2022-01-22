package main

import (
	"fmt"
	"go/format"
	"syscall/js"
)

func formatGoCode(this js.Value, args []js.Value) interface{} {
	if len(args) != 1 {
		fmt.Println("invalid number of arguments")
		return ""
	}

	input := args[0].String()

	result, err := format.Source([]byte(input))
	if err != nil {
		fmt.Println("failed to format input: ", err)
		return ""
	}

	return js.ValueOf(string(result))
}

func main() {
	wait := make(chan struct{}, 0)
	js.Global().Set("formatGoCode", js.FuncOf(formatGoCode))
	<-wait
}
