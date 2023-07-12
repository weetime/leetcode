use std::cell::RefCell;
use std::rc::Rc;

// Definition for a binary tree node.
#[derive(Debug, PartialEq, Eq)]
pub struct TreeNode {
    pub val: i32,
    pub left: Option<Rc<RefCell<TreeNode>>>,
    pub right: Option<Rc<RefCell<TreeNode>>>,
}

impl TreeNode {
    #[inline]
    pub fn new(val: i32) -> Self {
        TreeNode {
            val,
            left: None,
            right: None,
        }
    }
    pub fn from_array(arr: Vec<i32>) -> Option<Rc<RefCell<TreeNode>>> {
        let mut arr = arr;
        arr.sort_by(|a, b| a.cmp(b));
        let length = arr.len();
        Self::to_bst(arr, 0, length as i32)
    }

    pub fn to_bst(arr: Vec<i32>, s: i32, e: i32) -> Option<Rc<RefCell<TreeNode>>> {
        if s == e {
            return None;
        }
        let m = (s + e) >> 1;
        let mut t = TreeNode::new(arr[m as usize]);
        t.left = Self::to_bst(arr.to_vec(), s, m);
        t.right = Self::to_bst(arr.to_vec(), m + 1, e);

        Some(Rc::new(RefCell::new(t)))
    }
}
