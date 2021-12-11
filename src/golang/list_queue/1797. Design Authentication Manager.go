package list_queue

type AuthenticationManager struct {
	ttl    int
	record map[string]int
}

func Constructor1797(timeToLive int) AuthenticationManager {
	var obj AuthenticationManager
	obj.ttl = timeToLive
	obj.record = make(map[string]int)
	return obj
}

func (this *AuthenticationManager) Generate(tokenId string, currentTime int) {
	this.record[tokenId] = currentTime + this.ttl
}

func (this *AuthenticationManager) Renew(tokenId string, currentTime int) {
	if time, ok := this.record[tokenId]; ok {
		if time > currentTime {
			this.record[tokenId] = currentTime + this.ttl
		}
	}
}

func (this *AuthenticationManager) CountUnexpiredTokens(currentTime int) int {
	var cnt int = 0
	for _, v := range this.record {
		if v > currentTime {
			cnt++
		}
	}
	return cnt
}
