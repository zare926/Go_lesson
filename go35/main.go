package main

import (
	"database/sql"
	// "fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

type Person struct {
	Name string
	Age  int
}

func main() {
	Db, _ := sql.Open("sqlite3", "./example.sql")

	defer Db.Close()

	// CREATE　テーブル作成
	// cmd := `CREATE	TABLE IF NOT EXISTS persons (
	// 	name STRING,
	// 	age INT)`
	// _, err := Db.Exec(cmd)

	// INSERT カラムに値を追加
	// cmd := "INSERT INTO persons (name,age) VALUES (?,?)"
	// _, err := Db.Exec(cmd, "hanako", 19)

	// UPDATE 更新
	// cmd := "UPDATE persons SET age = ? WHERE name = ?"
	// _, err := Db.Exec(cmd, 35, "tarou")

	// INSERT データの取得
	// cmd := "SELECT * FROM persons WHERE age = ?"

	// // 第荷引数がwhereの値
	// row := Db.QueryRow(cmd, 35)
	// var p Person
	// err := row.Scan(&p.Name, &p.Age)
	// if err != nil {
	// 	if err == sql.ErrNoRows {
	// 		log.Println("No row")
	// 	} else {
	// 		log.Println(err)
	// 	}
	// }

	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// 全て取得するパターン
	// cmd := "SELECT * FROM persons"
	// rows, _ := Db.Query(cmd)
	// defer rows.Close()

	// var pp []Person
	// for rows.Next() {
	// 	var p Person
	// 	err := rows.Scan(&p.Name, &p.Age)
	// 	if err != nil {
	// 		log.Println(err)
	// 	}
	// 	pp = append(pp, p)
	// }

	// for _, p := range pp {
	// 	fmt.Println(p.Name, p.Age)
	// }

	cmd := "DELETE FROM persons WHERE name = ?"
	_, err := Db.Exec(cmd, "hanako")
	if err != nil {
		log.Fatalln(err)
	}

	// fmt.Println(p.Name, p.Age)
}
