package main

import (
	"fmt"
)

type CompanyProfit struct {
	ID         string `json:"CompanyID"`
	Profit     int    `json:"Profit"`
	Department string `json:"Department"`
}

type DepartmentProfit struct {
	Department        string `json:"Department"`
	DepartmentProfit  int    `json:"DepartmentProfit"`
}

type CompanyTotalProfit struct {
	CompanyID   string             `json:"CompanyID"`
	TotalProfit int                `json:"TotalProfit"`
	Departments []DepartmentProfit `json:"Department"`
}

func calculateTotalProfit(companyProfits []CompanyProfit) []CompanyTotalProfit {
	companyMap := make(map[string]map[string]int)

	for _, cp := range companyProfits {
		if _, ok := companyMap[cp.ID]; !ok {
			companyMap[cp.ID] = make(map[string]int)
		}
		companyMap[cp.ID][cp.Department] += cp.Profit
	}

	var result []CompanyTotalProfit

	for companyID, departmentMap := range companyMap {
		var departments []DepartmentProfit
		totalProfit := 0

		for department, departmentProfit := range departmentMap {
			departments = append(departments, DepartmentProfit{
				Department:       department,
				DepartmentProfit: departmentProfit,
			})
			totalProfit += departmentProfit
		}

		result = append(result, CompanyTotalProfit{
			CompanyID:   companyID,
			TotalProfit: totalProfit,
			Departments: departments,
		})
	}

	return result
}

func printResults(results []CompanyTotalProfit) {
	fmt.Println("[")
	for i, result := range results {
		fmt.Printf("    {\n        CompanyID: \"%s\",\n        TotalProfit: %d,\n        Department: []company.DepartmentProfit{\n",
			result.CompanyID, result.TotalProfit)

		for j, dept := range result.Departments {
			fmt.Printf("            {\n                Department: \"%s\",\n                DepartmentProfit: %d,\n            }",
				dept.Department, dept.DepartmentProfit)

			if j < len(result.Departments)-1 {
				fmt.Print(",")
			}
			fmt.Println()
		}

		fmt.Print("        },\n    }")

		if i < len(results)-1 {
			fmt.Print(",")
		}
		fmt.Println()
	}
	fmt.Println("]")
}

func main() {
	companyProfits := []CompanyProfit{
		{ID: "1", Profit: 10, Department: "A"},
		{ID: "1", Profit: 5, Department: "A"},
		{ID: "1", Profit: 5, Department: "B"},
		{ID: "2", Profit: 5, Department: "A"},
	}

	result := calculateTotalProfit(companyProfits)

	printResults(result)
}
