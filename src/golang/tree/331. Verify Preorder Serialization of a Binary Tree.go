package tree

import "strings"

//2 = leaf
//1 = half leaf
//0 = not leaf
//"9,3,4,#,#,1,#,#,2,#,6,#,#",
func IsValidSerialization(preorder string) bool {
	var record []string = strings.Split(preorder,",")
	var stack []string
	l := len(record)
	for i := 0;i < l;i++{
		len_s := len(stack)
		if record[i] == "#"{
			for len_s > 0 && stack[len_s - 1] == "#" {
				stack = stack[:len_s - 1]
				len_s = len(stack)
				if len_s == 0{
					return false
				}
				stack = stack[:len_s - 1]
				len_s = len(stack)
			}
		}
		stack = append(stack,record[i])
	}
	return len(stack) == 1 && stack[0] == "#"
}