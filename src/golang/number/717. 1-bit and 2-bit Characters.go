package number

func dp_isOneBitCharacter(bits []int,l int,pos int)bool{
	if pos >= l{
		return false
	}
	if pos == l - 1{
		return true
	}
	if bits[pos] == 1{
		return dp_isOneBitCharacter(bits,l,pos + 2)
	}else{
		return dp_isOneBitCharacter(bits,l,pos + 1)
	}
}

func isOneBitCharacter(bits []int) bool {
	var l int = len(bits)
	if bits[l - 1] == 1{
		return true
	}
	return dp_isOneBitCharacter(bits,l,0)
}