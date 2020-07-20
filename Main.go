package main

import "time"

func main() {
	//PrGetData()
	//DataProcess(PrintData)
	//PrRoute()
	doEvery(3*time.Second, runPr)
}

func doEvery(d time.Duration, f func(time.Time)) {
	for x := range time.Tick(d) {
		f(x)
	}
}
func runPr(t time.Time) {
	//CheckPrinter("192.168.1.20")
	PrGetData()
	//PrRoute()
}
