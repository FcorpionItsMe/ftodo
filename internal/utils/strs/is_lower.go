package strs

import "unicode"

func IsLower(str string) bool {
	for _, v := range str {
		if v == '_' {
			continue
		}
		if !unicode.IsLower(v) {
			return false
		}
	}
	return true
}
