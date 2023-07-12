use crate::list::list_node::ListNode;

pub struct Solution {}
impl Solution {
    /**
     * 回文链表一般有三种解题思路
     * 1.遍历一遍链表放到数组中，采用双指针，一个头，一个尾依次取出比对
     * 2、clone一份链表并反转，然后依次比对两个链表节点的值
     * 3. 快慢指针，找到链表的中点，然后反转后一部分，再遍历反转后的链表和原始链表，比对节点值（这里涉及到奇数节点的情况，处理起来比较复杂，就不实现了）
     */

    // 做一次后续遍历，由于Rust可修改的全局变量不好实现，故而改成传指针的形式
    pub fn is_palindrome(head: Option<Box<ListNode>>) -> bool {
        if head.is_none() {
            return true;
        }
        Self::dfs(head.clone(), &mut head.as_ref())
    }

    fn dfs(head_a: Option<Box<ListNode>>, head_b: &mut Option<&Box<ListNode>>) -> bool {
        match head_a {
            None => true,
            Some(h1) => {
                let mut res = Self::dfs(h1.next, head_b);
                res = res && h1.val == head_b.unwrap().val;
                *head_b = head_b.unwrap().next.as_ref();
                res
            }
        }
    }

    // 这种解法就比较容易看懂，反转一下链表，再比对
    pub fn is_palindrome2(head: Option<Box<ListNode>>) -> bool {
        if head.is_none() {
            return true;
        }
        let new = Self::reverse(head.clone());
        Self::check(head, new)
    }

    fn check(hea_a: Option<Box<ListNode>>, head_b: Option<Box<ListNode>>) -> bool {
        match (hea_a, head_b) {
            (None, None) => true,
            (None, Some(_)) => false,
            (Some(_), None) => false,
            (Some(l1), Some(l2)) => {
                if l1.val != l2.val {
                    false
                } else {
                    Self::check(l1.next, l2.next)
                }
            }
        }
    }

    // 反转其实是头插法构建链表
    fn reverse(mut head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        let mut new = None;
        while let Some(mut node) = head.take() {
            head = node.next.take();
            node.next = new;
            new = Some(node);
        }
        new
    }
}

#[test]
fn test() {
    let head_a = ListNode::from_array_by_head(vec![1, 2, 3, 2, 1]);
    let res = Solution::is_palindrome(head_a.clone());
    assert_eq!(res, true);

    let res = Solution::is_palindrome2(head_a.clone());
    assert_eq!(res, true);
}
