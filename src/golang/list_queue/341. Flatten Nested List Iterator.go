package list_queue


type NestedInteger struct {
}

func (this NestedInteger) IsInteger() bool {return true}

func (this NestedInteger) GetInteger() int {return 0}

func (n *NestedInteger) SetInteger(value int) {}

func (this *NestedInteger) Add(elem NestedInteger) {}

func (this NestedInteger) GetList() []*NestedInteger {return []*NestedInteger{}}


type NestedIterator struct {
	data []int
}

func dfs(nestedList []*NestedInteger,data *[]int){
	for  i := 0;i < len(nestedList);i++{
		if nestedList[i].IsInteger(){
			*data = append(*data,nestedList[i].GetInteger())
		}else{
		 	childlist := nestedList[i].GetList()
			dfs(childlist,data)
		}
	}
}

func Constructor341(nestedList []*NestedInteger) *NestedIterator {
	var obj NestedIterator
	dfs(nestedList,&obj.data);
	return &obj
}

func (this *NestedIterator) Next() int {
	res := this.data[0]
	this.data = this.data[1:]
	return res
}

func (this *NestedIterator) HasNext() bool {
	return len(this.data) > 0
}