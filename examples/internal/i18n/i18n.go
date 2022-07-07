package i18n

import (
	"github.com/meilihao/goi18n/v2"
)

var (
	Arg = struct {
		Mobile *goi18n.Elem
		Parse  *goi18n.Elem
	}{
		Mobile: &goi18n.Elem{
			Key: "Arg.Mobile",
			Map: map[string]string{
				"zh": `{0}长度不等于11位或{1}格式错误!`,
				"en": `{0} length need 11 or {1} format invalid!`,
			},
		}, Parse: &goi18n.Elem{
			Key: "Arg.Parse",
			Map: map[string]string{
				"zh": `%s`,
				"en": `%s`,
			},
		},
	}

	Token = struct {
		Empty *goi18n.Elem
	}{
		Empty: &goi18n.Elem{
			Key: "Token.Empty",
			Map: map[string]string{
				"zh": `token为空`,
				"en": `token empty`,
			},
		},
	}
)

var (
	Mapper = map[string]string{

		"Arg.Mobile_zh": `{0}长度不等于11位或{1}格式错误!`,
		"Arg.Mobile_en": `{0} length need 11 or {1} format invalid!`,

		"Arg.Parse_zh": `%s`,
		"Arg.Parse_en": `%s`,

		"Token.Empty_zh": `token为空`,
		"Token.Empty_en": `token empty`,
	}
)
