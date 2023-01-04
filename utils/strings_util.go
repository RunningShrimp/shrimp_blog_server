package utils

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strings"
)

var title = cases.Title(language.English)

func ToUpperCamel(name string) (result string) {
	if strings.Contains(name, "_") {
		n := strings.Split(name, "_")
		for _, s := range n {
			if strings.Contains(s, "id") {
				result += strings.ToUpper(s)
				continue
			}

			result += title.String(s)
		}
	} else {
		if strings.Contains(name, "id") {
			result = strings.ToUpper(name)
		} else {
			result = title.String(name)
		}
	}
	return result
}

func ToLowerCamel(name string) (result string) {
	if strings.Contains(name, "_") {
		n := strings.Split(name, "_")
		result += n[0]
		for _, s := range n[1:] {
			result += title.String(s)
		}
	} else {
		result = name
	}
	return
}
