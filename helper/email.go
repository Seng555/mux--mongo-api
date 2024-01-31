package helper

import (
	"regexp"
)

func IsValidEmail(email string) bool {
	// Use a regular expression to validate email format
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(emailRegex, email)
	return match
}
