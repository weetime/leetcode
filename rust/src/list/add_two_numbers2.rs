use crate::list::list_node::ListNode;
pub fn add_two_numbers1(
    l1: Option<Box<ListNode>>,
    l2: Option<Box<ListNode>>,
) -> Option<Box<ListNode>> {
    let (mut l1, mut l2) = (l1, l2);
    let mut stack1: Vec<i32> = vec![];
    let mut stack2: Vec<i32> = vec![];
    let mut dummy = None;
    while let Some(mut node) = l1.take() {
        l1 = node.next.take();
        stack1.push(node.val);
    }
    while let Some(mut node) = l2.take() {
        l2 = node.next.take();
        stack2.push(node.val);
    }

    let mut carry: i32 = 0;
    while !stack1.is_empty() || !stack2.is_empty() || carry > 0 {
        let mut sum: i32 = 0;
        sum += carry;
        if !stack1.is_empty() {
            match stack1.pop() {
                Some(n) => sum += n,
                None => todo!(),
            }
        }
        if !stack2.is_empty() {
            match stack2.pop() {
                Some(n) => sum += n,
                None => todo!(),
            }
        }
        carry = sum / 10;
        sum = sum % 10;

        let mut new_node = Box::new(ListNode::new(sum));
        new_node.next = dummy;
        dummy = Some(new_node);
    }
    dummy
}

pub fn add_two_numbers(
    l1: Option<Box<ListNode>>,
    l2: Option<Box<ListNode>>,
) -> Option<Box<ListNode>> {
    let (mut l1, mut l2) = (l1, l2);
    let mut stack1: Vec<i32> = vec![];
    let mut stack2: Vec<i32> = vec![];
    let mut dummy = None;
    while l1.is_some() {
        l1 = l1.and_then(|x| {
            stack1.push(x.val);
            x.next
        });
    }

    while l2.is_some() {
        l2 = l2.and_then(|x| {
            stack2.push(x.val);
            x.next
        });
    }

    let mut carry: i32 = 0;
    while !stack1.is_empty() || !stack2.is_empty() || carry > 0 {
        let mut sum: i32 = 0;
        sum += carry;
        if !stack1.is_empty() {
            match stack1.pop() {
                Some(n) => sum += n,
                None => todo!(),
            }
        }
        if !stack2.is_empty() {
            match stack2.pop() {
                Some(n) => sum += n,
                None => todo!(),
            }
        }
        carry = sum / 10;
        sum = sum % 10;

        let mut new_node = Box::new(ListNode::new(sum));
        new_node.next = dummy;
        dummy = Some(new_node);
    }
    dummy
}
