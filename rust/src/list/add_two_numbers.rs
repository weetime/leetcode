use crate::list::list_node::ListNode;
pub fn add_two_numbers3(
    l1: Option<Box<ListNode>>,
    l2: Option<Box<ListNode>>,
) -> Option<Box<ListNode>> {
    let mut dummy = Some(Box::new(ListNode { val: 0, next: None }));
    let mut ptr = dummy.as_mut();
    let (mut l1, mut l2) = (l1, l2);
    let (mut p1, mut p2) = (&mut l1, &mut l2);
    let mut carry: i32 = 0;
    while p1.is_some() || p2.is_some() {
        let (mut n1, mut n2) = (0, 0);
        if p1.is_some() {
            n1 = p1.as_ref().unwrap().val;
            p1 = &mut p1.as_mut().unwrap().next;
        }
        if p2.is_some() {
            n2 = p2.as_ref().unwrap().val;
            p2 = &mut p2.as_mut().unwrap().next;
        }
        let mut sum = n1 + n2 + carry;
        (sum, carry) = (sum % 10, sum / 10);
        ptr.as_mut().unwrap().next = Some(Box::new(ListNode {
            val: sum,
            next: None,
        }));
        ptr = ptr.unwrap().next.as_mut();
    }
    if carry > 0 {
        ptr.as_mut().unwrap().next = Some(Box::new(ListNode {
            val: carry,
            next: None,
        }));
    }
    dummy.unwrap().next
}

pub fn add_two_numbers2(
    l1: Option<Box<ListNode>>,
    l2: Option<Box<ListNode>>,
) -> Option<Box<ListNode>> {
    let mut dummy = ListNode { val: 0, next: None };
    let mut ptr = &mut dummy;
    let (mut l1, mut l2) = (l1, l2);
    let (mut p1, mut p2) = (&mut l1, &mut l2);
    let mut carry: i32 = 0;
    while p1.is_some() || p2.is_some() {
        let (mut n1, mut n2) = (0, 0);
        if p1.is_some() {
            n1 = p1.as_ref().unwrap().val;
            p1 = &mut p1.as_mut().unwrap().next;
        }
        if p2.is_some() {
            n2 = p2.as_ref().unwrap().val;
            p2 = &mut p2.as_mut().unwrap().next;
        }
        let mut sum = n1 + n2 + carry;
        (sum, carry) = (sum % 10, sum / 10);
        ptr.next = Some(Box::new(ListNode {
            val: sum,
            next: None,
        }));
        ptr = ptr.next.as_mut().unwrap().as_mut();
    }
    if carry > 0 {
        ptr.next = Some(Box::new(ListNode {
            val: carry,
            next: None,
        }));
    }
    dummy.next
}

pub fn add_two_numbers(
    l1: Option<Box<ListNode>>,
    l2: Option<Box<ListNode>>,
) -> Option<Box<ListNode>> {
    let mut dummy = ListNode { val: 0, next: None };
    let mut ptr = &mut dummy;
    let (mut p1, mut p2) = (l1, l2);
    let mut carry: i32 = 0;

    while p1.is_some() || p2.is_some() || carry > 0 {
        let (mut n1, mut n2) = (0, 0);
        if p1.is_some() {
            n1 = p1.as_ref()?.val;
            p1 = p1?.next.take();
        }
        if p2.is_some() {
            n2 = p2.as_ref()?.val;
            p2 = p2?.next.take();
        }
        let mut sum = n1 + n2 + carry;
        carry = sum / 10;
        sum = sum % 10;
        ptr.next = Some(Box::new(ListNode {
            val: sum,
            next: None,
        }));
        ptr = ptr.next.as_mut()?.as_mut();
    }

    dummy.next
}

pub fn add_two_numbers4(
    l1: Option<Box<ListNode>>,
    l2: Option<Box<ListNode>>,
) -> Option<Box<ListNode>> {
    let mut vec = vec![];
    let mut add_one = 0;
    let mut l1 = l1;
    let mut l2 = l2;

    while l1.is_some() || l2.is_some() || add_one > 0 {
        l1 = l1.and_then(|x| {
            add_one += x.val;
            x.next
        });
        l2 = l2.and_then(|x| {
            add_one += x.val;
            x.next
        });
        let curr = add_one % 10;
        add_one /= 10;
        vec.push(ListNode::new(curr));
    }

    let mut tail = None;
    while let Some(mut node) = vec.pop() {
        node.next = tail;
        tail = Some(Box::new(node));
    }
    tail
}
