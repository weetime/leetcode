use crate::tree::tree::TreeNode;
use std::cell::RefCell;
use std::rc::Rc;

pub fn search_bst(root: Option<Rc<RefCell<TreeNode>>>, val: i32) -> Option<Rc<RefCell<TreeNode>>> {
    let mut curr = root;
    while let Some(node) = curr {
        let x = node.borrow().val;
        match x {
            x if x > val => curr = node.borrow_mut().left.take(),
            x if x < val => curr = node.borrow_mut().right.take(),
            _ => return Some(node),
        }
    }
    None
}
