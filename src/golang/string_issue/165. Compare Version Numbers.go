package string_issue

import (
	"strconv"
	"strings"
)

//If version1 > version2 return 1; if version1 < version2 return -1;otherwise return 0.
//Input: version1 = "1.0.1", version2 = "1"
//Output: 1
//Input: version1 = "1.01", version2 = "1.001"
//Output: 0
//Explanation: Ignoring leading zeroes, both “01” and “001" represent the same number “1”
//Input: version1 = "1.0", version2 = "1.0.0"
//Output: 0
//Explanation: The first version number does not have a third level revision number, which means its third level revision number is default to "0"
func CompareVersion(version1 string, version2 string) int {
	var v1 []string = strings.Split(version1,".")
	var v2 []string = strings.Split(version2,".")
	var l1 int = len(v1)
	var l2 int = len(v2)
	var pos1 int = 0
	var pos2 int = 0
	for pos1 < l1 && pos2 < l2{
		if n1,err1 := strconv.Atoi(v1[pos1]);err1 == nil{
			if n2,err2 := strconv.Atoi(v2[pos2]);err2 == nil{
				if n1 < n2 {
					return -1
				}else if n1 > n2{
					return 1
				}
			}
		}
		pos1++
		pos2++
	}
	for pos1 < l1{
		if n1,err1 := strconv.Atoi(v1[pos1]);err1 == nil{
			if n1 != 0{
				return 1
			}
		}
		pos1++
	}
	for pos2 < l2{
		if n2,err2 := strconv.Atoi(v2[pos2]);err2 == nil{
			if n2 != 0{
				return -1
			}
		}
		pos2++
	}
	return 0
}