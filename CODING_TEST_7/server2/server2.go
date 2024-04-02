package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"os"
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

const EncryptedPassword = "arkised"

func receiveData(w http.ResponseWriter, r *http.Request) {
	var jsonData Data
	err := json.NewDecoder(r.Body).Decode(&jsonData)
	if err != nil {
		http.Error(w, "Unable to decode JSON data", http.StatusBadRequest)
		return
	}

	if jsonData.Password != EncryptedPassword {
		http.Error(w, "Incorrect password", http.StatusUnauthorized)
		return
	}

	var filename string
	if jsonData.C2BPaymentQueryResult.ResultCode == "2" {
		filename = "success.xml" 
	} else {
		filename = "failed.xml"
	}

	file, err := os.Create(filename)
	if err != nil {
		http.Error(w, "Unable to save data", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	xmlBytes, err := xml.Marshal(jsonData.C2BPaymentQueryResult)
	if err != nil {
		http.Error(w, "Unable to marshal XML data", http.StatusInternalServerError)
		return
	}

	_, err = file.Write(xmlBytes)
	if err != nil {
		http.Error(w, "Unable to write XML data to file", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Data received and saved successfully\n")
}

func main() {
	http.HandleFunc("/receive", receiveData)
	fmt.Println("Server 2 listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
