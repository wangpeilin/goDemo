package logic

import "strconv"

// IsValidInput 验证输入的字符串是否为有效的整数
func IsValidInput(a, b string) bool {
	_, err1 := strconv.Atoi(a)
	_, err2 := strconv.Atoi(b)
	return err1 == nil && err2 == nil
}
