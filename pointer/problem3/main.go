package main

import (
	"fmt"
)

func main() {
	var a = Student{}

	for i := 0; i < 5; i++ {
		var name string
		fmt.Print("Input Student's Name : ")
		fmt.Scan(&name)
		a.name = append(a.name, name)

		var score int
		fmt.Print("Input " + name + " Score : ")
		fmt.Scan(&score)
		a.score = append(a.score, score)
	}

	fmt.Println("\n\nAverage Score Students is", a.Average())
	scoreMax, nameMax := a.Max()
	fmt.Println("Max Score Student is : "+nameMax+" (", scoreMax, ")")
	scoreMin, nameMin := a.Min()
	fmt.Println("Min Score Student is : "+nameMin+" (", scoreMin, ")")
}

type Student struct {
	name  []string
	score []int
}

func (s Student) Average() float64 {
	var average float64

	for i := 0; i < len(s.score); i++ {
		average = average + float64(s.score[i])
	}

	average = average / float64(len(s.score))

	return average
}

func (s Student) Min() (min int, name string) {
	min = s.score[0]
	name = s.name[0]

	for i, a := range s.score {
		if a < min {
			min = a
			name = s.name[i]
		}
	}

	return min, name
}

func (s Student) Max() (max int, name string) {
	max = 0
	name = s.name[0]

	for i, a := range s.score {
		if a > max {
			max = a
			name = s.name[i]
		}
	}

	return max, name
}
