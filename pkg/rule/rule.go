package rule

import (
	"fmt"
	"os"
	"reflect"
)

type KeyWord struct {
	Surname              []string // Surname 姓氏
	GivenName            []string // GivenName  名
	FirstLetterSurname   []string // FirstLetterSurname 姓氏拼音首字母
	FirstLetterGivenName []string // FirstLetterGivenName 名拼音首字母
	Connector            []string // Connector   拼接字符
	WeakPass             []string // WeakPass 弱口令
	Year                 []string // Year 年份
	Company              []string // Company 公司
	Email                []string // Email 邮箱
	Phone                []string // Phone 电话
	City                 []string // City 城市
	Keyboard             []string // Keyboard 键盘弱密码
}

// FieldsWithValues 获取KeyWords中的存在值的字段名
// FieldsWithValues 使用反射来检查 KeyWord 结构体中每个字段的值是否非空
func FieldsWithValues(kw *KeyWord) []string {
	var fieldNames []string
	v := reflect.ValueOf(kw).Elem() // 获取reflect.Value，.Elem()获取指针指向的值
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.Kind() == reflect.Slice && !field.IsNil() && field.Len() > 0 {
			fieldNames = append(fieldNames, t.Field(i).Name)
		}
	}
	return fieldNames
}

// SelectFixedLengthPermutations 从给定的关键字列表中选择固定数量的关键字作为一个排列
func SelectFixedLengthPermutations(keywords []string, length int, prefix []string, combinations *[][]string) {
	if len(prefix) == length {
		//fmt.Println(prefix)
		*combinations = append(*combinations, prefix)
		return
	}
	for _, keyword := range keywords {
		SelectFixedLengthPermutations(keywords, length, append(prefix, keyword), combinations)
	}
}

// GenDir 生成字典
func GenDir(key *KeyWord, rules [][]string) []string {
	//combinations := make(map[string]int)
	var combinations []string
	// 遍历所有规则
	for _, rule := range rules {
		// 根据当前规则生成键
		generateByRule(key, rule, "", &combinations)
	}

	return combinations
}

// generateByRule 根据给定的规则生成键
func generateByRule(key *KeyWord, rule []string, prefix string, combinations *[]string) {
	// 选择当前规则的第一个字段
	firstField := rule[0]

	// 根据字段类型选择对应的字符串数组
	var strs []string
	switch firstField {
	case "Surname":
		strs = key.Surname
	case "GivenName":
		strs = key.GivenName
	case "FirstLetterSurname":
		strs = key.FirstLetterSurname
	case "FirstLetterGivenName":
		strs = key.FirstLetterGivenName
	case "Connector":
		strs = key.Connector
	case "WeakPass":
		strs = key.WeakPass
	case "Year":
		strs = key.Year
	case "Company":
		strs = key.Company
	case "Email":
		strs = key.Email
	case "Phone":
		strs = key.Phone
	case "City":
		strs = key.City
	case "Keyboard":
		strs = key.Keyboard
	default:
		return // 无效的规则
	}

	// 对于数组中的每个字符串，递归生成后续规则的键
	for _, str := range strs {
		newPrefix := prefix + str
		if len(rule) == 1 {
			// 如果是最后一个字段，添加到结果中
			*combinations = append(*combinations, newPrefix)
			//(*combinations)[newPrefix] = 1
		} else {
			// 否则，递归处理剩余的规则
			generateByRule(key, rule[1:], newPrefix, combinations)
		}
	}
}

// Deduplicate 去除重复元素
func Deduplicate(slice []string) []string {
	unique := make(map[string]bool)
	var result []string

	for _, v := range slice {
		if _, ok := unique[v]; !ok {
			unique[v] = true
			result = append(result, v)
		}
	}

	return result
}

