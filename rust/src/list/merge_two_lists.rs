use crate::list::list_node::ListNode;

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
                l1.next = merge_two_lists(l1.next, Some(l2));
                Some(l1)
            } else {
                l2.next = merge_two_lists(Some(l1), l2.next);
                Some(l2)
            }
        }
    }
}

pub fn merge_two_lists2(
    list1: Option<Box<ListNode>>,
    list2: Option<Box<ListNode>>,
) -> Option<Box<ListNode>> {
    let mut new = Some(Box::new(ListNode { val: 0, next: None }));
    let mut cur = &mut new;
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
    new.unwrap().next
}
