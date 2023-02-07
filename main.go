package main

import (
	logLib "github.com/madeindra/trainocate-library/pkg/log"
	mathLib "github.com/madeindra/trainocate-library/pkg/math/v2"
)

type List interface {
	Sum(ch chan<- int64)
	Average(ch chan<- float64)
	IsTooLong(list interface{}) (bool, error)
}

type Numbers struct {
	intList   []int64
	floatList []float64
}

func (n Numbers) Sum(ch chan<- int64) {
	mathLib.Sum(n.intList, ch)
}

func (n Numbers) Average(ch chan<- float64) {
	mathLib.Average(n.floatList, ch)
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

	sumChan := make(chan int64)
	avgChan := make(chan float64)

	go num.Sum(sumChan)
	go num.Average(avgChan)

	for i := 0; i < 2; i++ {
		select {
		case sum := <-sumChan:
			logLib.Print("Sum is", sum)
		case avg := <-avgChan:
			logLib.Print("Average is", avg)
		}
	}

}
