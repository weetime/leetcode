use crate::list::list_node::ListNode;

pub fn rotate_right(head: Option<Box<ListNode>>, k: i32) -> Option<Box<ListNode>> {
    if head.is_none() || head.as_ref().unwrap().next.is_none() {
        return head;
    }
    let mut head: Option<Box<ListNode>> = head;
    let mut length: i32 = 0;
    let mut ptr = head.as_ref();
    while let Some(node) = ptr {
        ptr = node.next.as_ref().take();
        length += 1;
    }
    if k % length == 0 {
        return head;
    }
    let k: i32 = length - (k % length);
    let mut ptr = &mut head;
    for _ in 1..k {
        ptr = &mut ptr.as_mut().unwrap().next;
    }

    let mut new = ptr.as_mut().unwrap().next.take();
    let mut tail = &mut new;
    while tail.is_some() && tail.as_ref().unwrap().next.is_some() {
        tail = &mut tail.as_mut().unwrap().next;
    }
    tail.as_mut().unwrap().next = head;

    new
}
