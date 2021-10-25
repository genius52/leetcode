package string_issue

import (
	"strconv"
)

//输入：names = ["onepiece","onepiece(1)","onepiece(2)","onepiece(3)","onepiece"]
//输出：["onepiece","onepiece(1)","onepiece(2)","onepiece(3)","onepiece(4)"]
//解释：当创建最后一个文件夹时，最小的正有效 k 为 4 ，文件名变为 "onepiece(4)"。
func GetFolderNames(names []string) []string {
	var exist map[string]int = make(map[string]int)
	var l int = len(names)
	var res []string = make([]string,l)
	for i := 0;i < l;i++{
		if val,ok := exist[names[i]];!ok{
			res[i] = names[i]
			exist[names[i]] = 0
		}else{
			var k int = val + 1
			for true{
				name2 := names[i] + "(" + strconv.Itoa(k) + ")"
				if _,ok := exist[name2];!ok {
					res[i] = name2
					exist[names[i]] = k
					exist[name2] = 0
					break
				}
				k++
			}
		}
	}
	return res
}