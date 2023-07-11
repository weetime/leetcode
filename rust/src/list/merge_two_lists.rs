use crate::list::list_node::ListNode;

pub struct Solution {}
impl Solution {
    // 合并两个有序链表变成一个新的有序链表
    pub fn merge_two_lists(
        list1: Option<Box<ListNode>>,
        list2: Option<Box<ListNode>>,
    ) -> Option<Box<ListNode>> {
        match (list1, list2) {
            (None, None) => None,
            (None, l2) => l2,
            (l1, None) => l1,
            (Some(mut l1), Some(mut l2)) => {
                if l1.val <= l2.val {
                    l1.next = Self::merge_two_lists(l1.next, Some(l2));
                    Some(l1)
                } else {
                    l2.next = Self::merge_two_lists(Some(l1), l2.next);
                    Some(l2)
                }
            }
        }
    }

    pub fn merge_two_lists1(
        list1: Option<Box<ListNode>>,
        list2: Option<Box<ListNode>>,
    ) -> Option<Box<ListNode>> {
        let mut dummy = Some(Box::new(ListNode { val: 0, next: None }));
        let mut cur = &mut dummy;
        let (mut p1, mut p2) = (list1, list2);
        while p1.is_some() && p2.is_some() {
            let mut temp;
            if p1.as_ref().unwrap().val < p2.as_ref().unwrap().val {
                temp = p1.take(); // 类似于实现了copy
                p1 = temp.as_mut().unwrap().next.take();
            } else {
                temp = p2.take();
                p2 = temp.as_mut().unwrap().next.take();
            }
            cur.as_mut().unwrap().next = temp;
            cur = &mut cur.as_mut().unwrap().next;
        }
        if p1.is_some() {
            cur.as_mut().unwrap().next = p1;
        } else {
            cur.as_mut().unwrap().next = p2;
        }
        dummy.unwrap().next
    }

    pub fn merge_two_lists2(
        list1: Option<Box<ListNode>>,
        list2: Option<Box<ListNode>>,
    ) -> Option<Box<ListNode>> {
        let mut dummy = ListNode { val: 0, next: None };
        let mut cur = &mut dummy;
        let (mut p1, mut p2) = (list1, list2);
        while p1.is_some() && p2.is_some() {
            let mut temp;
            if p1.as_ref().unwrap().val < p2.as_ref().unwrap().val {
                temp = p1.take(); // 类似于实现了copy
                p1 = temp.as_mut().unwrap().next.take();
            } else {
                temp = p2.take();
                p2 = temp.as_mut().unwrap().next.take();
            }
            cur.next = temp;
            cur = cur.next.as_mut().unwrap().as_mut();
        }
        if p1.is_some() {
            cur.next = p1;
        } else {
            cur.next = p2;
        }
        dummy.next
    }
}

#[test]
fn test() {
    let head_a = ListNode::from_array_by_head(vec![1, 3, 5]);
    let head_b = ListNode::from_array_by_head(vec![2, 4, 6]);
    let target = ListNode::from_array_by_head(vec![1, 2, 3, 4, 5, 6]);
    let res = Solution::merge_two_lists(head_a.clone(), head_b.clone());
    // println!("{:?}", res);
    assert_eq!(res, target);

    let res = Solution::merge_two_lists1(head_a.clone(), head_b.clone());
    assert_eq!(res, target);

    let res = Solution::merge_two_lists2(head_a.clone(), head_b.clone());
    assert_eq!(res, target);
}
