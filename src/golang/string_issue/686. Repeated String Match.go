package string_issue

import (
	"bytes"
	"strings"
)

func repeatedStringMatch(A string, B string) int {
	for i,_ := range B{
		if !strings.Contains(A,B[i:i+1]){
			return -1
		}
	}
	la := len(A)
	lb := len(B)
	cur_lena := la
	var buffer bytes.Buffer
	i := 1
	var find bool = false
	for ; i <= lb; i++ {
		buffer.WriteString(A)
		cur_lena += la
		if cur_lena < lb{
			continue
		}
		dupA := buffer.String()
		if strings.Contains(dupA,B){
			find = true
			break
		}
	}
	if !find{
		return -1
	}
	return i
}