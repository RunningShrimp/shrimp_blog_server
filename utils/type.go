package utils

import (
	"fmt"
	"time"
)

func ToString(e any) string {
	switch e.(type) {
	case time.Time:
		t := e.(time.Time)
		return t.Format("2006-01-02 15:04:05")
	default:
		return fmt.Sprintf("%v", e)
	}
}
