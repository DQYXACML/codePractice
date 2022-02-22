package main

import "math"

// https://www.nowcoder.com/questionTerminal/a5f542742fe24181b28f7d5b82e2e49a

type AnyTree struct {
	Val   int
	nexts []*AnyTree
}

func maxHappy(root *AnyTree) int {
	maxInfo := happyProcess(root)
	return int(math.Max(float64(maxInfo.no), float64(maxInfo.yes)))
}

type happyInfo struct {
	yes int // 来
	no  int // 不来
}

func happyProcess(root *AnyTree) *happyInfo {
	if root == nil {
		return &happyInfo{
			yes: 0,
			no:  0,
		}
	}

	var yes int
	var no int

	for _, next := range root.nexts {
		nextInfo := happyProcess(next)
		yes = nextInfo.no
		no = int(math.Max(float64(nextInfo.yes), float64(nextInfo.no)))
	}

	return &happyInfo{
		yes: yes,
		no:  no,
	}
}
