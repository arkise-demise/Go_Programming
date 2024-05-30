package main

import (
	"fmt"
	"math"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    if len(nums1) > len(nums2) {
        return findMedianSortedArrays(nums2, nums1)
    }

    m, n := len(nums1), len(nums2)
    imin, imax, half_len := 0, m, (m+n+1)/2

    for imin <= imax {
        i := (imin + imax) / 2
        j := half_len - i

        if i < m && nums1[i] < nums2[j-1] {
            imin = i + 1 // i is too small, must increase it
        } else if i > 0 && nums1[i-1] > nums2[j] {
            imax = i - 1 // i is too big, must decrease it
        } else { // i is perfect
            max_of_left := 0
            if i == 0 {
                max_of_left = nums2[j-1]
            } else if j == 0 {
                max_of_left = nums1[i-1]
            } else {
                max_of_left = int(math.Max(float64(nums1[i-1]), float64(nums2[j-1])))
            }

            if (m+n)%2 == 1 {
                return float64(max_of_left)
            }

            min_of_right := 0
            if i == m {
                min_of_right = nums2[j]
            } else if j == n {
                min_of_right = nums1[i]
            } else {
                min_of_right = int(math.Min(float64(nums1[i]), float64(nums2[j])))
            }

            return (float64(max_of_left) + float64(min_of_right)) / 2.0
        }
    }

    return 0.0
}

func main() {
    // Test cases
    nums1 := []int{1, 3}
    nums2 := []int{2}
    fmt.Printf("Median of arrays: %f\n", findMedianSortedArrays(nums1, nums2)) // Output: 2.0

    nums1 = []int{1, 2}
    nums2 = []int{3, 4}
    fmt.Printf("Median of arrays: %f\n", findMedianSortedArrays(nums1, nums2)) // Output: 2.5
}
