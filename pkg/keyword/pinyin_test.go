package keyword

import (
	"github.com/projectdiscovery/gologger"
	"testing"
)

func TestTitleCase(t *testing.T) {
	input := []string{"da", "zhang", "wei"}
	output := TitleCase(input[1:])
	gologger.Info().Msgf("首字母大写: %s", output) // 输出: Golang Is An Excellent Language

	//println(NameCase(input))
}

func TestPinyin(t *testing.T) {
	hans := "大张伟"
	namepinyin := Pinyin(hans)
	//gologger.Info().Msgf("%s", a)
	//gologger.Info().Msgf("%s", b)
	//namepinyin := &NamePinyin{
	//	Surname:              a[0],
	//	GivenName:            strings.Join(a[1:], ""),
	//	FirstLetterSurname:   b[0],
	//	FirstLetterGivenName: strings.Join(b[1:], ""),
	//}
	gologger.Info().Msgf("%+v", namepinyin.Surname)
	gologger.Info().Msgf("%+v", namepinyin)

	////
	////默认
	//a := pinyin.NewArgs()
	//gologger.Info().Msgf("%s", pinyin.Pinyin(hans, a))
	// [[zhong] [guo] [ren]]
	//
	//// 包含声调
	//a.Style = pinyin.Tone
	//gologger.Info().Msgf("%s", pinyin.Pinyin(hans, a))
	//// [[zhōng] [guó] [rén]]
	//
	//// 声调用数字表示
	//a.Style = pinyin.Tone2
	//gologger.Info().Msgf("%s", pinyin.Pinyin(hans, a))
	//// [[zho1ng] [guo2] [re2n]]

	// 开启多音字模式
	//a = pinyin.NewArgs()
	//a.Heteronym = true
	//gologger.Info().Msgf("%s", pinyin.Pinyin(hans, a))
	//// [[zhong zhong] [guo] [ren]]
	//a.Style = pinyin.Tone2
	//gologger.Info().Msgf("%s", pinyin.Pinyin(hans, a))
	//// [[zho1ng zho4ng] [guo2] [re2n]]
	//
	//gologger.Info().Msgf("%s", pinyin.LazyPinyin(hans, pinyin.NewArgs()))
	//// [zhong guo ren]
	//
	//gologger.Info().Msgf("%s", pinyin.Convert(hans, nil))
	//// [[zhong] [guo] [ren]]
	//
	//gologger.Info().Msgf("%s", pinyin.LazyConvert(hans, nil))
	//// [zhong guo ren]
	//
	////a.Style = pinyin.Initials
	////gologger.Info().Msgf("声母风格，只返回各个拼音的声母部分: %s", pinyin.Pinyin(hans, a))
	//
	//a.Style = pinyin.FirstLetter
	//gologger.Info().Msgf("首字母风格，只返回拼音的首字母部分: %s", pinyin.LazyPinyin(hans, a))
}
