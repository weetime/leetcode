use crate::list::list_node::ListNode;

pub fn remove_elements(head: Option<Box<ListNode>>, val: i32) -> Option<Box<ListNode>> {
    if head.is_none() {
        return head;
    }
    let mut sentinel = ListNode { val: 0, next: head };
    let mut prev = &mut sentinel;
    while let Some(node) = prev.next.take() {
        if node.val == val {
            prev.next = node.next;
        } else {
            prev.next = Some(node);
            prev = prev.next.as_mut().unwrap();
        }
    }

    sentinel.next
}

pub fn remove_elements2(head: Option<Box<ListNode>>, val: i32) -> Option<Box<ListNode>> {
    if head.is_none() {
        return head;
    }

    let mut sentinel = ListNode { val: 0, next: head };
    let mut prev = &mut sentinel;

    while let Some(ref mut node) = prev.next {
        if node.val == val {
            prev.next = node.next.take();
        } else {
            prev = prev.next.as_mut().unwrap();
        }
    }

    sentinel.next
}

// dfs 方式
pub fn remove_elements3(head: Option<Box<ListNode>>, val: i32) -> Option<Box<ListNode>> {
    match head {
        Some(mut head) => {
            let next_node = remove_elements3(head.next.take(), val);
            if head.val == val {
                next_node
            } else {
                head.next = next_node;
                Some(head)
            }
        }
        None => None,
    }
}
