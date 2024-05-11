package keyword

import (
	"github.com/mozillazg/go-pinyin"
	"github.com/projectdiscovery/gologger"
	"strings"
)

type NamePinyin struct {
	Surname              string // Surname 姓氏
	GivenName            string // GivenName  名
	FirstLetterSurname   string // FirstLetterSurname 姓氏拼音首字母
	FirstLetterGivenName string // FirstLetterGivenName 名拼音首字母
	SurnameBa            string // SurnameBa 姓氏拼音首字母大写
	GivenNameBa          string // GivenNameBa  名拼音首字母大写
	SurnameBaF           string // SurnameBaF 姓氏拼音首字母大写
	GivenNameBaF         string // GivenNameBaF  名拼音首字母大写
}

// Pinyin 汉字转拼音 取首字母 和 全拼
func Pinyin(name string) *NamePinyin {
	//namepinyin := &NamePinyin{}
	a := pinyin.NewArgs()
	qp := pinyin.LazyConvert(name, nil)
	gologger.Info().Msgf("汉字转拼音: %s", qp)
	a.Style = pinyin.FirstLetter
	szm := pinyin.LazyPinyin(name, a)
	gologger.Info().Msgf("首字母风格，只返回拼音的首字母部分: %s", szm)

	SurnameBa := make([]string, len(qp))
	copy(SurnameBa, qp)
	SurnameBa = TitleCase(SurnameBa)
	SurnameBaF := make([]string, len(szm))
	copy(SurnameBaF, szm)
	SurnameBaF = TitleCase(SurnameBaF)
	namepinyin := &NamePinyin{
		Surname:              qp[0],
		GivenName:            strings.Join(qp[1:], ""),
		FirstLetterSurname:   szm[0],
		FirstLetterGivenName: strings.Join(szm[1:], ""),
		SurnameBa:            SurnameBa[0],
		GivenNameBa:          strings.Join(SurnameBa[1:], ""),
		SurnameBaF:           SurnameBaF[0],
		GivenNameBaF:         strings.Join(SurnameBaF[1:], ""),
	}
	return namepinyin
}

// TitleCase 将字符串转换为标题大小写
func TitleCase(words []string) []string {

	for i, word := range words {
		// 使用strings.ToUpper转换第一个字母为大写
		// 然后加上剩余的小写字母
		words[i] = strings.ToUpper(string(word[0])) + word[1:]
	}
	//gologger.Info().Msgf("首字母大写: %s", words)
	return words
}

// NameCase 将字符串转换为标题大小写
func NameCase(word string) string {
	if len(word) == 0 {
		return ""
	}
	if len(word) == 1 {
		word = strings.ToUpper(string(word[0]))
	} else {
		word = strings.ToUpper(string(word[0])) + word[1:]
	}

	return word
}
