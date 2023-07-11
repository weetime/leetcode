pub mod list;
pub mod tree;
use list::list_node::ListNode;
use list::list_node::Solution as listSolution;
use std::cell::RefCell;
use std::rc::Rc;
use std::vec;
use tree::tree::Solution as treeSolution;
use tree::tree::TreeNode;

fn main() {}

#[test]
fn test_reverse_list() {
    let res = listSolution::reverse_list(ListNode::from_array_by_head(vec![1, 2, 3]));
    println!("{:?}", res);
}

#[test]
fn test_has_cycle() {
    let res = listSolution::has_cycle(ListNode::from_array_by_head(vec![1, 2, 3]));
    println!("{:?}", res);
}

#[test]
fn test_remove_elements() {
    let res = listSolution::remove_elements(ListNode::from_array_by_head(vec![1, 1, 2, 3]), 1);
    println!("{:?}", res);
}

#[test]
fn test_is_palindrome() {
    let res = listSolution::is_palindrome(ListNode::from_array_by_head(vec![1, 2, 3, 2, 1]));
    println!("{:?}", res);
}

#[test]
fn middle_node() {
    let res = listSolution::middle_node(ListNode::from_array_by_head(vec![1, 2, 3, 4]));
    println!("{:?}", res);
}

#[test]
fn get_decimal_value() {
    let res = listSolution::get_decimal_value(ListNode::from_array_by_head(vec![1, 0, 1]));
    println!("{:?}", res);
}

#[test]
fn remove_nth_from_end() {
    let res = listSolution::remove_nth_from_end(ListNode::from_array_by_head(vec![1, 2, 3, 4]), 1);
    println!("{:?}", res);
}

#[test]
fn sort_list() {
    let res = listSolution::sort_list(ListNode::from_array_by_head(vec![5, 4, 3, 2, 1]));
    println!("{:?}", res);
}

#[test]
fn insertion_sort_list() {
    let res = listSolution::insertion_sort_list(ListNode::from_array_by_head(vec![5, 4, 3, 2, 1]));
    println!("{:?}", res);
}

#[test]
fn odd_even_list() {
    let res = listSolution::odd_even_list(ListNode::from_array_by_head(vec![1, 2, 3, 4, 5, 6]));
    println!("{:?}", res);
}

#[test]
fn merge_in_between() {
    let res = listSolution::merge_in_between(
        ListNode::from_array_by_head(vec![1, 2, 3, 4, 5, 6]),
        2,
        2,
        ListNode::from_array_by_head(vec![7, 8, 9]),
    );
    println!("{:?}", res);
}

#[test]
fn delete_middle() {
    let res: Option<Box<ListNode>> =
        listSolution::delete_middle(ListNode::from_array_by_head(vec![1, 2, 3, 4]));
    println!("{:?}", res);
}

#[test]
fn add_two_numbers() {
    let res: Option<Box<ListNode>> = listSolution::add_two_numbers(
        ListNode::from_array_by_head(vec![7, 5, 6]),
        ListNode::from_array_by_head(vec![7, 5, 6]),
    );
    println!("{:?}", res);
}

#[test]
fn add_two_numbers2() {
    let res: Option<Box<ListNode>> = listSolution::add_two_numbers2(
        ListNode::from_array_by_head(vec![7, 5, 6]),
        ListNode::from_array_by_head(vec![7, 5, 6]),
    );
    println!("{:?}", res);
}

#[test]
fn swap_pairs() {
    let res: Option<Box<ListNode>> =
        listSolution::swap_pairs(ListNode::from_array_by_head(vec![1, 2, 3, 4, 5, 6]));
    println!("{:?}", res);
}

#[test]
fn rotate_right() {
    let res: Option<Box<ListNode>> =
        listSolution::rotate_right(ListNode::from_array_by_head(vec![1, 2, 3, 4, 5]), 1);
    println!("{:?}", res);
}

#[test]
fn pair_sum() {
    let res: i32 = listSolution::pair_sum(ListNode::from_array_by_tail(vec![1, 2, 3, 4, 5, 6]));
    println!("{:?}", res);
}

#[test]
fn partition() {
    let res: Option<Box<ListNode>> =
        listSolution::partition(ListNode::from_array_by_head(vec![3, 1, 2, 5, 4]), 1);
    println!("{:?}", res);
}
// -------------------- tree --------------------------------
#[test]
fn search_bst() {
    let res: Option<Rc<RefCell<TreeNode>>> =
        treeSolution::search_bst(TreeNode::from_array(vec![1, 2, 3, 4, 5, 6]), 1);
    println!("{:?}", res);
}
