package go_string

import (
	"testing"
)

func TestStringClass_DesensitizeMobile(t *testing.T) {
	result := StringUtilInstance.DesensitizeMobile(`18317042247`)
	if result != `183****2247` {
		t.Error(`18317042247 error`)
	}

	result1 := StringUtilInstance.DesensitizeMobile(`4042247`)
	if result1 != `40****47` {
		t.Error(`4042247 error`)
	}
}

func TestStringClass_DesensitizeEmail(t *testing.T) {
	result := StringUtilInstance.MustDesensitizeEmail(`pefish@qq.com`)
	if result != `pefi****@qq.com` {
		t.Error(`pefish@qq.com error`)
	}

	result1 := StringUtilInstance.MustDesensitizeEmail(`abc@qq.com`)
	if result1 != `abc****@qq.com` {
		t.Error(`abc@qq.com error`)
	}

	result2 := StringUtilInstance.MustDesensitizeEmail(`a@163.com`)
	if result2 != `a****@163.com` {
		t.Error(`a@163.com error`)
	}
}

func TestStringClass_SpanLeft(t *testing.T) {
	if StringUtilInstance.MustSpanLeft(`1234`, 6, `0`) != `001234` {
		t.Error()
	}
}
