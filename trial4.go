package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DbConnection *sql.DB

type Person struct {
	Name string
	Age  int
}

func main() {
	DbConnection, _ := sql.Open("sqlite3", "./example.sql")
	defer DbConnection.Close()

	cmd := `CREATE TABLE IF NOT EXISTS person(
		name STRING,
		age INT)`
	_, err := DbConnection.Exec(cmd)
	if err != nil {
		log.Fatal(err)
	}

	cmd = "INSERT INTO person (name, age) VALUES (?, ?)"
	_, err = DbConnection.Exec(cmd, "Jirorian", 19) //SQLコマンド（cmd）のVALUESの？に値を入れている。
	//?を使うのはsqlインジェクション対策
	if err != nil {
		log.Fatal(err)
	}

	// cmd = "UPDATE person SET age = ? where name = ?"
	// _, err = DbConnection.Exec(cmd, 25, "Taro") //SQLコマンド（cmd）のVALUESの？に値を入れている。
	// //?を使うのはsqlインジェクション対策
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// cmd = "SELECT * FROM person"
	// rows, _ := DbConnection.Query(cmd)//rowsにはDBから取ってきた複数件のデータが格納される
	// defer rows.Close()
	// var pp []Person //Person構造体のスライスを定義
	// for rows.Next() {
	// 	var p Person
	// 	err := rows.Scan(&p.Name, &p.Age) //Scan()が構造体の値(p)にDB から取ってきたデータ（rows）を格納してくれる
	// 	if err != nil {
	// 		log.Println(err)
	// 	}
	// 	pp = append(pp, p) //Person型のGoの構造体のスライスに
	// 	//pの値（DBからとってきたpersonの値のうちの一件）をappendした。
	// }
	// //以下のようなエラーの書き方もある。
	// // err := rows.Err()
	// // if err != nil{
	// // 	log.Fatal(err)
	// // }

	// for _, p := range pp { //ppとはDBから取ってきたデータがGoの構造体（p）に変換されたデータの集まり（スライス）
	// 	fmt.Println(p.Name, p.Age) //ppの個数分、その中の一個一個の要素であるpの値を出力している
	// }

	cmd = "SELECT * FROM person WHERE age = ?"
	row := DbConnection.QueryRow(cmd, 19)
	var p Person
	err = row.Scan(&p.Name, &p.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No Row")
		} else {
			log.Println(err)
		}
	}
	fmt.Println(p.Name, p.Age)

}
