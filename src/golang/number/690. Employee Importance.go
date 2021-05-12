package number


type Employee struct {
	Id int
	Importance int
	Subordinates []int
}

func recursive_getImportance(id int,data map[int]*Employee)int{
	if _,ok := data[id];!ok{
		return 0
	}
	var res int = data[id].Importance
	for _,subid := range data[id].Subordinates{
		res += recursive_getImportance(subid,data)
	}
	return res
}

func getImportance(employees []*Employee, id int) int {
	var data map[int]*Employee = make(map[int]*Employee)
	for _,e := range employees{
		data[e.Id] = e
	}
	return recursive_getImportance(id,data)
}