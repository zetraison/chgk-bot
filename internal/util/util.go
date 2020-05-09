package util

import "strings"

func CompareStrings(s1 string, s2 string) bool {
	s1 = strings.ToLower(s1)
	s2 = strings.ToLower(s2)

	if s1 == s2 {
		return true
	}

	//trim all special characters
	specialCharacters := ".,'\"-_!?:;+=()*[]@#$%^&"

	s1 = strings.TrimSpace(s1)
	s1 = strings.Trim(s1, specialCharacters)
	s2 = strings.TrimSpace(s2)
	s2 = strings.Trim(s2, specialCharacters)

	if s1 == s2 {
		return true
	}

	return false
}
