package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func manageErr(e error) {
	if e != nil {
		fmt.Println("에러 발생 : ", e)
		panic(e)
	}
}

var mysqlUrl string

func main() {
	//<내 계정이름>:<비번>@<프로토콜 방식>(<주소>)/<데이터베이스 이름>
	db, err := sql.Open("mysql", mysqlUrl)
	manageErr(err)
	defer db.Close()

	getData(db)
}

func getMySqlUrl() {
	file, err := os.Open("config.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		mysqlUrl = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

//READ
func getData(db *sql.DB) {

	var id int
	var title string
	var _type string
	var grade float32
	var actor string

	row, err := db.Query("select * from movie")
	manageErr(err)
	defer row.Close()
	for row.Next() {
		err := row.Scan(&id, &title, &_type, &grade, &actor)
		manageErr(err)

		fmt.Println("------------------------------")
		fmt.Println("id : ", id)
		fmt.Println("제목 : ", title)
		fmt.Println("장르 : ", _type)
		fmt.Println("평점 : ", grade)
		fmt.Println("배우 : ", actor)

	}
}

//CREATE
func addData(db *sql.DB) {
	stmt, err := db.Prepare("insert Movie set title=?, type=?, grade=?, actor=?")
	manageErr(err)

	var _title string
	var _type string
	var _grade float32
	var _actor string

	fmt.Println("영화 데이터 추가를 시작합니다.(띄어쓰기를 피해주세요.)")

	fmt.Print("영화 제목 : ")
	fmt.Scanln(&_title)
	fmt.Print("영화 장르 : ")
	fmt.Scanln(&_type)
	fmt.Print("영화 평점 : ")
	fmt.Scanln(&_grade)
	fmt.Print("출연 배우 : ")
	fmt.Scanln(&_actor)

	input, err := stmt.Exec(&_title, &_type, &_grade, &_actor)
	manageErr(err)

	id, err := input.LastInsertId()
	manageErr(err)

	fmt.Println("성공적으로 영화가 추가되었습니다. => ", id)
}

//DELETE
func deleteData(db *sql.DB) {
	stmt, err := db.Prepare("delete from Movie where id=?")
	manageErr(err)

	var id int

	fmt.Print("삭제하고자 하는 영화 ID : ")
	fmt.Scanln(&id)

	input, err := stmt.Exec(&id)
	manageErr(err)

	result, _ := input.RowsAffected()

	fmt.Println("삭제된 영화 개수 : ", result, "개")
}

//UPDATE
func updateData(db *sql.DB) {
	stmt, err := db.Prepare("update Movie set title=?, type=?, grade=?, actor=? where id=?")
	manageErr(err)

	var _id int
	var _title string
	var _type string
	var _grade float32
	var _actor string

	fmt.Println("영화 데이터 수정을 시작합니다.(띄어쓰기를 피해주세요.)")

	//TODO : 이거 일일히 이렇게 하는거 귀찮으니까 GORM사용해서 간단하게 하자...
	fmt.Print("영화 id : ")
	fmt.Scanln(&_id)
	fmt.Print("영화 제목 : ")
	fmt.Scanln(&_title)
	fmt.Print("영화 장르 : ")
	fmt.Scanln(&_type)
	fmt.Print("영화 평점 : ")
	fmt.Scanln(&_grade)
	fmt.Print("출연 배우 : ")
	fmt.Scanln(&_actor)

	input, err := stmt.Exec(&_title, &_type, &_grade, &_actor, &_id)
	manageErr(err)

	affect, err := input.RowsAffected()
	manageErr(err)

	fmt.Println(affect)

}
