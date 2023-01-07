package utils

func JoinIntSlice(sep string, s ...any) string {
	var result = ""

	for i, e := range s {
		result += ToString(e)
		if len(s)-1 != i {
			result += sep
		}
	}

	return result
}
