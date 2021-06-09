package string_issue

//Input:
//s = "heeellooo"
//words = ["hello", "hi", "helo"]
//Output: 1
//Explanation:
//We can extend "e" and "o" in the word "hello" to get "heeellooo".
//We can't extend "helo" to get "heeellooo" because the group "ll" is not size 3 or more.
func check_expressiveWords(S string,word string)bool{
	var l1 int = len(S)
	var l2 int = len(word)
	if l2 > l1{
		return false
	}
	var left1 int = 0
	var left2 int = 0
	var extend bool = false
	for left1 < l1 && left2 < l2{
		if S[left1] != word[left2]{
			return false
		}
		var right1 int = left1
		for right1 < l1 && S[left1] == S[right1]{
			right1++
		}
		dis1 := right1 - left1
		var right2 int = left2
		for right2 < l2 && word[left2] == word[right2]{
			right2++
		}
		dis2 := right2 - left2
		if dis2 > dis1 || (dis1 <= 2 && dis2 < dis1){
			return false
		}
		if dis1 > 2 && dis1 >= dis2{
			extend = true
		}
		left1 = right1
		left2 = right2
	}
	if left1 < l1 || left2 < l2{
		return false
	}
	if extend{
		return true
	}
	return false
}

func ExpressiveWords(S string, words []string) int {
	var l int = len(words)
	var res int = 0
	for i := 0;i < l;i++{
		if check_expressiveWords(S,words[i]){
			res++
		}
	}
	return res
}