use crate::list::list_node::ListNode;
pub fn remove_nth_from_end(head: Option<Box<ListNode>>, n: i32) -> Option<Box<ListNode>> {
    let mut head = head;
    let mut n = n;
    let mut sentinel = Some(Box::new(ListNode {
        val: 0,
        next: head.clone(), // 这里不优雅 clone了一份
    }));
    let mut curr = head.as_mut();
    let mut prev = &mut sentinel;
    while n > 0 {
        // todo 这里可以改用成for in 的模式 for _ in 0..n{}
        curr = curr.unwrap().next.as_mut();
        n -= 1;
    }
    while let Some(node) = curr.take() {
        curr = node.next.as_mut();
        prev = &mut prev.as_mut().unwrap().next;
    }

    prev.as_mut().unwrap().next = prev.as_mut().unwrap().next.take().unwrap().next;
    sentinel.unwrap().next
}

pub fn remove_nth_from_end2(head: Option<Box<ListNode>>, n: i32) -> Option<Box<ListNode>> {
    let mut dynamic = Box::new(ListNode { val: 0, next: head });
    unsafe {
        // 绑定两个指针，就不用clone了
        let mut fast = &mut dynamic as *mut Box<ListNode>;
        let mut slow = &mut dynamic as *mut Box<ListNode>;
        for _ in 0..n {
            fast = (*fast).next.as_mut().unwrap();
        }
        while (*fast).next.is_some() {
            fast = (*fast).next.as_mut().unwrap();
            slow = (*slow).next.as_mut().unwrap();
        }
        (*slow).next = (*slow).next.take().unwrap().next;
    }
    dynamic.next
}
