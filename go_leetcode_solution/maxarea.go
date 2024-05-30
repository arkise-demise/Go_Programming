package main

import (
	"fmt"
	"math"
)

func maxArea(height []int) int {
    left, right := 0, len(height) - 1
    maxArea := 0

    for left < right {
        // Calculate the area
        width := right - left
        h := int(math.Min(float64(height[left]), float64(height[right])))
        currentArea := width * h

        // Update max area if the current area is greater
        if currentArea > maxArea {
            maxArea = currentArea
        }

        // Move the pointers
        if height[left] < height[right] {
            left++
        } else {
            right--
        }
    }

    return maxArea
}

func main() {
    height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
    fmt.Printf("Max area: %d\n", maxArea(height)) // Output: 49
}
