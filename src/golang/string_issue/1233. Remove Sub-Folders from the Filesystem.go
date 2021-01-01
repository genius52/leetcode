package string_issue

import "sort"

//Input: folder = ["/a","/a/b/c","/a/b/d"]
//Output: ["/a"]
func exist_removeSubfolders(folder string,record map[string]bool)bool{
	var l int = len(folder)
	var i int = 0
	var tag_cnt int = 0
	for i < l{
		if folder[i] == '/'{
			tag_cnt++
			if tag_cnt > 1{
				sub := folder[:i]
				if _,ok := record[sub];ok{
					return true
				}
			}
		}
		i++
	}
	return false
}

func RemoveSubfolders(folder []string) []string {
	var record map[string]bool = make(map[string]bool)
	sort.Strings(folder)
	record[folder[0]] = true
	var res []string = []string{folder[0]}
	var l int = len(folder)
	for i := 1;i < l;i++{
		if !exist_removeSubfolders(folder[i],record){
			res = append(res,folder[i])
		}
		record[folder[i]] = true
	}
	return res
}