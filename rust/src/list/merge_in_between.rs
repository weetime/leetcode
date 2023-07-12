use crate::list::list_node::ListNode;

pub struct Solution {}
impl Solution {
    pub fn merge_in_between(
        list1: Option<Box<ListNode>>,
        a: i32,
        b: i32,
        list2: Option<Box<ListNode>>,
    ) -> Option<Box<ListNode>> {
        let (mut list1, mut list2) = (list1, list2);
        let mut l1 = &mut list1;

        for _ in 0..a - 1 {
            l1 = &mut l1.as_mut().unwrap().next;
        }

        // let mut next = l1.as_mut().unwrap().next.take();
        let mut l2 = l1.as_mut().unwrap().next.take();

        for _ in 0..b - a {
            l2 = l2.as_mut().unwrap().next.take();
        }

        let mut l3 = &mut list2;
        while l3.as_ref().unwrap().next.is_some() {
            l3 = &mut l3.as_mut().unwrap().next;
        }

        l3.as_mut().unwrap().next = l2.as_mut().unwrap().next.take();
        l1.as_mut().unwrap().next = list2;

        list1
    }
}

#[test]
fn test() {
    // 注意a是0开始计数
    let head_a = ListNode::from_array_by_head(vec![1, 2, 3, 8]);
    let head_b = ListNode::from_array_by_head(vec![5, 6, 7]);
    let target = ListNode::from_array_by_head(vec![1, 5, 6, 7, 8]);
    let res = Solution::merge_in_between(head_a.clone(), 1, 2, head_b.clone());
    assert_eq!(res, target);
}
