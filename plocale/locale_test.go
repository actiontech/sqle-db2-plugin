package plocale

import (
	"testing"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func TestBundle(t *testing.T) {
	if Bundle == nil {
		t.Fatal("Bundle should not be nil")
	}

	cases := []struct {
		name     string
		lang     language.Tag
		msg      *i18n.Message
		expected string
	}{
		{
			name:     "Chinese rule description",
			lang:     language.Chinese,
			msg:      DB2DQL1Desc,
			expected: "禁止不带where条件的查询",
		},
		{
			name:     "English rule description",
			lang:     language.English,
			msg:      DB2DQL1Desc,
			expected: "Prohibit queries without WHERE clause",
		},
		{
			name:     "Chinese rule category",
			lang:     language.Chinese,
			msg:      DB2RuleCategorySQLConvention,
			expected: "SQL规范",
		},
		{
			name:     "English rule category",
			lang:     language.English,
			msg:      DB2RuleCategorySQLConvention,
			expected: "SQL convention",
		},
		{
			name:     "Chinese audit message",
			lang:     language.Chinese,
			msg:      DB2DQL1Message,
			expected: "禁止不带where条件的查询",
		},
		{
			name:     "English audit message",
			lang:     language.English,
			msg:      DB2DQL1Message,
			expected: "Prohibit queries without WHERE clause",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := Bundle.LocalizeMsgByLang(tc.lang, tc.msg)
			if got != tc.expected {
				t.Errorf("expected %q, got %q", tc.expected, got)
			}
		})
	}
}
