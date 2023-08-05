package stringutils

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

// AppendString 拼接字符串
func AppendString(prefix string, content string, suffix string) string {
	var stringBuilder bytes.Buffer
	stringBuilder.WriteString(prefix)
	stringBuilder.WriteString(content)
	stringBuilder.WriteString(suffix)
	return stringBuilder.String()
}

// RangeString 遍历字符串
func RangeString() {
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

// SearchString 查找目标字符串在目标字符串中位置
func SearchString(content string, pattern string) int {
	return strings.Index(content, pattern)
}

// Switch2Int 字符串转换为整型数字
func Switch2Int(content string) (int, error) {
	return strconv.Atoi(content)
}
