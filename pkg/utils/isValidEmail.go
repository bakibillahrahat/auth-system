package utils

import "strings"

func IsVAlidEmail(email string) bool {
	// must contain exactly one "@"
	if strings.Count(email, "@") != 1 {
		return false
	}

	parts := strings.Split(email, "@")
	local := parts[0]
	domain := parts[1]

	// local and domain must not be empty
	if local == "" || domain == "" {
		return false
	}

	// domain must contain at least one "."
	if !strings.Contains(domain, "."){
		return false
	}

	// domain should not start or end with "."
	if strings.HasPrefix(domain, ".") || strings.HasSuffix(domain, ".") {
		return false
	}
	return true
}