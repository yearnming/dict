package keyword

import (
	"bufio"
	"github.com/projectdiscovery/gologger"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var WeakPassKeyboardFile = "passfolder/weak_pass_keyboard_walk.txt"    // 常用键盘密码
var WeakPassKeyboard4File = "passfolder/4_keyboard_walk.txt"           //  4键盘密码
var WeakPassTop100File = "passfolder/weak_pass_top100.txt"             // 弱密码
var ConnectKeywords = []string{"and", "or", "not", "(", ")", " ", ","} // 连接关键词特殊字符
// specialChars := collectSpecialChars()

// CollectSpecialChars 通过ASCII码收集指定范围内的特殊字符到一个字符串切片中
func CollectSpecialChars() []string {
	var specialChars []string
	// 遍历指定的ASCII码范围
	for _, rangeSpec := range []struct{ start, end int }{
		{32, 47},   // (space) -> /
		{58, 64},   // : -> @
		{91, 96},   // [ -> 、
		{123, 126}, // { -> `
	} {
		for i := rangeSpec.start; i <= rangeSpec.end; i++ {
			// 不排除排除空格字符
			//gologger.Info().Msgf("特殊字符:%d - %s", i, string(rune(i)))
			specialChars = append(specialChars, string(rune(i)))
		}
	}
	return specialChars
}

// FileLoad 函数读取文本文件，将每一行作为字符串切片的元素返回。
func FileLoad(filePath string) ([]string, error) {
	absPath, err := filepath.Abs(filePath)

	if err != nil {
		gologger.Error().Msgf("无法获取文件路径：%s", err)
		return nil, err
	}
	gologger.Info().Msgf("获取文件路径：%s", absPath)
	// 打开文件
	file, err := os.Open(absPath)
	if err != nil {
		return nil, err
	}
	defer file.Close() // 确保文件在函数返回时关闭

	// 创建 bufio.Reader 对象用于逐行读取
	reader := bufio.NewReader(file)

	var lines []string // 用于存储每一行的字符串

	// 逐行读取文件
	//for {
	//	line, err := reader.ReadString('\n') // 最后一行会遗漏
	//	if err != nil {
	//		if err == io.EOF {
	//			break // 文件结束
	//		}
	//		return nil, err // 读取错误
	//	}
	//
	//	// 添加去除换行符后的行到切片中
	//	lines = append(lines, line[:len(line)-1])
	//}

	// 逐行读取文件
	for {
		line, isPrefix, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break // 文件结束
			}
			return nil, err // 读取错误
		}
		// isPrefix 表示是否需要更多数据来组成完整的行
		// 如果 isPrefix 为 false，说明 line 已经是完整的行数据

		// 去除可能的换行符和空白字符
		trimmedLine := strings.TrimSpace(string(line))
		// 将读取的行添加到切片中，去掉可能的换行符
		//lines = append(lines, string(line))

		if len(trimmedLine) > 0 {
			lines = append(lines, trimmedLine)
		}
		// 如果 isPrefix 为 true，说明行数据可能被截断，需要继续读取
		if !isPrefix {
			continue
		}
	}

	return lines, nil
}
