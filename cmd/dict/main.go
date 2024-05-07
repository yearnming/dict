package main

import (
	"github.com/projectdiscovery/gologger"
	"github.com/yearnming/dict/runner"
)

func main() {

	// 接收参数
	options := runner.ParseOptions()

	// 创建一个生成器实例
	NewRunner, err := runner.NewRunner(options)
	if err != nil {
		gologger.Fatal().Msgf("创建生成器实例错误: %s\n", err)
	}

	// 运行生成器
	err = NewRunner.Run()
	if err != nil {
		gologger.Fatal().Msgf("运行生成器错误: %s\n", err)
	}
}
