package bfs

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
最近公共祖先
1. 节点为空时,退出, 说明没找到
2. 节点都在右子树,遍历至root.right,否则,遍历左子树
3. 节点分别在两侧子树，则找到最近公共祖先
*/

func solution(root, p, q *TreeNode) *TreeNode {
	for root != nil {
		if root.Val < p.Val && root.Val < q.Val {
			root = root.Right
			continue
		} else if root.Val > p.Val && root.Val > q.Val {
			root = root.Left
			continue
		}
		break
	}
	return root
}
