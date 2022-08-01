package diagram

import "container/list"

func closestMeetingNode(edges []int, node1 int, node2 int) int {
	var l int = len(edges)
	var dis1 []int = make([]int,l)
	for i := 0;i < l;i++{
		dis1[i] = 2147483647
	}
	dis1[node1] = 0
	var dis2 []int = make([]int,l)
	for i := 0;i < l;i++{
		dis2[i] = 2147483647
	}
	dis2[node2] = 0
	var q1 list.List
	q1.PushBack(node1)
	var steps int = 1
	for q1.Len() > 0{
		var next int = edges[q1.Front().Value.(int)]
		q1.Remove(q1.Front())
		if next != -1 && steps < dis1[next]{
			q1.PushBack(next)
			dis1[next] = steps
		}
		steps++
	}
	var q2 list.List
	q2.PushBack(node2)
	steps = 1
	for q2.Len() > 0{
		var next int = edges[q2.Front().Value.(int)]
		q2.Remove(q2.Front())
		if next != -1 && steps < dis2[next]{
			q2.PushBack(next)
			dis2[next] = steps
		}
		steps++
	}
	var res int = 2147483647
	var node int = -1
	for i := 0;i < l;i++{
		if dis1[i] != 2147483647 && dis2[i] != 2147483647{
			if max_int(dis1[i],dis2[i]) < res{
				res = max_int(dis1[i],dis2[i])
				node = i
			}
		}
	}
	return node
}