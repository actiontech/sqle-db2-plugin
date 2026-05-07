package plocale

import "github.com/nicksnyder/go-i18n/v2/i18n"

// 规则描述
var (
	DB2DQL1Desc = &i18n.Message{ID: "DB2DQL1Desc", Other: "禁止不带where条件的查询"}
)

// 规则分类
var (
	DB2RuleCategorySQLConvention = &i18n.Message{ID: "DB2RuleCategorySQLConvention", Other: "SQL规范"}
)

// 审核结果消息
var (
	DB2DQL1Message = &i18n.Message{ID: "DB2DQL1Message", Other: "禁止不带where条件的查询"}
)
