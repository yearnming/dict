package runner

import (
	"fmt"
	"github.com/projectdiscovery/goflags"
	"os"
)

type Options struct {
	DictType string
	KeyWord  string
	Rules    string
	Length   string
	Output   string
}

// ParseOptions 解析应用程序的命令行选项
func ParseOptions() *Options {
	options := &Options{}

	flagSet := goflags.NewFlagSet()
	flagSet.SetDescription(`dict 是一个快速生成字典工具。`)

	flagSet.CreateGroup("input", "输入",
		flagSet.StringVarP(&options.DictType, "dicttype", "dt", "wifi", "输入字典类型"),
		flagSet.StringVarP(&options.KeyWord, "keyword", "key", "", "输入关键词"),
		flagSet.StringVarP(&options.Rules, "rule", "r", "", "组合规则"),
		flagSet.StringVarP(&options.Length, "length", "l", "", "长度限制"),
	)
	flagSet.CreateGroup("output", "输出",
		flagSet.StringVarP(&options.Output, "output", "o", "", "保存文件"),
	)
	flagSet.SetCustomHelpText("使用示例:\ngo run cmd/dict/main.go -dicttype wifi -keyword qwer -rule 3 -length 8-10")

	if err := flagSet.Parse(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return options
}
