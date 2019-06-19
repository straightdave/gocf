package tests

//go:generate gocf -t ../templates/create_table.tmpl -f $GOFILE -o z_models.sql

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
