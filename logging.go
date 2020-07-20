package main

import (
	"fmt"
	"os"
	"time"
)

func logging(msg string) {
	//write to file
	currentTime := time.Now()
	finalTime := currentTime.Format("2006-01-02")
	finalTimeHr := currentTime.Format("2006-01-02 15:04:05")
	fileLoc := "log/err-log-" + finalTime + ".txt"
	//write log
	f, err := os.OpenFile(fileLoc, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0660)
	os.Chmod(fileLoc, 0660)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	newLine :=
		"\r" + finalTimeHr + " Msg = " + msg + "\r-------------\r"
	_, err = fmt.Fprintln(f, newLine)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
}
