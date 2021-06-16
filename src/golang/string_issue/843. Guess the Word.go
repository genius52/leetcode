package string_issue

import (
	"math/rand"
	"time"
)

type Master struct {

}
func (this *Master) Guess(word string) int {
	return 0
}

func check_same(s1 string,s2 string)int{
	var l int = len(s1)
	var res int = 0
	for i := 0;i < l;i++{
		if s1[i] == s2[i]{
			res++
		}
	}
	return res
}

func findSecretWord(wordlist []string, master *Master) {
	var choose []string = wordlist
	for i := 0;i < 10;i++{
		rand.Seed(time.Now().UnixNano())
		index := rand.Intn(len(choose)) % len(choose)
		correct_cnt := master.Guess(choose[index])
		if correct_cnt == 6{
			return
		}
		var next []string
		for j := 0;j < len(choose);j++{
			if j == index{
				continue
			}
			same_cnt := check_same(choose[index],choose[j])
			if same_cnt == correct_cnt{
				next = append(next,choose[j])
			}
		}
		choose = next
	}
}