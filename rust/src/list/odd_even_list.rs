use crate::list::list_node::ListNode;
pub struct Solution {}
impl Solution {
    /**
     * 解题思路，双指针，双链表，奇链表起始于head，偶链表起始于奇链表的后驱节点
     * 奇指针和偶指针分别指向奇链表和偶链表的头部
     * 如果偶指针后驱节点存在，奇指针指向偶指针的后驱节点
     * 奇指针前进一步
     * 如果奇指针后驱节点为空，遍历停止
     * 如果奇指针后驱节点不为空，偶指针指向奇指针的后驱节点
     * 偶指针前进一步
     * 最后将奇指针的后驱节点指向偶节点
     * 返回奇节点
     * 这种题目更多是考察链表指针的灵活移动，如果能快速写出来，说明链表基本操作已经掌握
     * */
    pub fn odd_even_list(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        let mut odd_node = head?;
        if odd_node.next.is_none() {
            return Some(odd_node);
        }

        let mut even_node = odd_node.next.take().unwrap();
        let (mut odd, mut even) = (&mut odd_node, &mut even_node);
        while even.next.is_some() {
            odd.next = even.next.take();
            odd = odd.next.as_mut().unwrap();
            if odd.next.is_none() {
                break;
            }
            even.next = odd.next.take();
            even = even.next.as_mut().unwrap();
        }
        odd.next = Some(even_node);
        Some(odd_node)
    }
}

#[test]
fn test() {
    let head_a = ListNode::from_array_by_head(vec![1, 2, 3, 4, 5, 6]);
    let target = ListNode::from_array_by_head(vec![1, 3, 5, 2, 4, 6]);
    let res = Solution::odd_even_list(head_a.clone());
    assert_eq!(res, target);
}
