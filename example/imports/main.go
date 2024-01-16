package main

import (
	"log"

	"github.com/septemhill/goastutil"
)

func main() {
	astFile, err := goastutil.NewFile("./sample.go")
	if err != nil {
		log.Fatal("error: ", err)
	}

	// Append `errors` into the astFile
	astFile.AddImports([]string{"errors"})

	// Remove `fmt` from the astFile
	astFile.RemoveImports([]string{"fmt"})

	// After remove `fmt`, astFile should be like:
	// package main
	//
	// import "errors"
	//
	// func printInt(x int) {
	//         fmt.Println(x)
	// }
	println(astFile.String())
}
