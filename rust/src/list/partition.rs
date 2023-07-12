use crate::list::list_node::ListNode;
pub struct Solution {}
impl Solution {
    /**
     * 这个题目只要读懂意思就好办
     * 借助于大小链表和双指针，给定一个数x，依次遍历链表，遇到节点的值大于等于x就放到大链表中，如果遇到小于x的就放到小链表中
     * 遍历一遍后，将大链表拼接到小链表上就可以了
     * 最后返回小链表即可 时间复杂度O(n),空间复杂度O(1)
     */
    pub fn partition(head: Option<Box<ListNode>>, x: i32) -> Option<Box<ListNode>> {
        if head.is_none() || head.as_ref().unwrap().next.is_none() {
            return head;
        }
        let mut head = head;
        let mut small = ListNode::new(0);
        let mut large = ListNode::new(0);
        let (mut s, mut l) = (&mut small, &mut large);
        while let Some(mut node) = head {
            head = node.next.take();
            match x {
                x if x <= node.val => {
                    l.next = Some(node);
                    l = l.next.as_mut().unwrap().as_mut();
                }
                _ => {
                    s.next = Some(node);
                    s = s.next.as_mut().unwrap().as_mut();
                }
            }
        }
        s.next = large.next;
        small.next
    }
}

#[test]
fn test() {
    let head_a = ListNode::from_array_by_head(vec![2, 5, 8, 2, 5, 7]);
    let target = ListNode::from_array_by_head(vec![2, 2, 5, 8, 5, 7]);
    let res = Solution::partition(head_a.clone(), 5);
    assert_eq!(res, target);
}
