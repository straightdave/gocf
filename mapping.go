package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/straightdave/lesphina"
)

var (
	rMyTag = regexp.MustCompile(`gocf:"(.+?)"`)
)

// Table ...
type Table struct {
	Name    string
	Columns []*Column
}

// Column ...
type Column struct {
	Name      string
	Type      string
	Modifiers []string

	RawTag string
}

// MetaToTables ...
func MetaToTables(gofile string) ([]*Table, error) {
	var res []*Table

	les, err := lesphina.Read(gofile)
	if err != nil {
		return nil, err
	}

	if les.Meta.NumStruct == 0 {
		return res, nil
	}

	for _, s := range les.Meta.Structs {
		tbl := &Table{
			Name: strings.ToLower(s.Name),
		}

		for _, fld := range s.Fields {
			col := &Column{
				Name:   fld.Name,
				Type:   fld.RawType,
				RawTag: getTag(fld.RawTag),
			}

			if col.RawTag == "" {
				break
			}

			// tag format:
			// `(name),(type), (modifier-1), (modifier-2), ...`
			splits := strings.Split(col.RawTag, ",")

			// get specified name
			name := strings.TrimSpace(splits[0])
			if name != "" {
				col.Name = name
			}

			// get specified type
			if len(splits) >= 2 {
				typ := strings.TrimSpace(splits[1])
				if typ != "" {
					col.Type = typ
				}
			}

			// get modifiers
			if len(splits) >= 3 {
				for _, m := range splits[2:] {
					if strings.TrimSpace(m) != "" {
						col.Modifiers = append(col.Modifiers, m)
					}
				}
			}

			tbl.Columns = append(tbl.Columns, col)
		}

		res = append(res, tbl)
	}

	return res, nil
}

func print(tables []*Table) {
	for _, t := range tables {
		fmt.Println("table:", t.Name)
		for _, col := range t.Columns {
			fmt.Printf("Name: %s, Type: %s, RawTag: %s\n", col.Name, col.Type, col.RawTag)
			fmt.Printf("-> modifiers: %q\n", col.Modifiers)
		}
	}
}

func getTag(rawTag string) string {
	m := rMyTag.FindStringSubmatch(rawTag)
	if m != nil {
		if len(m) >= 2 {
			return m[1]
		}
	}
	return ""
}
