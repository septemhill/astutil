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

	// Append `func fnDecl` into the astFile.
	astFile.AddDecls("func fnDecl() {}")

	// Append multiple decls: `const constDecl = 1` and `var varDecl = "varDecl"` into the astFile.
	astFile.AddDecls(`
		const constDecl = 1
		var varDecl = "varDecl"
	`)

	// Remove `existFunc`, `constDecl` (which just created above) and Base's methods from the astFile.
	// But the `Base.PointerReceiver` wouldn't be removed, we should use `*Base.PointerReceiver`
	astFile.RemoveDecls([]string{"existFunc", "constDecl", "Base.PointerReceiver", "Base.ConcreteReceiver"})

	// After operations above, astFile should be like:
	// package main
	//
	// var varDecl = "varDecl"
	//
	// type Base struct {
	// 		Name string
	// }
	//
	// func (b *Base) PointerReceiver() {
	// }
	//
	// func fnDecl() {
	// }
	println(astFile.String())
}
