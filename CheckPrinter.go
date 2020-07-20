package main

import (
	"fmt"
	"net"
	"time"
)

//CheckPrinter for checking Printer online
func CheckPrinter(ip string) {

	timeout := time.Duration(60 * time.Second)
	//t1 := time.Now()
	_, err := net.DialTimeout("tcp", ip+":443", timeout)
	//fmt.Println("waist time :", time.Now().Sub(t1))

	if err != nil {
		fmt.Println("Site unreachable, error: ", err)
		return
	} else {
		fmt.Println("OK ")
	}
}
