package main

import "fmt"

type stack struct {
	size int
	top  int
	data []int
}

var map1 = make(map[int]int)

func main() {
	nums1 := []int{4, 1, 2}
	nums2 := []int{1, 3, 4, 2}

	output := nextGreaterElement(nums1, nums2)
	fmt.Println(output)

}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	s := stack{}
	s.size = len(nums2)
	for i := len(nums2) - 1; i >= 0; i-- {
		s.push(nums2[i])
	}
	fmt.Println(map1)
	output := []int{}
	for j := 0; j < len(nums1); j++ {
		temp := map1[nums1[j]]
		output = append(output, temp)

	}
	return output
}

func (s *stack) push(x int) {

	if !s.isFull() {
		for s.top >= len(s.data) {
			if s.top == 0 {
				s.data = append(s.data, x)
				s.top++
				map1[x] = -1
				break
			} else {
				fmt.Println(s.top-1, len(s.data))

				if x > s.data[s.top-1] {

					s.pop()

					map1[x] = -1
				} else {
					map1[x] = s.data[s.top-1]
					s.data = append(s.data, x)
					s.top++
					break
				}

			}
		}

	} else {
		fmt.Println("overflow")
	}

}

func (s *stack) isFull() bool {

	if s.top == s.size {
		return true
	} else {
		return false
	}
}
func (s *stack) pop() {

	if !s.isEmpty() {
		s.data = append(s.data[:s.top-1], s.data[s.top-1+1:]...)
		s.top--
	} else {
		fmt.Println("underflow")
	}

}

func (s *stack) isEmpty() bool {

	if s.top == 0 {
		return true
	} else {
		return false
	}

}
