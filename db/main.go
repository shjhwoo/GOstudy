package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// 기본 연결
func main() {
	db, err := sqlx.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/hackerank")
	if err != nil {
		fmt.Println(err, "...1")
		return
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err, "...2")
		return
	}

	fmt.Println("connection established")
}

/*
sqlx is a library which provides a set of extensions
on go's standard database/sql library.
The sqlx versions of sql.DB, sql.TX, sql.Stmt, et al.
all leave the underlying interfaces untouched,
so that their interfaces are a superset on the standard ones.
This makes it relatively painless to integrate existing codebases using database/sql with sqlx.

Major additional concepts are:

Marshal rows into structs (with embedded struct support), maps, and slices
Named parameter support including prepared statements
Get and Select to go quickly from query to struct/slice

sqlx는 database/sql의 상위 패키지로서 익스텐션을 제공

기본설치
비번 초기화: https://dev-whoan.xyz/48
*/
