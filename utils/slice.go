package utils

func JoinIntSlice(s []any, sep string) string {
	var result = ""

	for i := range s {
		result += ToString(ToString(i))
		if len(s)-1 != i {
			result += sep
		}
	}

	return result
}
