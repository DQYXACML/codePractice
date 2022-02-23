package main

import (
	"sort"
)

// https://leetcode.com/problems/meeting-rooms-ii/

type Arrange struct {
	Begin int
	End   int
}

type arranges []Arrange

func (a arranges) Len() int {
	return len(a)
}

func (a arranges) Less(i, j int) bool {
	return a[j].End > a[i].End
}

func (a arranges) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func bestArrangeMeeting(arr *arranges) int {
	ans := 0
	timeline := 0
	sort.Sort(arr)
	for i := 0; i < len(*arr); i++ {
		if (*arr)[i].Begin >= timeline {
			ans++
			timeline = (*arr)[i].End
		}
	}
	return ans
}
