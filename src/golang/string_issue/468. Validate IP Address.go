package string_issue

import (
	"strconv"
	"strings"
)

//A valid IPv4 address is an IP in the form "x1.x2.x3.x4" where 0 <= xi <= 255 and xi cannot contain leading zeros.
//For example, "192.168.1.1" and "192.168.1.0" are valid IPv4 addresses but "192.168.01.1", "192.168.1.00" and "192.168@1.1" are invalid IPv4 adresses.
//
//A valid IPv6 address is an IP in the form "x1:x2:x3:x4:x5:x6:x7:x8" where:
//
//1 <= xi.length <= 4
//xi is hexadecimal string whcih may contain digits, lower-case English letter ('a' to 'f') and/or upper-case English letters ('A' to 'F').
//Leading zeros are allowed in xi.
func checkipv4(ipv4 []string)bool{
	for _,s := range ipv4{
		var l int = len(s)
		if l > 3 || l == 0{
			return false
		}
		if l > 1 && s[0] == '0'{
			return false
		}
		n,err := strconv.Atoi(s);
		if err != nil{
			return false
		}
		if n > 255 || n < 0{
			return false
		}
	}
	return true
}

func checkipv6(ipv6 []string)bool{
	for _,s := range ipv6{
		var l int = len(s)
		if l == 0 || l > 4{
			return false
		}
		for _,c := range s{
			if (c >= '0' && c <= '9') || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F'){
				continue
			}
			return false
		}
	}
	return true
}

func ValidIPAddress(IP string) string {
	var v4 []string = strings.Split(IP,".")
	var l int = len(v4)
	if l == 4{
		if checkipv4(v4){
			return "IPv4"
		}
	}else{
		var v6 []string = strings.Split(IP,":")
		var l6 int = len(v6)
		if l6 == 8 && checkipv6(v6){
			return "IPv6"
		}
	}
	return "Neither"
}