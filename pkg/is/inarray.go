package is

func InArrayString(arr []string, s ...string) bool {
	if len(arr) == 0 || len(s) == 0 {
		return false
	}
	for _, ar := range arr {
		for _, v := range s {
			if ar == v {
				return true
			}
		}
	}
	return false
}
