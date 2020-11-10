package main

import (
	"fmt"
	"log"
	"testing"
)

// TestIntervalMerge to test output
func TestIntervalMerge(t *testing.T) {
	var input = []interval{newInterval(25, 30), newInterval(2, 19), newInterval(14, 23), newInterval(4, 8)}
	output := MergeSort(input)

	mergedOutput := MergeInterval(output)
	log.Print(fmt.Sprintf("[%v]", mergedOutput))
	if len(mergedOutput) != 2 {
		t.Errorf("length of output %d; want 2", len(mergedOutput))
	}
}
