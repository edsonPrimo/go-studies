package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	stringConnection := "golang:golang@/devbook?charset=utf8&parseTime=true&loc=Local"
	db, err := sql.Open("mysql", stringConnection)
	if err != nil {
		fmt.Print("Dentro do SQL.open")
		log.Fatal(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		fmt.Println("Dentro do ping")
		log.Fatal(err)
	}

	fmt.Println("Conexão está aberta!")

	linhas, err := db.Query("select * from usuarios")
	if err != nil {
		log.Fatal(err)
	}
	defer linhas.Close()

	fmt.Println(linhas)
}
