package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{3, 3, 5, 4, 4, 4, 1, 2, 9, 7, 7, 6}
	target := 8
	arr := make([][]int, len(nums))
	for i := range nums {
		arr[i] = []int{i, nums[i]}
	}
	sort.Slice(arr, func(i, j int) bool { return arr[i][1] < arr[j][1] })
	var resValue, resIndex []int
	l, r := 0, len(nums)-1
	for l < r {
		if arr[l][1]+arr[r][1] < target {
			l++
		} else if arr[l][1]+arr[r][1] > target {
			r--
		} else {
			count1, count2 := 0, 0
			for l < r && arr[l][1] == arr[l+1][1] {
				l++
				count1++
			}
			for l < r && arr[r][1] == arr[r-1][1] {
				r--
				count2++
			}
			if arr[l][1] == arr[r][1] && count1 >= 2 {
				for i := l - count1; i < r; i++ {
					for j := i + 1; j <= r; j++ {
						resValue = append(resValue, arr[i][1], arr[j][1])
						resIndex = append(resIndex, arr[i][0], arr[j][0])
					}
				}
				break
			}
			for i := 0; i <= count1; i++ {
				for j := 0; j <= count2; j++ {
					if l-i != r+j {
						resValue = append(resValue, arr[l-i][1], arr[r+j][1])
						resIndex = append(resIndex, arr[l-i][0], arr[r+j][0])
					}
				}
			}
			l++
			r--
		}
	}
	fmt.Println("resValue=", resValue) // resValue= [1 7 1 7 2 6 3 5 3 5 4 4 4 4 4 4]
	fmt.Println("resIndex=", resIndex) // resIndex= [6 9 6 10 7 11 1 2 0 2 3 4 3 5 4 5]
}
