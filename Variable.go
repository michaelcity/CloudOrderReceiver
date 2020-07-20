package main

//PrData Data
type PrData struct {
	Name       string
	Amount     string
	OrderTable string
	OrderTime  string
	PrGroup    string
}

//PrDataArray Data
type PrDataArray struct {
	DBdata []PrData
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
