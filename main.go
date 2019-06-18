package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path"
	"strings"
	"text/template"
)

var (
	fOutout   = flag.String("o", "", "output file")
	fTmpl     = flag.String("t", "", "template file name")
	fGofile   = flag.String("f", "", "Go file name")
	fDebugLes = flag.Bool("debug:lesphina", false, "debug to show lesphina data")
)

func main() {
	flag.Parse()

	if *fDebugLes {
		if tables, err := MetaToTables(*fGofile); err != nil {
			fmt.Println(err)
		} else {
			for _, t := range tables {
				fmt.Println("table:", t.Name)
				for _, col := range t.Columns {
					fmt.Printf("Name: %s, Type: %s, RawTag: %s\n", col.Name, col.Type, col.RawTag)
					fmt.Printf("-> modifiers: %q\n", col.Modifiers)
				}
			}
		}
		return
	}

	if strings.TrimSpace(*fTmpl) == "" {
		fmt.Println("blank name")
		return
	}

	tmpl, err := template.
		New(path.Base(*fTmpl)).
		Funcs(template.FuncMap{"join": strings.Join}).
		ParseFiles(*fTmpl)

	if err != nil {
		fmt.Println("failed to parse template:", err)
		return
	}

	tmplData, err := MetaToTables(*fGofile)
	if err != nil {
		fmt.Println("failed to map data")
		return
	}

	var output io.Writer

	if fOutout == nil || len(*fOutout) == 0 {
		output = os.Stdout
	} else {
		wf, err := os.Create(*fOutout)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer wf.Close()
		output = wf
	}

	for _, oneTableData := range tmplData {
		if err := tmpl.Execute(output, oneTableData); err != nil {
			fmt.Println(err)
			return
		}
	}
}