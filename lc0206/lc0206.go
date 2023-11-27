package lc0206

var Target = test

func test() {
	t5 := new(ListNode)
	t5.Val = 5

	t4 := new(ListNode)
	t4.Val = 4
	t4.Next = t5

	t3 := new(ListNode)
	t3.Val = 3
	t3.Next = t4

	t2 := new(ListNode)
	t2.Val = 2
	t2.Next = t3

	t1 := new(ListNode)
	t1.Val = 1
	t1.Next = t2

	reverseList(t1)
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var i *ListNode = head
	var j *ListNode = head.Next

	i.Next = nil

	for j != nil {
		var k = j.Next

		j.Next = i

		i = j
		j = k
	}

	return i
}
