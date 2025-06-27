package main

import (
	"fmt"
	"strconv"
	"strings"
)

// TreeNode 二叉树节点定义
type TreeNode struct {
    Val   int       // 节点值
    Left  *TreeNode // 左子节点
    Right *TreeNode // 右子节点
}

// BuildTree 根据层序遍历数组构建二叉树
// 参数：nums - 层序遍历的数组，-1表示空节点
// 返回值：构建好的二叉树根节点
func BuildTree(nums []int) *TreeNode {
    // 处理空数组或根节点为空的情况
    if len(nums) == 0 || nums[0] == -1 {
        return nil
    }

    // 创建根节点
    root := &TreeNode{Val: nums[0]}
    queue := []*TreeNode{root}  // 使用队列辅助构建
    i := 1  // 数组索引

    // 循环处理队列中的节点
    for len(queue) > 0 && i < len(nums) {
        node := queue[0]
        queue = queue[1:]  // 出队

        // 处理左子节点
        if i < len(nums) && nums[i] != -1 {
            node.Left = &TreeNode{Val: nums[i]}
            queue = append(queue, node.Left)  // 左子节点入队
        }
        i++

        // 处理右子节点
        if i < len(nums) && nums[i] != -1 {
            node.Right = &TreeNode{Val: nums[i]}
            queue = append(queue, node.Right)  // 右子节点入队
        }
        i++
    }

    return root
}

// printTree 以树形结构打印二叉树
func printTree(root *TreeNode, prefix string, isLeft bool) {
    if root == nil {
        return
    }
    
    // 打印当前节点
    fmt.Print(prefix)
    if isLeft {
        fmt.Print("├── ")
    } else {
        fmt.Print("└── ")
    }
    fmt.Println(root.Val)
    
    // 递归打印子树
    newPrefix := prefix
    if isLeft {
        newPrefix += "│   "
    } else {
        newPrefix += "    "
    }
    printTree(root.Left, newPrefix, true)
    printTree(root.Right, newPrefix, false)
}

func main() {
    fmt.Println("===== 二叉树构建程序 =====")
    fmt.Println("功能：根据层序遍历数组构建二叉树")
    fmt.Println("输入说明：用空格分隔数字，-1表示空节点")
    fmt.Println("示例输入：1 2 3 -1 4 5")
    fmt.Println("请输入层序遍历数组：")
    
    // 读取输入
    var input string
    fmt.Scanln(&input)
    
    // 处理输入
    strNums := strings.Fields(input)
    var nums []int
    for _, s := range strNums {
        num, err := strconv.Atoi(s)
        if err != nil {
            fmt.Println("输入包含非数字内容，请重新运行程序")
            return
        }
        nums = append(nums, num)
    }
    
    // 构建二叉树
    tree := BuildTree(nums)
    
    // 打印结果
    fmt.Println("\n构建的二叉树结构：")
    printTree(tree, "", false)
    
    // 示例
    fmt.Println("\n示例输入：1 -1 2 -1 -1 3")
    fmt.Println("构建结果：")
    printTree(BuildTree([]int{1, -1, 2, -1, -1, 3}), "", false)
}