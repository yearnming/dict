package runner

import (
	"github.com/projectdiscovery/gologger"
	"github.com/yearnming/dict/pkg/rule"
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
	gologger.Info().Msgf("关键字数量: %v\n", keywords)
	var rules [][]string
	rule.SelectFixedLengthPermutations(keywords, 3, []string{}, &rules)
	rules = rule.RemoveSlicesWithDuplicates(rules)
	gologger.Info().Msgf("规则数量: %d\n", len(rules))
	err := rule.SaveCombinationsToFiles(rules)
	if err != nil {
		gologger.Error().Msgf("保存组合规则失败:%s", err)
	} else {
		gologger.Info().Msgf("组合保存到 rules.txt")
	}
	combinations := rule.GenDir(key, rules)
	combinations = rule.Deduplicate(combinations)
	gologger.Info().Msgf("字典长度: %d\n", len(combinations))
	err = rule.SaveCombinationsToFile(combinations)
	if err != nil {
		gologger.Error().Msgf("保存字典失败: %s", err)
	} else {
		gologger.Info().Msgf("字典保存到 字典.txt")
	}

	return nil
}
