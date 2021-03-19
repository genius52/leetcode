package number

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guess(num int) int{
	return 0
}

func guessNumber(n int) int {
	var low int = 1
	var high int = n
	for low <= high{
		var mid int = low + (high - low)/2
		ret := guess(mid)
		if ret == 0{
			return mid
		}else if ret == 1{
			low = mid + 1
		}else if ret == -1{
			high = mid - 1
		}
	}
	return low
}