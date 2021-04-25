package string_issue

import (
	"regexp"
	"strings"
)

//Input: paths = ["root/a 1.txt(abcd) 2.txt(efgh)","root/c 3.txt(abcd)","root/c/d 4.txt(efgh)",
//"root 4.txt(efgh)"]
//Output: [["root/a/2.txt","root/c/d/4.txt","root/4.txt"],["root/a/1.txt","root/c/3.txt"]]
func FindDuplicate(paths []string) [][]string {
	var l int = len(paths)
	var content_path map[string][]string = make(map[string][]string)
	r1 := regexp.MustCompile(`\(.*\)`)
	//r2 := regexp.MustCompile(`(.*)\(`)
	for i := 0;i < l;i++{
		var p []string = strings.Split(paths[i]," ")
		var root string = p[0]
		var sub_len int = len(p)
		for j := 1;j < sub_len;j++{
			content := r1.FindAllString(p[j], -1)
			abs_path := p[j][0:strings.IndexByte(p[j],'(')]
			full_path := root + "/" + abs_path
			content_path[content[0]] = append(content_path[content[0]],full_path)
		}
	}
	var res [][]string
	for _,c := range content_path{
		if len(c) > 1{
			res = append(res,c)
		}
	}
	return res
}