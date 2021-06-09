package string_issue

import (
	"strconv"
	"strings"
)

//811
//Input:
//["900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"]
//Output:
//["901 mail.com","50 yahoo.com","900 google.mail.com","5 wiki.org","5 org","1 intel.mail.com","951 com"]
func SubdomainVisits(cpdomains []string) []string {
	var res []string
	var records map[string]int = make(map[string]int);
	for _,ss := range cpdomains{
		item := strings.Split(ss," ")
		cnt,ok := strconv.Atoi(item[0])
		if ok != nil{
			return res
		}
		domain := item[1]
		for len(domain) > 0{
			if _,ok := records[domain];ok{
				records[domain] += cnt
			}else{
				records[domain] = cnt
			}
			pos := strings.Index(domain,".")
			if pos == -1{
				break
			}
			domain = domain[pos+1:]
		}
	}
	for domain,cnt := range records{
		s := strconv.Itoa(cnt) + " " + domain
		res = append(res,s)
	}
	return res
}
