package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// 기본 연결
func main() {
	//연결 객체가 아닌 DB를 대표하는 객체를 생성한다
	username := "root"
	password := "1234"
	host := "127.0.0.1"
	port := "3307"
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/", username, password, host, port) //"root:1234@tcp(127.0.0.1:3307)/"

	//내부적으로 커넥션 풀을 유지한다
	db, err := sqlx.Open("mysql", dataSourceName)
	if err != nil {
		fmt.Println(err, "...db 객체 생성")
		return
	}

	//커넥션 풀!
	//쿼리문 준비, 쿼리 실행은 연결을 필요로 함
	//db 객체는 이런 커넥션 풀을 관리함
	/*
		주의) 커넥션 풀을 쓴다는 건 두개의 sql문이 하나의 테이블에 각자 작업을 할수도 있다는 소리다.
		따라서 예상과 다른 결과가 나올 수 있음. 이 경우, LOCK TABLES 구문을 붙여 방지할 수 있다.
		기본적으로는 풀에 있는 커넥션 갯수는 무한정이다.
		풀에 커넥션이 다 차면 새로운 커넥션을 만든다는 소리다.
		이를 제한하기 위해 다음의 메서드를 사용한다
		sql 패키지에서 제공하는 기능이기 때문에 사용 가능하다.
		앞서 말했듯이 sqlx는 sql의 익스텐션이므로 기존 패키지 기능을 상속받아 쓸 수 있는 것임
		http://go-database-sql.org/connection-pool.html

		커넥션은 빠르게 재사용된다. (다쓰고 유휴 상태로 안남기려고 반납한단 의미)
		maxIdleConns를 높게 잡을 경우 이런 과정이 줄어들고, 유연한 재사용이 가능함
	*/
	db.SetMaxIdleConns(0) //놀고 있는 커넥션이 없어야 한다는 뜻이다. 최대 연결 개수 이상일 수 없다.	//0은 커넥션 타임아웃 시 사용가능한 방법임
	db.SetMaxOpenConns(1) //맺을 수 있는 커넥션의 최대 개수
	//db.SetConnMaxIdleTime(1 * time.Second) //풀에 있는 커넥션의 최대 유휴 시간. 시간 다되면 닫힘. 0으로 하면 안 닫힘
	db.SetConnMaxLifetime(5 * time.Second)
	//연결이 긴 시간 동안 맺어져 있는데 이를 오래 쓸 경우 네트웍 문제 발생가능. 안 쓰는/파기된 연결은 닫는다.
	//5초 이내에는 해당 연결 재사용 가능하다. 0으로 하면 커넥션 안 닫힘

	/*
		커넥션 사용 완료 후에는 항상 이를 반납하여야 한다.
		Row 하나하나씩 모두 Scan()한다
		Next()함수를 통해서 모든 Row를 읽어야 한다. (필요한 만큼의 row만 셀렉하는 것도 방법이다)
		Close() 호출
		트랜잭션의 경우 Commit(), Rollback() 연산을 통해 반납한다
		무시할 경우, 커넥션은 가비지 컬렉션에 들어가기 전까지 유휴 상태가 된다.
		그래서 자꾸 새로 만들게 되는 불상사가 생긴다
		Rows.Close()를 적절히 사용해서 이를 방지하자
	*/

	//연결 수립
	err = db.Ping()
	if err != nil {
		fmt.Println(err, "...연결 오류")
		return
	}

	fmt.Println("con1: 연결이 확인되었습니다")

	//db 생성 동시에 연결도 할 수 있다
	// db2, err = sqlx.Connect("mysql", dataSourceName)
	// if err != nil {
	// 	fmt.Println(err, "...2")
	// 	return
	// }
	// fmt.Println("con2: connection established")
	//연결 실패시 패닉
	// db3 := sqlx.MustConnect("mysql", dataSourceName)
	// fmt.Println("con3: connect OK")
	// createDBQuery3 := `Create database if not exists sqlx`
	// _, err = db3.Exec(createDBQuery3)
	// if err != nil {
	// 	fmt.Println(err, "...3")
	// 	return
	// }

	//db 생성 및 선택
	createDBQuery := `Create database if not exists mydb`
	_, err = db.Exec(createDBQuery)
	if err != nil {
		fmt.Println(err, "...3")
		return
	}

	selectDBQuery := `Use mydb`
	_, err = db.Exec(selectDBQuery)
	if err != nil {
		fmt.Println(err, "...4")
		return
	}

	//쿼리 작성하기
	type PlaceTable struct {
		Country  string `db:"country"`
		City     string `db:"city"`
		Telcode  int    `db:"telcode"`
		Database string `db:"-"`
	}

	sql1 := fmt.Sprintf(`Create Table if not exists %s (
		country text,
		city text NULL,
		telcode integer
	)`, "mydb.place")

	//커넥션 풀에서 커넥션 받아와서 지정된 쿼리를 실행함
	//result가 나오기 전에 쿼리 실행 완료 후 커넥션이 반납된다.
	_, err = db.Exec(sql1)
	if err != nil {
		fmt.Println(err, "...5")
		return
	}

	cityState := `INSERT INTO mydb.place (country, telcode) VALUES (?, ?)`
	countryCity := `INSERT INTO mydb.place (country, city, telcode) VALUES (?, ?, ?)`
	db.MustExec(cityState, "Hong Kong", 852)
	db.MustExec(cityState, "Singapore", 65)
	db.MustExec(countryCity, "South Africa", "Johannesburg", 27)

	/*
		bind variables , ? 중요함.
		항상 이런 방식으로 값을 전달해주는 것이 중요하다
		장점: sql 인젝션 공격을 막을 수 있다(sql이 악의적인 쿼리 검증을 안하기 때문이다)
		하드코딩과 변수에 비유할 수 있다.
		(반복되는 코드를 함수화하는 것에 비유가능)
		적절히 사용할 경우 성능 향상을 꾀할 수 있음
		하지만 자료가 검색조건에 따라 양이 천지차이가 나는 경우엔,
		 각 상황에 맞게끔 쿼리를 하드코딩하는게 맞을 수 있다.
		https://pinokio0702.tistory.com/54
		1. 구문오류체크
		2. 공유영역에서 해당구문검색(동일 구문 있으면 6번으로 건너뜀 => 더 빠름)
		3. 권한체크
		4. 실행계획수립
		5. 실행계획 공유영역에 저장
		6. 쿼리실행
		파라미터와 함께 글자 그대로 전달된다
		주의사항: SQL 구문의 구조를 바꾸기 위해 써서는 안된다
	*/

	//잘못된 예: 컬럼, 테이블 명에 대해서는 사용 불가하다
	//https://stackoverflow.com/questions/610056/is-it-possible-to-refer-to-column-names-via-bind-variables-in-oracle
	//실헁계획 수립이 불가하기 떄문에 써서는 안된다. 직접 하드코딩해주는 것이 필수
	//또한 이렇게 할 경우 SQL 인젝션 공격에도 취약해진다
	// // doesn't work
	// db.Query("SELECT * FROM ?", "mytable")
	// // also doesn't work
	// db.Query("SELECT ?, ? FROM people", "name", "location")
	//쿼리: 이 메서드로도 쿼리 실행이 가능하다
	//rows는 실제 데이터가 아니고 데이터베이스 커서임
	rows, err := db.Query("SELECT * FROM  mydb.place")
	if err != nil {
		fmt.Println(err) //커넥션 풀에서 안 좋은 커넥션을 가지고왔거나, sql 문법이 오류가있거나..
		return
	}

	//next를 통해서 큰 결과값에도 메모리 사용량을 줄일 수 있음(이말이 이해가 안됨)
	for rows.Next() {
		var country string
		var city sql.NullString // note that city can be NULL, so we use the NullString type
		var telcode int
		err = rows.Scan(&country, &city, &telcode)
		//reflect 메서드를 써서 조회한 결과값을 go의 타입으로 변경함
		//필요 이상으로 rows를 읽어온 경우 반드시 커넥션을 반납해줘야 하기 때문에,
		//꼭 필요한 만큼만 읽어오는 것이 중요
		//아무튼 모든 row를 다 Next로 작업하게 되면 커넥션을 반납하게 된다
	}
	// check the error from rows
	err = rows.Err()

	rows2, err2 := db.Query("SELECT * FROM mydb.place")
	if err2 != nil {
		fmt.Println(err2)
		return
	}

	fmt.Println(rows2, "커넥션 반납 안할거야.")

	//일부러 시간 끌어서 사용중인 커넥션으로 강제로 만들어줌
	//time.Sleep(10 * time.Second)

	rows2.Close() //반납한다면?

	//이 상태에서 다시 쿼리 실행을 시도한다면?
	rows3, err4 := db.Query("SELECT * FROM mydb.place")
	if err4 != nil {
		fmt.Println(err4)
		return
	}

	time.Sleep(10 * time.Second) //유휴 상태
	fmt.Println(db, "객체")

	_, err41 := db.Query("SELECT * FROM mydb.place")
	if err4 != nil {
		fmt.Println(err41, "ㅋㅋㅋ")
		return
	}

	fmt.Println("두번쨰 연결 생성했습니다: ", rows3)

	/////스캐닝 방법
	type Place struct {
		Country       string
		City          sql.NullString
		TelephoneCode int `db:"telcode"` //매칭되는 칼럼명을 명시해줘야 한다
	}

	//이렇게 읽어오는 방법도 있다.
	rows5, err := db.Queryx("SELECT * FROM mydb.place")
	for rows2.Next() {
		var p Place
		err = rows5.StructScan(&p)
	}

	//Query랑 달리 에러를 리턴하지 않는다. 대신에  쿼리 수행 시 발생한 에러는
	//다음 scan 과정에서 리턴된다
	row := db.QueryRow("SELECT * FROM mydb.place WHERE telcode=?", 852)
	var telcode int
	err = row.Scan(&telcode)

	//비슷한 역할
	var p Place
	err = db.QueryRowx("SELECT city, telcode FROM mydb.place LIMIT 1").StructScan(&p)
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
