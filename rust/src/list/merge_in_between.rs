use crate::list::list_node::ListNode;
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