// RemoveSlicesWithDuplicates 删除包含重复元素的二维切片 slices [][]string
func RemoveSlicesWithDuplicates(slices [][]string) [][]string {
	result := [][]string{}
	seen := make(map[string]int) // 用于跟踪每个元素的出现次数

	// 遍历所有切片
	for _, slice := range slices {
		elementCounts := make(map[string]int) // 用于当前切片的元素计数
		hasDuplicates := false                // 标记当前切片是否有重复元素

		// 计算当前切片中每个元素的出现次数
		for _, element := range slice {
			elementCounts[element]++
			if elementCounts[element] > 1 {
				hasDuplicates = true // 发现重复元素
				break
			}
		}

		// 如果当前切片没有重复元素，加入结果切片
		if !hasDuplicates {
			result = append(result, slice)
		}

		// 更新全局的seen映射
		for element, count := range elementCounts {
			seen[element] += count
		}
	}

	// 返回没有重复元素的切片
	return result
}

// SaveCombinationsToFiles 二维切片保存文件 将组合保存到文件
func SaveCombinationsToFiles(combinations [][]string) error {
	file, err := os.Create("rules.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	for _, combination := range combinations {
		// 将组合转换为字符串，并写入文件
		comboStr := fmt.Sprintf("%v", combination)  // 将组合转换为字符串表示
		_, err := file.WriteString(comboStr + "\n") // 写入组合并换行
		if err != nil {
			return err
		}
	}

	return nil
}

