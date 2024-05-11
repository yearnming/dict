package rule

import (
	"fmt"
	"testing"
)

func TestFilterStringsByLength(t *testing.T) {
	key := &KeyWord{
		Surname:   []string{"li"},
		GivenName: []string{"si"},
		WeakPass:  []string{"123456", "qwer"},
		Connector: []string{"@", ".", "#"},
	}
	keywords := FieldsWithValues(key)
	//var rules [][]string
	//SelectFixedLengthPermutations(keywords, 3, []string{}, &rules)
	rules := GenRulesLength(keywords, 3, 3)
	rules = RemoveSlicesWithDuplicates(rules)
	fmt.Printf("规则长度: %d\n", len(rules))
	filepathrule := "output.txt"
	err := SaveCombinationsToFiles(rules, filepathrule)
	if err != nil {
		fmt.Println("保存组合规则失败:", err)
	} else {
		fmt.Println("组合保存到 ", filepathrule)
	}
	combinations := GenDir(key, rules)
	combinations = Deduplicate(combinations)
	fmt.Printf("字典数量: %d\n", len(combinations))

	combinations = FilterStringsByLength(combinations, -1, -1)
	fmt.Printf("限制长度后字典数量: %d\n", len(combinations))
	filepath := "output.txt"
	err = SaveCombinationsToFile(combinations, filepath)
	if err != nil {
		fmt.Println("保存字典失败:", err)
	} else {
		fmt.Println("字典保存到 ", filepath)
	}
}
