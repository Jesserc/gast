package hex

import (
	"regexp"
)

// WithOrWithout0xPrefix checks if a string is hex with or without `0x` prefix using regular expression: `^(0x)?[0-9a-fA-F]+$`
func WithOrWithout0xPrefix(data string) bool {
	pattern := `^(0x)?[0-9a-fA-F]+$`
	matched, _ := regexp.MatchString(pattern, data)
	return matched
}
