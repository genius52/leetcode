package number

//Input: k = 7
//Output: 2
//Explanation: The Fibonacci numbers are: 1, 1, 2, 3, 5, 8, 13, ...
//For k = 7 we can use 2 + 5 = 7.
func recursive_findMinFibonacciNumbers(target int,record []int,h int)int{
	i := h
	for ;i >= 0;i-- {
		if record[i] < target{
			target = target - record[i]
			break
		}
		if record[i] == target{
			return 1
		}
	}
	return 1 + recursive_findMinFibonacciNumbers(target,record,i - 1)
}

func FindMinFibonacciNumbers(k int) int {
	var record []int
	var first int = 1
	var second int = 1
	for second <= k{
		if second == k{
			return 1
		}
		tmp := first + second
		first = second
		second = tmp
		record = append(record,first)
	}
	return recursive_findMinFibonacciNumbers(k,record,len(record) - 1)
}
