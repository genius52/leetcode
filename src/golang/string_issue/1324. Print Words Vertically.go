package string_issue

import "strings"

//Input: s = "CONTEST IS COMING"
//Output: ["CIC","OSO","N M","T I","E N","S G","T"]
func PrintVertically(s string) []string {
	var record []string = strings.Split(s," ")//each word
	var word_len int = len(record)
	var rows int = 0
	for i := 0;i < word_len;i++{
		cur_len :=  len(record[i])
		if cur_len > rows{
			rows = cur_len
		}
	}
	var res []string = make([]string,rows)//vertical word
	for pos := 0;pos < rows;pos++{
		for j := 0;j < word_len;j++{
			if pos >= len(record[j]){
				res[pos] += " "
			}else{
				res[pos] += string(record[j][pos])
			}
		}
		res[pos] = strings.TrimRight(res[pos]," ")
	}
	return res
}