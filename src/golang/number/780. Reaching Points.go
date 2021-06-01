package number


func reachingPoints(sx int, sy int, tx int, ty int) bool{
	if sx > tx || sy > ty{
		return false
	}
	for sx < tx && sy < ty{
		if tx > ty{
			tx %= ty
		}else{
			ty %= tx
		}
	}
	return (sx == tx && (ty - sy) % sx == 0) || (sy == ty && (tx - sx) % sy == 0)
}


//func dfs_reachingPoints(sx int64, sy int64, tx int64, ty int64)bool{
//	if(sx > tx || sy > ty){
//		return false
//	}
//	return dfs_reachingPoints(sx + sy,sy,tx,ty) || dfs_reachingPoints(sx,sx + sy,tx,ty)
//}
//
//func reachingPoints(sx int, sy int, tx int, ty int) bool {
//	return dfs_reachingPoints(int64(sx),int64(sy),int64(tx),int64(ty))
//}

//type point struct{
//	x int64
//	y int64
//}
//
//func reachingPoints(sx int, sy int, tx int, ty int) bool {
//	var q list.List
//	var p point
//	var targetx int64 = int64(tx)
//	var targety int64 = int64(ty)
//	p.x = int64(sx)
//	p.y = int64(sy)
//	q.PushBack(p)
//	for q.Len() > 0{
//		var l int = q.Len()
//		for i := 0;i < l;i++{
//			var cur point = q.Front().Value.(point)
//			q.Remove(q.Front())
//			if cur.x == targetx && cur.y == targety{
//				return true
//			}
//			var p1 point
//			p1.x = cur.x + cur.y
//			p1.y = cur.y
//			if p1.x <= targetx && p1.y <= targety{
//				q.PushBack(p1)
//			}
//			var p2 point
//			p2.x = cur.x
//			p2.y = cur.y + cur.x
//			if p2.x <= targetx && p2.y <= targety{
//				q.PushBack(p2)
//			}
//		}
//	}
//	return false
//}