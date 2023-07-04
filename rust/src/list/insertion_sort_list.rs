use crate::list::list_node::ListNode;
pub fn insertion_sort_list(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
    if head.is_none() {
        return head;
    }
    let mut head = head;
    let mut dummy = ListNode { val: 0, next: None };

    while let Some(mut curr) = head {
        head = curr.next.take();
        let mut p = &mut dummy;

        while p.next.is_some() && p.next.as_ref().unwrap().val < curr.val {
            p = p.next.as_mut().unwrap();
        }

        // while let Some(ref mut prev) = p.next {
        //     if prev.val >= curr.val {
        //         break;
        //     }
        //     p = p.next.as_mut().unwrap();
        // }
        curr.next = p.next.take();
        p.next = Some(curr);
    }
    dummy.next.take()
}
