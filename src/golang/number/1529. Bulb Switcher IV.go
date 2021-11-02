package number

func minFlips2(s string) int{
	s = "0" + s
	var l int = len(s)
	var steps int = 0
	var left int = 0
	for left < l{
		var right int = left + 1
		for right < l && s[left] == s[right]{
			right++
		}
		if right == l{
			break
		}
		steps++
		left = right
	}
	return steps
}

func MinFlips(target string) int {
	var l int = len(target)
	var pos int = 0
	var find_one bool = false
	var blocks int = 0
	for(pos < l){
		if(!find_one){
			if(target[pos] == '1'){
				find_one = true
				blocks++
			}
			pos++
		}else{
			if(target[pos] != target[pos-1]){
				blocks++
			}
			pos++
		}
	}
	return blocks
}