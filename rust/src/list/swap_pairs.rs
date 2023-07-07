use crate::list::list_node::ListNode;
pub fn swap_pairs1(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
    let mut remain = head;
    let mut result = ListNode::new(0);
    let mut tail = &mut result;
    while let Some(mut n1) = remain {
        remain = n1.next.take(); // take()将n1打断，这样n1只有一个值，返回值是除n1节点外的剩余节点
        if let Some(mut n2) = remain {
            remain = n2.next.take(); // take()将n2打断，n2只有一个值，返回值是除n2节点外的剩余节点
            n2.next = Some(n1);
            tail.next = Some(n2);
            tail = tail.next.as_mut().unwrap().next.as_mut().unwrap();
        } else {
            tail.next = Some(n1);
            tail = tail.next.as_mut().unwrap();
        }
    }
    result.next
}
pub fn swap_pairs(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
    let mut dummy = Box::new(ListNode { val: 0, next: None });
    let mut temp = dummy.as_mut();
    let mut head = head;
    while let Some(mut node1) = head {
        head = node1.next.take();
        if let Some(mut node2) = head {
            head = node2.next.take();
            node2.next = Some(node1);
            temp.next = Some(node2);
            temp = temp.next.as_mut().unwrap().next.as_mut().unwrap();
        } else {
            temp.next = Some(node1);
            temp = temp.next.as_mut().unwrap();
        }
    }
    dummy.next
}
