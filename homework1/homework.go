package main

import (
	"slices"
)

// 1. 只出现一次的数字
// 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
func SingleNumber2(nums []int) int {
	maps := make(map[int]int)
	for _, value := range nums {
		maps[value] = maps[value] + 1
	}
	for key, value := range maps {
		if value == 1 {
			return key
		}
	}
	return 0
}
func SingleNumber(nums []int) int {
	result := 0
	for _, num := range nums {
		result ^= num
	}
	return result
}

// 2. 回文数
// 判断一个整数是否是回文数
func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	aa := make([]int, 0)
	for {
		if x == 0 {
			break
		}
		a := x % 10
		x = x / 10
		aa = append(aa, a)
	}
	left, right := 0, len(aa)-1
	for left <= right {
		if aa[left] != aa[right] {
			return false
		}
		left++
		right--
	}

	return true
}

// 3. 有效的括号
// 给定一个只包括 '(', ')', '{', '}', '[', ']' 的字符串，判断字符串是否有效
func IsValid(s string) bool {
	stack := make([]rune, 0)
	for _, v := range s {
		if v == '(' || v == '[' || v == '{' {
			stack = append(stack, v)
		} else {
			temp := stack[len(stack)-1]
			if (v == ')' && temp == '(') || (v == ']' && temp == '[') || (temp == '{' && v == '}') {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return true
}

// 4. 最长公共前缀
// 查找字符串数组中的最长公共前缀
func LongestCommonPrefix(strs []string) string {
	ans := strs[0]
	for _, value := range strs {
		for i := 0; i < len(value) && i < len(ans); i++ {
			if ans[i] == value[i] {
				continue
			}
			ans = ans[:i]
			break
		}
	}
	return ans
}

// 5. 加一
// 给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一
func PlusOne(digits []int) []int {
	j := 0
	i := len(digits) - 1
	if digits[i]+1 >= 10 {
		digits[i] = 0
		j = 1
	} else {
		digits[i]++
	}
	i--
	for ; i >= 0; i-- {
		if digits[i]+j >= 10 {
			j = 1
			digits[i] = 0
		} else {
			digits[i] += j
			j = 0
		}
	}
	if j > 0 {
		digitss := make([]int, 0)
		digitss = append(digitss, j)
		digitss = append(digitss, digits[:len(digits)]...)
		return digitss
	}
	return digits
}

// 6. 删除有序数组中的重复项
// 给你一个有序数组 nums ，请你原地删除重复出现的元素，使每个元素只出现一次，返回删除后数组的新长度。
// 不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
func RemoveDuplicates(nums []int) int {
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
		fast++
	}
	return slow + 1
}

// 7. 合并区间
// 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
// 请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
func Merge(intervals [][]int) [][]int {
	slices.SortFunc(intervals, func(e []int, e2 []int) int {
		return e[0] - e2[0]
	})
	ans := make([][]int, 0)
	for _, value := range intervals {
		m := len(ans)
		if m > 0 && value[0] <= ans[m-1][1] {
			ans[m-1][1] = max(ans[m-1][1], value[1])
		} else {
			ans = append(ans, value)
		}
	}

	return ans
}

// 8. 两数之和
// 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数
func TwoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
