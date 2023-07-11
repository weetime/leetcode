use crate::list::list_node::ListNode;
pub struct Solution {}
impl Solution {
    /**
     * 解题的关键是需要引入一个哑节点，剩下的和203题类似了
     * 不同的地方在于，值相同的时候是否保留一个节点，一个都不保留的时候，需要持续遍历，直到遇到不同的值为止
     * 这种写法就是代码看起来不优雅
     * delete_duplicates 和 delete_duplicates1 没有区别，只是为了练习所有权
     */
    pub fn delete_duplicates(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        if head.is_none() {
            return head;
        }
        let mut dummy = Some(Box::new(ListNode { val: 0, next: head }));
        let mut curr = &mut dummy;
        while curr.as_ref().unwrap().next.as_ref().is_some()
            && curr.as_ref().unwrap().next.as_ref().unwrap().next.is_some()
        {
            if curr.as_ref().unwrap().next.as_ref().unwrap().val
                == curr
                    .as_ref()
                    .unwrap()
                    .next
                    .as_ref()
                    .unwrap()
                    .next
                    .as_ref()
                    .unwrap()
                    .val
            {
                let val = curr.as_ref().unwrap().next.as_ref().unwrap().val;
                while curr.as_ref().unwrap().next.is_some()
                    && curr.as_ref().unwrap().next.as_ref().unwrap().val == val
                {
                    curr.as_mut().unwrap().next = curr.as_mut().unwrap().next.take().unwrap().next;
                }
            } else {
                curr = &mut curr.as_mut().unwrap().next;
            }
        }

        dummy.unwrap().next
    }

    pub fn delete_duplicates1(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        if head.is_none() {
            return head;
        }
        let mut dummy = Some(Box::new(ListNode { val: 0, next: head }));
        let mut cur = dummy.as_mut();
        while cur.as_ref().unwrap().next.is_some()
            && cur.as_ref().unwrap().next.as_ref().unwrap().next.is_some()
        {
            if cur.as_ref().unwrap().next.as_ref().unwrap().val
                == cur
                    .as_ref()
                    .unwrap()
                    .next
                    .as_ref()
                    .unwrap()
                    .next
                    .as_ref()
                    .unwrap()
                    .val
            {
                let x = cur.as_ref().unwrap().next.as_ref().unwrap().val;
                while cur.as_ref().unwrap().next.is_some()
                    && cur.as_ref().unwrap().next.as_ref().unwrap().val == x
                {
                    cur.as_mut().unwrap().next =
                        cur.as_mut().unwrap().next.as_mut().unwrap().next.take();
                }
            } else {
                cur = cur.unwrap().next.as_mut();
            }
        }
        dummy.unwrap().next
    }

    // 另外的一个解题思路就是重建一个新的链表，只保留值不重复的节点
    pub fn delete_duplicates2(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        if head.is_none() {
            return head;
        }
        let mut head = head;
        let mut dummy = ListNode { val: 0, next: None };
        let mut curr = &mut dummy;
        let mut pre_v = -101;
        while let Some(mut node) = head {
            head = node.next.take();
            if (head.is_some() && head.as_ref().unwrap().val == node.val) || node.val == pre_v {
                pre_v = node.val;
            } else {
                pre_v = node.val;
                curr.next = Some(node);
                curr = curr.next.as_mut().unwrap();
            }
        }
        dummy.next
    }
}

#[test]
fn test() {
    let head_a = ListNode::from_array_by_head(vec![1, 1, 2, 3, 3, 4]);
    let target = ListNode::from_array_by_head(vec![2, 4]);
    let res = Solution::delete_duplicates(head_a.clone());
    // println!("{:?}", res);
    assert_eq!(res, target);

    let res = Solution::delete_duplicates1(head_a.clone());
    assert_eq!(res, target);

    let res = Solution::delete_duplicates2(head_a.clone());
    assert_eq!(res, target);
}
