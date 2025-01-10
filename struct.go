package main

import "fmt"

func main() {
	std1 := Student{"Likhil", []int{80, 90}}
	std1.addMark(100)
	std1.CalculateAverage()
}

type Student struct {
	Name  string
	Marks []int
}

func (s Student) addMark(mark int) {
	s.Marks = append(s.Marks, mark)
	fmt.Println(s)
}
func (s Student) CalculateAverage() float64 {

	return (s.CalculateAverage())
}
