package main

import (
	"go/format"
	"log"
	"syscall/js"
)

func formatGoCode() js.Func {
	formatFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 1 {
			log.Println("invalid number of arguments")
			return ""
		}

		input := args[0].String()

		result, err := format.Source([]byte(input))
		if err != nil {
			log.Println("failed to format input: ", err)
			return ""
		}

		return string(result)
	})
	return formatFunc
}

func main() {
	js.Global().Set("formatGoCode", formatGoCode())
	<-make(chan bool)
}
