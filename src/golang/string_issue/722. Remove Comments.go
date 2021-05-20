package string_issue

import "strings"

//The line by line code is visualized as below:
///*Test program */
//int main()
//{
//  // variable declaration
//int a, b, c;
///* This is a test
//   multiline
//   comment for
//   testing */
//a = b + c;
//}
func RemoveComments(source []string) []string {
	var l int = len(source)
	var res []string
	var block bool = false
	var this_line string
	for i := 0;i < l;i++{
		if strings.Contains(source[i],"//"){
			if !block{
				start_pos := strings.Index(source[i],"//")
				start_pos2 := strings.Index(source[i],"/*")
				if start_pos2 == -1 || start_pos < start_pos2{
					if start_pos > 0{
						res = append(res,source[i][:start_pos])
					}
					continue
				}
			}
		}
		start_pos := strings.Index(source[i],"/*")
		if start_pos != -1{
			if !block{
				this_line += source[i][:start_pos]
				block = true
			}
		}

		end_pos := strings.LastIndex(source[i],"*/")
		if end_pos != -1 && start_pos != (end_pos - 1){
			if block{
				this_line += source[i][end_pos + 2:]
			}else{
				this_line = source[i]
			}
			block = false
		}else{
			if !block{
				this_line += source[i]
			}
		}
		if !block{
			if len(this_line) > 0{
				res = append(res,this_line)
			}
			this_line = ""
		}
	}
	return res
}