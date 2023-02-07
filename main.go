package main

import (
	"fmt"

	mathLib "github.com/madeindra/trainocate-library/pkg/math"
)

type List interface {
	Sum(number []int64) int64
	Average(number []float64) float64
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

func main() {
	num := Numbers{
		intList:   []int64{1, 2, 3, 4, 5},
		floatList: []float64{1, 2, 3, 4, 5},
	}

	sum := num.Sum()
	avg := num.Average()

	fmt.Println("Sum:", sum, "Avg:", avg)
}
