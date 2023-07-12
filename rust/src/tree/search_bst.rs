use crate::tree::tree::TreeNode;
use std::cell::RefCell;
use std::rc::Rc;

pub struct Solution {}
impl Solution {
    pub fn search_bst(
        root: Option<Rc<RefCell<TreeNode>>>,
        val: i32,
    ) -> Option<Rc<RefCell<TreeNode>>> {
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
}

#[test]
fn test() {
    let tree = TreeNode::from_array(vec![1, 2, 3, 4, 5, 6]);
    let target = TreeNode::from_array(vec![1, 2, 3]);
    let res: Option<Rc<RefCell<TreeNode>>> = Solution::search_bst(tree.clone(), 2);
    // println!("{:?}", res);
    assert_eq!(res, target);
}
