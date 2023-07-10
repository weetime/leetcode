pub mod list;
pub mod tree;
use list::list_node::ListNode;
use list::list_node::Solution as listSolution;
use std::cell::RefCell;
use std::rc::Rc;
use tree::tree::Solution as treeSolution;
use tree::tree::TreeNode;
fn main() {}

#[test]
fn test_reverse_list() {
    let res = listSolution::reverse_list(ListNode::sample());
    println!("{:?}", res);
}

#[test]
fn test_merge_two_lists() {
    let res = listSolution::merge_two_lists(ListNode::sample(), ListNode::sample());
    println!("{:?}", res);
}

#[test]
fn test_delete_duplicates() {
    let res = listSolution::delete_duplicates(ListNode::case01());
    println!("{:?}", res);
}

#[test]
fn test_delete_duplicates2() {
    let res = listSolution::delete_duplicates2(ListNode::case03());
    println!("{:?}", res);
}

#[test]
fn test_has_cycle() {
    let res = listSolution::has_cycle(ListNode::case01());
    println!("{:?}", res);
}

#[test]
fn test_ref() {
    let node = ListNode { val: 1, next: None };
    let a = Some(Box::new(node));
    let b = a.as_ref();
    let c = a.as_ref();
    assert_eq!(b, c);
    let b = a.as_ref().unwrap().as_ref();
    let c = a.as_ref().unwrap().as_ref();
    println!("{:p}, {:p}", b, c)
}

#[test]
fn test_remove_elements() {
    let res = listSolution::remove_elements(ListNode::case01(), 1);
    println!("{:?}", res);
}

#[test]
fn test_is_palindrome() {
    let res = listSolution::is_palindrome(ListNode::case02());
    println!("{:?}", res);
}

#[test]
fn middle_node() {
    let res = listSolution::middle_node(ListNode::sample());
    println!("{:?}", res);
    let res = listSolution::middle_node(ListNode::case01());
    println!("{:?}", res);
    let res = listSolution::middle_node(ListNode::case02());
    println!("{:?}", res);
}

#[test]
fn get_decimal_value() {
    let res = listSolution::get_decimal_value(ListNode::sample());
    println!("{:?}", res);
}

#[test]
fn remove_nth_from_end() {
    let res = listSolution::remove_nth_from_end(ListNode::sample(), 1);
    println!("{:?}", res);
}

#[test]
fn sort_list() {
    let res = listSolution::sort_list(ListNode::case03());
    println!("{:?}", res);
}

#[test]
fn insertion_sort_list() {
    let res = listSolution::insertion_sort_list(ListNode::case03());
    println!("{:?}", res);
}

#[test]
fn odd_even_list() {
    let res = listSolution::odd_even_list(ListNode::case03());
    println!("{:?}", res);
}

#[test]
fn merge_in_between() {
    let res = listSolution::merge_in_between(ListNode::case03(), 2, 2, ListNode::case03());
    println!("{:?}", res);
}

#[test]
fn delete_middle() {
    let res: Option<Box<ListNode>> = listSolution::delete_middle(ListNode::case03());
    println!("{:?}", res);
}

#[test]
fn add_two_numbers() {
    let res: Option<Box<ListNode>> =
        listSolution::add_two_numbers(ListNode::case03(), ListNode::case03());
    println!("{:?}", res);
}

#[test]
fn add_two_numbers2() {
    let res: Option<Box<ListNode>> =
        listSolution::add_two_numbers2(ListNode::case03(), ListNode::case03());
    println!("{:?}", res);
}

#[test]
fn swap_pairs() {
    let res: Option<Box<ListNode>> = listSolution::swap_pairs(ListNode::case03());
    println!("{:?}", res);
}

#[test]
fn rotate_right() {
    let res: Option<Box<ListNode>> = listSolution::rotate_right(ListNode::case03(), 1);
    println!("{:?}", res);
}

#[test]
fn search_bst() {
    let res: Option<Rc<RefCell<TreeNode>>> =
        treeSolution::search_bst(TreeNode::from_array(vec![1, 2, 3, 4, 5, 6]), 1);
    println!("{:?}", res);
}
