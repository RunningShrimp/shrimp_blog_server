package utils

import "strconv"

func JoinIntSlice(s []int, sep string) string {
	var result = ""

	for i := range s {
		result += strconv.Itoa(s[i])
		if len(s)-1 != i {
			result += sep
		}
	}

	return result
}
