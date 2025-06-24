package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Deduplicate 对整数切片进行去重
// 参数：nums - 待去重的整数切片
// 返回值：去重后的新切片（保持原顺序）
func Deduplicate(nums []int) []int {
    // 使用map记录已出现的元素
    seen := make(map[int]bool)
    result := []int{}
    
    // 遍历原切片
    for _, num := range nums {
        // 如果元素未出现过
        if !seen[num] {
            seen[num] = true  // 标记为已出现
            result = append(result, num)  // 加入结果
        }
    }
    return result
}

func main() {
    fmt.Println("===== 数组去重程序 =====")
    fmt.Println("功能：去除整数数组中的重复元素")
    fmt.Println("请输入整数数组，用空格分隔（如：1 2 2 3）：")
    
    // 读取整行输入
    var input string
    fmt.Scanln(&input)
    
    // 分割字符串为切片
    strNums := strings.Fields(input)
    var nums []int
    
    // 转换字符串为整数
    for _, s := range strNums {
        num, err := strconv.Atoi(s)
        if err != nil {
            fmt.Println("输入包含非数字内容，请重新运行程序")
            return
        }
        nums = append(nums, num)
    }
    
    // 调用去重函数
    uniqueNums := Deduplicate(nums)
    
    // 输出结果
    fmt.Println("原始数组：", nums)
    fmt.Println("去重结果：", uniqueNums)
    
    // 示例
    fmt.Println("\n示例：输入 5 3 5 2 3 5")
    fmt.Println("去重结果：", Deduplicate([]int{5, 3, 5, 2, 3, 5}))
}