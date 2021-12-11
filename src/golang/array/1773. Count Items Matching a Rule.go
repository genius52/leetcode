package array

//给你一个数组 items ，其中 items[i] = [typei, colori, namei] ，描述第 i 件物品的类型、颜色以及名称。
//
//另给你一条由两个字符串 ruleKey 和 ruleValue 表示的检索规则。
//
//如果第 i 件物品能满足下述条件之一，则认为该物品与给定的检索规则 匹配 ：
//
//ruleKey == "type" 且 ruleValue == typei 。
//ruleKey == "color" 且 ruleValue == colori 。
//ruleKey == "name" 且 ruleValue == namei 。

func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	var l int = len(items)
	var cnt int = 0
	var check_idx int = 0
	if ruleKey == "color" {
		check_idx = 1
	} else if ruleKey == "name" {
		check_idx = 2
	}
	for i := 0; i < l; i++ {
		if items[i][check_idx] == ruleValue {
			cnt++
		}
	}
	return cnt
}
