package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

//PrGetData retrieve DB data
func PrGetData() {
	db4Cache, _ := sql.Open("mysql", dbForCache())
	dbLocal, _ := sql.Open("mysql", dbcon())
	db4Cache.SetMaxOpenConns(2000)
	db4Cache.SetMaxIdleConns(500)
	var bgCtx = context.Background()
	var ctx2SecondTimeout, cancelFunc2SecondTimeout = context.WithTimeout(bgCtx, time.Second*2)
	defer cancelFunc2SecondTimeout()
	//ping the cloud DB
	if err := db4Cache.PingContext(ctx2SecondTimeout); err != nil {
		fmt.Println(err.Error())
		logging(err.Error())
	} else {
		stmt, _ := db4Cache.Prepare("Select ID,OrData from orderRec")
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("Connection error %s\r\n", err)
			}
		}()
		rows, err := stmt.Query()
		for rows.Next() {
			var ID string
			var Data string
			err = rows.Scan(&ID, &Data)
			checkErr(err)
			DataProcess2(Data, dbLocal, db4Cache, ID)
		}
		rows.Close()
		stmt.Close()
	}
	db4Cache.Close()
	dbLocal.Close()
}

//DataProcess2 for distributing data to related printer group
func DataProcess2(DBData string, conn *sql.DB, conn4cache *sql.DB, DelID string) {
	var layout []PrData
	Getdata := []byte(DBData)
	json.Unmarshal(Getdata, &layout)
	for i := range layout {

		InsertDB2(layout[i].Name, layout[i].Amount, layout[i].OrderTable, layout[i].OrderTime, layout[i].PrGroup, conn)
		DeleteDB(DelID, conn4cache)
	}
}

//InsertDB2 retrieve DB data and insert to Local DB
func InsertDB2(name string, amount string, table string, ordertime string, prgroup string, conn *sql.DB) {
	stmt, err := conn.Prepare("Insert PrData set Name=?, Amount=?, OrderTable=?, OrderTime=?, PrGroup=?")
	stmt2, err := conn.Prepare("Insert OrderHistory set Name=?, Amount=?, OrderTable=?, OrderTime=?")

	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Insert error %s\r\n", err)
		}
	}()
	checkErr(err)
	stmt.Exec(name, amount, table, ordertime, prgroup)
	stmt2.Exec(name, amount, table, ordertime)
	stmt.Close()
	stmt2.Close()
}

//DeleteDB for delete DB Cache
func DeleteDB(ID string, conn *sql.DB) {
	stmt, err := conn.Prepare("Delete From orderRec where ID = ?")
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Delete error %s\r\n", err)
		}
	}()

	checkErr(err)
	stmt.Exec(ID)
	stmt.Close()
}
