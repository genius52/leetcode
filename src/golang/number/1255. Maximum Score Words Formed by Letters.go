package number

//Input: words = ["dog","cat","dad","good"], letters = ["a","a","c","d","d","d","g","o","o"], score = [1,0,9,5,0,0,3,0,0,0,0,0,0,0,2,0,0,0,0,0,0,0,0,0,0,0]
//Output: 23
//Explanation:
//Score  a=1, c=9, d=5, g=3, o=2
//Given letters, we can form the words "dad" (5+1+5) and "good" (3+2+2+5) with a score of 23.
//Words "dad" and "dog" only get a score of 21.
func dfs_maxScoreWords(words []string,l int,cur_pos int,can_choose [26]int,score []int)int{
	if cur_pos >= l{
		return 0
	}
	var meet bool = true
	var cur_char_cnt [26]int
	for _,c := range words[cur_pos]{
		cur_char_cnt[c - 'a']++
		if cur_char_cnt[c - 'a'] > can_choose[c - 'a']{
			meet = false
			break
		}
	}
	//choose current word
	var profit1 int = 0
	if meet{
		cur_score := 0
		for i := 0;i < 26;i++{
			cur_score += (score[i] * cur_char_cnt[i])
			can_choose[i] -= cur_char_cnt[i]
		}
		profit1 += cur_score + dfs_maxScoreWords(words,l,cur_pos + 1,can_choose,score)
		for i := 0;i < 26;i++{
			can_choose[i] += cur_char_cnt[i]
		}
	}
	//skip current word
	var profit2 int = dfs_maxScoreWords(words,l,cur_pos + 1,can_choose,score)
	return max_int(profit1,profit2)
}

func MaxScoreWords(words []string, letters []byte, score []int) int {
	var can_choose [26]int
	for i := 0;i < len(letters);i++{
		can_choose[letters[i] - 'a']++
	}
	return dfs_maxScoreWords(words,len(words),0,can_choose,score)
}