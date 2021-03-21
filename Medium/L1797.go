package Medium

type AuthenticationManager struct {
	mp map[string]int
	life int
}


func Constructor(timeToLive int) AuthenticationManager {
	return AuthenticationManager{
		mp:   make(map[string]int),
		life: timeToLive,
	}
}


func (this *AuthenticationManager) Generate(tokenId string, currentTime int)  {
	this.mp[tokenId] = currentTime + this.life
}


func (this *AuthenticationManager) Renew(tokenId string, currentTime int)  {
	if _, ok := this.mp[tokenId]; ok {
		if this.mp[tokenId] > currentTime {
			this.mp[tokenId] = currentTime + this.life
		}
	}
}


func (this *AuthenticationManager) CountUnexpiredTokens(currentTime int) int {
	cnt := 0
	for _, v := range this.mp {
		if v > currentTime {
			cnt++
		}
	}
	
	return cnt
}


/**
 * Your AuthenticationManager object will be instantiated and called as such:
 * obj := Constructor(timeToLive);
 * obj.Generate(tokenId,currentTime);
 * obj.Renew(tokenId,currentTime);
 * param_3 := obj.CountUnexpiredTokens(currentTime);
 */
