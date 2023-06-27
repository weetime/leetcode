use crate::list::list_node::ListNode;

pub fn is_palindrome(head: Option<Box<ListNode>>) -> bool {
    if head.is_none() {
        return true;
    }
    self::dfs(head.clone(), &mut head.as_ref())
}

fn dfs(head_a: Option<Box<ListNode>>, head_b: &mut Option<&Box<ListNode>>) -> bool {
    match head_a {
        None => true,
        Some(h1) => {
            let mut res = self::dfs(h1.next, head_b);
            res = res && h1.val == head_b.unwrap().val;
            *head_b = head_b.unwrap().next.as_ref();
            res
        }
    }
}

pub fn is_palindrome2(head: Option<Box<ListNode>>) -> bool {
    if head.is_none() {
        return true;
    }
    let new = self::reverse(head.clone());
    self::check(head, new)
}

fn check(hea_a: Option<Box<ListNode>>, head_b: Option<Box<ListNode>>) -> bool {
    match (hea_a, head_b) {
        (None, None) => true,
        (None, Some(_)) => false,
        (Some(_), None) => false,
        (Some(l1), Some(l2)) => {
            if l1.val != l2.val {
                false
            } else {
                self::check(l1.next, l2.next)
            }
        }
    }
}

fn reverse(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
    if head.is_none() {
        return head;
    }
    let mut curr = head;
    let mut new = None;
    while let Some(mut node) = curr.take() {
        curr = node.next;
        node.next = new;
        new = Some(node);
    }
    new
}
