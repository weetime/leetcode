use crate::list::list_node::ListNode;
pub struct Solution {}
impl Solution {
    // 典型的解题思路就是采用快慢双指针
    pub fn middle_node(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        let mut slow_ptr = head.as_ref();
        let mut fast_ptr = head.as_ref();

        while fast_ptr.is_some() && fast_ptr.as_ref()?.next.is_some() {
            fast_ptr = fast_ptr?.next.as_ref()?.next.as_ref();
            slow_ptr = slow_ptr?.next.as_ref();
        }

        slow_ptr.cloned()
    }

    // clone一次head
    pub fn middle_node2(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        let mut slow_ptr = head.clone();
        let mut fast_ptr = head;

        while fast_ptr.is_some() && fast_ptr.as_ref().unwrap().next.is_some() {
            fast_ptr = fast_ptr.unwrap().next.unwrap().next;
            slow_ptr = slow_ptr.unwrap().next;
        }
        slow_ptr
    }

    // 也可以用递归的方式
    pub fn middle_node3(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        Self::helper(head.clone(), head.clone())
    }

    pub fn helper(
        head_a: Option<Box<ListNode>>,
        head_b: Option<Box<ListNode>>,
    ) -> Option<Box<ListNode>> {
        match (head_a, head_b) {
            (None, None) => None,
            (None, slow_ptr) => slow_ptr,
            (Some(_), None) => None,
            (Some(fast_ptr), Some(slow_ptr)) => {
                if fast_ptr.next.is_some() {
                    Self::helper(fast_ptr.next.unwrap().next, slow_ptr.next)
                } else {
                    Some(slow_ptr)
                }
            }
        }
    }
}

#[test]
fn test() {
    let head_a = ListNode::from_array_by_head(vec![1, 2, 3, 3, 2, 1]);
    let target = ListNode::from_array_by_head(vec![3, 2, 1]);
    let res = Solution::middle_node(head_a.clone());
    assert_eq!(res, target);

    let head_a = ListNode::from_array_by_head(vec![1, 2, 3, 2, 1]);
    let target = ListNode::from_array_by_head(vec![3, 2, 1]);
    let res = Solution::middle_node2(head_a.clone());
    assert_eq!(res, target);

    let res = Solution::middle_node3(head_a.clone());
    assert_eq!(res, target);
}
