package string_issue

import "strings"

func SimplifyPath(path string) string {
	path_list := strings.Split(path,"/")
	var new_path []string
	var upper int = 0
	for _,p := range path_list{
		if p == "" || p == "."{
			continue
		}
		if p == ".."{
			if len(new_path) > 0{
				new_path = new_path[:len(new_path) - 1]
			}else{
				upper++
			}
			continue
		}
		new_path = append(new_path,p)
	}
	new_path = new_path[upper:]
	return strings.Join(new_path,"/")
}