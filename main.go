package main

import (
	"fmt"

	mathLib "github.com/madeindra/trainocate-library/pkg/math"
)

type List interface {
	Sum() int64
	Average() float64
	IsTooLong(list interface{}) (bool, error)
}

type Numbers struct {
	intList   []int64
	floatList []float64
}

func (n Numbers) Sum() int64 {
	return mathLib.Sum(n.intList)
}

func (n Numbers) Average() float64 {
	return mathLib.Average(n.floatList)
}

func (n Numbers) IsTooLong(list interface{}) (bool, error) {
	return mathLib.IsTooLong(list)
}

func main() {
	num := Numbers{
		intList:   []int64{1, 2, 3, 4, 5},
		floatList: []float64{1, 2, 3, 4, 5},
	}

	if tooLong, err := num.IsTooLong(num.intList); tooLong || err != nil {
		panic("sum operation not possible")
	}

	if tooLong, err := num.IsTooLong(num.floatList); tooLong || err != nil {
		panic("average operation not possible")
	}

	sum := num.Sum()
	avg := num.Average()

	fmt.Println("Sum:", sum, "Avg:", avg)
}
