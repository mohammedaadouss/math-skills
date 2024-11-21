package dvma

import "strings"

func NameChecker(fileName string) bool {
	if strings.HasSuffix(fileName, ".txt") && strings.HasPrefix(fileName, "data") && len(fileName) == 8 {
		return true
	}
	return false
}
