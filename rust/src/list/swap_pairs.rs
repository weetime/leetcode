use crate::list::list_node::ListNode;
pub struct Solution {}
impl Solution {
    /**
     * 题意是需要两两交换
     * 解题思路:
     * 需要一个dummy链表存储转换后的节点，需要一个指针操作dummy
     * 从head中两个两个的取节点(每一轮操作取两个节点)
     * 取到的两个节点，轮换后，采用尾插法放到dummy中即可
     * 最后返回dummy.next
     */
    pub fn swap_pairs(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        let mut dummy = ListNode { val: 0, next: None };
        let mut tail = &mut dummy;
        let mut head = head;
        while let Some(mut node1) = head {
            head = node1.next.take();
            if let Some(mut node2) = head {
                head = node2.next.take();
                node2.next = Some(node1);
                tail.next = Some(node2);
                tail = tail.next.as_mut().unwrap().next.as_mut().unwrap();
            } else {
                tail.next = Some(node1);
                tail = tail.next.as_mut().unwrap(); // 最后一步可以不要
            }
        }
        dummy.next
    }
}

#[test]
fn test() {
    let head_a = ListNode::from_array_by_head(vec![1, 2, 3, 4, 5]);
    let target = ListNode::from_array_by_head(vec![2, 1, 4, 3, 5]);
    let res = Solution::swap_pairs(head_a.clone());
    assert_eq!(res, target);
}
