package diagram

import "sort"

func maxArea(h int, w int, horizontalCuts []int, verticalCuts []int) int {
	sort.Ints(horizontalCuts)
	sort.Ints(verticalCuts)
	var l1 int = len(horizontalCuts)
	var l2 int = len(verticalCuts)
	var max_h int = 0
	pre_h := 0
	for i := 0;i < l1;i++{
		max_h = max_int(max_h,horizontalCuts[i] - pre_h)
		pre_h = horizontalCuts[i]
	}
	max_h = max_int(max_h,h - pre_h)
	var max_v int = 0
	pre_v := 0
	for i := 0;i < l2;i++{
		max_v = max_int(max_v,verticalCuts[i] - pre_v)
		pre_v = verticalCuts[i]
	}
	max_v = max_int(max_v,w - pre_v)
	return max_h * max_v % (1e9 + 7)
}

func maxArea2(h int, w int, horizontalCuts []int, verticalCuts []int) int{
	//horizontalCuts = append(horizontalCuts,0)
	horizontalCuts = append(horizontalCuts,h)
	//verticalCuts = append(verticalCuts,0)
	verticalCuts = append(verticalCuts,w)
	sort.Ints(horizontalCuts)
	sort.Ints(verticalCuts)
	var l1 int = len(horizontalCuts)
	var l2 int = len(verticalCuts)
	var max_h int = 0
	var max_w int = 0
	var pre_h int = 0
	var pre_w int = 0
	for i := 0;i < l1;i++{
		if horizontalCuts[i] - pre_h > max_h{
			max_h = horizontalCuts[i] - pre_h
		}
		pre_h = horizontalCuts[i]
	}
	for i := 0;i < l2;i++{
		if verticalCuts[i] - pre_w > max_w{
			max_w = verticalCuts[i] - pre_w
		}
		pre_w = verticalCuts[i]
	}
	return max_h * max_w % (1e9 + 7)
}