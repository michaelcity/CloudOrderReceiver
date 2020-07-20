package main

import (
	"encoding/json"

	_ "github.com/go-sql-driver/mysql"
)

//DataProcess for distributing data to related printer group
func DataProcess(DBData string) {
	var layout []PrData
	Getdata := []byte(DBData)
	json.Unmarshal(Getdata, &layout)
	for i := range layout {
		InsertDB(layout[i].Name, layout[i].Amount, layout[i].OrderTable, layout[i].OrderTime, layout[i].PrGroup)
	}

}
