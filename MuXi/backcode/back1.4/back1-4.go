package main

import (
	"fmt"
	"strconv"
	"strings"
)

// merge 合并两个有序数组到nums1中
// 参数：
//   nums1 - 第一个数组（有足够空间）
//   m - nums1中的元素数量
//   nums2 - 第二个数组
//   n - nums2中的元素数量
func merge(nums1 []int, m int, nums2 []int, n int) {
    // 设置三个指针：
    // i - nums1有效元素末尾
    // j - nums2末尾
    // k - 合并后数组末尾
    i, j, k := m-1, n-1, m+n-1

    // 从后向前合并
    for j >= 0 {
        // 比较两个数组当前元素
        if i >= 0 && nums1[i] > nums2[j] {
            nums1[k] = nums1[i]
            i--
        } else {
            nums1[k] = nums2[j]
            j--
        }
        k--
    }
}

func main() {
    fmt.Println("===== 有序数组合并程序 =====")
    fmt.Println("功能：合并两个有序数组到第一个数组中")
    fmt.Println("注意：第一个数组需要有足够的空间")
    
    // 输入nums1
    fmt.Println("\n请输入第一个有序数组（包含预留空间），用空格分隔")
    fmt.Println("示例：1 3 5 0 0 0 表示有3个有效元素和3个预留空间")
    fmt.Print("请输入：")
    var input1 string
    fmt.Scanln(&input1)
    
    // 处理nums1输入
    strNums1 := strings.Fields(input1)
    var nums1 []int
    for _, s := range strNums1 {
        num, err := strconv.Atoi(s)
        if err != nil {
            fmt.Println("输入包含非数字内容，请重新运行程序")
            return
        }
        nums1 = append(nums1, num)
    }
    
    // 输入有效元素数m
    fmt.Print("请输入第一个数组的有效元素数：")
    var m int
    _, err := fmt.Scanln(&m)
    if err != nil || m < 0 || m > len(nums1) {
        fmt.Println("输入错误，有效元素数必须小于等于数组长度")
        return
    }
    
    // 输入nums2
    fmt.Print("请输入第二个有序数组，用空格分隔：")
    var input2 string
    fmt.Scanln(&input2)
    
    // 处理nums2输入
    strNums2 := strings.Fields(input2)
    var nums2 []int
    for _, s := range strNums2 {
        num, err := strconv.Atoi(s)
        if err != nil {
            fmt.Println("输入包含非数字内容，请重新运行程序")
            return
        }
        nums2 = append(nums2, num)
    }
    
    n := len(nums2)
    
    // 检查nums1空间是否足够
    if len(nums1) < m+n {
        fmt.Println("错误：第一个数组空间不足")
        return
    }
    
    // 合并前打印
    fmt.Println("\n合并前：")
    fmt.Println("nums1有效部分：", nums1[:m])
    fmt.Println("nums2：", nums2)
    
    // 执行合并
    merge(nums1, m, nums2, n)
    
    // 输出结果
    fmt.Println("\n合并后：")
    fmt.Println(nums1)
    
    // 示例
    fmt.Println("\n示例：")
    ex1 := make([]int, 6)
    copy(ex1, []int{1, 3, 5})
    ex2 := []int{2, 4, 6}
    fmt.Println("输入：nums1 = [1 3 5 0 0 0], m = 3, nums2 = [2 4 6], n = 3")
    merge(ex1, 3, ex2, 3)
    fmt.Println("输出：", ex1)
}