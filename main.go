package main

import (
	"fmt"

	"github.com/spf13/pflag"
        "github.com/taisan11/nanx/file"

)

func main() {
	// Define flags
	f := pflag.BoolP("foo", "f", false, "option foo")
	b := pflag.BoolP("bar", "b", false, "option bar")
	text := pflag.StringP("text", "t", "", "text to print")
	pflag.Parse()

	fmt.Println("foo = ", *f)
	fmt.Println("bar = ", *b)
	SaveFile("file.txt", *text)
}
