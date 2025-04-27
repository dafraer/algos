func addBinary(a string, b string) string {
	next := 0
	if len(b) > len(a) {
		a, b = b, a
	}
	ans := make([]byte, len(a)+1)
	for i := 1; i <= len(a); i++ {
		an := 0
		bn := 0
		if a[len(a)-i] == '1' {
			an = 1
		}
		if len(b)-i >= 0 && b[len(b)-i] == '1' {
			bn = 1
		}
		switch an + bn + next {
		case 3:
			ans[len(ans)-i] = '1'
			next = 1
		case 2:
			ans[len(ans)-i] = '0'
			next = 1
		case 1:
			ans[len(ans)-i] = '1'
			next = 0
		case 0:
			ans[len(ans)-i] = '0'
			next = 0
		}
	}
	if next == 1 {
		ans[0] = '1'
		return string(ans)
	}
	return string(ans[1:])
}
