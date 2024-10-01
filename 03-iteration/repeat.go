package iteration

import "strings"

func Repeat(char string, count int) string {
	var builder strings.Builder
	for i := 0; i < count; i++ {
		builder.WriteString(char)
	}
	return builder.String()
}
