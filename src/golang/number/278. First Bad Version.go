package number

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func isBadVersion(version int) bool{
	return true
}

func firstBadVersion(n int) int {
	var low int = 0
	var high int = n
	for low < high{
		mid := low + (high - low)/2
		if isBadVersion(mid){
			high = mid
		}else{
			low = mid + 1
		}
	}
	return low
}

//func firstBadVersion(n int) int {
//	var low int = 0
//	var high int = n
//	for low <= high{
//		mid := low + (high - low)/2
//		if isBadVersion(mid){
//			if !isBadVersion(mid-1){
//				return mid
//			}
//			high = mid - 1
//		}else{
//			if isBadVersion(mid + 1){
//				return mid + 1
//			}
//			low = mid + 1
//		}
//	}
//	return low
//}