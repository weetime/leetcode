use crate::list::list_node::ListNode;
pub struct Solution {}
impl Solution {
    /**
     * 两种解题思路
     * 1.后续遍历，然后pow(2,n)
     * 2.每读取链表的一个节点值，可以认为读到的节点值是当前二进制数的最低位；读到下一个节点值的时候，需要将已经读到的结果乘以2，再将新读到的节点值当作当前二进制数的最低位；
     */

    // 时间复杂度O(n),空间复杂度O(n)
    pub fn get_decimal_value(head: Option<Box<ListNode>>) -> i32 {
        Self::helper(head, &mut 0)
    }

    pub fn helper(head: Option<Box<ListNode>>, ticket: &mut u32) -> i32 {
        match head {
            Some(head) => {
                let mut res = Self::helper(head.next, ticket);
                // res += head.val * (2_i32.pow(*ticket));
                res += head.val << *ticket;
                *ticket += 1;
                res
            }
            None => 0,
        }
    }

    // 时间复杂度O(n),空间复杂度O(1)
    pub fn get_decimal_value1(head: Option<Box<ListNode>>) -> i32 {
        if head.is_none() {
            return 0;
        }
        let mut curr = head;
        let mut res = 0;
        while let Some(node) = curr {
            curr = node.next;
            res = res * 2 + node.val;
        }
        res
    }
}

#[test]
fn test() {
    let head_a = ListNode::from_array_by_head(vec![1, 1, 0]);
    let res = Solution::get_decimal_value(head_a.clone());
    assert_eq!(res, 6);

    let res = Solution::get_decimal_value1(head_a.clone());
    assert_eq!(res, 6);
}
