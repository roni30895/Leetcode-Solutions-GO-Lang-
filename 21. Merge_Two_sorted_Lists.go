/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    
    if list1 == nil && list2 != nil { return list2 }
    if list1 != nil && list2 == nil { return list1 }
    if list1 == nil && list2 == nil { return nil }
    newNode := new(ListNode)
    if list1.Val >= list2.Val {
        newNode.Val = list2.Val
        list2 = list2.Next
        newNode.Next = mergeTwoLists(list1, list2)
    } else {
        newNode.Val = list1.Val
        list1 = list1.Next
        newNode.Next = mergeTwoLists(list1, list2)
    }
    return newNode
 }