// SaveCombinationsToFile 保存切片到文件 将组合保存到文件
func SaveCombinationsToFile(combinations []string) error {
	file, err := os.Create("字典.txt")
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

// --------------------------------------------------------使用上面的方法，下面的方法已淘汰--------------------------------------------------------

// 获取KeyWords中的字段名
func getKeyWords() []string {
	var keywords []string
	// 接收any 返回传入any的类型和值的信息
	// reflect.Value 是一个能够持有任何类型值的接口
	// 这里返回一个代表该实例的 Value 对象，这个对象包含了关于 kw 的类型和值的信息
	s := reflect.TypeOf(KeyWord{})
	//NumField 返回结构类型的字段计数。如果类型的 Kind 不是 Struct，它会崩溃
	for i := 0; i < s.NumField(); i++ {
		//Field 返回结构类型的第 i 个字段
		field := s.Field(i)
		keywords = append(keywords, field.Name)
	}
	return keywords
}

// 生成给定长度组合的所有可能组合
func genCom(words []string, k int64) [][]string {
	// 初始化空列表用于存储组合
	combinations := make([][]string, 0)

	// 生成所有可能的组合
	n := int64(len(words))
	for mask := int64(1); mask < (1 << uint64(n)); mask++ {
		if int64(bitCount(mask)) == k {
			combination := make([]string, 0)
			for i := int64(0); i < n; i++ {
				if (mask>>uint64(i))&1 == 1 {
					combination = append(combination, words[i])
				}
			}
			combinations = append(combinations, combination)
		}
	}

	return combinations
}

// 计算整数的位数
func bitCount(n int64) int {
	count := 0
	for n > 0 {
		count += int(n & 1)
		n >>= 1
	}
	return count
}

// 反转一个字符串切片的函数
func reverseSlice(slice []string) []string {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}

// 将反转前后的切片都添加到新切片的函数
func addReversedAndOriginal(genCom [][]string) [][]string {
	newGenCom := make([][]string, 0, len(genCom)*2) // 预分配足够的空间

	for _, slice := range genCom {
		// 添加原始切片
		newGenCom = append(newGenCom, slice)

		// 反转当前切片并添加到新切片
		reversedSlice := make([]string, len(slice))
		copy(reversedSlice, slice) // 复制原始切片，以便可以安全地反转
		reversedSlice = reverseSlice(reversedSlice)
		newGenCom = append(newGenCom, reversedSlice)
	}

	return newGenCom
}

// 根据关键字和长度生成组合
func generateCombinations(keywords []string, minLen, maxLen int) []string {
	var combinations []string
	for i := minLen; i <= maxLen; i++ {
		generateCombinationsUtil(keywords, i, 0, "", &combinations)
	}
	return combinations
}

// keywords []string: 这是一个包含所有可能关键词的切片，表示可用于组合的所有关键词列表。
// k int: 这是一个整数参数，表示当前要生成的组合中包含的关键词数量。
// start int: 这是一个整数参数，表示在关键词列表中开始遍历的索引位置。
// current string: 这是一个字符串参数，表示当前已经生成的部分组合。
// combinations *[]string: 这是一个字符串切片指针，用于存储生成的所有组合。
// 递归生成组合
func generateCombinationsUtil(keywords []string, k, start int, current string, combinations *[]string) {
	if k == 0 {
		*combinations = append(*combinations, current)
		return
	}

	for i := start; i <= len(keywords)-k; i++ {
		newCurrent := current
		if current == "" {
			newCurrent = keywords[i]
		} else {
			newCurrent = current + " " + keywords[i]
		}
		generateCombinationsUtil(keywords, k-1, i+1, newCurrent, combinations)
	}
}

func generateCombinationss(keywords []string, minLen, maxLen int) [][]string {
	// 二维切片，外层切片的每个元素是一个组合，每个组合是一个字符串切片
	var allCombinations [][]string

	// 遍历 minLen 到 maxLen 之间的每个 k 值
	for k := minLen; k <= maxLen; k++ {
		// 为当前的 k 值初始化一个新切片来收集当前长度的所有组合
		combinationsForK := make([][]string, 0)

		// 使用 generateCombinationsUtil 填充 combinationsForK
		generateCombinationsUtils(keywords, k, 0, []string{}, &combinationsForK)

		// 将当前 k 值的所有组合添加到 allCombinations 中
		allCombinations = append(allCombinations, combinationsForK...)
	}

	return allCombinations
}

func generateCombinationsUtils(keywords []string, k int, start int, current []string, combinations *[][]string) {
	if k == 0 {
		// 复制当前的组合并添加到结果中
		newCombination := make([]string, len(current))
		copy(newCombination, current)
		*combinations = append(*combinations, newCombination)
		return
	}

	for i := start; i <= len(keywords)-k; i++ {
		newCurrent := make([]string, len(current)+1)
		copy(newCurrent, current)              // 复制当前组合
		newCurrent[len(current)] = keywords[i] // 添加新的关键词
		generateCombinationsUtils(keywords, k-1, i+1, newCurrent, combinations)
	}
}

// 递归生成全排列
func generatePermutations(keywords []string, k int, start int, current []string, permutations *[][]string) {
	if len(current) == k {
		p := make([]string, k)
		copy(p, current)
		*permutations = append(*permutations, p)
		return
	}

	for i := start; i < len(keywords); i++ {
		// 递归生成全排列
		current = append(current, keywords[i])
		generatePermutations(keywords, k, i+1, current, permutations)
		// 回溯，移除最后一个添加的关键词
		current = current[:len(current)-1]
	}
}

// 5 的阶乘（表示为 5!）等于 5 × 4 × 3 × 2 × 1 = 120。
// 1. 首先，我们调用 `factorial(5)`。
// 2. 在 `factorial(5)` 函数中，它会调用 `factorial(4)` 来计算 4 的阶乘，并将结果乘以 5，因为 5! = 5 × 4!。
// 3. 在 `factorial(4)` 函数中，它会调用 `factorial(3)` 来计算 3 的阶乘，并将结果乘以 4，因为 4! = 4 × 3!。
// 4. 在 `factorial(3)` 函数中，它会调用 `factorial(2)` 来计算 2 的阶乘，并将结果乘以 3，因为 3! = 3 × 2!。
// 5. 在 `factorial(2)` 函数中，它会调用 `factorial(1)` 来计算 1 的阶乘，并将结果乘以 2，因为 2! = 2 × 1!。
// 6. 在 `factorial(1)` 函数中，因为 1 的阶乘等于 1，所以它直接返回 1。
// 7. 然后 `factorial(2)` 函数将返回 2 × 1 = 2。
// 8. 接着 `factorial(3)` 函数将返回 3 × 2 = 6。
// 9. `factorial(4)` 函数返回 4 × 6 = 24。
// 10. 最后 `factorial(5)` 函数返回 5 × 24 = 120。
func factorial(n int) int {
	// 基本情况：当 n 等于 0 或 1 时，返回 1
	if n == 0 || n == 1 {
		return 1
	}
	// 递归情况：计算 n 乘以 (n-1) 的阶乘
	// fmt.Println("5的阶乘是：", factorial(5))
	return n * factorial(n-1)
}
