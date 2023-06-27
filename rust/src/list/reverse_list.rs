use crate::list::list_node::ListNode;
pub fn reverse_list(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
    if head.is_none() {
        return head;
    }
    let mut curr = head;
    let mut new = None;
    while let Some(mut node) = curr.take() {
        curr = node.next;
        node.next = new;
        new = Some(node);
    }
    new
}

pub fn reverse_list2(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
    if head.is_none() {
        return head;
    }
    let mut new = None;
    let mut curr = head;
    while let Some(mut node) = curr {
        let next = node.next;
        node.next = new;
        new = Some(node);
        curr = next;
    }
    return new;
}
