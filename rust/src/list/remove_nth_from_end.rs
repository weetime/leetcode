use crate::list::list_node::ListNode;
pub struct Solution {}
impl Solution {
    /**
     * 按位置删除和按值删除都是非常常规的操作
     * 解题思路：思路其实很简单，采用快慢指针，快指针先移动n步，然后快慢指针再同时移动
     * 当快指针到达结尾的时候，慢指针指向后驱节点的下一个节点
     * 由于Rust语言的特性，写起来不太优雅，这种常见的双指针可以采用unsafe模式，这样和go版本类似了
     */
    pub fn remove_nth_from_end(head: Option<Box<ListNode>>, n: i32) -> Option<Box<ListNode>> {
        let mut head = head;
        let mut n = n;
        let mut dummy = Some(Box::new(ListNode {
            val: 0,
            next: head.clone(), // 这里不优雅 clone了一份
        }));
        let mut curr = head.as_mut();
        let mut prev = &mut dummy;
        while n > 0 {
            // todo 这里可以改用成for in 的模式 for _ in 0..n{}
            curr = curr.unwrap().next.as_mut();
            n -= 1;
        }
        while let Some(node) = curr.take() {
            curr = node.next.as_mut();
            prev = &mut prev.as_mut().unwrap().next;
        }

        prev.as_mut().unwrap().next = prev.as_mut().unwrap().next.take().unwrap().next;
        dummy.unwrap().next
    }

    // 用快慢指针来解决，快指针先跑N步之后，慢指针开始跑，等快指针跑到头，慢指针刚好到了倒是第N个节点
    pub fn remove_nth_from_end1(head: Option<Box<ListNode>>, n: i32) -> Option<Box<ListNode>> {
        let mut dummy = Box::new(ListNode { val: 0, next: head });
        unsafe {
            // 绑定两个指针，就不用clone了
            let mut fast = &mut dummy as *mut Box<ListNode>;
            let mut slow = &mut dummy as *mut Box<ListNode>;
            for _ in 0..n {
                fast = (*fast).next.as_mut().unwrap();
            }
            while (*fast).next.is_some() {
                fast = (*fast).next.as_mut().unwrap();
                slow = (*slow).next.as_mut().unwrap();
            }
            (*slow).next = (*slow).next.take().unwrap().next;
        }
        dummy.next
    }
}

#[test]
fn test() {
    let head_a = ListNode::from_array_by_head(vec![1, 1, 2, 3, 4]);
    let target = ListNode::from_array_by_head(vec![1, 1, 2, 4]);
    let res = Solution::remove_nth_from_end(head_a.clone(), 2);
    // println!("{:?}", res);
    assert_eq!(res, target);

    let res = Solution::remove_nth_from_end1(head_a.clone(), 2);
    assert_eq!(res, target);
}
