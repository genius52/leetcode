package string_issue

func NextGreatestLetter(letters []byte, target byte) byte {
	var low int = 0
	var high int = len(letters) - 1
	if target < letters[low] || target > letters[high]{
		return letters[low]
	}
	for low < high{
		var mid int = (low + high)/2
		if letters[mid] <= target{
			low = mid + 1
		}else{
			high = mid
		}
	}
	return letters[low]
}