package number

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