// https://www.hackerrank.com/challenges/two-arrays/problem

package main

import (
	"fmt"
	"sort"
)

func twoArrays(k int32, A []int32, B []int32) string {
	sort.Slice(A, func(i, j int) bool {
		return A[i] < A[j]
	})

	sort.Slice(B, func(i, j int) bool {
		return B[i] > B[j]
	})

	for i := 0; i < len(A); i++ {
		if A[i] + B[i] < k {
			return "NO"
		}
	}

	return "YES"
}

func main() {
	fmt.Println(twoArrays(10, []int32{2, 1, 3}, []int32{7, 8, 9}) == "YES")
	fmt.Println(twoArrays(5, []int32{1, 2, 2, 1}, []int32{3, 3, 3, 4}) == "NO")
}
