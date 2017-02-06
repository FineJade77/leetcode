import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func Max(first int, args ...int) int {
	for _, v := range args {
		if v > first {
			first = v
		}
	}

	return first
}

func GetLength(node *ListNode) int {
	length := 0
	for node != nil {
		node = node.Next
		length += 1
	}

	return length
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var carry, index, length int
	node := new(ListNode)
	head := node
	list := make([]int, Max(GetLength(l1), GetLength(l2))+1)

	for l1 != nil || l2 != nil {
		var val int

		if l1 == nil {
			l1 = new(ListNode)
			l1.Val = 0
			l1.Next = nil
		}

		if l2 == nil {
			l2 = new(ListNode)
			l2.Val = 0
			l2.Next = nil
		}

		sum := l1.Val + l2.Val + carry
		val, carry = sum%10, sum/10

		list[index] = val
		index += 1

		l1, l2 = l1.Next, l2.Next
	}

	if carry != 0 {
		list[index] = carry
		length = len(list)
	} else {
		length = len(list) - 1
	}

	for i := 0; i < length; i++ {
		temp := new(ListNode)
		temp.Val = list[i]

		node.Next = temp
		node = temp
	}

	return head.Next
}

func main() {
	l1 := new(ListNode)
	temp := new(ListNode)
	l2 := new(ListNode)
	l1.Val = 2
	temp.Val = 5
	temp.Next = nil
	l1.Next = temp

	l2.Val = 8
	l2.Next = nil

	node := addTwoNumbers(l1, l2)

	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}