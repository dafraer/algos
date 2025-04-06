func convertDateToBinary(date string) string {
	y, _ := strconv.Atoi(string(date[0:4]))
	m, _ := strconv.Atoi(string(date[5:7]))
	d, _ := strconv.Atoi(string(date[8:]))
	return fmt.Sprintf("%s-%s-%s", strconv.FormatInt(int64(y), 2), strconv.FormatInt(int64(m), 2), strconv.FormatInt(int64(d), 2))
}
