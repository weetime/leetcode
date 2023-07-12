use crate::list::list_node::ListNode;

pub struct Solution {}
impl Solution {
    /**
     * 解题思路利用数组，采用出栈的方式遍历，然后采用头插法的构建新的链表
     */
    pub fn add_two_numbers(
        l1: Option<Box<ListNode>>,
        l2: Option<Box<ListNode>>,
    ) -> Option<Box<ListNode>> {
        let (mut l1, mut l2) = (l1, l2);
        let mut stack1: Vec<i32> = vec![];
        let mut stack2: Vec<i32> = vec![];
        let mut head = None;
        while l1.is_some() {
            l1 = l1.and_then(|x| {
                stack1.push(x.val);
                x.next
            });
        }

        while l2.is_some() {
            l2 = l2.and_then(|x| {
                stack2.push(x.val);
                x.next
            });
        }

        let mut carry: i32 = 0;
        while !stack1.is_empty() || !stack2.is_empty() || carry > 0 {
            if !stack1.is_empty() {
                match stack1.pop() {
                    Some(n) => carry += n,
                    None => todo!(),
                }
            }
            if !stack2.is_empty() {
                match stack2.pop() {
                    Some(n) => carry += n,
                    None => todo!(),
                }
            }

            let mut node = Box::new(ListNode::new(carry % 10));
            node.next = head;
            head = Some(node);
            carry = carry / 10;
        }
        head
    }
}

#[test]
fn test() {
    let head_a = ListNode::from_array_by_head(vec![7, 5, 3]);
    let head_b = ListNode::from_array_by_head(vec![2, 5, 4]);
    let target = ListNode::from_array_by_head(vec![1, 0, 0, 7]);
    let res = Solution::add_two_numbers(head_a.clone(), head_b.clone());
    assert_eq!(res, target);
}
