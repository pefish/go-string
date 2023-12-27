package go_string

import (
	"github.com/pkg/errors"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type StringUtil struct {
}

var StringUtilInstance = StringUtil{}

func (su *StringUtil) Desensitize(str string) string {
	index := strings.Index(str, `@`)
	if index == -1 {
		return su.DesensitizeMobile(str)
	} else {
		return su.MustDesensitizeEmail(str)
	}
}

/*
>7        前3后4中间4个*
<=7 && >4 前2后2中间4个*
<=4 && >2 前1后1中间2个*
*/
func (su *StringUtil) DesensitizeMobile(str string) string {
	result := ``
	length := len(str)
	if length > 7 {
		result = str[:3] + `****` + str[length-4:]
	} else if length <= 7 && length > 4 {
		result = str[:2] + `****` + str[length-2:]
	} else if length <= 4 && length > 2 {
		result = str[:1] + `**` + str[length-1:]
	} else {
		result = "*"
	}
	return result
}

func (su *StringUtil) MustDesensitizeEmail(str string) string {
	result, err := su.DesensitizeEmail(str)
	if err != nil {
		panic(err)
	}
	return result
}

/*
@前字符串长度>3   前4 中4个* 后@后面所有
@前字符串长度<=3  前@前面所有 中4个* 后@后面所有
*/
func (su *StringUtil) DesensitizeEmail(str string) (string, error) {
	result := ``
	index := strings.Index(str, `@`)
	if index == -1 {
		return "", errors.New(`Not email.`)
	}
	preAt := str[:index]
	if len(preAt) > 3 {
		result = str[:4] + `****` + str[index:]
	} else {
		result = preAt + `****` + str[index:]
	}
	return result, nil
}

func (su *StringUtil) RemoveLast(str string, num int) string {
	return str[:len(str)-num]
}

func (su *StringUtil) RemoveFirst(str string, num int) string {
	return str[num:]
}

func (su *StringUtil) Reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func (su *StringUtil) ReplaceAll(str string, oldStr string, newStr string) (result string) {
	return strings.Replace(str, oldStr, newStr, -1)
}

func (su *StringUtil) MustSpanLeft(str string, length int, fillChar string) string {
	result, err := su.SpanLeft(str, length, fillChar)
	if err != nil {
		panic(err)
	}
	return result
}

func (su *StringUtil) SpanLeft(str string, length int, fillChar string) (string, error) {
	if len(str) > length {
		return "", errors.New(`Length is too small.`)
	}
	if len(fillChar) != 1 {
		return "", errors.New(`Length of fillChar must be 1.`)
	}
	result := ``
	for i := 0; i < length-len(str); i++ {
		result += fillChar
	}
	return result + str, nil
}

func (su *StringUtil) MustSpanRight(str string, length int, fillChar string) string {
	result, err := su.SpanRight(str, length, fillChar)
	if err != nil {
		panic(err)
	}
	return result
}

func (su *StringUtil) SpanRight(str string, length int, fillChar string) (string, error) {
	if len(str) > length {
		return "", errors.New(`Length is too small.`)
	}
	if len(fillChar) != 1 {
		return "", errors.New(`Length of fillChar must be 1.`)
	}
	result := str
	for i := 0; i < length-len(str); i++ {
		result += fillChar
	}
	return result, nil
}

func (su *StringUtil) StartWith(str string, substr string) bool {
	return strings.HasPrefix(str, substr)
}

func (su *StringUtil) EndWith(str string, substr string) bool {
	return strings.HasSuffix(str, substr)
}

func (su *StringUtil) UserIdToInviteCode(userId uint64, length int) (string, error) {
	userIdStr := strconv.FormatUint(userId, 10)
	if len(userIdStr) > length {
		return "", errors.New("Length is too small.")
	}
	result := ""
	for _, a := range userIdStr {
		result += string(a + 20)
	}
	if len(userIdStr) != length {
		r, err := randomStringFromDic("ABCDEFGHIJKLMNOPQRSTUVWXYZ", int32(length-len(userIdStr)))
		if err != nil {
			return "", err
		}
		result += r
	}
	return result, nil
}

func randomStringFromDic(dictionary string, count int32) (string, error) {
	b := make([]byte, count)
	l := len(dictionary)

	_, err := rand.New(rand.NewSource(time.Now().UnixNano())).Read(b)

	if err != nil {
		return "", err
	}
	for i, v := range b {
		b[i] = dictionary[v%byte(l)]
	}

	return string(b), nil
}
