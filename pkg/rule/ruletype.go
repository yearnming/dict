package rule

type DictType struct {
	Wifi      []string
	Directory []string
	Url       []string
	WeakPwd   []string
	Other     []string
}

//预定义一些特定场景下的规则 减少无效字典生成 精准字典 需要变量存在、例如WiFi可能变量名字  但是lif的字典目前没发现变量会影响这个字典

// RuleName GenRulesLength[][]名字相关关键词 不能同时出现两个姓氏相关的关键词OR两个名相关的关键词
func RuleName(GenRulesLength [][]string) [][]string {
	FirstName := []string{"Surname", "FirstLetterSurname", "SurnameBa", "SurnameBaF"}
	LastName := []string{"GivenName", "FirstLetterGivenName", "GivenNameBa", "GivenNameBaF"}
	GenRulesLength = RemoveSubSlicesWithDuplicateKeywords(GenRulesLength, FirstName)
	GenRulesLength = RemoveSubSlicesWithDuplicateKeywords(GenRulesLength, LastName)

	return GenRulesLength
}

// RuleWifi 预定义规则
// 输入参数 名字、密码长度
// 返回值是一个规则
func RuleWifi() {
	// 满足一下条件可视为Wifi密码的生成规则
	// KeyWordLength = 2,3
	// 8<= 密码Length <= 10
	// 从规则里面找出 	GenRulesLength[][0] = 名字相关关键词 || GenRulesLength[][]名字相关关键词 不能同时出现两个姓氏相关的关键词OR两个名相关的关键词
	// 从规则里面找出 拼接字符 KeyWordLength = 2
	//

}

// Directory 目录字典

// login

// xss

// sql

// rce

// lfi

// ssrf

// xxe

// api

// 文件包含

// 命令执行

// 文件上传

// weak 端口 弱密码

// apply 应用 设备 弱密码

// 学校 特定学校 特定学院 特定专业 特定班级 学号

// 电话号码

// 身份证

// 年份

// 特定 特定时间 特定日期

// 子域名

// cms 不同框架 的 字典 [][]string

// windows linux 不同系统下的 字典 [][]string
