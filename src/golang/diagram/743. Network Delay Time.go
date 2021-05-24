package diagram

//You are given a network of n nodes, labeled from 1 to n. You are also given times,
//a list of travel times as directed edges times[i] = (ui, vi, wi), where ui is the source node,
//vi is the target node, and wi is the time it takes for a signal to travel from source to target.
//
//We will send a signal from a given node k.
//Return the time it takes for all the n nodes to receive the signal.
//If it is impossible for all the n nodes to receive the signal, return -1.

//times = [[2,1,1],[2,3,1],[3,4,1]], n = 4, k = 2

type node_duration struct{
	num int
	durantion int
}

func NetworkDelayTime(times [][]int, N int, K int) int {
	var distance []int = make([]int,N + 1)//the distance from K to current node
	for i := 1;i <= N;i++{
		distance[i] = 2147483647
	}
	distance[K] = 0
	var l int = len(times)
	for i := 1;i <= N;i++{
		for j := 0;j < l;j++{
			var start int = times[j][0]
			var end int = times[j][1]
			var duration int = times[j][2]
			if distance[start] == 2147483647{
				continue
			}
			if (distance[start] + duration) < distance[end]{
				distance[end] = distance[start] + duration
			}
		}
	}
	var res int = 0
	for i := 1;i <= N;i++{
		if distance[i] == 2147483647{
			return -1
		}
		if res < distance[i]{
			res = distance[i]
		}
	}
	return res
}