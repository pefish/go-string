package p_string

import (
	"errors"
	"strings"
)

type StringClass struct {
}

var String = StringClass{}

func (this *StringClass) RemoveLast(str string, num int) string {
	return str[:len(str)-num]
}

func (this *StringClass) RemoveFirst(str string, num int) string {
	return str[num:]
}

func (this *StringClass) Reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func (this *StringClass) ReplaceAll(str string, oldStr string, newStr string) (result string) {
	return strings.Replace(str, oldStr, newStr, -1)
}

func (this *StringClass) SpanLeft(str string, length int, fillChar string) string {
	if len(str) > length {
		panic(errors.New(`length is too small`))
	}
	if len(fillChar) != 1 {
		panic(errors.New(`length of fillChar must be 1`))
	}
	result := ``
	for i := 0; i < length-len(str); i++ {
		result += fillChar
	}
	return result + str
}

func (this *StringClass) SpanRight(str string, length int, fillChar string) string {
	if len(str) > length {
		panic(errors.New(`length is too small`))
	}
	if len(fillChar) != 1 {
		panic(errors.New(`length of fillChar must be 1`))
	}
	result := str
	for i := 0; i < length-len(str); i++ {
		result += fillChar
	}
	return result
}

func (this *StringClass) StartWith(str string, substr string) bool {
	return strings.HasPrefix(str, substr)
}

func (this *StringClass) EndWith(str string, substr string) bool {
	return strings.HasSuffix(str, substr)
}
