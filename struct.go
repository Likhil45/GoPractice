package main

import "fmt"

func main() {
	std1 := Student{"Likhil", []int{80, 90}}
	std1.addMark(100)
	std1.CalculateAverage()
	fmt.Println(std1)
	fmt.Printf("Average = %f", std1.CalculateAverage())

}

type Student struct {
	Name  string
	Marks []int
}

func (s *Student) addMark(mark int) {
	s.Marks = append(s.Marks, mark)

}
func (s *Student) CalculateAverage() float64 {
	var tot float64
	var avg float64
	for _, val := range s.Marks {
		tot = tot + float64(val)
	}
	avg = tot / float64((len(s.Marks)))
	return avg
}
