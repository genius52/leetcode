package number

// che,xiang,queen
func MinMovesToCaptureTheQueen(a int, b int, c int, d int, e int, f int) int {
	if a == e { //che eat queen
		if c != e {
			return 1
		} else {
			che_xiang := abs_int(b - d)
			che_queen := abs_int(b - f)
			xiang_queen := abs_int(d - f)
			if che_xiang+xiang_queen != che_queen {
				return 1
			}
		}
	}
	if b == f {
		if d != f {
			return 1
		} else {
			che_xiang := abs_int(a - c)
			che_queen := abs_int(a - e)
			xiang_queen := abs_int(c - e)
			if che_xiang+xiang_queen != che_queen {
				return 1
			}
		}
	}
	d1 := abs_int(c - e)
	d2 := abs_int(d - f)
	if d1 == d2 {
		//check che
		var h_dir int = 0 //水平方向
		var v_dir int = 0 //垂直方向
		if c < e {
			v_dir = 1
		} else {
			v_dir = -1
		}
		if d < f {
			h_dir = 1
		} else {
			h_dir = -1
		}
		start_v := c + v_dir
		start_h := d + h_dir
		end_v := e
		end_h := f
		var conflict bool = false
		for start_v != end_v && start_h != end_h {
			if start_v == a && start_h == b {
				conflict = true
				break
			}
			start_v += v_dir
			start_h += h_dir
		}
		if !conflict {
			return 1
		}
	}
	return 2
}
