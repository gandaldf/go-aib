package main

import (
	"flag"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	fileName := flag.String("fn", "version.go", "Go source file name")
	patchName := flag.String("pn", "versionPatch", "Patch var/const name")
	timestampName := flag.String("tn", "versionTimestamp", "Timestamp var/const name")
	timestampFormat := flag.String("tf", "0601021504", "Timestamp output format")

	flag.Parse()

	setFile := token.NewFileSet()
	astFile, err := parser.ParseFile(setFile, *fileName, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	var foundPatch bool
	var foundTimestamp bool

	ast.Inspect(astFile, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.Ident:
			switch x.Name {
			case *patchName:
				foundPatch = true
			case *timestampName:
				foundTimestamp = true
			}
		case *ast.BasicLit:
			if foundPatch == true {
				foundPatch = false

				patch, err := strconv.Atoi(strings.Replace(x.Value, "\"", "", -1))
				if err != nil {
					panic(err)
				}

				x.Value = "\"" + strconv.Itoa(patch+1) + "\""
			}

			if foundTimestamp == true {
				foundTimestamp = false

				x.Value = "\"" + time.Now().UTC().Format(*timestampFormat) + "\""
			}
		}

		return true
	})

	f, err := os.Create(*fileName)
	defer f.Close()

	if err := printer.Fprint(f, setFile, astFile); err != nil {
		panic(err)
	}
}
