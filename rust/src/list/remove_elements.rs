use crate::list::list_node::ListNode;

pub struct Solution {}
impl Solution {
    pub fn remove_elements(head: Option<Box<ListNode>>, val: i32) -> Option<Box<ListNode>> {
        if head.is_none() {
            return head;
        }
        let mut dummy = ListNode { val: 0, next: head };
        let mut prev = &mut dummy;
        while let Some(mut node) = prev.next.take() {
            match node.val {
                x if x == val => {
                    prev.next = node.next.take();
                }
                _ => {
                    prev.next = Some(node);
                    prev = prev.next.as_mut().unwrap();
                }
            }
            // match 等价于 if 的写法
            // if node.val == val {
            //     prev.next = node.next.take();
            // } else {
            //     prev.next = Some(node);
            //     prev = prev.next.as_mut().unwrap();
            // }
        }

        dummy.next
    }

    // dfs 方式 其实是后序遍历
    pub fn remove_elements1(head: Option<Box<ListNode>>, val: i32) -> Option<Box<ListNode>> {
        match head {
            Some(mut head) => {
                let next_node = Self::remove_elements1(head.next.take(), val);
                if head.val == val {
                    // 如果在回归的时候值相等，则不组装当前node节点
                    next_node
                } else {
                    // 如果回归的时候值不等，则组装当前的node节点
                    head.next = next_node;
                    Some(head)
                }
            }
            None => None,
        }
    }

    // 下面这个解题思路纯粹是为了练习match，熟悉所有权的实现，主要是while和match 如果有let绑定，会转移所有权
    pub fn remove_elements2(head: Option<Box<ListNode>>, val: i32) -> Option<Box<ListNode>> {
        if head.is_none() {
            return head;
        }
        let mut dummy = Some(Box::new(ListNode { val: 0, next: head }));
        let mut prev = &mut dummy;
        while prev.as_ref().unwrap().next.is_some() {
            match prev.as_ref().unwrap().next.as_ref().unwrap().val {
                x if x == val => {
                    prev.as_mut().unwrap().next =
                        prev.as_mut().unwrap().next.as_mut().unwrap().next.take();
                }
                _ => {
                    prev = &mut prev.as_mut().unwrap().next;
                }
            }
        }

        dummy.unwrap().next
    }
}

#[test]
fn test() {
    let head_a = ListNode::from_array_by_head(vec![1, 1, 2, 3, 3, 4]);
    let target = ListNode::from_array_by_head(vec![1, 1, 3, 3, 4]);
    let res = Solution::remove_elements(head_a.clone(), 2);
    // println!("{:?}", res);
    assert_eq!(res, target);

    let res = Solution::remove_elements1(head_a.clone(), 2);
    assert_eq!(res, target);

    let res = Solution::remove_elements2(head_a.clone(), 2);
    assert_eq!(res, target);
}
