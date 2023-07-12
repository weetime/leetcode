use crate::list::list_node::ListNode;
pub struct Solution {}
impl Solution {
    /**
     * 解题思路，正常遍历两个链表，用相加结果res构建新的链表，如果res大于10，则需要携带进位，如果遍历到链表结尾，进位依然存在，则需要再追加一个节点
     */
    pub fn add_two_numbers(
        l1: Option<Box<ListNode>>,
        l2: Option<Box<ListNode>>,
    ) -> Option<Box<ListNode>> {
        let mut dummy = ListNode { val: 0, next: None };
        let mut ptr = &mut dummy;
        let (mut p1, mut p2) = (l1, l2);
        let mut carry: i32 = 0;

        while p1.is_some() || p2.is_some() || carry > 0 {
            let (mut n1, mut n2) = (0, 0);
            if p1.is_some() {
                n1 = p1.as_ref()?.val;
                p1 = p1?.next.take();
            }
            if p2.is_some() {
                n2 = p2.as_ref()?.val;
                p2 = p2?.next.take();
            }
            let sum = n1 + n2 + carry;
            carry = sum / 10;
            ptr.next = Some(Box::new(ListNode {
                val: sum % 10,
                next: None,
            }));
            ptr = ptr.next.as_mut()?.as_mut();
        }

        dummy.next
    }

    // 用Option的特性，练习and_then的用法
    pub fn add_two_numbers1(
        l1: Option<Box<ListNode>>,
        l2: Option<Box<ListNode>>,
    ) -> Option<Box<ListNode>> {
        let mut vec = vec![];
        let mut add_one = 0;
        let mut l1 = l1;
        let mut l2 = l2;

        while l1.is_some() || l2.is_some() || add_one > 0 {
            l1 = l1.and_then(|x| {
                add_one += x.val;
                x.next
            });
            l2 = l2.and_then(|x| {
                add_one += x.val;
                x.next
            });
            let curr = add_one % 10;
            add_one /= 10;
            vec.push(ListNode::new(curr));
        }

        // 采用头插法构建链表
        let mut tail = None;
        while let Some(mut node) = vec.pop() {
            node.next = tail;
            tail = Some(Box::new(node));
        }
        tail
    }
}

#[test]
fn test() {
    let head_a = ListNode::from_array_by_head(vec![7, 5, 3]);
    let head_b = ListNode::from_array_by_head(vec![2, 5, 4]);
    let target = ListNode::from_array_by_head(vec![9, 0, 8]);
    let res = Solution::add_two_numbers(head_a.clone(), head_b.clone());
    assert_eq!(res, target);

    let res = Solution::add_two_numbers1(head_a.clone(), head_b.clone());
    assert_eq!(res, target);
}
