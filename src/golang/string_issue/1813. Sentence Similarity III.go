package string_issue

import "strings"

func AreSentencesSimilar(sentence1 string, sentence2 string) bool {
	var l1 int = len(sentence1)
	var l2 int = len(sentence2)
	if strings.Compare(sentence1, sentence2) == 0 ||
		(strings.HasPrefix(sentence1, sentence2) && sentence1[l2] == ' ') ||
		(strings.HasSuffix(sentence1, sentence2) && sentence1[l1-1-l2] == ' ') ||
		(strings.HasPrefix(sentence2, sentence1) && sentence2[l1] == ' ') ||
		(strings.HasSuffix(sentence2, sentence1) && sentence2[l2-1-l1] == ' ') {
		return true
	}
	var record1 []string = strings.Split(sentence1, " ")
	var record2 []string = strings.Split(sentence2, " ")
	var left1 int = 0
	var right1 int = len(record1) - 1
	var left2 int = 0
	var right2 int = len(record2) - 1
	var find_left bool = false
	for left1 <= right1 && left2 < right2 {
		if record1[left1] == record2[left2] {
			find_left = true
			left1++
			left2++
		} else {
			break
		}
	}
	var find_right bool = false
	for right1 >= left1 && right2 >= left2 {
		if record1[right1] == record2[right2] {
			right1--
			right2--
			find_right = true
		} else {
			if right1 == left1 || right2 == left2 {
				return false
			}
			break
		}
	}
	if !find_left || !find_right {
		return false
	}
	if left1 >= right1 || left2 >= right2 {
		return true
	}
	return false
}
