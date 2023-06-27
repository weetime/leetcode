use crate::list::list_node::ListNode;
pub fn get_decimal_value(head: Option<Box<ListNode>>) -> i32 {
    self::helper(head, &mut 0)
}

pub fn helper(head: Option<Box<ListNode>>, ticket: &mut u32) -> i32 {
    match head {
        Some(head) => {
            let mut res = self::helper(head.next, ticket);
            // res += head.val * (2_i32.pow(*ticket));
            res += head.val << *ticket;
            *ticket += 1;
            res
        }
        None => 0,
    }
}

pub fn get_decimal_value1(head: Option<Box<ListNode>>) -> i32 {
    if head.is_none() {
        return 0;
    }
    let mut curr = head;
    let mut res = 0;
    while let Some(node) = curr {
        curr = node.next;
        res = res * 2 + node.val;
    }
    res
}
