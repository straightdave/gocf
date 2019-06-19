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
	fDebugLes = flag.Bool("debug:les", false, "debug printing lesphina data")
)

var (
	crlf = []byte("\r\n")
)

func main() {
	flag.Parse()

	if len(*fGofile) == 0 {
		fmt.Println("Go source file name (-f) was not provided.")
		return
	}

	// just display meta data read by lesphina
	if *fDebugLes {
		if tables, err := MetaToTables(*fGofile); err != nil {
			fmt.Println(err)
		} else {
			print(tables)
		}
		return
	}

	if len(*fTmpl) == 0 {
		fmt.Println("Template file name (-t) was not provided.")
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
		output.Write(crlf)
	}
}
