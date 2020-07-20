package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

//InsertDB retrieve DB data
func InsertDB(name string, amount string, table string, ordertime string, prgroup string) {
	db, _ := sql.Open("mysql", dbcon())
	stmt, err := db.Prepare("Insert PrData set Name=?, Amount=?, OrderTable=?, OrderTime=?, PrGroup=?")
	checkErr(err)
	defer stmt.Close()
	defer db.Close()
	stmt.Exec(name, amount, table, ordertime, prgroup)
}
