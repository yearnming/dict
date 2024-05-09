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
