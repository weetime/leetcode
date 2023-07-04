use crate::list::list_node::ListNode;
pub fn odd_even_list(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
    let mut odd_node = head?;
    if odd_node.next.is_none() {
        return Some(odd_node);
    }

    let mut even_node = odd_node.next.take().unwrap();
    let (mut odd, mut even) = (&mut odd_node, &mut even_node);
    while even.next.is_some() {
        odd.next = even.next.take();
        odd = odd.next.as_mut().unwrap();
        if odd.next.is_none() {
            break;
        }
        even.next = odd.next.take();
        even = even.next.as_mut().unwrap();
    }
    odd.next = Some(even_node);
    Some(odd_node)
}
