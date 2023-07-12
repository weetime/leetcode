use crate::list::list_node::ListNode;
pub struct Solution {}
impl Solution {
    /**
     * 最大孪生和，解题思路：
     * 通过快慢指针找到中间节点(注意题目明确说明了节点数为偶数，所以链表一分为二的时候等长，但这里的解法并没有改变head的数据）
     * 反转中间节点后的链表
     * 依次遍历反转后的节点和head，取最大和即可
     */
    pub fn pair_sum(head: Option<Box<ListNode>>) -> i32 {
        let mut max = i32::MIN;
        let (mut fast, mut slow) = (head.as_ref(), head.as_ref());
        while fast.is_some() && fast.as_ref().unwrap().next.is_some() {
            fast = fast.unwrap().next.as_ref().unwrap().next.as_ref();
            slow = slow.unwrap().next.as_ref();
        }
        let mut middle = slow.unwrap().next.clone();
        let mut new_head = None;
        while let Some(mut node) = middle {
            middle = node.next.take();
            node.next = new_head;
            new_head = Some(node);
        }

        let mut curr = head.as_ref();
        while let Some(mut node) = new_head {
            new_head = node.next.take();
            max = max.max(node.val + curr.unwrap().val);
            curr = curr.unwrap().next.as_ref();
        }
        max
    }
}

#[test]
fn test() {
    let head_a = ListNode::from_array_by_head(vec![2, 5, 8, 2, 5, 7]);
    let target = 10;
    let res = Solution::pair_sum(head_a.clone());
    assert_eq!(res, target);
}
