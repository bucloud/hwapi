package hwapi

func uniqueSlice(sl []string) []string {
	r := []string{}
	c := map[string]bool{}
	for _, v := range sl {
		if _, value := c[v]; !value {
			c[v] = true
			r = append(r, v)
		}
	}
	return r
}
