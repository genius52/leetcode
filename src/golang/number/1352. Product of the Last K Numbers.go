package number

type ProductOfNumbers struct {
	prefix []int
}

func Constructor() ProductOfNumbers {
	var obj ProductOfNumbers
	obj.prefix = []int{1}
	return obj
}

//Adds the number num to the back of the current list of numbers.
func (this *ProductOfNumbers) Add(num int)  {
	var l int = len(this.prefix)
	if num == 0{
		this.prefix = []int{1}
	}else{
		this.prefix = append(this.prefix,this.prefix[l - 1] * num)
	}
}

//Returns the product of the last k numbers in the current list.
//You can assume that always the current list has at least k numbers.
func (this *ProductOfNumbers) GetProduct(k int) int {
	var l int = len(this.prefix)
	if k >= l{
		return 0
	}
	return this.prefix[l - 1]/this.prefix[l - 1 - k]
}