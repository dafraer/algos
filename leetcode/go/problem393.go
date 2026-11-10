func validUtf8(data []int) bool {
	a := data
	i := 0
	for i < len(data) {
		if (1<<7)&a[i] == 0 {
			i++
			continue
		} else if (31<<3)&a[i] == (31 << 3) {
			return false
		} else if (30<<3)&a[i] == (30 << 3) {
			if i >= len(a)-3 {
				return false
			}
			for j := i + 1; j < i+4; j++ {
				if (2<<6)&a[j] != (2<<6) || (3<<6)&a[j] == (3<<6) {
					return false
				}
			}
			i += 4
		} else if (14<<4)&a[i] == (14 << 4) {
			if i >= len(a)-2 {
				return false
			}
			for j := i + 1; j < i+3; j++ {
				if (2<<6)&a[j] != (2<<6) || (3<<6)&a[j] == (3<<6) {
					return false
				}
			}
			i += 3
		} else if (6<<5)&a[i] == (6 << 5) {
			if i >= len(a)-1 || (2<<6)&a[i+1] != (2<<6) || (3<<6)&a[i+1] == (3<<6) {
				return false
			}
			i += 2
		} else {
			return false
		}
	}
	return true
}