package rule

import (
	"fmt"
	"github.com/projectdiscovery/gologger"
	"os"
)

// FilterStringsByLength 输出 关键词长度限制，生成的密码长度
func FilterStringsByLength(strings []string, min int, max int) []string {
	var result []string
	for _, s := range strings {
		length := len(s)
		if (min == -1 || length >= min) && (max == -1 || length <= max) {
			result = append(result, s)
		}
	}
	return result
}

// SaveCombinationsToFiles 二维切片保存文件 将组合的规则保存到规则文件
func SaveCombinationsToFiles(combinations [][]string, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			gologger.Error().Msgf("关闭文件失败: %s", err)
		}
	}(file)

	for _, combination := range combinations {
		// 将组合转换为字符串，并写入文件
		comboStr := fmt.Sprintf("%v", combination)  // 将组合转换为字符串表示
		_, err := file.WriteString(comboStr + "\n") // 写入组合并换行
		if err != nil {
			gologger.Error().Msgf("将组合转换为字符串，并写入文件失败: %s", err)
			return err
		}
	}

	return nil
}

// SaveCombinationsToFile 保存切片到文件 将生成的字典切片保存到文件
func SaveCombinationsToFile(combinations []string, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, combination := range combinations {
		_, err := file.WriteString(combination + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}
