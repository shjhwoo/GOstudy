package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// 기본 연결
func main() {
	//연결 객체가 아닌 DB를 대표하는 객체를 생성.
	//내부적으로 커넥션 풀을 유지한다
	db, err := sqlx.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/hackerank")
	if err != nil {
		fmt.Println(err, "...1")
		return
	}

	//연결 수립
	err = db.Ping()
	if err != nil {
		fmt.Println(err, "...2")
		return
	}

	fmt.Println("connection established")

	//db 생성 동시에 연결도 할 수 있다
	db, err = sqlx.Connect("mysql", "root:1234@tcp(127.0.0.1:3306)/hackerank")
	if err != nil {
		fmt.Println(err, "...2")
		return
	}

	//연결 실패시 패닉
	db = sqlx.MustConnect("mysql", "root:1234@tcp(127.0.0.1:3306)/hackerank")
	fmt.Println("connection still OK")

	//커넥션 풀!

	//쿼리 작성하기
	// schema := `Create Table place (
	// 	country text,
	// 	city text NULL,
	// 	telcode integer
	// )`

	//커넥션 풀에서 커넥션 받아와서 지정된 쿼리를 실행함
	//result가 나오기 전에 쿼리 실행 완료 후 커넥션이 반납된다.
	// result, err := db.Exec(schema)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// id, err := result.LastInsertId()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(id)

	cityState := `INSERT INTO place (country, telcode) VALUES (?, ?)`
	countryCity := `INSERT INTO place (country, city, telcode) VALUES (?, ?, ?)`
	db.MustExec(cityState, "Hong Kong", 852)
	db.MustExec(cityState, "Singapore", 65)
	db.MustExec(countryCity, "South Africa", "Johannesburg", 27)

	/*
		bind variables , ? 중요함.
		항상 이런 방식으로 값을 전달해주는 것이 중요하다
		장점: sql 인젝션 공격을 막을 수 있다(sql이 악의적인 쿼리 검증을 안하기 때문이다)
		하드코딩과 변수에 비유할 수 있다.
		(반복되는 코드를 함수화하는 것에 비유가능)
		https://pinokio0702.tistory.com/54
		파라미터와 함께 글자 그대로 전달된다
		드라이버가 이런 인터페이스를 제공하지 않는 경우 쿼리 실행 전에 쿼리문이 다 완성된다.
		따라서 bindvar는 데이터베이스 특이적으로 작동함
		무슨말이냐..
		MySQL uses the ? variant shown above
		PostgreSQL uses an enumerated $1, $2, etc bindvar syntax
		SQLite accepts both ? and $1 syntax
		Oracle uses a :name syntax

		주의사항: SQL 구문의 구조를 바꾸기 위해 써서는 안된다
	*/

	//잘못된 예: 컬럼, 테이블 명에 대해서는 사용 불가하다
	//https://stackoverflow.com/questions/610056/is-it-possible-to-refer-to-column-names-via-bind-variables-in-oracle
	//실헁계획 수립이 불가하기 떄문에 써서는 안된다. 직접 하드코딩해주는 것이 필수
	//또한 이렇게 할 경우 SQL 인젝션 공격에도 취약해진다
	// doesn't work
	db.Query("SELECT * FROM ?", "mytable")

	// also doesn't work
	db.Query("SELECT ?, ? FROM people", "name", "location")

	//쿼리: 이 메서드로도 쿼리 실행이 가능하다
	// fetch all places from the db
	//rows는 실제 데이터가 아니고 데이터베이스 커서임
	rows, err := db.Query("SELECT * FROM place")
	if err != nil {
		fmt.Println(err) //커넥션 풀에서 안 좋은 커넥션을 가지고왔거나, sql 문법이 오류가있거나..
		return
	}

	// iterate over each row
	//next를 통해서 큰 결과값에도 메모리 사용량을 줄일 수 있음(이말이 이해가 안됨)
	fmt.Println(rows.Next(), "확인")
	for rows.Next() {
		var country string
		// note that city can be NULL, so we use the NullString type
		var city sql.NullString
		var telcode int
		err = rows.Scan(&country, &city, &telcode) //reflect 메서드를 써서 조회한 결과값을 go의 타입으로 변경함
		fmt.Println(err)
		fmt.Println(country, city, telcode)
		//필요 이상으로 rows를 읽어온 경우 반드시 커넥션을 반납해줘야 하기 때문에,
		//꼭 필요한 만큼만 읽어오는 것이 중요
		//아무튼 모든 row를 다 Next로 작업하게 되면 커넥션을 반납하게 된다
	}
	// check the error from rows
	err = rows.Err()

	type Place struct {
		Country       string
		City          sql.NullString
		TelephoneCode int `db:"telcode"` //매칭되는 칼럼명을 명시해줘야 한다
	}

	//이렇게 읽어오는 방법도 있다.
	rows2, err := db.Queryx("SELECT * FROM place")
	for rows2.Next() {
		var p Place
		err = rows2.StructScan(&p)
	}

	//Query랑 달리 에러를 리턴하지 않는다. 대신에  쿼리 수행 시 발생한 에러는
	//다음 scan 과정에서 리턴된다
	row := db.QueryRow("SELECT * FROM place WHERE telcode=?", 852)
	var telcode int
	err = row.Scan(&telcode)

	//비슷한 역할
	var p Place
	err = db.QueryRowx("SELECT city, telcode FROM place LIMIT 1").StructScan(&p)
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
