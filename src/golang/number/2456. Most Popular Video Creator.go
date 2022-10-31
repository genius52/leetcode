package number

import "sort"

//输入：creators = ["alice","bob","alice","chris"], ids = ["one","two","three","four"], views = [5,10,5,4]
//输出：[["alice","one"],["bob","two"]]
type id_view struct{
	id string
	view int
}

func mostPopularCreator(creators []string, ids []string, views []int) [][]string {
	var l int = len(creators)
	var creator_cnt map[string]int = make(map[string]int)
	var creator_ids map[string]id_view = make(map[string]id_view)
	var max_cnt int = 0
	for i := 0;i < l;i++{
		creator_cnt[creators[i]] += views[i]
		if creator_cnt[creators[i]] > max_cnt{
			max_cnt = creator_cnt[creators[i]]
		}
		//creator_ids[creators[i]] = append(creator_ids[creators[i]],ids[i])
		if pre,ok := creator_ids[creators[i]];ok{
			//creator_ids[creators[i]] = make(map[string]int)
			if views[i] > pre.view{
				creator_ids[creators[i]] = id_view{ids[i],views[i]}
			}else if views[i] == pre.view{
				if ids[i] < pre.id{
					creator_ids[creators[i]] = id_view{ids[i],views[i]}
				}
			}
		}else{
			creator_ids[creators[i]] = id_view{ids[i],views[i]}
		}
	}
	var res [][]string
	for k,v := range creator_cnt{
		if v == max_cnt{
			res = append(res,[]string{k,creator_ids[k].id})
		}
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i][0] < res[j][0]
	})
	return res
}