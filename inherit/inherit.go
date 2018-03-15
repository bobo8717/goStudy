package inherit

import (
	"fmt"
)

type M1 struct {
	T1 int
	T2 int
}

func CreateM1() (m *M1) {
	return &M1{T1: 100}
}
func (m *M1) ShowT1() {
	fmt.Printf("M1: T1->%d ", m.T1)
}

type M2 struct {
	*M1
	Val int
}

func CreateM2() (m *M2) {
	//res := &M2{value: 1000}
	res := &M2{CreateM1(), 1000}
	return res
}
func (m *M2) Test2() {
	fmt.Printf("Val-2:%d\n", m.Val)
}
