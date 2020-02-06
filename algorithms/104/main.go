package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// maxDepth 二叉树最大深度
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

// max 取最大数
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	t := createTree(arr)
	fmt.Println(maxDepth(t))
}

// createTree 根据数组创建一个二叉树
func createTree(arr []int) *TreeNode {
	var temp []*TreeNode
	for _, v := range arr {
		temp = append(temp, &TreeNode{
			Val: v,
		})
	}
	if len(temp) > 0 {
		for i := 0; i < len(temp)/2-1; i++ {
			if 2*i+1 < len(temp)-1 {
				temp[i].Left = temp[2*i+1]
			}
			if 2*i+2 < len(temp)-1 {
				temp[i].Right = temp[2*i+2]
			}
		}
		last := len(arr)/2 - 1
		temp[last].Left = temp[last*2+1]
		if len(arr)%2 == 1 {
			temp[last].Right = temp[last*2+2]
		}
	}
	if len(temp) > 0 {
		return temp[0]
	}
	return nil
}

// print 打印一个二叉树
func print(node *TreeNode) {
	if node != nil {
		fmt.Println(node.Val)
		print(node.Right)
		print(node.Left)
	}
}
