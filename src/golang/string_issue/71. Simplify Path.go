package string_issue

import "strings"

//Input: "/a/../../b/../c//.//"
//Output: "/c"
func is_same(path string)bool{
	var l int = len(path)
	for i := 0;i < l;i++{
		if(path[i] != '/'){
			return false
		}
	}
	return true
}

func SimplifyPath(path string) string {
	var len_path int = len(path)
	if(len_path == 0){
		return path
	}
	path = strings.ReplaceAll(path,"//","/")
	if(is_same(path)){
		return "/"
	}
	if(path[0] == '/'){
		path = path[1:]
	}
	len_path = len(path)
	if(len_path == 0){
		return path
	}
	if(path[len_path - 1] == '/'){
		path = path[:len_path - 1]
	}
	var data []string = strings.Split(path,"/")
	var folder []string
	var l int = len(data)
	var res string
	for i := 0;i < l;i++{
		if(data[i] == ".."){
			if(len(folder) > 0){
				folder = append(folder[:len(folder) - 1],folder[len(folder):]...)
			}
		}else if(data[i] == "."){

		}else if(data[i] != ""){
			folder = append(folder,data[i])
		}
	}
	for i := 0;i < len(folder);i++ {
		res = res + "/" + folder[i]
	}
	if(res == ""){
		return "/"
	}
	return res
}


func OLD_SimplifyPath(path string) string {
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