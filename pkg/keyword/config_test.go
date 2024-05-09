package keyword

import (
	"fmt"
	"github.com/projectdiscovery/gologger"
	folderutil "github.com/projectdiscovery/utils/folder"
	"github.com/yearnming/dict/pkg/rule"
	"testing"
)

func TestFileLoad(t *testing.T) {
	folderutil.PrintStdDirs("dict")
	//gologger.Info().Msgf("程序当前的目录: %s", configDir)
	load, err := FileLoad("../../passfolder/weak_pass_top100.txt")
	if err != nil {
		gologger.Error().Msgf("读取文件错误: %s\n", err)
	}
	fmt.Printf("文件长度: %d\n", len(load))
	err = rule.SaveCombinationsToFile(load)
	if err != nil {
		return
	}
}

func TestAscii(t *testing.T) {
	//collectSpecialChars()
	//fmt.Println("Ascii码特殊字符:", collectSpecialChars())
	//gologger.Info().Msgf("Ascii码特殊字符: %s", collectSpecialChars())
	specialChars := CollectSpecialChars()
	gologger.Info().Msgf("Ascii码特殊字符: %s", specialChars)

}
