package string_issue

import (
	"strings"
)

func ReverseVowels(s string) string {
	var data []string = strings.Split(s,"")
	var l int = len(data)
	var low int = 0
	var high int = l - 1
	for low <= high{
		for low < high && data[low] != "a" && data[low] != "e" && data[low] != "i" && data[low] != "o" && data[low] != "u" &&
			data[low] != "A" && data[low] != "E" && data[low] != "I" && data[low] != "O" && data[low] != "U"{
			low++
		}
		for low < high && data[high] != "a" && data[high] != "e" && data[high] != "i" && data[high] != "o" && data[high] != "u" &&
			data[high] != "A" && data[high] != "E" && data[high] != "I" && data[high] != "O" && data[high] != "U"{
			high--
		}
		if low < high{
			data[low],data[high] = data[high],data[low]
		}
		low++
		high--
	}
	res := strings.Join(data,"")
	return res
}