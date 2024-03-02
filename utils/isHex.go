package utils

import (
	"fmt"
	"regexp"
)

func IsHexWithOrWithout0xPrefix(data string) bool {
	pattern := `^(0x)?[0-9a-fA-F]+$`
	matched, _ := regexp.MatchString(pattern, data)
	fmt.Println("is hex?:", matched)
	return matched
}
