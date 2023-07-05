use crate::list::list_node::ListNode;
pub fn delete_middle(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
    let mut dummy = Box::new(ListNode { val: 0, next: head });
    unsafe {
        // 绑定两个指针，就不用clone了
        let mut fast = &mut dummy as *mut Box<ListNode>;
        let mut slow = &mut dummy as *mut Box<ListNode>;
        while (*fast).next.is_some() && (*fast).next.as_ref().unwrap().next.is_some() {
            fast = (*fast).next.as_mut().unwrap().next.as_mut().unwrap();
            slow = (*slow).next.as_mut().unwrap();
        }
        (*slow).next = (*slow).next.take().unwrap().next;
    }
    dummy.next
}

pub fn delete_middle1(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
    let mut head = head;
    let mut prev = std::ptr::null_mut();
    let mut slow = &mut head as *mut Option<Box<ListNode>>;
    let mut fast = head.as_ref();
    while fast.is_some() && fast.as_ref().unwrap().next.is_some() {
        fast = fast.as_ref().unwrap().next.as_ref().unwrap().next.as_ref();
        prev = slow;
        slow = unsafe {
            &mut slow.as_mut().unwrap().as_mut().unwrap().next as *mut Option<Box<ListNode>>
        };
    }

    unsafe { (*prev).as_mut().unwrap().next = slow.as_mut().unwrap().as_mut().unwrap().next.take() }
    head
}
