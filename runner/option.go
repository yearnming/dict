package runner

import (
	"fmt"
	"github.com/projectdiscovery/goflags"
	"github.com/projectdiscovery/gologger"
	"github.com/yearnming/dict/pkg/keyword"
	"os"
	"regexp"
	"strconv"
	"time"
)

type Options struct {
	DictType             string              // DictType 输入字典类型
	KeyWord              string              // KeyWord 输入关键词
	Rules                string              // Rules 自定义关键词组合规则
	Length               goflags.StringSlice // Length 关键词长度限制，生成的密码长度
	LengthMin            int                 // KeyWordLength 关键词长度限制，生成的密码长度，最少
	LengthMax            int                 // KeyWordLength 关键词长度限制，生成的密码长度，最大
	Output               string              // Output 保存文件
	KeyWordLength        goflags.StringSlice // KeyWordLength 关键词规则数量限制，两个关键词组合，三个关键词组合
	KeyWordLengthMin     int                 // KeyWordLength 关键词规则数量限制，最少
	KeyWordLengthMax     int                 // KeyWordLength 关键词规则数量限制，最大
	OutputRule           string              // Output 保存规则文件
	Surname              goflags.StringSlice // Surname 姓氏
	GivenName            goflags.StringSlice // GivenName  名
	FirstLetterSurname   goflags.StringSlice // FirstLetterSurname 姓氏拼音首字母
	FirstLetterGivenName goflags.StringSlice // FirstLetterGivenName 名拼音首字母
	SurnameBa            goflags.StringSlice // Surname 姓氏拼音首字母大写
	GivenNameBa          goflags.StringSlice // GivenName  名拼音首字母大写
	SurnameBaF           goflags.StringSlice // FirstLetterSurname 姓氏拼音首字母拼音大写
	GivenNameBaF         goflags.StringSlice // FirstLetterGivenName 名拼音首字母拼音大写
	Connector            goflags.StringSlice // Connector   拼接字符
	WeakPass             goflags.StringSlice // WeakPass 弱口令
	Year                 goflags.StringSlice // Year 年份
	Company              goflags.StringSlice // Company 公司
	Email                goflags.StringSlice // Email 邮箱
	Phone                goflags.StringSlice // Phone 电话
	City                 goflags.StringSlice // City 城市
	Keyboard             goflags.StringSlice // Keyboard 键盘弱密码
	ChineseName          string              // 中文名字
}

