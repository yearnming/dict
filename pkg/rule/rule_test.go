package rule

import (
	"fmt"
	"testing"
)

func TestGeneratePermutations(t *testing.T) {
	key := &KeyWord{
		Surname:   []string{"zhang", "li"},
		GivenName: []string{"san", "si"},
		City:      []string{"beijing", "shanghai"},
	}
	keywords := fieldsWithValues(key)
	//fmt.Println(existingFields) // 输出: ["Surname" "GivenName" "City"]
	fmt.Println("KeyWord结构体:", keywords)

	var combinations [][]string
	//keywords := []string{"关键词1", "关键词2", "关键词3", "关键词4"}
	//keywords := getKeyWords()
	SelectFixedLengthPermutations(keywords, 3, []string{}, &combinations)
	fmt.Println("combinations长度:", len(combinations))

	combinations = removeSlicesWithDuplicates(combinations)
	//fmt.Println("After removal:", combinations)
	fmt.Println("去除重复规则长度:", len(combinations))
	err := saveCombinationsToFiles(combinations)
	if err != nil {
		fmt.Println("保存组合规则失败:", err)
	} else {
		fmt.Println("组合保存到 combinations.txt")
	}
	//keywords := []string{"关键词1", "关键词2", "关键词3", "关键词4"}
	//k := 3
	//var permutations [][]string
	//
	//generatePermutations(keywords, k, 0, []string{}, &permutations)
	//for _, comb := range permutations {
	//	fmt.Println(comb)
	//}
	//err := saveCombinationsToFiles(permutations)
	//if err != nil {
	//	fmt.Printf("Error saving permutations to file: %v\n", err)
	//} else {
	//	fmt.Println("Permutations have been successfully saved to the file.")
	//}
}

func TestAbc(t *testing.T) {
	//keywords := getKeyWords()
	keywords := []string{"关键词1", "关键词2", "关键词3", "关键词4"}
	//fmt.Println(keywords)
	combinations := generateCombinationss(keywords, 3, 3)
	fmt.Printf("共有 %d 个组合\n", len(combinations))
	newGenCom := addReversedAndOriginal(combinations)
	fmt.Printf("原始切片和反转切片总和:%d\n", len(newGenCom))
	for _, comb := range combinations {
		fmt.Println(comb)
	}
}

func TestNameRule(t *testing.T) {
	//keywords := getKeyWords()
	//fmt.Println(keywords)
	keywords := []string{"关键词1", "关键词2", "关键词3", "关键词4", "关键词5"}
	combinations := generateCombinations(keywords, 3, 3)
	fmt.Println("combinations长度:", len(combinations))
	err := saveCombinationsToFile(combinations)
	if err != nil {
		fmt.Println("组合规则保存失败:", err)
	} else {
		fmt.Println("组合规则保存到 combinations.txt")
	}
}

func TestFactorial(t *testing.T) {
	// 5 的阶乘（表示为 5!）等于 5 × 4 × 3 × 2 × 1 = 120。
	//1. 首先，我们调用 `factorial(5)`。
	//2. 在 `factorial(5)` 函数中，它会调用 `factorial(4)` 来计算 4 的阶乘，并将结果乘以 5，因为 5! = 5 × 4!。
	//3. 在 `factorial(4)` 函数中，它会调用 `factorial(3)` 来计算 3 的阶乘，并将结果乘以 4，因为 4! = 4 × 3!。
	//4. 在 `factorial(3)` 函数中，它会调用 `factorial(2)` 来计算 2 的阶乘，并将结果乘以 3，因为 3! = 3 × 2!。
	//5. 在 `factorial(2)` 函数中，它会调用 `factorial(1)` 来计算 1 的阶乘，并将结果乘以 2，因为 2! = 2 × 1!。
	//6. 在 `factorial(1)` 函数中，因为 1 的阶乘等于 1，所以它直接返回 1。
	//7. 然后 `factorial(2)` 函数将返回 2 × 1 = 2。
	//8. 接着 `factorial(3)` 函数将返回 3 × 2 = 6。
	//9. `factorial(4)` 函数返回 4 × 6 = 24。
	//10. 最后 `factorial(5)` 函数返回 5 × 24 = 120。
	//fmt.Println("5的阶乘是：", factorial(5))

	// 列表关键词数量
	words := []string{"关键词1", "关键词2", "关键词3", "关键词4", "关键词5"}
	//words := getKeyWords()

	// 组合的关键字数量
	k := int64(2)

	// 计算组合数量
	combinations := genCom(words, k)

	// 打印所有可能的组合
	fmt.Printf("关键词列表包含 %d 个关键词，计算 %d 个关键词组成的所有可能组合为：\n", len(words), k)
	fmt.Printf("共有 %d 个组合\n", len(combinations))
	// 调用函数并打印结果
	newGenCom := addReversedAndOriginal(combinations)
	fmt.Printf("原始切片和反转切片总和:%d\n", len(newGenCom))
	for _, comb := range newGenCom {
		fmt.Println(comb)
	}
	//sort.Reverse()
}
