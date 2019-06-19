# gocf
Golang tool similar to Code-First (MS)

## Work with `go generate`

The Go source file:
```golang
package tests

//go:generate gocf -t <path/to/template/file> -f $GOFILE -o <output/file/name>.sql

// User ...
type User struct {
	ID   int    `gocf:"id,int,not null,auto_increment,primary key"`
	Name string `gocf:"name11111,varchar(50),not null"`
	Age  int    `gocf:"age,tinyint,not null,default 0"`
}

// Student ...
type Student struct {
	ID   int    `gocf:"id,int,not null,auto_increment,primary key"`
	Name string `gocf:"name,varchar(100),unique"`
	Age  int
}
```

Then you can just use `go generate` in this folder (see go-generate documents for details):
```
$ go generate
```

## Used in console

To see meta data read by [Lesphina](https://github.com/straightdave/lesphina):
```
$ ./gocf -f tests/a.go -debug:les
table: user
Name: name11111, Type: varchar(50), RawTag: name11111,varchar(50),not null
-> modifiers: ["not null"]
Name: age, Type: tinyint, RawTag: age,tinyint,not null,default 0
-> modifiers: ["not null" "default 0"]
table: student
Name: name, Type: varchar(100), RawTag: name,varchar(100),unique
-> modifiers: ["unique"]
```

To generate SQL creation script and print to Standard output:
```
$ ./gocf -f tests/a.go -t templates/create_table.tmpl
CREATE TABLE `user` (
    `name11111` varchar(50) not null,
    `age` tinyint not null default 0
) DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

CREATE TABLE `student` (
    `name` varchar(100) unique
) DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

```

> Use `-o <file name>` to output to a file.

## Your custom templates

You may want to use your own templates to generate.
`gocf` uses below data structure for the data in templates:
```golang
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
```

The data fields of the object `Table` are used as you can see in the `templates/create_table.tmpl` in this repo.
For example:
```
{{ .Name }} is Table#Name,
{{ .Columns }} is Table#Columns,

use:
{{ range $index, $col in .Columns }}
to iterate columns.

... and so on ...
```
