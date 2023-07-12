use crate::list::list_node::ListNode;
pub struct Solution {}
impl Solution {
    /**
     * 插入排序，依次遍历head取节点的val，然后遍历已经排好序的链表对比，找到恰好比val大的一个前驱节点，将val节点插入进去即可
     * 插入排序最坏时间复杂度O(n²)，因为存的是指针空间复杂度O(1)
     * 可以采用归并排序，优化时间复杂度为O(nlogn)，空间复杂度可以控制在O(1)
     */
    pub fn insertion_sort_list(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        if head.is_none() {
            return head;
        }
        let mut head = head;
        let mut dummy = ListNode { val: 0, next: None };

        while let Some(mut curr) = head {
            head = curr.next.take();
            let mut p = &mut dummy;

            while p.next.is_some() && p.next.as_ref().unwrap().val < curr.val {
                p = p.next.as_mut().unwrap();
            }

            // while let Some(ref mut prev) = p.next {
            //     if prev.val >= curr.val {
            //         break;
            //     }
            //     p = p.next.as_mut().unwrap();
            // }
            curr.next = p.next.take();
            p.next = Some(curr);
        }
        dummy.next.take()
    }
}

#[test]
fn test() {
    let head_a = ListNode::from_array_by_head(vec![3, 1, 2, 5, 4]);
    let target = ListNode::from_array_by_head(vec![1, 2, 3, 4, 5]);
    let res = Solution::insertion_sort_list(head_a.clone());
    assert_eq!(res, target);
}
