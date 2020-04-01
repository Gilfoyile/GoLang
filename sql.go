package main

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)
type User struct{
	Name string `json:"name"`
}
func main(){
	fmt.Println("Sql Connecting Method")

// establising the connectio using to the database

	db,err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/newDB")

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("Successfully Connected to mysql database")

// inserting the data into mysql database

	insert, err := db.Query("INSERT INTO STUDENTS VALUES('RACHARANN','YADAV')")

	if err != nil {
		panic(err.Error())

	}
	defer insert.Close()
	fmt.Println("Data Successfully Inserted into the database")

// selecting the data from mysql  database

	results,err := db.Query("SELECT firstname FROM STUDENTS")
	if err != nil {
		panic(err.Error())
	}
	for results.Next(){
		var user User
		err = results.Scan(&user.Name)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(user.Name)
	}
}
