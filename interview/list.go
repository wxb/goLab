package main

import "fmt"

type ListNode struct {
	Value string
	Next  *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	current := head

	for current != nil {
		next := current.Next // 临时保存下一个节点
		current.Next = prev  // 将当前节点的next指向前一个节点
		prev = current       // 移动prev指向当前节点
		current = next       // 移动current到下一个节点
	}

	return prev // 返回反转后的头结点,此时prev已经指向最后一个节点
}

// 打印链表
func printList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Println("--->", current.Value)
		current = current.Next
	}
}

// 检查链表是否有环

// 哈希表法：遍历链表，用哈希表记录节点地址，若遇到重复地址则有环。
func checkLoopHash(head *ListNode) bool {
	current := head
	visited := map[*ListNode]bool{}

	for current != nil {
		if _, ok := visited[current]; ok {
			return true
		}

		visited[current] = true
		current = current.Next
	}

	return false
}

// 快慢指针法（最优）：快指针每次走 2 步，慢指针每次走 1 步，若相遇则有环；若快指针走到 null 则无环。
func checkLoop(head *ListNode) bool {
	// 边界检查：空链表或只有一个节点
	if head == nil || head.Next == nil {
		return false
	}

	// 初始化：两个指针都从头开始
	slow, fast := head, head

	// 循环条件：fast能继续走两步
	for fast != nil && fast.Next != nil {
		slow = slow.Next      // 慢指针走1步
		fast = fast.Next.Next // 快指针走2步
		// 如果相遇，说明有环
		if slow == fast {
			return true
		}
	}

	// fast到达nil，说明无环
	return false
}

// 计算环的长度
func getCycleLength(head *ListNode) int {
	if head == nil || head.Next == nil {
		return 0
	}

	slow, fast := head, head

	// 找到相遇点
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			// 从相遇点开始，走一圈计算长度
			length := 1
			current := slow.Next
			for current != slow {
				length++
				current = current.Next
			}
			return length
		}
	}

	return 0
}

// 删除倒数第 N 个节点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 使用虚拟头节点，处理删除头节点的情况
	dummy := &ListNode{Value: "", Next: head}
	slow, fast := dummy, dummy

	// 快指针先走n步
	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	// 快慢指针同时移动，直到快指针走到末尾
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	// 删除节点
	slow.Next = slow.Next.Next

	return dummy.Next
}

func main() {
	// 创建一个链表：a -> b -> c -> nil
	c := &ListNode{Value: "c", Next: nil}
	b := &ListNode{Value: "b", Next: c}
	a := &ListNode{Value: "a", Next: b}

	fmt.Println("是否存在：", getCycleLength(a))

	c.Next = a
	fmt.Println("是否存在：", getCycleLength(a))
}
