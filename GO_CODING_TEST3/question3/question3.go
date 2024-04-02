package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Employee struct {
	Name string `json:"name"`
	Salary float64 `json:"salary"`
}

func calculateAverageSalary(employees []Employee)float64 {
	totalsalary :=0.0
	for _, emp := range employees {
		totalsalary += emp.Salary
	}
	return totalsalary / float64(len(employees))
}

func main() {
	inputFile := "employees.json"
	outputFile :="updated_employees.json"

	file, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	var employees []Employee
	if err := json.Unmarshal(file, &employees); err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	averageSalary := calculateAverageSalary(employees)
	for i := range employees {
		if employees[i].Salary < averageSalary {
			employees[i].Salary *= 1.1
		}
	}

		updatedData, err := json.MarshalIndent(employees, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	err = ioutil.WriteFile(outputFile, updatedData, 0644)
	if err != nil {
		fmt.Println("Error writing output file:", err)
		return
	}

	fmt.Println("Employee data updated", outputFile)
 
}
