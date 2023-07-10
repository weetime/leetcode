use crate::list::list_node::ListNode;
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
