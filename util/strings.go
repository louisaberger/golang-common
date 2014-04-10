package util

import (
	"fmt"
	"strings"
	"time"
	"unicode"
)

// @return s with 'prefix' removed from beginning of string if present
func RemovePrefix(s, prefix string) string {
	if strings.HasPrefix(s, prefix) {
		return s[len(prefix):]
	}
	return s
}

// @return s with 'suffix' removed from end of string if present
func RemoveSuffix(s, suffix string) string {
	if strings.HasSuffix(s, suffix) {
		return s[:len(s)-len(suffix)]
	}
	return s
}

// @return s with its first letter in lower case.
func Untitle(s string) string {
	if s == "" {
		return s
	}
	a := []rune(s)
	a[0] = unicode.ToLower(a[0])
	return string(a)
}

func Truncate(input string, outputLength int) string {
	if len(input) <= outputLength {
		return input
	}
	return input[:outputLength]
}

func DateAsString(when time.Time, layout string) string {
	return when.Format(layout)
}

func Sprintln(message string, args ...interface{}) {
	fmt.Println(fmt.Sprintf(message, args...))
}
