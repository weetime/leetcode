use crate::list::list_node::ListNode;
pub struct Solution {}
impl Solution {
    /**
     * 解题思路：
     * 后序遍历，然后把值push到数组中即可
     * 需要注意的是传的arr是全局变量，并且是可变引用
     */
    pub fn reverse_print(head: Option<Box<ListNode>>) -> Vec<i32> {
        let mut arr: Vec<i32> = Vec::new();
        Self::dfs(head, &mut arr);
        arr
    }

    pub fn dfs(head: Option<Box<ListNode>>, arr: &mut Vec<i32>) {
        match head {
            Some(node) => {
                Self::dfs(node.next, arr);
                arr.push(node.val);
            }
            None => {}
        }
    }
}

#[test]
fn test() {
    let head_a = ListNode::from_array_by_head(vec![1, 2, 3, 4, 5]);
    let target = vec![5, 4, 3, 2, 1];
    let res = Solution::reverse_print(head_a.clone());
    assert_eq!(res, target);
}
