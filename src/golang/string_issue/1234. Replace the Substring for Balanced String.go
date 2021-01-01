package string_issue

//Input: s = "QQQW"
//Output: 2
func BalancedString(s string) int {
	var l int = len(s)
	var target int = l / 4
	var record map[uint8]int = make(map[uint8]int)
	for i := 0;i < l;i++{
		record[s[i]]++
	}
	var need map[uint8]int = make(map[uint8]int)
	var exist bool = false
	for c,cnt := range record{
		if cnt > target{
			need[c] = cnt - target
			exist = true
		}
	}
	if !exist{
		return 0
	}
	var res int = l
	var begin int = 0
	var end int = 0
	for end < l{
		if _,ok := need[s[end]];ok{
			need[s[end]]--
		}
		var meet bool = true
		for _,cnt := range need{
			if cnt > 0{
				meet = false
				break
			}
		}
		if meet{
			if end - begin + 1 < res{
				res = end - begin + 1
			}
			for begin < end{
				if _,ok := need[s[begin]];ok{
					need[s[begin]]++
					if need[s[begin]] > 0{
						begin++
						break
					}
				}
				begin++
				if end - begin + 1 < res{
					res = end - begin + 1
				}
			}
		}
		end++
	}
	return res
}