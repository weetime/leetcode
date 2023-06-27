use crate::list::list_node::ListNode;

pub fn has_cycle(head: Option<Box<ListNode>>) -> bool {
    let mut slow = head.as_ref();
    let mut fast = head.as_ref();
    while fast.as_ref().unwrap().next.is_some()
        && fast.as_ref().unwrap().next.as_ref().unwrap().next.is_some()
    {
        fast = fast.unwrap().next.as_ref().unwrap().next.as_ref();
        slow = slow.unwrap().next.as_ref();
        if fast.unwrap().as_ref() == slow.unwrap().as_ref() {
            return true;
        }
    }
    return false;
}
