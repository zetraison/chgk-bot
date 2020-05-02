package main

import "strings"

func compareString(s1 string, s2 string) bool {
	s1 = strings.ToLower(s1)
	s2 = strings.ToLower(s2)

	if s1 == s2 {
		return true
	}

	//trim all special characters
	specialCharacters := []string{".", ",", "'", "\"", "-", "_", "!", "?", ":", ";", "+", "=", "(", ")", "*", "[", "]", "@", "#", "$", "%", "^", "&"}

	s1 = strings.TrimSpace(s1)
	for _, c := range specialCharacters {
		s1 = strings.Trim(s1, c)
	}
	s2 = strings.TrimSpace(s2)
	for _, c := range specialCharacters {
		s2 = strings.Trim(s2, c)
	}

	if s1 == s2 {
		return true
	}

	return false
}
