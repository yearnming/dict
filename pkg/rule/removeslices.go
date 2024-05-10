package rule

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

// RemoveSlicesWithDuplicates 函数是一个清理函数，它的输入参数是一个二维切片，其中每个元素是一个子切片。
// 这个函数的作用是，找出那些子切片中存在有重复元素的子切片，并将它们从原切片中移除。
// 函数的输出是一个新的二维切片，其中所有的子切片都不包含任何的重复元素。slices [][]string
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

// 对比两个切片是否相等
func compareSlices(slice1, slice2 []string) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i, v := range slice1 {
		if v != slice2[i] {
			return false
		}
	}
	return true
}

// RemoveDuplicateSlices 函数是一个清理函数，移除存在重复的子切片，它的输入参数是一个二维切片，其中每个元素是一个子切片。
// 这个函数的作用是，找出存在有重复的子切片，并将它们从原切片中移除。
// 函数的输出是一个新的二维切片，其中不会存在重复的子切片。slices [][]string
func RemoveDuplicateSlices(slices [][]string) [][]string {
	result := [][]string{}

	for i, slice1 := range slices {
		isDuplicate := false

		for _, slice2 := range slices[i+1:] {
			if compareSlices(slice1, slice2) {
				isDuplicate = true
				break
			}
		}

		// 如果当前切片不是重复的，将它加入结果切片中
		if !isDuplicate {
			result = append(result, slice1)
		}
	}

	return result
}
