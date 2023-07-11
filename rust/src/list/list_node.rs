use crate::list::add_two_numbers::add_two_numbers;
use crate::list::add_two_numbers2::add_two_numbers as add_two_numbers2;
use crate::list::delete_middle::delete_middle;
use crate::list::get_decimal_value::get_decimal_value;
use crate::list::has_cycle::has_cycle;
use crate::list::insertion_sort_list::insertion_sort_list;
use crate::list::is_palindrome::is_palindrome;
use crate::list::merge_in_between::merge_in_between;
// use crate::list::merge_two_lists::merge_two_lists;
use crate::list::middle_node::middle_node;
use crate::list::odd_even_list::odd_even_list;
use crate::list::pair_sum::pair_sum;
use crate::list::partition::partition;
use crate::list::remove_elements::remove_elements;
use crate::list::remove_nth_from_end::remove_nth_from_end;
use crate::list::reverse_list::reverse_list;
use crate::list::rotate_right::rotate_right;
use crate::list::sort_list::sort_list;
use crate::list::swap_pairs::swap_pairs;

#[derive(PartialEq, Eq, Clone, Debug)]
pub struct ListNode {
    pub val: i32,
    pub next: Option<Box<ListNode>>,
}

/**
 * 内置常见的链表操作
 * 1.构建链表的两种形式，头插法和尾插法
 * 2.取链表的长度
 * 3.找中点
 * 4.反转链表
 * 5.合并两个链表
 * 6.按索引删除节点
 * 7.按值删除节点
 * 8.排序
 */
impl ListNode {
    #[inline]
    pub fn new(val: i32) -> Self {
        ListNode { next: None, val }
    }

    // 头插法构建链表 [1,2,3,4,5] => 1->2->3->4->5
    pub fn from_array_by_head(arr: Vec<i32>) -> Option<Box<ListNode>> {
        let mut arr = arr;
        let mut head = None;
        while let Some(x) = arr.pop() {
            let mut node = Box::new(Self::new(x));
            node.next = head;
            head = Some(node);
        }
        head
    }

    // 尾插法构建链表 [1,2,3,4,5] => 1->2->3->4->5
    pub fn from_array_by_tail(arr: Vec<i32>) -> Option<Box<ListNode>> {
        let mut dummy = ListNode { val: 0, next: None };
        let mut tail = &mut dummy;
        for x in arr.iter() {
            let node = Box::new(ListNode::new(*x));
            tail.next = Some(node);
            tail = tail.next.as_mut().unwrap().as_mut();
        }
        dummy.next
    }
}

pub struct Solution {}
impl Solution {
    pub fn reverse_list(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        reverse_list(head)
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

    pub fn sort_list(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        sort_list(head)
    }

    pub fn insertion_sort_list(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        insertion_sort_list(head)
    }

    pub fn odd_even_list(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        odd_even_list(head)
    }

    pub fn merge_in_between(
        list1: Option<Box<ListNode>>,
        a: i32,
        b: i32,
        list2: Option<Box<ListNode>>,
    ) -> Option<Box<ListNode>> {
        merge_in_between(list1, a, b, list2)
    }

    pub fn delete_middle(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        delete_middle(head)
    }

    pub fn add_two_numbers(
        l1: Option<Box<ListNode>>,
        l2: Option<Box<ListNode>>,
    ) -> Option<Box<ListNode>> {
        add_two_numbers(l1, l2)
    }

    pub fn add_two_numbers2(
        l1: Option<Box<ListNode>>,
        l2: Option<Box<ListNode>>,
    ) -> Option<Box<ListNode>> {
        add_two_numbers2(l1, l2)
    }

    pub fn swap_pairs(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        swap_pairs(head)
    }

    pub fn rotate_right(head: Option<Box<ListNode>>, k: i32) -> Option<Box<ListNode>> {
        rotate_right(head, k)
    }

    pub fn partition(head: Option<Box<ListNode>>, x: i32) -> Option<Box<ListNode>> {
        partition(head, x)
    }

    pub fn pair_sum(head: Option<Box<ListNode>>) -> i32 {
        pair_sum(head)
    }
}
