package main

import (
	"fmt"
	"log"
)

type interval struct {
	Low  int
	High int
}

func main() {
	var input = []interval{newInterval(25, 30), newInterval(2, 19), newInterval(14, 23), newInterval(4, 8)}
	output := mergeSort(input)
	log.Print(fmt.Sprintf("input: %s", output))

	low := output[0].Low
	high := output[0].High
	for i := 0; i < len(input); i++ {
		// check if it is end of list
		if i+1 < len(input) {
			// next interval overlaps current interval
			if output[i+1].Low < high {
				if output[i+1].High > high {
					high = output[i+1].High
				}
			} else {
				// use next interval
				log.Print(fmt.Sprintf("[%d,%d]", low, high))
				low = output[i+1].Low
				high = output[i+1].High
			}
		}
	}
	log.Print(fmt.Sprintf("[%d,%d]", low, high))
}

//simple divide and conquer sort
func mergeSort(input []interval) []interval {
	length := len(input)
	if length == 1 {
		return input
	}

	middle := int(length / 2)
	left := make([]interval, middle)
	right := make([]interval, length-middle)

	for i := 0; i < length; i++ {
		if i < middle {
			left[i] = input[i]
		} else {
			right[i-middle] = input[i]
		}
	}

	return merge(mergeSort(left), mergeSort(right))
}

func merge(left, right []interval) (result []interval) {
	result = make([]interval, len(left)+len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0].Low < right[0].Low {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}

	return
}

//convinient function to create interval
func newInterval(low int, high int) interval {
	return interval{low, high}
}
