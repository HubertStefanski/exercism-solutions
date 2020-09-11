package encode

import (
	"fmt"
	"strings"
)

func RunLengthEncode(input string) string {
	if input != "" {
		var outString string

		stringSplit := strings.Split(input, "")

	one:
		for index, letter := range stringSplit {
			count := 0
		two:
			for i := 1; i <= len(stringSplit); i++ {
				if i >= 0 && (i+index) <= len(stringSplit) {
					if stringSplit[index+i] == letter {
						count++
						continue two
					} else if count == 0 {
						outString = fmt.Sprintf("%v%v", outString, letter)
						continue one
					}
					if count != 0 {
						outString = fmt.Sprintf("%v%v%v", outString, count, letter)
						continue one
					}
				}
			}
		}

		return outString
	}

	return ""
}

func RunLengthDecode(input string) string {
	return ""
}

func remove(slice []string, s int) []string {
	slice = nil
	return append(slice[:s], slice[s+1:]...)
}
