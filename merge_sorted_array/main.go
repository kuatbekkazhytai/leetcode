package main

import "fmt"

func main() {
	nums1 := []int{2, 0}
	nums2 := []int{1}
	m := 1
	n := 1
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	index1 := 0
	index2 := 0
	cp := make([]int, m+n)
	if len(nums2) == 0 {
		copy(cp, nums1)
	}

	for i := 0; i < m+n; i++ {
		if index2 == len(nums2) {
			cp[i] = nums1[index1]
			index1++
			continue
		}
		if nums1[index1] < nums2[index2] {
			if nums1[index1] == 0 && index1 > m-1 {
				cp[i] = nums2[index2]
				index2++
				continue
			}

			cp[i] = nums1[index1]
			index1++
			continue

		}
		cp[i] = nums2[index2]
		index2++
	}
	copy(nums1, cp)
}
