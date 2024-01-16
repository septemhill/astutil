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

	// Prepend a stmt to the first function body
	astFile.FuncDecls()[0].Body().Stmts()[0].PrependStmt("fmt.Println(a, b)")

	// Prepend multiple stmts to the first function body (which is `fmt.Println(a, b)`)
	astFile.FuncDecls()[0].Body().Stmts()[0].PrependStmt(`
		const prepend = "Septem"
		var append = 123456
	`)

	// After operations above, astFile should be like:
	// package main
	//
	// func existFunc(a int, b int) int {
	// 		const prepend = "Septem"
	// 		append := 123456
	// 		fmt.Println(a, b)
	// 		return a + b
	// }
	println(astFile.String())
}
