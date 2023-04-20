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

//func merge(nums1 []int, m int, nums2 []int, n int) {
//	index1 := 0
//	index2 := 0
//	cp := make([]int, len(nums1))
//	copy(cp, nums1)
//
//	for i := 0; i < m+n; i++ {
//		if index2 == len(nums2) {
//			break
//		}
//		if nums1[index1] <= nums2[index2] && nums1[index1] != 0 {
//			if i == len(nums1) {
//				nums1 = append(nums1, nums1[index1])
//			} else {
//				nums1[i] = nums1[index1]
//			}
//
//			index1++
//		} else {
//			if nums1[index1] == 0 && i < m && nums2[index2] > 0 {
//				if i == len(nums1) {
//					nums1 = append(nums1, nums1[index1])
//				} else {
//					nums1[i] = nums1[index1]
//				}
//				index1++
//				continue
//			}
//
//			x := nums1[i:]
//			x = x[:len(x)-1]
//			temp := append([]int{nums2[index2]}, x...)
//			nums1 = append(nums1[:i], temp...)
//			index1++
//			index2++
//		}
//	}
//	fmt.Println(nums1)
//}