// ParseOptions 解析应用程序的命令行选项
func ParseOptions() *Options {
	options := &Options{}

	flagSet := goflags.NewFlagSet()
	flagSet.SetDescription(`passfolder 是一个快速生成字典工具。`)

	flagSet.CreateGroup("input", "关键词",
		flagSet.StringVarP(&options.ChineseName, "ChineseName", "chn", "", "中文名字"),
		flagSet.StringSliceVarP(&options.Surname, "surname", "sn", nil, "姓氏", goflags.NormalizedStringSliceOptions),
		flagSet.StringSliceVarP(&options.GivenName, "givenName", "gn", nil, "名", goflags.NormalizedStringSliceOptions),
		flagSet.StringSliceVarP(&options.FirstLetterSurname, "firstlettersurname", "flsn", nil, "姓氏拼音首字母", goflags.NormalizedStringSliceOptions),
		flagSet.StringSliceVarP(&options.FirstLetterGivenName, "firstlettergivenname", "flgn", nil, "名拼音首字母", goflags.NormalizedStringSliceOptions),
		flagSet.StringSliceVarP(&options.Connector, "connector", "cn", nil, "拼接字符", goflags.NormalizedStringSliceOptions),
		flagSet.StringSliceVarP(&options.WeakPass, "weakpass", "wp", nil, "弱口令", goflags.NormalizedStringSliceOptions),
		flagSet.StringSliceVarP(&options.Year, "year", "y", nil, "年份", goflags.NormalizedStringSliceOptions),
		flagSet.StringSliceVarP(&options.Company, "company", "cp", nil, "公司", goflags.NormalizedStringSliceOptions),
		flagSet.StringSliceVarP(&options.Email, "email", "em", nil, "邮箱", goflags.NormalizedStringSliceOptions),
		flagSet.StringSliceVarP(&options.Phone, "phone", "ph", nil, "电话", goflags.NormalizedStringSliceOptions),
		flagSet.StringSliceVarP(&options.City, "city", "ct", nil, "城市", goflags.NormalizedStringSliceOptions),
		flagSet.StringSliceVarP(&options.Keyboard, "keyboard", "kb", nil, "键盘弱密码", goflags.NormalizedStringSliceOptions),
	)
	flagSet.CreateGroup("configuration", "配置",
		flagSet.StringVarP(&options.DictType, "dicttype", "dt", "wifi", "输入字典类型"),
		//flagSet.StringVarP(&options.KeyWord, "keyword", "key", "", "输入关键词"),
		flagSet.StringVarP(&options.Rules, "rule", "r", "", "自定义关键词组合规则"),
		flagSet.StringSliceVarP(&options.Length, "length", "l", nil, "生成的密码长度限制,-1不限制最大长度,例如：8,10 | 8,-1", goflags.NormalizedStringSliceOptions),
		flagSet.StringSliceVarP(&options.KeyWordLength, "KeyWordLength", "kwl", nil, "关键词规则数量限制，两个关键词组合，三个关键词组合", goflags.NormalizedStringSliceOptions),
	)
	flagSet.CreateGroup("output", "输出",
		flagSet.StringVarP(&options.Output, "output", "o", "", "保存文件"),
		flagSet.StringVarP(&options.OutputRule, "outputrule", "ol", "", "保存规则文件文件"),
	)
	flagSet.SetCustomHelpText("使用示例:\ngo run cmd/dict/main.go -chn 张大伟 -kwl 2,3")

	if err := flagSet.Parse(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if err := options.ValidateOptions(); err != nil {
		gologger.Fatal().Msgf("检查参数时出错: %s\n", err)
	}
	return options
}
func (options *Options) ValidateOptions() error {
	// connector weakpass keyboard
	if options.ChineseName == "" && options.Surname == nil && options.GivenName == nil && options.FirstLetterSurname == nil && options.FirstLetterGivenName == nil && options.Year == nil && options.Company == nil && options.Email == nil && options.Phone == nil && options.City == nil {
		//gologger.Error().Msgf("需要提供关键词才能生成字典: \n")
		return fmt.Errorf("需要提供关键词才能生成字典")
	}

	if options.ChineseName != "" {
		namepinyin := keyword.Pinyin(options.ChineseName)
		options.Surname = append(options.Surname, namepinyin.Surname)
		options.GivenName = append(options.GivenName, namepinyin.GivenName)
		options.FirstLetterSurname = append(options.FirstLetterSurname, namepinyin.FirstLetterSurname)
		options.FirstLetterGivenName = append(options.FirstLetterGivenName, namepinyin.FirstLetterGivenName)
		options.SurnameBa = append(options.SurnameBa, namepinyin.SurnameBa)
		options.GivenNameBa = append(options.GivenNameBa, namepinyin.GivenNameBa)
		options.SurnameBaF = append(options.SurnameBaF, namepinyin.SurnameBaF)
		options.GivenNameBaF = append(options.GivenNameBaF, namepinyin.GivenNameBaF)

	}

	if options.WeakPass == nil {
		WeakPass, err := keyword.FileLoad(keyword.WeakPassTop100File)
		if err != nil {
			gologger.Error().Msgf("加载弱密码文件时出错: %s\n", err)
			return err
		}
		options.WeakPass = WeakPass
	}

	if options.Keyboard == nil {
		Keyboard, err := keyword.FileLoad(keyword.WeakPassKeyboard4File)
		if err != nil {
			gologger.Error().Msgf("加载弱密码文件时出错: %s\n", err)
			return err
		}
		options.Keyboard = Keyboard
	}

	if options.Connector == nil {
		options.Keyboard = keyword.CollectSpecialChars()
		//gologger.Info().Msgf("使用特殊字符作为连接符: %v\n", options.Keyboard)
	}

	if options.Output == "" {
		// 获取当前时间戳
		timestamp := time.Now().Unix()

		// 将时间戳添加到文件名中，例如 "exampleDict_1673886400.txt"
		timestampedFilename := "字典_" + strconv.FormatInt(timestamp, 10) + ".txt"
		options.Output = timestampedFilename
	}
	if options.OutputRule == "" {
		// 获取当前时间戳
		timestamp := time.Now().Unix()

		// 将时间戳添加到文件名中，例如 "exampleDict_1673886400.txt"
		timestampedFilename := "生成规则_" + strconv.FormatInt(timestamp, 10) + ".txt"
		options.OutputRule = timestampedFilename
	}

	err := options.parseKeyWordLength(options.KeyWordLength)
	if err != nil {
		gologger.Error().Msgf("解析长度限制时出错: %s\n", err)
		panic(err)
	}
	if options.KeyWordLength == nil {
		options.KeyWordLengthMin = 3
		options.KeyWordLengthMax = 3
	} else {
		if len(options.KeyWordLength) == 1 {
			num1, _ := strconv.Atoi(options.KeyWordLength[0])
			options.KeyWordLengthMin = num1
			options.KeyWordLengthMax = num1
		} else {
			num1, _ := strconv.Atoi(options.KeyWordLength[0])
			num2, _ := strconv.Atoi(options.KeyWordLength[1])
			options.KeyWordLengthMin = num1
			options.KeyWordLengthMax = num2
		}
	}

	//gologger.Info().Msgf("关键词规则数量限制: %v", options.KeyWordLength)
	//gologger.Info().Msgf("关键词规则数量限制: \"%d\"到\"%d\"", options.KeyWordLengthMin, options.KeyWordLengthMax)

	err = options.parseKeyWordLength(options.Length)
	if err != nil {
		gologger.Error().Msgf("解析长度限制时出错: %s\n", err)
		panic(err)
	}

	if options.Length == nil {
		options.LengthMin = -1
		options.LengthMax = -1
	} else {
		if len(options.Length) == 1 {
			num1, _ := strconv.Atoi(options.Length[0])
			options.LengthMin = num1
			options.LengthMax = num1
		} else {
			num1, _ := strconv.Atoi(options.Length[0])
			num2, _ := strconv.Atoi(options.Length[1])
			options.LengthMin = num1
			options.LengthMax = num2
		}
	}
	//gologger.Info().Msgf("密码长度限制: %v", options.Length)
	//gologger.Info().Msgf("密码长度限制: \"%d\"到\"%d\"", options.LengthMin, options.LengthMax)

	return nil
}

func isNumeric(s string) bool {
	// 正则表达式匹配一个或多个数字
	re := regexp.MustCompile(`^\d+$`)
	return re.MatchString(s)
}

func (options *Options) parseKeyWordLength(Length []string) error {
	switch len(Length) {
	case 0:
		Length = []string{"3"}
	case 1, 2:
		if Length[1] == "-1" {
			break
		}
		for _, str := range Length {
			if !isNumeric(str) || str == "-1" {
				return fmt.Errorf("长度得为 1 或 2，且为数字 you -1，例如: \"3\", \"2,3\"")
			}
		}
		num1, _ := strconv.Atoi(Length[0])
		num2 := num1
		if len(Length) == 2 {
			num2, _ = strconv.Atoi(Length[1])
		}
		if num1 > num2 {
			return fmt.Errorf("KeyWordLength 长度得为 1 或 2，且为数字，例如: \"3\", \"2,3\"")
		}
	default:
		return fmt.Errorf("KeyWordLength 长度得为 1 或 2，且为数字，例如: \"3\", \"2,3\"")
	}
	return nil
}
