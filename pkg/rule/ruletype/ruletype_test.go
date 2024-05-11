package ruletype

import (
	"fmt"
	"github.com/projectdiscovery/gologger"
	"github.com/yearnming/dict/pkg/keyword"
	"github.com/yearnming/dict/pkg/rule"
	"github.com/yearnming/dict/runner"
	"testing"
)

//	Surname              []string // Surname 姓氏
//	GivenName            []string // GivenName  名
//	FirstLetterSurname   []string // FirstLetterSurname 姓氏拼音首字母
//	FirstLetterGivenName []string // FirstLetterGivenName 名拼音首字母
//	Connector            []string // Connector   拼接字符
//	WeakPass             []string // WeakPass 弱口令
//	Year                 []string // Year 年份
//	Company              []string // Company 公司
//	Email                []string // Email 邮箱
//	Phone                []string // Phone 电话
//	City                 []string // City 城市
//	Keyboard             []string // Keyboard 键盘弱密码

func TestKwl(t *testing.T) {
	minLen, maxLen := 4, 4
	//keywords := []string{"姓氏", "名", "姓氏拼音首字母", "名拼音首字母", "拼接字符", "弱口令", "年份", "公司", "邮箱", "电话", "城市", "键盘"}
	keywords := []string{"姓氏", "名", "公司", "邮箱", "电话", "城市", "键盘"}

	//var combinations [][]string
	//SelectFixedLengthPermutations(keywords, 4, []string{}, &combinations)
	//rule.SelectFixedLengthPermutations(keywords, 4, []string{}, &combinations)
	combinations := rule.GenRulesLength(keywords, minLen, maxLen)
	gologger.Info().Msgf("去重前长度为: %d\n", len(combinations))
	combinations = rule.RemoveSlicesWithDuplicates(combinations)
	gologger.Info().Msgf("清理子切片中存在的重复元素后的长度 : %d\n", len(combinations))
	combinations = rule.RemoveDuplicateSlices(combinations)
	gologger.Info().Msgf("移除存在重复的子切片后的长度: %d\n", len(combinations))
	filepathrule := "WIFI规则.txt"
	err := rule.SaveCombinationsToFiles(combinations, filepathrule)
	if err != nil {
		fmt.Println("保存组合规则失败:", err)
	} else {
		fmt.Println("组合保存到 ", filepathrule)
	}
}

func TestRuleName(t *testing.T) {
	minLen, maxLen := 2, 3
	hans := "大张伟"

	namepinyin := keyword.Pinyin(hans)
	r := &runner.Options{}
	r.Surname = append(r.Surname, namepinyin.Surname)
	r.GivenName = append(r.GivenName, namepinyin.GivenName)
	r.FirstLetterSurname = append(r.FirstLetterSurname, namepinyin.FirstLetterSurname)
	r.FirstLetterGivenName = append(r.FirstLetterGivenName, namepinyin.FirstLetterGivenName)

	key := &rule.KeyWord{
		Surname:              r.Surname,
		GivenName:            r.GivenName,
		FirstLetterSurname:   r.FirstLetterSurname,
		FirstLetterGivenName: r.FirstLetterGivenName,
		WeakPass:             r.WeakPass,
		Connector:            r.Connector,
		Year:                 r.Year,
		Company:              r.Company,
		Email:                r.Email,
		Phone:                r.Phone,
		City:                 r.City,
		Keyboard:             r.Keyboard,
	}

	keywords := rule.FieldsWithValues(key)

	combinations := rule.GenRulesLength(keywords, minLen, maxLen)
	gologger.Info().Msgf("去重前长度为: %d\n", len(combinations))
	combinations = rule.RemoveSlicesWithDuplicates(combinations)
	gologger.Info().Msgf("清理子切片中存在的重复元素后的长度 : %d\n", len(combinations))
	combinations = rule.RemoveDuplicateSlices(combinations)
	gologger.Info().Msgf("移除存在重复的子切片后的长度: %d\n", len(combinations))
	combinations = RuleName(combinations)
	gologger.Info().Msgf("添加规则名称后的长度: %d\n", len(combinations))
	filepathrule := "WIFI规则.txt"
	err := rule.SaveCombinationsToFiles(combinations, filepathrule)
	if err != nil {
		fmt.Println("保存组合规则失败:", err)
	} else {
		fmt.Println("组合保存到 ", filepathrule)
	}

}
