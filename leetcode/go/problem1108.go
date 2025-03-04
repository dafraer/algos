func defangIPaddr(address string) string {
	res := make([]byte, 0, len(address)+6)
	for i := 0; i < len(address); i++ {
		if address[i] == '.' {
			res = append(res, '[', address[i], ']')
		} else {
			res = append(res, address[i])
		}
	}
	return string(res)
}
