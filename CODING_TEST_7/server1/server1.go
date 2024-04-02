package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type XMLData struct {
	ResultCode    string `xml:"ResultCode"`
	ResultDesc    string `xml:"ResultDesc"`
	TransID       string `xml:"TransID"`
	BillRefNumber string `xml:"BillRefNumber"`
	UtilityName   string `xml:"UtilityName"`
	CustomerName  string `xml:"CustomerName"`
	Amount        string `xml:"Amount"`
}

type Data struct {
	C2BPaymentQueryResult XMLData `json:"C2BPaymentQueryResult"`
	Password              string  `json:"password"`
}

func sendDataToServer2(xmlString, password string) {
	jsonData := Data{
		C2BPaymentQueryResult: XMLData{
			ResultCode:    "2",
			ResultDesc:    "Failed",
			TransID:       "10111",
			BillRefNumber: "12233",
			UtilityName:   "sddd",
			CustomerName:  "wee",
			Amount:        "30",
		},
		Password: password,
	}

	payload, err := json.Marshal(jsonData)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	client := http.Client{
		Timeout: 10 * time.Second,
	}

	url := "http://localhost:8080/receive"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Unable to send request to Server 2:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Println("Data sent successfully to Server 2.")
	} else {
		fmt.Println("Server 2 failed to acknowledge. Status:", resp.Status)
	}
}

func main() {
	xmlString := `
	<soapenv:Envelope
	xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/"
	xmlns:c2b="http://cps.huawei.com/cpsinterface/c2bpayment">
	<soapenv:Header/>
	<soapenv:Body>
	<c2b:C2BPaymentQueryResult>
	<ResultCode>2</ResultCode>
	<ResultDesc>Failed</ResultDesc>
	<TransID>10111</TransID>
	<BillRefNumber>12233</BillRefNumber>
	<UtilityName>sddd</UtilityName>
	<CustomerName>wee</CustomerName>
	<Amount>30</Amount>
	</c2b:C2BPaymentQueryResult>
	</soapenv:Body>
	</soapenv:Envelope>
	`
	password := "arkised"

	sendDataToServer2(xmlString, password)
}
