package array

import (
	"container/list"
	"math"
	"strconv"
)
//Input: forbidden = [14,4,18,1,15], a = 3, b = 15, x = 9
//Output: 3

type trace struct {
	pos int
	canback bool
}

func MinimumJumps(forbidden []int, a int, b int, x int) int {
	if x == 0{
		return 0
	}
	var maxLimit int  = 2000 + 2 * b;
	var block map[int]bool = make(map[int]bool)
	for _,f := range forbidden{
		block[f] = true
		maxLimit = int(math.Max(float64(maxLimit), float64(f + 2 * b)));
	}
	var q list.List
	var visited map[string]bool = make(map[string]bool)
	var start trace
	start.pos = 0
	start.canback = true
	visited["0,1"] = true
	q.PushBack(start)
	var steps int = 0
	for q.Len() > 0{
		var l int = q.Len()
		for i := 0;i < l;i++{
			var cur trace = q.Front().Value.(trace)
			q.Remove(q.Front())
			if cur.pos == x{
				return steps
			}
			var forward trace
			forward.pos = cur.pos + a
			forward.canback = true
			if forward.pos <= maxLimit && forward.pos >= -2000 {
				if _,ok1 := block[forward.pos];!ok1{
					k1 := strconv.Itoa(forward.pos) + ",1"
					if _,ok2 := visited[k1];!ok2{
						visited[k1] = true
						q.PushBack(forward)
					}
				}
			}
			if cur.canback{
				var backward trace
				backward.pos = cur.pos - b
				backward.canback = false
				if backward.pos <= maxLimit && backward.pos >= 0{
					if _,ok1 := block[backward.pos];!ok1{
						k2 := strconv.Itoa(backward.pos) + ",0"
						if _,ok2 := visited[k2];!ok2{
							visited[k2] = true
							q.PushBack(backward)
						}
					}
				}
			}
		}
		steps++
	}
	return -1
}