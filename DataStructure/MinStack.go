package datastructure

import (
	"fmt"
)

func TestMinStack() {
	// minStack := ConstructorMS()
	// fmt.Println(minStack)
	// minStack.Push(-2)
	// fmt.Println(minStack)
	// minStack.Push(0)
	// fmt.Println(minStack)
	// minStack.Push(-3)
	// fmt.Println(minStack)
	// fmt.Println(minStack.GetMin())
	// minStack.Pop()
	// fmt.Println(minStack)
	// fmt.Println(minStack.Top())
	// fmt.Println(minStack.GetMin())

	obj := ConstructorMS()
	obj.Push(0)
	obj.Push(1)
	obj.Push(0)
	fmt.Println(obj.GetMin())
	obj.Pop()
	fmt.Println(obj.GetMin())
}

type MinStack struct {
	data         []int
	minValue     []int
	sizeData     int
	sizeMinValue int
}

func ConstructorMS() MinStack {
	var s []int
	var m []int
	return MinStack{s, m, 0, 0}
}

func (s *MinStack) Push(x int) {
	s.data = append(s.data, x)
	lOfMinValue := s.sizeMinValue
	if lOfMinValue == 0 {
		s.minValue = append(s.minValue, x)
		s.sizeMinValue++
	} else {
		if x <= s.minValue[lOfMinValue-1] {
			s.minValue = append(s.minValue, x)
			s.sizeMinValue++
		}
	}
	s.sizeData++
}

func (s *MinStack) Pop() {
	lData := s.sizeData
	lMinValue := s.sizeMinValue
	if lData > 0 {
		if s.minValue[lMinValue-1] == s.data[lData-1] {
			s.minValue = s.minValue[:lMinValue-1]
			s.sizeMinValue--
		}
		s.data = s.data[:lData-1]
	}
	s.sizeData--
}

func (s *MinStack) Top() int {
	return s.data[s.sizeData-1]
}

func (s *MinStack) GetMin() int {
	lOfMinValue := len(s.minValue)
	return s.minValue[lOfMinValue-1]
}
