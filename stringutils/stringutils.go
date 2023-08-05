package stringutils

import (
	"bytes"
	"fmt"
)

// 拼接字符串
func appendString(prefix string, content string, suffix string) string {
	var stringBuilder bytes.Buffer
	stringBuilder.WriteString(prefix)
	stringBuilder.WriteString(content)
	stringBuilder.WriteString(suffix)
	return stringBuilder.String()
}

// 遍历字符串
func rangeString() {
	var str1 string = "hello"
	var str2 string = "hello,码神之路"
	// 遍历
	for i := 0; i < len(str1); i++ {
		fmt.Printf("ASCII: %c %d\n", str1[i], str1[i])
	}
	for _, s := range str1 {
		fmt.Printf("Unicode: %c %d\n", s, s)
	}

	// 中文只能用for...range遍历
	for _, s := range str2 {
		fmt.Printf("Unicode: %c %d\n", s, s)
	}
}
