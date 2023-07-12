use crate::list::list_node::ListNode;

pub struct Solution {}
impl Solution {
    /**
     * 其实这道题目经典的解法，是构建一个环形链表，然后遍历k%length后，再断开链表即可
     * 这里的解题思路：
     * 为了避免遍历多次链表，首先需要计算链表的长度length，然后再计算k%length，如果取模后为0，则直接返回head
     * 根据题意，是需要找到合适的位置，断开链表，然后再把断开的后半部分拼接到head头上，返回拼接后的链表即可
     */
    pub fn rotate_right(head: Option<Box<ListNode>>, k: i32) -> Option<Box<ListNode>> {
        if head.is_none() || head.as_ref().unwrap().next.is_none() {
            return head;
        }
        let mut head: Option<Box<ListNode>> = head;
        let mut length: i32 = 0;
        let mut ptr = head.as_ref();
        while let Some(node) = ptr {
            ptr = node.next.as_ref().take();
            length += 1;
        }
        if k % length == 0 {
            return head;
        }
        let k: i32 = length - (k % length);
        let mut ptr = &mut head;
        for _ in 1..k {
            ptr = &mut ptr.as_mut().unwrap().next;
        }

        let mut new = ptr.as_mut().unwrap().next.take();
        let mut tail = &mut new;
        while tail.is_some() && tail.as_ref().unwrap().next.is_some() {
            tail = &mut tail.as_mut().unwrap().next;
        }
        tail.as_mut().unwrap().next = head;

        new
    }
}

#[test]
fn test() {
    let head_a = ListNode::from_array_by_head(vec![1, 2, 3, 4, 5]);
    let target = ListNode::from_array_by_head(vec![5, 1, 2, 3, 4]);
    let res = Solution::rotate_right(head_a.clone(), 101);
    assert_eq!(res, target);
}
