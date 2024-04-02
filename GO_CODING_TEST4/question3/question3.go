package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Person struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

func main(){
	file,err :=os.Open("data.json")
	if err !=nil {
		panic(err)
		
	}
	defer file.Close()
	data, err :=ioutil.ReadAll(file)
	if err !=nil {
		panic(err)
	}

	var person []Person

	err = json.Unmarshal(data,&person)
	if err != nil {
		panic(err)
	}
	fmt.Println("=================================")
	for _, p := range person {
		fmt.Printf("Name: %s, Age: %d\n", p.Name, p.Age)
	}
	fmt.Println("=================================")

}
