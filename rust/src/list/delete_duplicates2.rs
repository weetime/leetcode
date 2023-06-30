use crate::list::list_node::ListNode;
pub fn delete_duplicates(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
    if head.is_none() {
        return head;
    }
    let mut dummy = Some(Box::new(ListNode { val: 0, next: head }));
    let mut curr = &mut dummy;
    while curr.as_ref().unwrap().next.as_ref().is_some()
        && curr
            .as_ref()
            .unwrap()
            .next
            .as_ref()
            .unwrap()
            .next
            .as_ref()
            .is_some()
    {
        if curr.as_ref().unwrap().next.as_ref().unwrap().val
            == curr
                .as_ref()
                .unwrap()
                .next
                .as_ref()
                .unwrap()
                .next
                .as_ref()
                .unwrap()
                .val
        {
            let val = curr.as_ref().unwrap().next.as_ref().unwrap().val;
            while curr.as_ref().unwrap().next.is_some()
                && curr.as_ref().unwrap().next.as_ref().unwrap().val == val
            {
                curr.as_mut().unwrap().next = curr.as_mut().unwrap().next.take().unwrap().next;
            }
        } else {
            curr = &mut curr.as_mut().unwrap().next;
        }
    }

    dummy.unwrap().next
}

/* pub fn delete_duplicates1(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
    if head.is_none() {
        return head;
    }
    let mut head = head;
    let mut dummy = Some(Box::new(ListNode {
        val: 0,
        next: head.clone(),
    }));
    let mut fast = &mut head;
    let mut slow = &mut dummy;
    while let Some(mut node) = fast.take() {
        while node.next.as_ref().is_some()
            && node.val == node.next.as_ref().unwrap().next.as_ref().unwrap().val
        {
            node = node.next.take().unwrap();
        }
        if slow.as_ref().unwrap().next == fast.as_ref().unwrap().next {
            let mut temp = fast.take(); // 类似于实现了copy
            slow.as_mut().unwrap().next = temp.as_mut().unwrap().next.take();
        } else {
            slow = &mut slow.as_mut().unwrap().next;
        }
        fast = &mut fast.as_mut().unwrap().next;
    }
    dummy.unwrap().next
} */

pub fn delete_duplicates2(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
    let mut dummy = Some(Box::new(ListNode { val: 0, next: head }));
    let mut cur = dummy.as_mut();
    while cur.as_ref().unwrap().next.is_some()
        && cur.as_ref().unwrap().next.as_ref().unwrap().next.is_some()
    {
        if cur.as_ref().unwrap().next.as_ref().unwrap().val
            == cur
                .as_ref()
                .unwrap()
                .next
                .as_ref()
                .unwrap()
                .next
                .as_ref()
                .unwrap()
                .val
        {
            let x = cur.as_ref().unwrap().next.as_ref().unwrap().val;
            while cur.as_ref().unwrap().next.is_some()
                && cur.as_ref().unwrap().next.as_ref().unwrap().val == x
            {
                cur.as_mut().unwrap().next =
                    cur.as_mut().unwrap().next.as_mut().unwrap().next.take();
            }
        } else {
            cur = cur.unwrap().next.as_mut();
        }
    }
    dummy.unwrap().next
}

pub fn delete_duplicates3(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
    let mut head = head;
    let mut new = ListNode { val: 0, next: None };
    let mut curr = &mut new;
    let mut pre_v = head.as_ref().unwrap().val - 1;
    while let Some(mut node) = head {
        head = node.next.take();
        if (head.is_some() && head.as_ref().unwrap().val == node.val) || node.val == pre_v {
            pre_v = node.val;
        } else {
            pre_v = node.val;
            curr.next = Some(node);
            curr = curr.next.as_mut().unwrap();
        }
    }
    new.next
}
