package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//PrRoute for distributing data to related printer group
func PrRoute() {
	db, err := sql.Open("mysql", dbcon())
	stmt, err := db.Prepare("SELECT ID,Name,Amount,OrderTable,OrderTime,PrGroup from PrData")
	stmtDelete, err := db.Prepare("Delete From PrData where ID = ?")
	checkErr(err)
	rows, err := stmt.Query()
	for rows.Next() {
		var ID string
		var Name string
		var Amount string
		var OrderTable string
		var OrderTime string
		var PrGroup string
		statusSucc := true
		err = rows.Scan(&ID, &Name, &Amount, &OrderTable, &OrderTime, &PrGroup)
		//json.Unmarshal(temp, &layout)
		switch PrGroup {
		case "0":
			statusSucc = Printer1(Name, Amount, OrderTable, OrderTime)
			fmt.Println(statusSucc)
		case "1":
			statusSucc = Printer1(Name, Amount, OrderTable, OrderTime)
			fmt.Println(statusSucc)
		case "2":
			statusSucc = Printer1(Name, Amount, OrderTable, OrderTime)
			fmt.Println(statusSucc)
		case "3":
			statusSucc = Printer1(Name, Amount, OrderTable, OrderTime)
			fmt.Println(statusSucc)
		case "4":
			statusSucc = Printer1(Name, Amount, OrderTable, OrderTime)
			fmt.Println(statusSucc)
		case "5":
			statusSucc = Printer1(Name, Amount, OrderTable, OrderTime)
			fmt.Println(statusSucc)
		}
		if statusSucc {
			stmtDelete.Exec(ID)
		}
	}

}
