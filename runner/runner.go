package runner

import (
	"github.com/projectdiscovery/gologger"
	"github.com/yearnming/dict/pkg/rule"
	"github.com/yearnming/dict/pkg/rule/ruletype"
)

type Runner struct {
	options *Options
}

// NewRunner 创建一个新的 Runner 实例，并且初始化它的 options 字段
func NewRunner(options *Options) (*Runner, error) {

	runner := &Runner{options: options}

	return runner, nil
}

// Run 执行 Runner 的逻辑
func (r *Runner) Run() error {
	key := &rule.KeyWord{
		Surname:              r.options.Surname,
		GivenName:            r.options.GivenName,
		FirstLetterSurname:   r.options.FirstLetterSurname,
		FirstLetterGivenName: r.options.FirstLetterGivenName,
		SurnameBa:            r.options.SurnameBa,
		GivenNameBa:          r.options.GivenNameBa,
		SurnameBaF:           r.options.SurnameBaF,
		GivenNameBaF:         r.options.GivenNameBaF,
		WeakPass:             r.options.WeakPass,
		Connector:            r.options.Connector,
		Year:                 r.options.Year,
		Company:              r.options.Company,
		Email:                r.options.Email,
		Phone:                r.options.Phone,
		City:                 r.options.City,
		Keyboard:             r.options.Keyboard,
	}
	keywords := rule.FieldsWithValues(key)
	gologger.Info().Msgf("关键字一共有%d个: %v\n", len(keywords), keywords)
	//var rules [][]string
	//rule.SelectFixedLengthPermutations(keywords, r.options.KeyWordLength, []string{}, &rules)
	rules := rule.GenRulesLength(keywords, r.options.KeyWordLengthMin, r.options.KeyWordLengthMax)
	rules = rule.RemoveSlicesWithDuplicates(rules) // 所有的子切片都不包含任何的重复元素
	// 以防万一出现 存在同一子切片的情况可使用下面一行代码
	//rules = rule.RemoveDuplicateSlices(rules)
	rules = ruletype.RuleName(rules)
	gologger.Info().Msgf("规则数量: %d\n", len(rules))
	err := rule.SaveCombinationsToFiles(rules, r.options.OutputRule)
	if err != nil {
		gologger.Error().Msgf("保存组合规则失败:%s", err)
	} else {
		gologger.Info().Msgf("组合保存到 %s", r.options.OutputRule)
	}
	combinations := rule.GenDir(key, rules)
	combinations = rule.Deduplicate(combinations)
	gologger.Info().Msgf("字典长度: %d\n", len(combinations))
	err = rule.SaveCombinationsToFile(combinations, r.options.Output)
	if err != nil {
		gologger.Error().Msgf("保存字典失败: %s", err)
	} else {
		gologger.Info().Msgf("字典保存到 %s", r.options.Output)
	}

	return nil
}
