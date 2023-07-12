use crate::list::list_node::ListNode;

pub struct Solution {}
impl Solution {
    /**
     * 反转是非常常规的操作，直接采用迭代的方式遍历链表，然后采用头插法重建新链表就可以了
     * 有点遗憾，Rust用递归的方式很难实现链表的反转
     */
    pub fn reverse_list(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        if head.is_none() {
            return head;
        }
        let mut curr = head;
        let mut new = None;
        while let Some(mut node) = curr.take() {
            curr = node.next;
            node.next = new;
            new = Some(node);
        }
        new
    }

    // 类似于go语言风格 引入next临时变量
    pub fn reverse_list2(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        if head.is_none() {
            return head;
        }
        let mut new = None;
        let mut curr = head;
        while let Some(mut node) = curr {
            let next = node.next;
            node.next = new;
            new = Some(node);
            curr = next;
        }
        return new;
    }
}

#[test]
fn test() {
    let head_a = ListNode::from_array_by_head(vec![1, 2, 3, 4, 5]);
    let target = ListNode::from_array_by_head(vec![5, 4, 3, 2, 1]);
    let res = Solution::reverse_list(head_a.clone());
    assert_eq!(res, target);

    let res = Solution::reverse_list2(head_a.clone());
    assert_eq!(res, target);
}
