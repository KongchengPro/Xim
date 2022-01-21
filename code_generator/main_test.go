package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

func TestFindMarks(t *testing.T) {
	bytes, err := ioutil.ReadFile("./findmarks.txt")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	lines := strings.Split(string(bytes), "\n")
	marksLineIndex := FindMarks(lines)
	fmt.Println(marksLineIndex)
	if !AssertIntSliceEqual(marksLineIndex, []int{1, 3, 5, 7}) {
		t.Fail()
	}
}
func AssertIntSliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	if a == nil || b == nil {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
