func findWords(words []string) []string {
	res := make([]string, 0, len(words))
	rows := []string{"qwertyuiop", "asdfghjkl", "zxcvbnm"}

	for i, w := range words {
		w = strings.ToLower(w)
	loop:
		for _, r := range rows {
			for _, c := range w {
				if !strings.Contains(r, string(c)) {
					continue loop
				}
			}
			res = append(res, words[i])
		}
	}
	return res
}