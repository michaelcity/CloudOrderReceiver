package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"
)

//Printer1 for print group 1
func Printer1(Name string, Amount string, Table string, OrderTime string) (status bool) {

	//DataBody := []byte(`[{"Name":"三文魚","Amount": "1","Table":"1","OrderTime":"10:30"}]`)
	//var PrData1 []PrData
	//json.Unmarshal(DataBody, &PrData1)
	PrintingSucc := false
	var PrintingErrMsg string
	url := "http://192.168.1.20/cgi-bin/epos/service.cgi?devid=printer&timeout=10000"
	data := "<s:Envelope xmlns:s=\"http://schemas.xmlsoap.org/soap/envelope/\">\"" +
		"<s:Body>" +
		"<epos-print xmlns=\"http://www.epson-pos.com/schemas/2011/03/epos-print\">\"" +
		"<text lang=\"zh-tw\" smooth=\"true\" width=\"3\" height=\"3\" >" + Name + "&#10;</text>'" +
		"<text lang=\"zh-tw\" smooth=\"true\">數量: " + Amount + "&#10;</text>'" +
		"<text lang=\"zh-tw\" smooth=\"true\">台號: " + Table + "&#10;</text>'" +
		"<text lang=\"zh-tw\" smooth=\"true\">時間: " + OrderTime + "&#10;</text>'" +
		"<cut/>" +
		"</epos-print>" +
		"</s:Body>" +
		"</s:Envelope>"
	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(data)))
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Printer1 no response %s\r\n", err)
		}
	}()
	if err != nil {
		fmt.Println(err)
		logging("Printer1 no response")
		PrintingSucc = false
	} else {
		req.Header.Add("Content-Type", "text/xml; charset=utf-8")
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Println("HTTP Response Status:", resp.Status, http.StatusText(resp.StatusCode))
		//return "HTTP Response Status: 200 OK OK"

		if resp.Status == "200" {
			PrintingSucc = true
		} else {
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			bodyString := string(body)
			if err != nil {
				fmt.Println(err)
			}
			//Error handling
			EPTRCOVEROPEN, _ := regexp.Compile("EPTR_COVER_OPEN")
			EPTRRECEMPTY, _ := regexp.Compile("EPTR_REC_EMPTY")
			EXBADPORT, _ := regexp.Compile("EX_BADPORT")
			switch {
			case EPTRCOVEROPEN.MatchString(bodyString):
				fmt.Println("EPTRCOVEROPEN")
				PrintingSucc = false
				PrintingErrMsg = "Printer Cover Opened"
			case EPTRRECEMPTY.MatchString(bodyString):
				PrintingSucc = false
				PrintingErrMsg = "Empty Paper"
				fmt.Println(PrintingErrMsg)
			case EXBADPORT.MatchString(bodyString):
				PrintingSucc = false
				fmt.Println("Connect Connect Printer 1")
			}
			logging(PrintingErrMsg)
		}
	}
	fmt.Println(PrintingSucc)
	return PrintingSucc
}
