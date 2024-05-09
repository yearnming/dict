package keyword

import (
	"github.com/mozillazg/go-pinyin"
	"strings"
)

type NamePinyin struct {
	Surname              string // Surname 姓氏
	GivenName            string // GivenName  名
	FirstLetterSurname   string // FirstLetterSurname 姓氏拼音首字母
	FirstLetterGivenName string // FirstLetterGivenName 名拼音首字母
	//FirstLetter          string // FirstLetter 拼音首字母
}

// Pinyin 汉字转拼音 取首字母 和 全拼
func Pinyin(name string) *NamePinyin {
	//namepinyin := &NamePinyin{}
	a := pinyin.NewArgs()
	qp := pinyin.LazyConvert(name, nil)
	//gologger.Info().Msgf("汉字转拼音: %s", pinyin.LazyConvert(name, nil))
	a.Style = pinyin.FirstLetter
	//gologger.Info().Msgf("首字母风格，只返回拼音的首字母部分: %s", pinyin.LazyPinyin(name, a))
	szm := pinyin.LazyPinyin(name, a)

	namepinyin := &NamePinyin{
		Surname:              qp[0],
		GivenName:            strings.Join(qp[1:], ""),
		FirstLetterSurname:   szm[0],
		FirstLetterGivenName: strings.Join(szm[1:], ""),
	}
	return namepinyin
}
