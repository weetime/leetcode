#[derive(PartialEq, Eq, Clone, Debug)]
pub struct ListNode {
    pub val: i32,
    pub next: Option<Box<ListNode>>,
}

impl ListNode {
    #[inline]
    fn new(val: i32) -> Self {
        ListNode { next: None, val }
    }
}

pub fn sample() -> Option<Box<ListNode>> {
    let mut node1: ListNode = ListNode::new(1);
    let mut node2: ListNode = ListNode::new(2);
    let node3: ListNode = ListNode::new(3);
    node2.next = Some(Box::new(node3));
    node1.next = Some(Box::new(node2));
    return Some(Box::new(node1));
}
