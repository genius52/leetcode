package diagram

//输入：recipes = ["bread","sandwich"], ingredients = [["yeast","flour"],["bread","meat"]],
//supplies = ["yeast","flour","meat"]
//输出：["bread","sandwich"]
func FindAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
	var exist map[string]bool = make(map[string]bool)
	for _, s := range supplies {
		exist[s] = true
	}
	var l int = len(recipes)
	var visited []bool = make([]bool, l)
	var res []string
	for true {
		var add bool = false
		for i := 0; i < l; i++ {
			if visited[i] {
				continue
			}
			var find bool = true
			for j := 0; j < len(ingredients[i]); j++ {
				if _, ok := exist[ingredients[i][j]]; !ok {
					find = false
					break
				}
			}
			if find {
				exist[recipes[i]] = true
				visited[i] = true
				add = true
				res = append(res, recipes[i])
			}
		}
		if !add {
			break
		}
	}
	return res
}
