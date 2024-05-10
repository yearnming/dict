package ruletype

import (
	"github.com/projectdiscovery/gologger"
	"github.com/yearnming/dict/pkg/rule"
	"testing"
)

func TestWifi(t *testing.T) {
	minLen, maxLen := 2, 3
	keywords := []string{"1", "2", "3", "4"}
	//var rules [][]string
	//for i := minLen; i <= maxLen; i++ {
	//	GenRules(keywords, i, []string{}, &rules)
	//}
	//gologger.Info().Msgf("GenRules 长度为: %d\n规则: %v", len(rules), rules)

	combinations := rule.GenRulesLength(keywords, minLen, maxLen)

	//var combinations [][]string
	//for i := minLen; i <= maxLen; i++ {
	//	rule.SelectFixedLengthPermutations(keywords, i, []string{}, &combinations)
	//}
	gologger.Info().Msgf("GenRulesLength 长度为: %d\n规则: %v", len(combinations), combinations)
}

func GenRules(keywords []string, length int, prefix []string, combinations *[][]string) {
	if len(prefix) == length {
		//fmt.Println(prefix)
		*combinations = append(*combinations, prefix)
		return
	}
	for _, keyword := range keywords {
		GenRules(keywords, length, append(prefix, keyword), combinations)
	}
}
