package number

import (
	"sort"
	"strings"
)

func topStudents(positive_feedback []string, negative_feedback []string, report []string, student_id []int, k int) []int {
	var pos map[string]bool = make(map[string]bool)
	var neg map[string]bool = make(map[string]bool)
	for _,word := range positive_feedback{
		pos[word] = true
	}
	for _,word := range negative_feedback{
		neg[word] = true
	}
	var l int = len(student_id)
	var id_score [][2]int = make([][2]int,l)
	for i := 0;i < l;i++{
		var s []string = strings.Split(report[i]," ")
		var scores int = 0
		for _,word := range s{
			if _,ok := pos[word];ok{
				scores += 3
			}
			if _,ok := neg[word];ok{
				scores--
			}
		}
		id_score[i] = [2]int{student_id[i],scores}
	}
	sort.Slice(id_score, func(i, j int) bool {
		if id_score[i][1] == id_score[j][1]{
			return id_score[i][0] < id_score[j][0]
		}
		return id_score[i][1] > id_score[j][1]
	})
	var res []int
	var i int = 0
	for k > 0{
		res = append(res,id_score[i][0])
		i++
		k--
	}
	return res
}