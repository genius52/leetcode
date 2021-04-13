package string_issue

import "strings"

func findWords(words []string) []string {
	var char_pos map[string]int = make(map[string]int)
	char_pos["q"] = 1
	char_pos["w"] = 1
	char_pos["e"] = 1
	char_pos["r"] = 1
	char_pos["t"] = 1
	char_pos["y"] = 1
	char_pos["u"] = 1
	char_pos["i"] = 1
	char_pos["o"] = 1
	char_pos["p"] = 1

	char_pos["a"] = 2
	char_pos["s"] = 2
	char_pos["d"] = 2
	char_pos["f"] = 2
	char_pos["g"] = 2
	char_pos["h"] = 2
	char_pos["j"] = 2
	char_pos["k"] = 2
	char_pos["l"] = 2

	char_pos["z"] = 3
	char_pos["x"] = 3
	char_pos["c"] = 3
	char_pos["v"] = 3
	char_pos["b"] = 3
	char_pos["n"] = 3
	char_pos["m"] = 3

	var res []string

	for _,str := range words{
		var last_pos int = -1
		var equal bool = true
		s := strings.ToLower(str)
		for i := 0;i < len(s);i++{
			ele := string(s[i])
			if(last_pos == -1){
				last_pos = char_pos[ele]
			}else{
				if(last_pos != char_pos[ele]){
					equal = false
					break
				}
			}
		}
		if(equal){
			res = append(res, str)
		}
	}
	return res
}