package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var Db *sql.DB

var err error

type Person struct {
	Name string
	Age  int
}

func main() {
	Db, _ = sql.Open("postgres", "user=test_user dbname=testdb password=password sslmode=disable")
	if err != nil {
		log.Panicln(err)
	}
	defer Db.Close()

	// //C
	// cmd := "INSERT INTO persons (name, age) VALUES ($1, $2)"
	// _, err := Db.Exec(cmd, "hanako", 20)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	//R
	// cmd := "SELECT * FROM persons where age = $1"
	// //1レコードのみ取得
	// row := Db.QueryRow(cmd, 20)
	// var pp Person
	// err = row.Scan(&pp.Name, &pp.Age)
	// if err != nil {
	// 	// データがなかったら
	// 	if err == sql.ErrNoRows {
	// 		log.Println("No row")
	// 	} else {
	// 		log.Println(err)
	// 	}
	// }
	// fmt.Println(pp.Name, pp.Age)

	// cmd = "SELECT * FROM persons"
	// // 条件に合うものを全て取得
	// rows, _ := Db.Query(cmd)
	// defer rows.Close()
	// // structを作成
	// var ppp []Person
	// // 取得したデータをスライスに追加　 for rows.Next()
	// for rows.Next() {
	// 	var p Person

	// 	err := rows.Scan(&p.Name, &p.Age)

	// 	if err != nil {
	// 		log.Println(err)
	// 	}

	// 	ppp = append(ppp, p)
	// }

	// //まとめてエラーハンドリング
	// err = rows.Err()
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// for _, p := range ppp {
	// 	fmt.Println(p.Name, p.Age)
	// }

	// //U
	// cmd := "UPDATE persons SET age = $1 WHERE name = $2"
	// _, err := Db.Exec(cmd, 25, "hanako")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	//D
	cmd := "DELETE FROM persons WHERE name = $1"
	_, err := Db.Exec(cmd, "hanako")
	if err != nil {
		log.Fatalln(err)
	}
}
