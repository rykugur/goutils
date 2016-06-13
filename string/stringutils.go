package stringutils

import (
	"regexp"
)

func Matches(pattern string, input string) (bool, error) {
	matched, err := regexp.MatchString(pattern, input)
	if err != nil {
		return false, err
	}

	return matched, nil
}
