use crate::list::list_node::ListNode;
pub fn partition(head: Option<Box<ListNode>>, x: i32) -> Option<Box<ListNode>> {
    if head.is_none() || head.as_ref().unwrap().next.is_none() {
        return head;
    }
    let mut head = head;
    let mut small = ListNode::new(0);
    let mut large = ListNode::new(0);
    let (mut s, mut l) = (&mut small, &mut large);
    while let Some(mut node) = head {
        head = node.next.take();
        match x {
            x if x <= node.val => {
                l.next = Some(node);
                l = l.next.as_mut().unwrap().as_mut();
            }
            _ => {
                s.next = Some(node);
                s = s.next.as_mut().unwrap().as_mut();
            }
        }
    }
    s.next = large.next;
    small.next
}
