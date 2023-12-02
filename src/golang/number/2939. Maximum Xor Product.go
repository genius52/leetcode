package number

// Given three integers a, b, and n, return the maximum value of (a XOR x) * (b XOR x) where 0 <= x < 2n.
func MaximumXorProduct(a int64, b int64, n int) int {
	var MOD int64 = 1e9 + 7
	for i := n - 1; i >= 0; i-- {
		if ((a | (1 << i)) == a) && ((b | (1 << i)) == b) { //都是1
			continue
		}
		if ((a | (1 << i)) != a) && ((b | (1 << i)) != b) { //都是0
			a |= 1 << i
			b |= 1 << i
		} else {
			if a|(1<<i) == a { //a的i位是1
				if a > b {
					a ^= 1 << i
					b |= 1 << i
				}
			}
			if b|(1<<i) == b { //b的i位是1
				if a < b {
					a |= 1 << i
					b ^= 1 << i
				}
			}
		}
	}
	a %= MOD
	b %= MOD
	return int(a*b) % int(MOD)
}
