package diagram

//OrderedStream os= new OrderedStream(5);
//os.insert(3, "ccccc"); // Inserts (3, "ccccc"), returns [].
//os.insert(1, "aaaaa"); // Inserts (1, "aaaaa"), returns ["aaaaa"].
//os.insert(2, "bbbbb"); // Inserts (2, "bbbbb"), returns ["bbbbb", "ccccc"].
//os.insert(5, "eeeee"); // Inserts (5, "eeeee"), returns [].
//os.insert(4, "ddddd"); // Inserts (4, "ddddd"), returns ["ddddd", "eeeee"].
type OrderedStream struct {
	data []string
	pointer int
	last int
}

func Constructor1656(n int) OrderedStream {
	var obj OrderedStream
	obj.data = make([]string,n)
	obj.pointer = 0
	obj.last = 0
	return obj
}

func (this *OrderedStream) Insert(id int, value string) []string {
	this.data[id - 1] = value
	if id - 1 > this.last{
		this.last = id - 1
	}
	var res []string
	if this.pointer != id - 1{
		return res
	}
	res = append(res,this.data[this.pointer])
	var l int = len(this.data)
	var visit int = id
	for visit < l{
		if len(this.data[visit]) == 0{
			break
		}
		res = append(res,this.data[visit])
		visit++
	}
	this.pointer = visit
	return res
}