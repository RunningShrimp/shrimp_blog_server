package utils

import "testing"

func TestToString(t *testing.T) {
	var n uint
	n = 1
	t.Log(ToString(n))
}
