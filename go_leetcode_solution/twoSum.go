package main

import "fmt"

func twoSum(nums []int, target int) []int {
    length := len(nums)
    var index []int

    for i := range nums {
        for j := i+1; j < length; j++ {
            if nums[i] + nums[j] == target {
                index = append(index,i,j)
                return index
            }
        }
    }
    return index
}

func main() {

	nums := []int{2, 7, 11, 15}
    target := 9
    result := twoSum(nums, target)

    if len(result) == 0 {
        fmt.Println("No two sum solution found")
    } else {
        fmt.Printf("Indices of the numbers that add up to %d are: %v\n", target, result)
    }
}