package word

import (
	"strings"
	"unicode"
)

// 单词转大写
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// 单词转小写
func ToLower(s string) string {
	return strings.ToLower(s)
}

// 下划线单词转首字母大写驼峰单词
func UndersireToUpperCamelCase(s string) string {
	// -1表示更换的次数没有限制即替换所有
	s = strings.Replace(s, "_", " ", -1)
	s = strings.Title(s)
	return strings.Replace(s, " ", "", -1)
}

// 下划线单词转首字母小写驼峰单词
func UnderscoreToLowerCamelCase(s string) string {
	s = UndersireToUpperCamelCase(s)
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}

// 驼峰单词转下划线单词
func CamelCaseToUnderscore(s string) string {
	var output []rune
	for i, r := range s {
		if i == 0 {
			output = append(output, unicode.ToLower(r))
			continue
		}
		if unicode.IsUpper(r) {
			output = append(output, '_')
		}
		output = append(output, unicode.ToLower(r))
	}
	return string(output)
}
