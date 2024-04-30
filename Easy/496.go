package main

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	stack := []int{}
	m := make(map[int]int)

	for _, v := range nums2 {
		for len(stack) != 0 && v >= nums2[stack[len(stack)-1]] {

			m[nums2[stack[len(stack)-1]]] = v
			stack = stack[:len(stack)-1]
		}

		for i := range nums2 {
			if nums2[i] == v {
				stack = append(stack, i)
				break
			}
		}
	}
	res := make([]int, len(nums1))
	for i, num := range nums1 {
		if val, ok := m[num]; ok {
			res[i] = val
		} else {
			res[i] = -1
		}
	}
	return res
}
