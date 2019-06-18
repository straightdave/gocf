package tests

//go:generate gocf

// User ...
type User struct {
	Name string `gocf:"name11111,varchar(50),not null"`
	Age  int    `gocf:"age,tinyint,not null,default 0"`
}

// Student ...
type Student struct {
	Name string `gocf:"name,varchar(100),unique"`
	Age  int
}
