package main

import (
	"fmt"
	"sort"
)

func main() {
	var intArr = []int{5, 6, 7, 5, 6, 2, 3, 2, 3}
	result, flag := getOnceNum(intArr)
	if flag {
		fmt.Printf("切片中只出现一次的数字: %d\n", result)
		fmt.Println()
	} else {
		fmt.Println("切片中无只出现一次的数字")
	}

	num := 12121
	fmt.Printf("数字[%v]是否是回文数：%v", num, isPalindrome(num))
	fmt.Println()

	fmt.Printf("给定的括号字符串是否为有效括号：%v", isValidBraces("{{(([]))}}"))
	fmt.Println()

	strs := []string{"fixabc", "fie", "fixefg"}
	fmt.Printf("数组中最长前缀是：%v", longestCommonPrefix(strs))
	fmt.Println()

	dupliNums := []int{5, 5, 6, 7, 7, 8, 8, 8, 9}
	fmt.Println(delDuplicates(dupliNums))

	plusOneNums := []int{9, 9, 9, 9}
	fmt.Println(plusOne(plusOneNums))

	rangeArr := [][]int{{1, 5}, {9, 15}, {7, 12}}
	rangeArr = mergeRange(rangeArr)
	fmt.Println(rangeArr)

	nums := []int{2, 9, 8, 20}
	fmt.Println(getSumTarget(nums, 10))
}

/*
*
获取数组中只出现一次的数字
*/
func getOnceNum(intArr []int) (int, bool) {

	numCountMap := make(map[int]int)
	for _, v := range intArr {
		numCountMap[v]++
	}
	result := 0
	flag := false
	for k, v := range numCountMap {
		if v == 1 {
			result = k
			flag = true
			break
		}
	}
	return result, flag
}

/*
*
回文数
*/
func isPalindrome(num int) bool {
	// 如果数字小于0，则一定不是回文数
	if num < 0 {
		return false
	}
	// 数字反转
	revertNum := 0
	for revertNum < num {
		revertNum = revertNum*10 + num%10
		num = num / 10
	}
	return revertNum == num || num == revertNum/10
}

/*
*
有效的括号
*/
func isValidBraces(str string) bool {
	if len(str)%2 == 1 {
		return false
	}
	bracesMap := map[byte]byte{
		'{': '}',
		'[': ']',
		'(': ')',
	}
	var stack []byte
	for _, char := range str {
		if bracesMap[byte(char)] > 0 {
			// 如果是'(', '[', '{'中的一个，则入栈
			stack = append(stack, byte(char))
		} else if bracesMap[stack[len(stack)-1]] == byte(char) {
			// 如果切片中最后一个元素能和char元素匹配，则出栈
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

/*
*
最长公共前缀
*/
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		char := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if len(strs[j]) == i || strs[j][i] != char {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}

/*
*
删除排序数组中的重复项
*/
func delDuplicates(nums []int) []int {
	length := len(nums)
	if length == 0 {
		return nums
	}
	slow := 1
	for fast := 1; fast < length; fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return nums[:slow]
}

/*
*
给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一
*/
func plusOne(num []int) []int {
	for i := len(num) - 1; i >= 0; i-- {
		if num[i] != 9 {
			num[i]++
			for j := i + 1; j < len(num); j++ {
				num[j] = 0
			}
			return num
		}
	}
	newNum := make([]int, len(num)+1)
	newNum[0] = 1
	return newNum
}

/*
*
合并区间
*/
func mergeRange(num [][]int) [][]int {
	if len(num) <= 1 {
		return num
	}
	sort.SliceStable(num, func(i, j int) bool {
		return num[i][0] < num[j][0]
	})
	fmt.Println("排序后：", num)
	var result [][]int
	curr := num[0]
	for i := 1; i < len(num); i++ {
		if curr[1] < num[i][0] {
			result = append(result, curr)
			curr = num[i]
		} else {
			curr[1] = max(curr[1], num[i][1])
		}
	}
	result = append(result, curr)
	return result
}

/*
*
两数之和
*/
func getSumTarget(nums []int, target int) []int {
	tempMap := make(map[int]int)
	for _, num := range nums {
		if v, ok := tempMap[target-num]; ok {
			return []int{v, num}
		}
		tempMap[num] = num
	}
	return nil
}
