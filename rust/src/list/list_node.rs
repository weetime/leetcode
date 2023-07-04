use crate::list::delete_duplicates::delete_duplicates;
use crate::list::delete_duplicates2::delete_duplicates as delete_duplicates2;
use crate::list::get_decimal_value::get_decimal_value;
use crate::list::has_cycle::has_cycle;
use crate::list::insertion_sort_list::insertion_sort_list;
use crate::list::is_palindrome::is_palindrome;
use crate::list::merge_two_lists::merge_two_lists;
use crate::list::middle_node::middle_node;
use crate::list::odd_even_list::odd_even_list;
use crate::list::remove_elements::remove_elements;
use crate::list::remove_nth_from_end::remove_nth_from_end;
use crate::list::reverse_list::reverse_list;
use crate::list::sort_list::sort_list;

#[derive(PartialEq, Eq, Clone, Debug)]
pub struct ListNode {
    pub val: i32,
    pub next: Option<Box<ListNode>>,
}

impl ListNode {
    #[inline]
    fn new(val: i32) -> Self {
        ListNode { next: None, val }
    }
    pub fn sample() -> Option<Box<ListNode>> {
        let mut node1: ListNode = ListNode::new(1);
        let mut node2: ListNode = ListNode::new(2);
        let node3: ListNode = ListNode::new(3);
        node2.next = Some(Box::new(node3));
        node1.next = Some(Box::new(node2));
        return Some(Box::new(node1));
    }

    pub fn case01() -> Option<Box<ListNode>> {
        let mut node1: ListNode = ListNode::new(1);
        let mut node2: ListNode = ListNode::new(2);
        let mut node3: ListNode = ListNode::new(3);
        let node4: ListNode = ListNode::new(4);
        node3.next = Some(Box::new(node4));
        node2.next = Some(Box::new(node3));
        node1.next = Some(Box::new(node2));
        return Some(Box::new(node1));
    }

    pub fn case02() -> Option<Box<ListNode>> {
        let mut node1: ListNode = ListNode::new(1);
        let mut node2: ListNode = ListNode::new(1);
        let mut node3: ListNode = ListNode::new(2);
        let node4: ListNode = ListNode::new(3);
        node3.next = Some(Box::new(node4));
        node2.next = Some(Box::new(node3));
        node1.next = Some(Box::new(node2));
        return Some(Box::new(node1));
    }

    pub fn case03() -> Option<Box<ListNode>> {
        return Some(Box::new(ListNode {
            val: 3,
            next: Some(Box::new(ListNode {
                val: 1,
                next: Some(Box::new(ListNode {
                    val: 2,
                    next: Some(Box::new(ListNode {
                        val: 5,
                        next: Some(Box::new(ListNode {
                            val: 4,
                            next: None, // next: Some(Box::new(ListNode { val: 6, next: None })),
                        })),
                    })),
                })),
            })),
        }));
    }
}

pub struct Solution {}
impl Solution {
    pub fn reverse_list(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        reverse_list(head)
    }
    pub fn merge_two_lists(
        list1: Option<Box<ListNode>>,
        list2: Option<Box<ListNode>>,
    ) -> Option<Box<ListNode>> {
        merge_two_lists(list1, list2)
    }

    pub fn delete_duplicates(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        delete_duplicates(head)
    }

    pub fn has_cycle(head: Option<Box<ListNode>>) -> bool {
        has_cycle(head)
    }

    pub fn remove_elements(head: Option<Box<ListNode>>, val: i32) -> Option<Box<ListNode>> {
        remove_elements(head, val)
    }

    pub fn is_palindrome(head: Option<Box<ListNode>>) -> bool {
        is_palindrome(head)
    }

    pub fn middle_node(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        middle_node(head)
    }

    pub fn get_decimal_value(head: Option<Box<ListNode>>) -> i32 {
        get_decimal_value(head)
    }

    pub fn remove_nth_from_end(head: Option<Box<ListNode>>, n: i32) -> Option<Box<ListNode>> {
        remove_nth_from_end(head, n)
    }

    pub fn delete_duplicates2(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        delete_duplicates2(head)
    }

    pub fn sort_list(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        sort_list(head)
    }

    pub fn insertion_sort_list(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        insertion_sort_list(head)
    }

    pub fn odd_even_list(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        odd_even_list(head)
    }
}
